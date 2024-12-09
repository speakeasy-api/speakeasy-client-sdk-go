// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"io"
	"mockserver/internal/sdk/models/components"
)

type SuggestOpenAPIRegistryRequest struct {
	// Suggest options
	SuggestRequestBody *components.SuggestRequestBody `request:"mediaType=application/json"`
	NamespaceName      string                         `pathParam:"style=simple,explode=false,name=namespace_name"`
	// Tag or digest
	RevisionReference string `pathParam:"style=simple,explode=false,name=revision_reference"`
	XSessionID        string `header:"style=simple,explode=false,name=x-session-id"`
}

func (o *SuggestOpenAPIRegistryRequest) GetSuggestRequestBody() *components.SuggestRequestBody {
	if o == nil {
		return nil
	}
	return o.SuggestRequestBody
}

func (o *SuggestOpenAPIRegistryRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

func (o *SuggestOpenAPIRegistryRequest) GetRevisionReference() string {
	if o == nil {
		return ""
	}
	return o.RevisionReference
}

func (o *SuggestOpenAPIRegistryRequest) GetXSessionID() string {
	if o == nil {
		return ""
	}
	return o.XSessionID
}

type SuggestOpenAPIRegistryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// An overlay containing the suggested spec modifications.
	// The Close method must be called on this field, even if it is not used, to prevent resource leaks.
	Schema io.ReadCloser
}

func (o *SuggestOpenAPIRegistryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SuggestOpenAPIRegistryResponse) GetSchema() io.ReadCloser {
	if o == nil {
		return nil
	}
	return o.Schema
}
