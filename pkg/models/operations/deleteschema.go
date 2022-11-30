package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type DeleteSchemaPathParams struct {
	APIID      string `pathParam:"style=simple,explode=false,name=apiID"`
	RevisionID string `pathParam:"style=simple,explode=false,name=revisionID"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionID"`
}

type DeleteSchemaRequest struct {
	Retries    *utils.RetryConfig
	PathParams DeleteSchemaPathParams
}

type DeleteSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
