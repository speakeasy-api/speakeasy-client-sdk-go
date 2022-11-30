package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type RevokeEmbedAccessTokenPathParams struct {
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

type RevokeEmbedAccessTokenRequest struct {
	Retries    *utils.RetryConfig
	PathParams RevokeEmbedAccessTokenPathParams
}

type RevokeEmbedAccessTokenResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
