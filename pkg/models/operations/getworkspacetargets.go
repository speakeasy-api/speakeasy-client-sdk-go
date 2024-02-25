// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
)

type GetWorkspaceTargetsRequest struct {
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *GetWorkspaceTargetsRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetWorkspaceTargetsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TargetSDKList []shared.TargetSDK
}

func (o *GetWorkspaceTargetsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkspaceTargetsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkspaceTargetsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWorkspaceTargetsResponse) GetTargetSDKList() []shared.TargetSDK {
	if o == nil {
		return nil
	}
	return o.TargetSDKList
}
