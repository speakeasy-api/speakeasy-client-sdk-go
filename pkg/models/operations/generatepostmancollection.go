package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GeneratePostmanCollectionPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GeneratePostmanCollectionRequest struct {
	Retries    *utils.RetryConfig
	PathParams GeneratePostmanCollectionPathParams
}

type GeneratePostmanCollectionResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int64
}
