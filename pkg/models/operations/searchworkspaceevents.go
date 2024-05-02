// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type SearchWorkspaceEventsGlobals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *SearchWorkspaceEventsGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type SearchWorkspaceEventsRequest struct {
	// Unique identifier of the lint report digest.
	LintReportDigest *string `queryParam:"style=form,explode=true,name=lint_report_digest"`
	// Unique identifier of the openapi diff report digest.
	OpenapiDiffReportDigest *string `queryParam:"style=form,explode=true,name=openapi_diff_report_digest"`
	// Unique identifier of the source revision digest.
	SourceRevisionDigest *string `queryParam:"style=form,explode=true,name=source_revision_digest"`
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *SearchWorkspaceEventsRequest) GetLintReportDigest() *string {
	if o == nil {
		return nil
	}
	return o.LintReportDigest
}

func (o *SearchWorkspaceEventsRequest) GetOpenapiDiffReportDigest() *string {
	if o == nil {
		return nil
	}
	return o.OpenapiDiffReportDigest
}

func (o *SearchWorkspaceEventsRequest) GetSourceRevisionDigest() *string {
	if o == nil {
		return nil
	}
	return o.SourceRevisionDigest
}

func (o *SearchWorkspaceEventsRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type SearchWorkspaceEventsResponse struct {
	// Success
	CliEventBatch []shared.CliEvent
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *SearchWorkspaceEventsResponse) GetCliEventBatch() []shared.CliEvent {
	if o == nil {
		return nil
	}
	return o.CliEventBatch
}

func (o *SearchWorkspaceEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *SearchWorkspaceEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *SearchWorkspaceEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
