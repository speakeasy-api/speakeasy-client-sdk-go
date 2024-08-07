// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetLintingReportSignedURLRequest struct {
	// The checksum of the document to retrieve the signed access url for.
	DocumentChecksum string `pathParam:"style=simple,explode=false,name=documentChecksum"`
}

func (o *GetLintingReportSignedURLRequest) GetDocumentChecksum() string {
	if o == nil {
		return ""
	}
	return o.DocumentChecksum
}

// GetLintingReportSignedURLSignedAccess - OK
type GetLintingReportSignedURLSignedAccess struct {
	URL string `json:"url"`
}

func (o *GetLintingReportSignedURLSignedAccess) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type GetLintingReportSignedURLResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	SignedAccess *GetLintingReportSignedURLSignedAccess
}

func (o *GetLintingReportSignedURLResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLintingReportSignedURLResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLintingReportSignedURLResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLintingReportSignedURLResponse) GetSignedAccess() *GetLintingReportSignedURLSignedAccess {
	if o == nil {
		return nil
	}
	return o.SignedAccess
}
