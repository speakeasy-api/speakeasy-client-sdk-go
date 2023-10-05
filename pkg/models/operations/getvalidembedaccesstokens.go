// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"net/http"
)

type GetValidEmbedAccessTokensResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	EmbedTokens []shared.EmbedToken
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetValidEmbedAccessTokensResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetValidEmbedAccessTokensResponse) GetEmbedTokens() []shared.EmbedToken {
	if o == nil {
		return nil
	}
	return o.EmbedTokens
}

func (o *GetValidEmbedAccessTokensResponse) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *GetValidEmbedAccessTokensResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetValidEmbedAccessTokensResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
