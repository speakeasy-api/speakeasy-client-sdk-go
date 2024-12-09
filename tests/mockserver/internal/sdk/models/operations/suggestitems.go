// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type SuggestItemsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// One suggestion per item. Guaranteed to be the same length as the input items.
	Strings []string
}

func (o *SuggestItemsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SuggestItemsResponse) GetStrings() []string {
	if o == nil {
		return nil
	}
	return o.Strings
}
