// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

type GetWorkspaceEventsRequest struct {
	// Filter to only return events created after this timestamp
	AfterCreatedAt *time.Time `queryParam:"style=form,explode=true,name=after_created_at"`
	// Filter to only return events corresponding to a particular gen_lock_id (gen_lock_id uniquely identifies a target)
	GenerateGenLockID *string `queryParam:"style=form,explode=true,name=generate_gen_lock_id"`
	// Unique identifier of the workspace.
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (g GetWorkspaceEventsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetWorkspaceEventsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetWorkspaceEventsRequest) GetAfterCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AfterCreatedAt
}

func (o *GetWorkspaceEventsRequest) GetGenerateGenLockID() *string {
	if o == nil {
		return nil
	}
	return o.GenerateGenLockID
}

func (o *GetWorkspaceEventsRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetWorkspaceEventsResponse struct {
	// Success
	CliEventBatch []shared.CliEvent
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetWorkspaceEventsResponse) GetCliEventBatch() []shared.CliEvent {
	if o == nil {
		return nil
	}
	return o.CliEventBatch
}

func (o *GetWorkspaceEventsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWorkspaceEventsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWorkspaceEventsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
