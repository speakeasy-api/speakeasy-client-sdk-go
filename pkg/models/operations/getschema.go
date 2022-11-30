package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetSchemaPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetSchemaPathParams
}

type GetSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      *shared.Schema
	StatusCode  int64
}
