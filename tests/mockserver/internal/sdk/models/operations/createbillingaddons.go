// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type CreateBillingAddOnsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	OrganizationBillingAddOnResponse *components.OrganizationBillingAddOnResponse
}

func (o *CreateBillingAddOnsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateBillingAddOnsResponse) GetOrganizationBillingAddOnResponse() *components.OrganizationBillingAddOnResponse {
	if o == nil {
		return nil
	}
	return o.OrganizationBillingAddOnResponse
}
