// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
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
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	SignedAccess *GetLintingReportSignedURLSignedAccess
}

func (o *GetLintingReportSignedURLResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetLintingReportSignedURLResponse) GetSignedAccess() *GetLintingReportSignedURLSignedAccess {
	if o == nil {
		return nil
	}
	return o.SignedAccess
}
