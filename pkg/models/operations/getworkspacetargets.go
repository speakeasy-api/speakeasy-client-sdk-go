// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type GetWorkspaceTargetsGlobals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *GetWorkspaceTargetsGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetWorkspaceTargetsRequest struct {
	// Filter to only return targets with events created after this timestamp
	AfterLastEventCreatedAt *time.Time `queryParam:"style=form,explode=true,name=after_last_event_created_at"`
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (g GetWorkspaceTargetsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetWorkspaceTargetsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetWorkspaceTargetsRequest) GetAfterLastEventCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AfterLastEventCreatedAt
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
