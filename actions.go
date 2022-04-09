package incident

import (
	"context"
	"fmt"
)

// ActionsService handles communication with the actions related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/#tag/Actions
type ActionsService service

// ListActions list all actions for an organisation.
//
// API docs: https://api-docs.incident.io/#operation/Actions_List
func (s *ActionsService) ListActions(ctx context.Context, opts *ActionsListOptions) (*ActionsList, *Response, error) {
	u := "actions"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &ActionsList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// GetCustomField returns a single action.
//
// id represents the unique identifier for the action
//
// API docs: https://api-docs.incident.io/#operation/Actions_Show
func (s *ActionsService) GetAction(ctx context.Context, id string) (*ActionResponse, *Response, error) {
	u := fmt.Sprintf("actions/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO Should we return the action directly? Would be more userfriendly - Maybe talking to the Incident.io folks?
	v := &ActionResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
