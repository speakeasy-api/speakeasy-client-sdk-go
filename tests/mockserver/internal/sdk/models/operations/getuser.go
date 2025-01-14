// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	User *components.User
}

func (o *GetUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
