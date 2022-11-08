package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIEndpointPathParams struct {
	APIEndpointID string `pathParam:"style=simple,explode=false,name=apiEndpointID"`
	APIID         string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID     string `pathParam:"style=simple,explode=false,name=versionID"`
}

type UpsertAPIEndpointRequest struct {
	PathParams UpsertAPIEndpointPathParams
	Request    shared.APIEndpointInput `request:"mediaType=application/json"`
}

type UpsertAPIEndpointResponse struct {
	APIEndpoint *shared.APIEndpoint
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
