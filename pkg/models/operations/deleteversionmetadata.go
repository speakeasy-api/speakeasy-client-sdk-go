package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type DeleteVersionMetadataPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	MetaKey   string `pathParam:"style=simple,explode=false,name=metaKey"`
	MetaValue string `pathParam:"style=simple,explode=false,name=metaValue"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteVersionMetadataRequest struct {
	Retries    *utils.RetryConfig
	PathParams DeleteVersionMetadataPathParams
}

type DeleteVersionMetadataResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
