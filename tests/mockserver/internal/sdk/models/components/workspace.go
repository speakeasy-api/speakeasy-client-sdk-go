// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// Workspace - A speakeasy workspace
type Workspace struct {
	CreatedAt      time.Time `json:"created_at"`
	ID             string    `json:"id"`
	Inactive       *bool     `json:"inactive,omitempty"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organization_id"`
	Slug           string    `json:"slug"`
	// Deprecated. Use organization.telemetry_disabled instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	TelemetryDisabled *bool     `json:"telemetry_disabled,omitempty"`
	UpdatedAt         time.Time `json:"updated_at"`
	Verified          bool      `json:"verified"`
}

func (w Workspace) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(w, "", false)
}

func (w *Workspace) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &w, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Workspace) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Workspace) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Workspace) GetInactive() *bool {
	if o == nil {
		return nil
	}
	return o.Inactive
}

func (o *Workspace) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Workspace) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *Workspace) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *Workspace) GetTelemetryDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.TelemetryDisabled
}

func (o *Workspace) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}

func (o *Workspace) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}