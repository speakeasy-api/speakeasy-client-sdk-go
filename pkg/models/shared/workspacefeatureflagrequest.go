// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// WorkspaceFeatureFlagRequest - A request to add workspace feature flags
type WorkspaceFeatureFlagRequest struct {
	FeatureFlags []WorkspaceFeatureFlag `json:"feature_flags"`
}

func (o *WorkspaceFeatureFlagRequest) GetFeatureFlags() []WorkspaceFeatureFlag {
	if o == nil {
		return []WorkspaceFeatureFlag{}
	}
	return o.FeatureFlags
}