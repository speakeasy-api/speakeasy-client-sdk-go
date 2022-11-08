package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type UpsertAPIPathParams struct {
	APIID string `pathParam:"style=simple,explode=false,name=apiID"`
}

type UpsertAPIRequest struct {
	PathParams UpsertAPIPathParams
	Request    shared.APIInput `request:"mediaType=application/json"`
}

type UpsertAPIResponse struct {
	API         *shared.API
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
