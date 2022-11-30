package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetEmbedAccessTokenQueryParams struct {
	Description *string         `queryParam:"style=form,explode=true,name=description"`
	Duration    *int64          `queryParam:"style=form,explode=true,name=duration"`
	Filters     *shared.Filters `queryParam:"serialization=json,name=filters"`
}

type GetEmbedAccessTokenRequest struct {
	Retries     *utils.RetryConfig
	QueryParams GetEmbedAccessTokenQueryParams
}

type GetEmbedAccessTokenResponse struct {
	ContentType              string
	EmbedAccessTokenResponse *shared.EmbedAccessTokenResponse
	Error                    *shared.Error
	StatusCode               int64
}
