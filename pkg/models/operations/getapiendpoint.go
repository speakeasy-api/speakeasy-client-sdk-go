package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetAPIEndpointRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetAPIEndpointPathParams
}

type GetAPIEndpointResponse struct {
	APIEndpoint *shared.APIEndpoint
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
