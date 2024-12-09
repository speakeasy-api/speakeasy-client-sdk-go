// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"mockserver/internal/sdk/utils"
	"time"
)

// An UnboundedRequest represents the HAR content capture by Speakeasy when logging a request.
type UnboundedRequest struct {
	// Creation timestamp.
	CreatedAt time.Time `json:"created_at"`
	// The HAR content of the request.
	Har string `json:"har"`
	// The size of the HAR content in bytes.
	HarSizeBytes int64 `json:"har_size_bytes"`
	// The ID of this request.
	RequestID string `json:"request_id"`
	// The workspace ID this request was made to.
	WorkspaceID string `json:"workspace_id"`
}

func (u UnboundedRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UnboundedRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UnboundedRequest) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *UnboundedRequest) GetHar() string {
	if o == nil {
		return ""
	}
	return o.Har
}

func (o *UnboundedRequest) GetHarSizeBytes() int64 {
	if o == nil {
		return 0
	}
	return o.HarSizeBytes
}

func (o *UnboundedRequest) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

func (o *UnboundedRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
