package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type DownloadSchemaRevisionPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DownloadSchemaRevisionRequest struct {
	Retries    *utils.RetryConfig
	PathParams DownloadSchemaRevisionPathParams
}

type DownloadSchemaRevisionResponse struct {
	ContentType string
	Error       *shared.Error
	Schema      []byte
	StatusCode  int64
}
