package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GenerateOpenAPISpecPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecRequest struct {
	Retries    *utils.RetryConfig
	PathParams GenerateOpenAPISpecPathParams
}

type GenerateOpenAPISpecResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int64
}
