// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
	"time"
)

type GetPublishingTokenPublicMetadataRequest struct {
	// Unique identifier of the publishing token.
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

func (o *GetPublishingTokenPublicMetadataRequest) GetTokenID() string {
	if o == nil {
		return ""
	}
	return o.TokenID
}

// GetPublishingTokenPublicMetadataResponseBody - OK
type GetPublishingTokenPublicMetadataResponseBody struct {
	OrganizationID *string    `json:"organization_id,omitempty"`
	TargetID       *string    `json:"target_id,omitempty"`
	TargetResource *string    `json:"target_resource,omitempty"`
	ValidUntil     *time.Time `json:"valid_until,omitempty"`
	WorkspaceID    *string    `json:"workspace_id,omitempty"`
}

func (g GetPublishingTokenPublicMetadataResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPublishingTokenPublicMetadataResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPublishingTokenPublicMetadataResponseBody) GetOrganizationID() *string {
	if o == nil {
		return nil
	}
	return o.OrganizationID
}

func (o *GetPublishingTokenPublicMetadataResponseBody) GetTargetID() *string {
	if o == nil {
		return nil
	}
	return o.TargetID
}

func (o *GetPublishingTokenPublicMetadataResponseBody) GetTargetResource() *string {
	if o == nil {
		return nil
	}
	return o.TargetResource
}

func (o *GetPublishingTokenPublicMetadataResponseBody) GetValidUntil() *time.Time {
	if o == nil {
		return nil
	}
	return o.ValidUntil
}

func (o *GetPublishingTokenPublicMetadataResponseBody) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type GetPublishingTokenPublicMetadataResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Object *GetPublishingTokenPublicMetadataResponseBody
}

func (o *GetPublishingTokenPublicMetadataResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPublishingTokenPublicMetadataResponse) GetObject() *GetPublishingTokenPublicMetadataResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
