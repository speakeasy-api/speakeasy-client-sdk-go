package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type DeleteAPIPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteAPIRequest struct {
	PathParams DeleteAPIPathParams
}

type DeleteAPIResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
