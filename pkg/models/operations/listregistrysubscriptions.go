// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"net/http"
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
	SubscriptionType *shared.SubscriptionType `queryParam:"style=form,explode=true,name=subscription_type"`
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

func (o *ListRegistrySubscriptionsRequest) GetSubscriptionType() *shared.SubscriptionType {
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
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Classes []shared.RegistrySubscription
}

func (o *ListRegistrySubscriptionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRegistrySubscriptionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRegistrySubscriptionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListRegistrySubscriptionsResponse) GetClasses() []shared.RegistrySubscription {
	if o == nil {
		return nil
	}
	return o.Classes
}
