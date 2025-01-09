// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type SetWorkspaceFeatureFlagsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	WorkspaceFeatureFlagResponse *components.WorkspaceFeatureFlagResponse
}

func (o *SetWorkspaceFeatureFlagsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetWorkspaceFeatureFlagsResponse) GetWorkspaceFeatureFlagResponse() *components.WorkspaceFeatureFlagResponse {
	if o == nil {
		return nil
	}
	return o.WorkspaceFeatureFlagResponse
}
