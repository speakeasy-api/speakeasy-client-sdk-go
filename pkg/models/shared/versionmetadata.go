package shared

import (
	"time"
)

// VersionMetadataInput
// A set of keys and associated values, attached to a particular version of an Api.
type VersionMetadataInput struct {
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}

// VersionMetadata
// A set of keys and associated values, attached to a particular version of an Api.
type VersionMetadata struct {
	APIID       string    `json:"api_id"`
	CreatedAt   time.Time `json:"created_at"`
	MetaKey     string    `json:"meta_key"`
	MetaValue   string    `json:"meta_value"`
	VersionID   string    `json:"version_id"`
	WorkspaceID string    `json:"workspace_id"`
}
