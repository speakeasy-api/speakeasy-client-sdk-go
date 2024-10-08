// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type GetWorkspaceTargetsDeprecatedRequest struct {
	// Filter to only return targets with events created after this timestamp
	AfterLastEventCreatedAt *time.Time `queryParam:"style=form,explode=true,name=after_last_event_created_at"`
	// Unique identifier of the workspace.
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (g GetWorkspaceTargetsDeprecatedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetWorkspaceTargetsDeprecatedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetWorkspaceTargetsDeprecatedRequest) GetAfterLastEventCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AfterLastEventCreatedAt
}

func (o *GetWorkspaceTargetsDeprecatedRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GetWorkspaceTargetsDeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Success
	TargetSDKList []shared.TargetSDK
}

func (o *GetWorkspaceTargetsDeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkspaceTargetsDeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkspaceTargetsDeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWorkspaceTargetsDeprecatedResponse) GetTargetSDKList() []shared.TargetSDK {
	if o == nil {
		return nil
	}
	return o.TargetSDKList
}