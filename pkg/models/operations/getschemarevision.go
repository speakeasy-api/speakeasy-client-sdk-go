package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type GetSchemaRevisionPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type GetSchemaRevisionRequest struct {
	Retries    *utils.RetryConfig
	PathParams GetSchemaRevisionPathParams
}

type GetSchemaRevisionResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      *shared.Schema
	StatusCode  int64
}
