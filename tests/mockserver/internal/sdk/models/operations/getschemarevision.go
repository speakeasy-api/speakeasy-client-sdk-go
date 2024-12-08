// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetSchemaRevisionRequest struct {
	// The ID of the Api to retrieve schemas for.
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
	// The revision ID of the schema to retrieve.
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	// The version ID of the Api to delete metadata for.
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

func (o *GetSchemaRevisionRequest) GetAPIID() string {
	if o == nil {
		return ""
	}
	return o.APIID
}

func (o *GetSchemaRevisionRequest) GetRevisionID() string {
	if o == nil {
		return ""
	}
	return o.RevisionID
}

func (o *GetSchemaRevisionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

type GetSchemaRevisionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Schema *components.Schema
}

func (o *GetSchemaRevisionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSchemaRevisionResponse) GetSchema() *components.Schema {
	if o == nil {
		return nil
	}
	return o.Schema
}
