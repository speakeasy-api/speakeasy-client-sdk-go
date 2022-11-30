package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetRequestFromEventLogPathParams struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GetRequestFromEventLogRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetRequestFromEventLogPathParams
}

type GetRequestFromEventLogResponse struct {
	ContentType      string
	Error            *shared.Error
	StatusCode       int64
	UnboundedRequest *shared.UnboundedRequest
}
