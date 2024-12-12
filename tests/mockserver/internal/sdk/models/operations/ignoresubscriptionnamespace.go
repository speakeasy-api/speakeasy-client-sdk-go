// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type IgnoreSubscriptionNamespaceRequest struct {
	// The namespace name
	NamespaceName string `pathParam:"style=simple,explode=false,name=namespaceName"`
	// The existing subscription ID
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscriptionID"`
}

func (o *IgnoreSubscriptionNamespaceRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *IgnoreSubscriptionNamespaceRequest) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

type IgnoreSubscriptionNamespaceResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *IgnoreSubscriptionNamespaceResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
