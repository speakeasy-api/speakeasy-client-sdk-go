// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type CreateRequestBody struct {
	// URL to shorten
	URL string `json:"url"`
}

func (o *CreateRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type CreateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ShortURL *shared.ShortURL
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResponse) GetShortURL() *shared.ShortURL {
	if o == nil {
		return nil
	}
	return o.ShortURL
}

func (o *CreateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
