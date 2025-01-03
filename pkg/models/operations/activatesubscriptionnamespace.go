// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ActivateSubscriptionNamespaceRequest struct {
	// The namespace name
	NamespaceName string `pathParam:"style=simple,explode=false,name=namespaceName"`
	// The existing subscription ID
	SubscriptionID string `pathParam:"style=simple,explode=false,name=subscriptionID"`
}

func (o *ActivateSubscriptionNamespaceRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *ActivateSubscriptionNamespaceRequest) GetSubscriptionID() string {
	if o == nil {
		return ""
	}
	return o.SubscriptionID
}

type ActivateSubscriptionNamespaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ActivateSubscriptionNamespaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ActivateSubscriptionNamespaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ActivateSubscriptionNamespaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}