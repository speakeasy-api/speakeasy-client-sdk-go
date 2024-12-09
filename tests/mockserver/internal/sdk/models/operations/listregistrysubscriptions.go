// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type ListRegistrySubscriptionsGlobals struct {
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *ListRegistrySubscriptionsGlobals) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type ListRegistrySubscriptionsRequest struct {
	// The event type
	EventType *string `queryParam:"style=form,explode=true,name=event_type"`
	// The namespace name
	NamespaceName *string `queryParam:"style=form,explode=true,name=namespace_name"`
	// The subscription type
	SubscriptionType *components.SubscriptionType `queryParam:"style=form,explode=true,name=subscription_type"`
	// The tag
	Tag *string `queryParam:"style=form,explode=true,name=tag"`
	// The workspace ID
	WorkspaceID *string `pathParam:"style=simple,explode=false,name=workspace_id"`
}

func (o *ListRegistrySubscriptionsRequest) GetEventType() *string {
	if o == nil {
		return nil
	}
	return o.EventType
}

func (o *ListRegistrySubscriptionsRequest) GetNamespaceName() *string {
	if o == nil {
		return nil
	}
	return o.NamespaceName
}

func (o *ListRegistrySubscriptionsRequest) GetSubscriptionType() *components.SubscriptionType {
	if o == nil {
		return nil
	}
	return o.SubscriptionType
}

func (o *ListRegistrySubscriptionsRequest) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

func (o *ListRegistrySubscriptionsRequest) GetWorkspaceID() *string {
	if o == nil {
		return nil
	}
	return o.WorkspaceID
}

type ListRegistrySubscriptionsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Classes []components.RegistrySubscription
}

func (o *ListRegistrySubscriptionsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRegistrySubscriptionsResponse) GetClasses() []components.RegistrySubscription {
	if o == nil {
		return nil
	}
	return o.Classes
}
