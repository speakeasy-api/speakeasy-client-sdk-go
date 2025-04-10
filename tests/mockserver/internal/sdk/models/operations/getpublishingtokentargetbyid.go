// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetPublishingTokenTargetByIDRequest struct {
	// Unique identifier of the publishing token.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

func (o *GetPublishingTokenTargetByIDRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

// GetPublishingTokenTargetByIDResponseBody - OK
type GetPublishingTokenTargetByIDResponseBody struct {
}

type GetPublishingTokenTargetByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Object *GetPublishingTokenTargetByIDResponseBody
}

func (o *GetPublishingTokenTargetByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPublishingTokenTargetByIDResponse) GetObject() *GetPublishingTokenTargetByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
