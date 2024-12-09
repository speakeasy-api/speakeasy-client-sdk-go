// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreateOrganizationResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Organization *components.Organization
}

func (o *CreateOrganizationResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateOrganizationResponse) GetOrganization() *components.Organization {
	if o == nil {
		return nil
	}
	return o.Organization
}
