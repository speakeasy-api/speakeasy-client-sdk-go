package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetVersionMetadataRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetVersionMetadataPathParams
}

type GetVersionMetadataResponse struct {
	ContentType     string
	Error           *shared.Error
	StatusCode      int64
	VersionMetadata []shared.VersionMetadata
}
