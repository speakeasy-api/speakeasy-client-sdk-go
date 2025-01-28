// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"time"
)

type CompositeSpecMetadata struct {
	// The subscription ID for the remote source subscription, if applicable. This indicates that the namespace is created by a remote source and thus is composite.
	SubscriptionID       string                           `json:"subscription_id"`
	SubscriptionSettings RemoteSourceSubscriptionSettings `json:"subscription_settings"`
}

func (o *CompositeSpecMetadata) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

func (o *CompositeSpecMetadata) GetSubscriptionSettings() RemoteSourceSubscriptionSettings {
	if o == nil {
		return RemoteSourceSubscriptionSettings{}
	}
	return o.SubscriptionSettings
}

// Namespace - A namespace contains many revisions.
type Namespace struct {
	ArchivedAt            *time.Time             `json:"archived_at,omitempty"`
	CompositeSpecMetadata *CompositeSpecMetadata `json:"composite_spec_metadata,omitempty"`
	CreatedAt             time.Time              `json:"created_at"`
	// {organization_slug}/{workspace_slug}/{namespace_name}
	ID                     string                    `json:"id"`
	LatestRevisionMetadata *RevisionContentsMetadata `json:"latest_revision_metadata,omitempty"`
	// A human-readable name for the namespace.
	Name string `json:"name"`
	// Indicates whether the namespace is publicly accessible
	Public    *bool     `json:"public,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (n Namespace) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(n, "", false)
}

func (n *Namespace) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &n, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Namespace) GetArchivedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *Namespace) GetCompositeSpecMetadata() *CompositeSpecMetadata {
	if o == nil {
		return nil
	}
	return o.CompositeSpecMetadata
}

func (o *Namespace) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Namespace) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Namespace) GetLatestRevisionMetadata() *RevisionContentsMetadata {
	if o == nil {
		return nil
	}
	return o.LatestRevisionMetadata
}

func (o *Namespace) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Namespace) GetPublic() *bool {
	if o == nil {
		return nil
	}
	return o.Public
}

func (o *Namespace) GetUpdatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedAt
}
