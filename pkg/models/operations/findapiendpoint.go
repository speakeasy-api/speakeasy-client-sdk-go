package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type FindAPIEndpointPathParams struct {
	APIID       string `pathParam:"style=simple,explode=false,name=apiID"`
	DisplayName string `pathParam:"style=simple,explode=false,name=displayName"`
	VersionID   string `pathParam:"style=simple,explode=false,name=versionID"`
}

type FindAPIEndpointRequest struct {
	PathParams FindAPIEndpointPathParams
}

type FindAPIEndpointResponse struct {
	APIEndpoint *shared.APIEndpoint
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
