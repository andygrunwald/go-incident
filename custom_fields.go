package incident

import (
	"context"
	"fmt"
)

// CustomFieldsService handles communication with the custom field related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/#tag/Custom-Fields
type CustomFieldsService service

// List list all custom fields for an organisation.
//
// API docs: https://api-docs.incident.io/#operation/Custom%20Fields_List
func (s *CustomFieldsService) List(ctx context.Context) (*CustomFieldsList, *Response, error) {
	u := "custom_fields"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &CustomFieldsList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Get returns a single custom field.
//
// id represents the unique identifier for the custom field
//
// API docs: https://api-docs.incident.io/#operation/Custom%20Fields_Show
func (s *CustomFieldsService) Get(ctx context.Context, id string) (*CustomFieldResponse, *Response, error) {
	u := fmt.Sprintf("custom_fields/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO Should we return the custom field directly? Would be more userfriendly - Maybe talking to the Incident.io folks?
	v := &CustomFieldResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
