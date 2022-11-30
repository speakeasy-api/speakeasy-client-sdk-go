package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GenerateOpenAPISpecForAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GenerateOpenAPISpecForAPIEndpointRequest struct {
	Retries    *utils.RetryConfig
	PathParams GenerateOpenAPISpecForAPIEndpointPathParams
}

type GenerateOpenAPISpecForAPIEndpointResponse struct {
	ContentType             string
	Error                   *shared.Error
	GenerateOpenAPISpecDiff *shared.GenerateOpenAPISpecDiff
	StatusCode              int64
}
