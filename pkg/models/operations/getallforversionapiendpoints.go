package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetAllForVersionAPIEndpointsPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAllForVersionAPIEndpointsRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetAllForVersionAPIEndpointsPathParams
}

type GetAllForVersionAPIEndpointsResponse struct {
	APIEndpoints []shared.APIEndpoint
	ContentType  string
	Error        *shared.Error
	StatusCode   int64
}
