package incident

import (
	"context"
	"fmt"
)

// SeveritiesService handles communication with the severity related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/#tag/Severities
type SeveritiesService service

// List list all incident severities for an organisation.
//
// API docs: https://api-docs.incident.io/#operation/Severities_List
func (s *SeveritiesService) List(ctx context.Context) (*SeveritiesList, *Response, error) {
	u := "severities"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &SeveritiesList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Get returns a single incident severity.
//
// id represents the unique identifier for the severity
//
// API docs: https://api-docs.incident.io/#operation/Severities_Show
func (s *SeveritiesService) Get(ctx context.Context, id string) (*SeverityResponse, *Response, error) {
	u := fmt.Sprintf("severities/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO Should we return the severity directly? Would be more userfriendly - Maybe talking to the Incident.io folks?
	v := &SeverityResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
