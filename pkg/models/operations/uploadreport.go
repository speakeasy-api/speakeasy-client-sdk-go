// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type File struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// UploadReportRequestBody - The report file to upload provided as a multipart/form-data file segment.
type UploadReportRequestBody struct {
	Data shared.Report `multipartForm:"name=data,json"`
	File File          `multipartForm:"file"`
}

func (o *UploadReportRequestBody) GetData() shared.Report {
	if o == nil {
		return shared.Report{}
	}
	return o.Data
}

func (o *UploadReportRequestBody) GetFile() File {
	if o == nil {
		return File{}
	}
	return o.File
}

// UploadReportUploadedReport - OK
type UploadReportUploadedReport struct {
	URL string `json:"url"`
}

func (o *UploadReportUploadedReport) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type UploadReportResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UploadedReport *UploadReportUploadedReport
}

func (o *UploadReportResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UploadReportResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UploadReportResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UploadReportResponse) GetUploadedReport() *UploadReportUploadedReport {
	if o == nil {
		return nil
	}
	return o.UploadedReport
}
