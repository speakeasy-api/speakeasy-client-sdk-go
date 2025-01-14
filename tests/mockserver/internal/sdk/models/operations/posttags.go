// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PostTagsRequest struct {
	// A JSON representation of the tags to add
	AddTags       *components.AddTags `request:"mediaType=application/json"`
	NamespaceName string              `pathParam:"style=simple,explode=false,name=namespace_name"`
}

func (o *PostTagsRequest) GetAddTags() *components.AddTags {
	if o == nil {
		return nil
	}
	return o.AddTags
}

func (o *PostTagsRequest) GetNamespaceName() string {
	if o == nil {
		return ""
	}
	return o.NamespaceName
}

type PostTagsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *PostTagsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
