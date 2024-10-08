// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"net/http"
)

type GeneratePostmanCollectionRequest struct {
	// The ID of the Api to generate a Postman collection for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The version ID of the Api to generate a Postman collection for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GeneratePostmanCollectionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GeneratePostmanCollectionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GeneratePostmanCollectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	PostmanCollection io.ReadCloser
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GeneratePostmanCollectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GeneratePostmanCollectionResponse) GetPostmanCollection() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.PostmanCollection
}

func (o *GeneratePostmanCollectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GeneratePostmanCollectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
