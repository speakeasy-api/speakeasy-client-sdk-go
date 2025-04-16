// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type TargetResource string

const (
	TargetResourceDocument TargetResource = "document"
)

func (e TargetResource) ToPointer() *TargetResource {
	return &e
}
func (e *TargetResource) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "document":
		*e = TargetResource(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TargetResource: %v", v)
	}
}

// PublishingToken - A token used to publish to a target
type PublishingToken struct {
	CreatedAt      time.Time      `json:"created_at"`
	CreatedBy      string         `json:"created_by"`
	ID             string         `json:"id"`
	OrganizationID string         `json:"organization_id"`
	TargetID       string         `json:"target_id"`
	TargetResource TargetResource `json:"target_resource"`
	Token          string         `json:"token"`
	TokenName      string         `json:"token_name"`
	UpdatedAt      *time.Time     `json:"updated_at,omitempty"`
	UpdatedBy      *string        `json:"updated_by,omitempty"`
	ValidUntil     time.Time      `json:"valid_until"`
	WorkspaceID    string         `json:"workspace_id"`
}

func (p PublishingToken) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PublishingToken) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PublishingToken) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *PublishingToken) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *PublishingToken) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *PublishingToken) GetOrganizationID() string {
	if o == nil {
		return ""
	}
	return o.OrganizationID
}

func (o *PublishingToken) GetTargetID() string {
	if o == nil {
		return ""
	}
	return o.TargetID
}

func (o *PublishingToken) GetTargetResource() TargetResource {
	if o == nil {
		return TargetResource("")
	}
	return o.TargetResource
}

func (o *PublishingToken) GetToken() string {
	if o == nil {
		return ""
	}
	return o.Token
}

func (o *PublishingToken) GetTokenName() string {
	if o == nil {
		return ""
	}
	return o.TokenName
}

func (o *PublishingToken) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *PublishingToken) GetUpdatedBy() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedBy
}

func (o *PublishingToken) GetValidUntil() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.ValidUntil
}

func (o *PublishingToken) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}
