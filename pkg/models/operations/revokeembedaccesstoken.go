package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
)

type RevokeEmbedAccessTokenPathParams struct {
	TokenID string `pathParam:"style=simple,explode=false,name=tokenID"`
}

type RevokeEmbedAccessTokenRequest struct {
	PathParams RevokeEmbedAccessTokenPathParams
}

type RevokeEmbedAccessTokenResponses struct {
	Error *shared.Error
}

type RevokeEmbedAccessTokenResponse struct {
	ContentType string
	Responses   map[int64]map[string]RevokeEmbedAccessTokenResponses
	StatusCode  int64
}