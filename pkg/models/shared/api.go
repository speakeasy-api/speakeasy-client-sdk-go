package shared

import (
	"time"
)

// APIInput
// An Api is representation of a API (a collection of API Endpoints) within the Speakeasy Platform.
type APIInput struct {
	APIID       string              `json:"api_id"`
	Description string              `json:"description"`
	MetaData    map[string][]string `json:"meta_data,omitempty"`
	VersionID   string              `json:"version_id"`
}

// API
// An Api is representation of a API (a collection of API Endpoints) within the Speakeasy Platform.
type API struct {
	APIID       string              `json:"api_id"`
	CreatedAt   time.Time           `json:"created_at"`
	Description string              `json:"description"`
	Matched     *bool               `json:"matched,omitempty"`
	MetaData    map[string][]string `json:"meta_data,omitempty"`
	UpdatedAt   time.Time           `json:"updated_at"`
	VersionID   string              `json:"version_id"`
	WorkspaceID string              `json:"workspace_id"`
}
