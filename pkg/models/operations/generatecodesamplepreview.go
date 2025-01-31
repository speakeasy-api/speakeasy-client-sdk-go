// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GenerateCodeSamplePreviewResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UsageSnippets *shared.UsageSnippets
}

func (o *GenerateCodeSamplePreviewResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GenerateCodeSamplePreviewResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GenerateCodeSamplePreviewResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GenerateCodeSamplePreviewResponse) GetUsageSnippets() *shared.UsageSnippets {
	if o == nil {
		return nil
	}
	return o.UsageSnippets
}
