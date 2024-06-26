// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type OrganizationAccountType string

const (
	OrganizationAccountTypeFree       OrganizationAccountType = "free"
	OrganizationAccountTypeScaleUp    OrganizationAccountType = "scale-up"
	OrganizationAccountTypeEnterprise OrganizationAccountType = "enterprise"
)

func (e OrganizationAccountType) ToPointer() *OrganizationAccountType {
	return &e
}
func (e *OrganizationAccountType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "free":
		fallthrough
	case "scale-up":
		fallthrough
	case "enterprise":
		*e = OrganizationAccountType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OrganizationAccountType: %v", v)
	}
}

// Organization - A speakeasy organization
type Organization struct {
	AccountType       OrganizationAccountType `json:"account_type"`
	CreatedAt         *time.Time              `json:"created_at,omitempty"`
	ID                string                  `json:"id"`
	Name              string                  `json:"name"`
	Slug              *string                 `json:"slug,omitempty"`
	TelemetryDisabled bool                    `json:"telemetry_disabled"`
	UpdatedAt         *time.Time              `json:"updated_at,omitempty"`
}

func (o Organization) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *Organization) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Organization) GetAccountType() OrganizationAccountType {
	if o == nil {
		return OrganizationAccountType("")
	}
	return o.AccountType
}

func (o *Organization) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Organization) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Organization) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Organization) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *Organization) GetTelemetryDisabled() bool {
	if o == nil {
		return false
	}
	return o.TelemetryDisabled
}

func (o *Organization) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
