package incident

import (
	"context"
	"fmt"
)

// IncidentRolesService handles communication with the incident roles related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/#tag/Incident-Roles
type IncidentRolesService service

// ListIncidentRoles list all incident roles for an organisation.
//
// API docs: https://api-docs.incident.io/#operation/Incident%20Roles_List
func (s *IncidentRolesService) ListIncidentRoles(ctx context.Context) (*IncidentRolesList, *Response, error) {
	u := "incident_roles"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &IncidentRolesList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// GetIncidentRole returns a single incident role.
//
// id represents the unique identifier for the incident role
//
// API docs: https://api-docs.incident.io/#operation/Incident%20Roles_Show
func (s *IncidentRolesService) GetIncidentRole(ctx context.Context, id string) (*IncidentRoleResponse, *Response, error) {
	u := fmt.Sprintf("incident_roles/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO Should we return the incident role directly? Would be more userfriendly - Maybe talking to the Incident.io folks?
	v := &IncidentRoleResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}
