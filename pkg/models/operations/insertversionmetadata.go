package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type InsertVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type InsertVersionMetadataRequest struct {
	PathParams InsertVersionMetadataPathParams
	Request    shared.VersionMetadataInput `request:"mediaType=application/json"`
}

type InsertVersionMetadataResponse struct {
	ContentType     string
	Error           *shared.Error
	StatusCode      int64
	VersionMetadata *shared.VersionMetadata
}
