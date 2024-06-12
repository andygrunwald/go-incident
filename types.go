package incident

import (
	"time"
)

const (
	// Custom Field Types
	CustomFieldTypeLink         = "link"
	CustomFieldTypeMultiSelect  = "multi_select"
	CustomFieldTypeNumeric      = "numeric"
	CustomFieldTypeSingleSelect = "single_select"
	CustomFieldTypeText         = "text"

	// User Roles
	UserRoleAdministrator = "administrator"
	UserRoleOwner         = "owner"
	UserRoleResponder     = "responder"
	UserRoleViewer        = "viewer"

	// Incident Roles
	IncidentRoleCustom = "custom"
	IncidentRoleLead   = "lead"

	// Incident Status
	IncidentStatusClosed        = "closed"
	IncidentStatusDeclined      = "declined"
	IncidentStatusFixing        = "fixing"
	IncidentStatusInvestigating = "investigating"
	IncidentStatusMonitoring    = "monitoring"
	IncidentStatusTriage        = "triage"

	// Incident Type
	IncidentTypeReal     = "real"
	IncidentTypeTest     = "test"
	IncidentTypeTutorial = "tutorial"

	// Incident Visbility
	IncidentVisibilityPrivate = "private"
	IncidentVisibilityPublic  = "public"

	// External Issue Reference Provider
	ExternalIssueReferenceProviderClubhouse  = "clubhouse"
	ExternalIssueReferenceProviderGithub     = "github"
	ExternalIssueReferenceProviderJira       = "jira"
	ExternalIssueReferenceProviderJiraServer = "jira_server"
	ExternalIssueReferenceProviderLinear     = "linear"

	// Action Status
	ActionStatusCompleted   = "completed"
	ActionStatusDeleted     = "deleted"
	ActionStatusNotDoing    = "not_doing"
	ActionStatusOutstanding = "outstanding"
)

// IncidentsListOptions defines parameters for IncidentsService.List.
type IncidentsListOptions struct {
	// Number of records to return
	PageSize int `url:"page_size,omitempty"`

	// An incident's ID. This endpoint will return a list of incidents created after this incident.
	After string `url:"after,omitempty"`

	// Filter for incidents in these statuses
	Status []string `url:"status,omitempty"`
}

type Incident struct {
	// The call URL attached to this incident
	CallUrl string `json:"call_url,omitempty"`

	// When the incident was created
	CreatedAt time.Time `json:"created_at"`

	// Custom field entries for this incident
	CustomFieldEntries []CustomFieldEntry `json:"custom_field_entries"`

	// Unique identifier for the incident
	Id string `json:"id"`

	// A list of who is assigned to each role for this incident
	IncidentRoleAssignments []IncidentRoleAssignment `json:"incident_role_assignments"`

	// Explanation of the incident
	Name string `json:"name,omitempty"`

	// Description of the incident
	PostmortemDocumentUrl string `json:"postmortem_document_url,omitempty"`

	// Reference to this incident, as displayed across the product
	Reference string   `json:"reference"`
	Creator   Actor    `json:"creator"`
	Severity  Severity `json:"severity"`

	// ID of the Slack channel in the organisation Slack workspace
	SlackChannelId string `json:"slack_channel_id"`

	// Name of the slack channel
	SlackChannelName string `json:"slack_channel_name,omitempty"`

	// Current status of the incident
	Status string `json:"status"`

	// Detailed description of the incident
	Summary string `json:"summary,omitempty"`

	// Incident lifecycle events and when they last occurred
	Timestamps *[]IncidentTimestamp `json:"timestamps,omitempty"`

	// Whether the incident is real, a test, or a tutorial
	Type string `json:"type"`

	// When the incident was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Whether the incident is public or private
	Visibility string `json:"visibility"`

	//
	// Values below only need to be set when an incident is being created
	//

	SeverityId string `json:"severity_id"`

	// Unique string used to de-duplicate incident create requests
	IdempotencyKey string `json:"idempotency_key"`
}

type IncidentsList struct {
	Incidents      []Incident      `json:"incidents"`
	PaginationMeta *PaginationMeta `json:"pagination_meta,omitempty"`
}

type PaginationMeta struct {
	// If provided, were records after a particular ID
	After string `json:"after,omitempty"`

	// What was the maximum number of results requested
	PageSize int64 `json:"page_size"`

	// How many matching records were there in total
	TotalRecordCount int64 `json:"total_record_count"`
}

type CustomFieldEntry struct {
	CustomField CustomFieldTypeInfo `json:"custom_field"`

	// List of custom field values set on this entry
	Values []CustomFieldValue `json:"values"`

	//
	// Values below only need to be set when an incident is being created
	//

	// ID of the custom field this entry is linked against
	CustomFieldId string `json:"custom_field_id"`
}

type CustomFieldTypeInfo struct {
	// Description of the custom field
	Description string `json:"description"`

	// Type of custom field
	FieldType string `json:"field_type"`

	// Unique identifier for the custom field
	Id string `json:"id"`

	// Human readable name for the custom field
	Name string `json:"name"`

	// What options are available for this custom field, if this field has options
	Options []CustomFieldOption `json:"options"`
}

type CustomFieldOption struct {
	// ID of the custom field this option belongs to
	CustomFieldId string `json:"custom_field_id"`

	// Unique identifier for the custom field option
	Id string `json:"id"`

	// Sort key used to order the custom field options correctly
	SortKey int64 `json:"sort_key"`

	// Human readable name for the custom field option
	Value string `json:"value"`
}

type CustomFieldValue struct {
	// Link value
	ValueLink string `json:"value_link,omitempty"`

	// Numeric value
	ValueNumeric string             `json:"value_numeric,omitempty"`
	ValueOption  *CustomFieldOption `json:"value_option,omitempty"`

	// Text value
	ValueText string `json:"value_text,omitempty"`

	//
	// Values below only need to be set when an incident is being created
	//

	// Unique identifier for the custom field value
	Id string `json:"id,omitempty"`

	// ID of the custom field option
	ValueOptionId string `json:"value_option_id,omitempty"`
}

type IncidentRoleAssignment struct {
	Assignee *User        `json:"assignee,omitempty"`
	Role     IncidentRole `json:"role"`
}

type User struct {
	// Unique identifier of the user
	Id string `json:"id"`

	// Name of the user
	Name string `json:"name"`

	// Role of the user
	Role string `json:"role"`

	// Email of the user
	Email string `json:"email"`

	// Slack User ID of the user
	SlackUserID string `json:"slack_user_id"`
}

type IncidentRole struct {
	// When the action was created
	CreatedAt time.Time `json:"created_at"`

	// Describes the purpose of the role
	Description string `json:"description"`

	// Unique identifier for the role
	Id string `json:"id"`

	// Provided to whoever is nominated for the role
	Instructions string `json:"instructions"`

	// Human readable name of the incident role
	Name string `json:"name"`

	// Whether incident require this role to be set
	Required bool `json:"required"`

	// Type of incident role
	RoleType string `json:"role_type"`

	// Short human readable name for Slack
	Shortform string `json:"shortform"`

	// When the action was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type Actor struct {
	ApiKey *APIKey `json:"api_key,omitempty"`
	User   *User   `json:"user,omitempty"`
}

type Severity struct {
	// When the action was created
	CreatedAt time.Time `json:"created_at"`

	// Description of the severity
	Description string `json:"description"`

	// Unique identifier of the severity
	Id string `json:"id"`

	// Human readable name of the severity
	Name string `json:"name"`

	// Rank to help sort severities (lower numbers are less severe)
	Rank int64 `json:"rank"`

	// When the action was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type IncidentTimestamp struct {
	// When this last occurred, if it did
	LastOccurredAt time.Time `json:"last_occurred_at,omitempty"`

	// Name of the lifecycle event
	Name string `json:"name"`
}

type APIKey struct {
	// Unique identifier for this API key
	Id string `json:"id"`

	// The name of the API key, for the user's reference
	Name string `json:"name"`
}

type IncidentResponse struct {
	Incident Incident `json:"incident"`
}

type SeveritiesList struct {
	Severities []Severity `json:"severities"`
}

type SeverityResponse struct {
	Severity Severity `json:"severity"`
}

type IncidentRolesList struct {
	IncidentRoles []IncidentRole `json:"incident_roles"`
}

type IncidentRoleResponse struct {
	IncidentRole IncidentRole `json:"incident_role"`
}

type CustomField struct {
	// When the action was created
	CreatedAt time.Time `json:"created_at"`

	// Description of the custom field
	Description string `json:"description"`

	// Type of custom field
	FieldType string `json:"field_type"`

	// Unique identifier for the custom field
	Id string `json:"id"`

	// Human readable name for the custom field
	Name string `json:"name"`

	// What options are available for this custom field, if this field has options
	Options []CustomFieldOption `json:"options"`

	// Whether a custom field should be required in the incident close modal
	RequireBeforeClosure bool `json:"require_before_closure"`

	// Whether a custom field should be required in the incident creation modal
	RequireBeforeCreation bool `json:"require_before_creation"`

	// Whether a custom field should be shown in the incident close modal
	ShowBeforeClosure bool `json:"show_before_closure"`

	// Whether a custom field should be shown in the incident creation modal
	ShowBeforeCreation bool `json:"show_before_creation"`

	// When the action was last updated
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomFieldsList struct {
	CustomFields []CustomField `json:"custom_fields"`
}

type CustomFieldResponse struct {
	CustomField CustomField `json:"custom_field"`
}

type ActionsListOptions struct {
	// Find actions related to this incident
	IncidentId string `url:"incident_id,omitempty"`

	// Filter to actions marked as being follow up actions
	IsFollowUp bool `url:"is_follow_up,omitempty"`

	// Filter to actions from incidents of the given mode.
	// If not set, only actions from real incidents are returned
	// Enum: "real" "test" "tutorial"
	IncidentMode string `url:"incident_mode,omitempty"`
}

type ActionsList struct {
	Actions []Action `json:"actions"`
}

type Action struct {
	// When the action was completed
	CompletedAt time.Time `json:"completed_at,omitempty"`

	// When the action was created
	CreatedAt time.Time `json:"created_at"`

	// Description of the action
	Description            string                  `json:"description"`
	ExternalIssueReference *ExternalIssueReference `json:"external_issue_reference,omitempty"`

	// Whether an action is marked as follow-up
	FollowUp bool `json:"follow_up"`

	// Unique identifier for the action
	Id string `json:"id"`

	// Unique identifier of the incident the action belongs to
	IncidentId string `json:"incident_id"`

	// Status of the action
	Status string `json:"status"`

	// When the action was last updated
	UpdatedAt time.Time `json:"updated_at"`

	// Assignee of the action
	Assignee *User `json:"assignee,omitempty"`
}

type ExternalIssueReference struct {
	// Human readable ID for the issue
	IssueName string `json:"issue_name"`

	// URL linking directly to the action in the issue tracker
	IssuePermalink string `json:"issue_permalink"`

	// ID of the issue tracker provider
	Provider string `json:"provider"`
}

type ActionResponse struct {
	Action Action `json:"action"`
}
