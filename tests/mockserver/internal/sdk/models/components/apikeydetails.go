// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

type APIKeyDetails struct {
	AccountTypeV2   AccountType `json:"account_type_v2"`
	EnabledFeatures []string    `json:"enabled_features"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	FeatureFlags              []string `json:"feature_flags,omitempty"`
	GenerationAccessUnlimited *bool    `json:"generation_access_unlimited,omitempty"`
	OrgSlug                   string   `json:"org_slug"`
	TelemetryDisabled         bool     `json:"telemetry_disabled"`
	// Workspace creation timestamp.
	WorkspaceCreatedAt time.Time `json:"workspace_created_at"`
	WorkspaceID        string    `json:"workspace_id"`
	WorkspaceSlug      string    `json:"workspace_slug"`
}

func (a APIKeyDetails) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *APIKeyDetails) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *APIKeyDetails) GetAccountTypeV2() AccountType {
	if o == nil {
		return AccountType("")
	}
	return o.AccountTypeV2
}

func (o *APIKeyDetails) GetEnabledFeatures() []string {
	if o == nil {
		return []string{}
	}
	return o.EnabledFeatures
}

func (o *APIKeyDetails) GetFeatureFlags() []string {
	if o == nil {
		return nil
	}
	return o.FeatureFlags
}

func (o *APIKeyDetails) GetGenerationAccessUnlimited() *bool {
	if o == nil {
		return nil
	}
	return o.GenerationAccessUnlimited
}

func (o *APIKeyDetails) GetOrgSlug() string {
	if o == nil {
		return ""
	}
	return o.OrgSlug
}

func (o *APIKeyDetails) GetTelemetryDisabled() bool {
	if o == nil {
		return false
	}
	return o.TelemetryDisabled
}

func (o *APIKeyDetails) GetWorkspaceCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.WorkspaceCreatedAt
}

func (o *APIKeyDetails) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *APIKeyDetails) GetWorkspaceSlug() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceSlug
}
