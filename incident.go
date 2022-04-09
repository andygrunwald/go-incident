package incident

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/google/go-querystring/query"
)

const (
	// API URL
	// Right now, the API is only available in v1 and as a cloud version.
	apiURL = "https://api.incident.io/v1/"

	// User Agent that will be used for HTTP requests.
	// Should help to identify the source of the calls in case of emergency.
	userAgent = "go-incident"
)

var errNonNilContext = errors.New("context must be non-nil")

// A Client manages communication with the Incident.io API.
type Client struct {
	// clientMu protects the client during calls that modify the client.
	clientMu sync.Mutex
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	// We only have a cloud version of the Incident.io API.
	// However, we export it in case some companies run a Incident.io compatible API version.
	// BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL

	// API Key used for authentication against the API
	apiKey string

	// User agent used when communicating with the Incident.io API.
	UserAgent string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service

	// Services used for talking to different parts of the Incident.io API.
	Actions       *ActionsService
	CustomFields  *CustomFieldsService
	Severities    *SeveritiesService
	IncidentRoles *IncidentRolesService
	Incidents     *IncidentsService
}

type service struct {
	client *Client
}

// NewClient returns a new Incident.io API client.
// All endpoints require authentication, the apiKey should be set.
// If a nil httpClient is provided, a new http.Client will be used.
func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(apiURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		apiKey:    apiKey,
		UserAgent: userAgent,
	}
	c.common.client = c
	c.Actions = (*ActionsService)(&c.common)
	c.CustomFields = (*CustomFieldsService)(&c.common)
	c.Severities = (*SeveritiesService)(&c.common)
	c.IncidentRoles = (*IncidentRolesService)(&c.common)
	c.Incidents = (*IncidentsService)(&c.common)

	return c
}

// Client returns the http.Client used by this Incident.io client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.client
	return &clientCopy
}

// addOptions adds the parameters in opts as URL query parameters to s.
// opts must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if len(c.apiKey) > 0 {
		var bearer = "Bearer " + c.apiKey
		req.Header.Add("Authorization", bearer)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Response is a Incident.io API response. This wraps the standard http.Response
// returned from Incident.io. Right now, it only covers the native http response.
// We wrap it to enable future extension like providing convenient access to things like
// pagination information.
type Response struct {
	*http.Response
}

// newResponse creates a new Response for the provided http.Response.
// r must not be nil.
func newResponse(r *http.Response) *Response {
	response := &Response{Response: r}
	return response
}

// BareDo sends an API request and lets you handle the api response. If an error
// or API error occurs, the error will contain more information. Otherwise you
// are supposed to read and close the response's Body.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it is
// canceled or times out, ctx.Err() will be returned.
func (c *Client) BareDo(ctx context.Context, req *http.Request) (*Response, error) {
	if ctx == nil {
		return nil, errNonNilContext
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// If the error type is *url.Error, add the URL into the return val.
		if e, ok := err.(*url.Error); ok {
			if url, err := url.Parse(e.URL); err == nil {
				e.URL = url.String()
				return nil, e
			}
		}

		return nil, err
	}

	response := newResponse(resp)

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}
	return response, err
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred. If v implements the io.Writer interface,
// the raw response body will be written to v, without attempting to first
// decode it. If v is nil, and no error hapens, the response is returned as is.
//
// The provided ctx must be non-nil, if it is nil an error is returned. If it
// is canceled or times out, ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)
		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}
		if decErr != nil {
			err = decErr
		}
	}
	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range or equal to 202 Accepted.
// API error responses are expected to have response
// body, and a JSON response body that maps to ErrorResponse.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: r}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, errorResponse)
	}

	return errorResponse
}

// An ErrorResponse reports one or more errors caused by an API request.
//
// API docs: https://api-docs.incident.io/#section/Making-requests/Errors
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	// References the type of error
	Type string `json:"type"`
	// Contains the HTTP status
	Status int `json:"status"`
	// Request ID that can be provided to incident.io support to help debug questions with your API request
	RequestID string `json:"request_id"`

	// A list of individual errors, which go into detail about why the error occurred
	Errors []Error `json:"errors"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Type, r.RequestID, r.Errors)
}

// An Error reports more details on an individual error in an ErrorResponse.
//
// API docs: https://api-docs.incident.io/#section/Making-requests/Errors
type Error struct {
	// validation error code
	Code string `json:"code"`
	// Message describing the error.
	Message string      `json:"message"`
	Source  ErrorSource `json:"source"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v (code: %v)", e.Message, e.Code)
}

type ErrorSource struct {
	Field string `json:"field"`
}
