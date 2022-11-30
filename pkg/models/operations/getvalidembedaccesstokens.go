package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetValidEmbedAccessTokensRequest struct {
	Retries *utils.RetryConfig
}

type GetValidEmbedAccessTokensResponse struct {
	ContentType string
	EmbedTokens []shared.EmbedToken
	Error       *shared.Error
	StatusCode  int64
}
