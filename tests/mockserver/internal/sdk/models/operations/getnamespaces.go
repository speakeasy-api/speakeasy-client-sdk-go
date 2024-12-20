// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetNamespacesResponse struct {
	// OK
	GetNamespacesResponse *components.GetNamespacesResponse
	HTTPMeta              components.HTTPMetadata `json:"-"`
}

func (o *GetNamespacesResponse) GetGetNamespacesResponse() *components.GetNamespacesResponse {
	if o == nil {
		return nil
	}
	return o.GetNamespacesResponse
}

func (o *GetNamespacesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
