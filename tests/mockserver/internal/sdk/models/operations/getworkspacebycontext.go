// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetWorkspaceByContextResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	WorkspaceAndOrganization *components.WorkspaceAndOrganization
}

func (o *GetWorkspaceByContextResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWorkspaceByContextResponse) GetWorkspaceAndOrganization() *components.WorkspaceAndOrganization {
	if o == nil {
		return nil
	}
	return o.WorkspaceAndOrganization
}