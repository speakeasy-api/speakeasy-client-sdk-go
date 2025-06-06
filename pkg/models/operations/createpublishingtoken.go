// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

// CreatePublishingTokenRequestBody - The publishing token to create
type CreatePublishingTokenRequestBody struct {
	TargetID       string    `json:"target_id"`
	TargetResource string    `json:"target_resource"`
	TokenName      string    `json:"token_name"`
	ValidUntil     time.Time `json:"valid_until"`
}

func (c CreatePublishingTokenRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatePublishingTokenRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreatePublishingTokenRequestBody) GetTargetID() string {
	if o == nil {
		return ""
	}
	return o.TargetID
}

func (o *CreatePublishingTokenRequestBody) GetTargetResource() string {
	if o == nil {
		return ""
	}
	return o.TargetResource
}

func (o *CreatePublishingTokenRequestBody) GetTokenName() string {
	if o == nil {
		return ""
	}
	return o.TokenName
}

func (o *CreatePublishingTokenRequestBody) GetValidUntil() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ValidUntil
}

type CreatePublishingTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PublishingToken *shared.PublishingToken
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreatePublishingTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreatePublishingTokenResponse) GetPublishingToken() *shared.PublishingToken {
	if o == nil {
		return nil
	}
	return o.PublishingToken
}

func (o *CreatePublishingTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreatePublishingTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
