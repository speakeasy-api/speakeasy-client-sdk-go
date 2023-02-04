package shared

import (
	"time"
)

// Plugin
// A plugin is a short script that is run against ingested requests
type Plugin struct {
	Code        string     `json:"code"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	EvalHash    *string    `json:"eval_hash,omitempty"`
	PluginID    string     `json:"plugin_id"`
	Title       string     `json:"title"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	WorkspaceID string     `json:"workspace_id"`
}
