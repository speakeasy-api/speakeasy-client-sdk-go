// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type GenerateCodeSamplePreviewResponse struct {
	// Successfully returned codeSample overlay file
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredApplicationJSONResponseStream io.ReadCloser
	// Successfully returned codeSample overlay file
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	TwoHundredApplicationXYamlResponseStream io.ReadCloser
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GenerateCodeSamplePreviewResponse) GetTwoHundredApplicationJSONResponseStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONResponseStream
}

func (o *GenerateCodeSamplePreviewResponse) GetTwoHundredApplicationXYamlResponseStream() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationXYamlResponseStream
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
