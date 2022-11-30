package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GenerateRequestPostmanCollectionPathParams struct {
	RequestID string `pathParam:"style=simple,explode=false,name=requestID"`
}

type GenerateRequestPostmanCollectionRequest struct {
	Retries    *utils.RetryConfig
	PathParams GenerateRequestPostmanCollectionPathParams
}

type GenerateRequestPostmanCollectionResponse struct {
	ContentType       string
	Error             *shared.Error
	PostmanCollection []byte
	StatusCode        int64
}
