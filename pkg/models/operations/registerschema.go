package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/utils"
)

type RegisterSchemaPathParams struct {
	APIID     string `pathParam:"style=simple,explode=false,name=apiID"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionID"`
}

type RegisterSchemaRequestBodyFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type RegisterSchemaRequestBody struct {
	File RegisterSchemaRequestBodyFile `multipartForm:"file"`
}

type RegisterSchemaRequest struct {
	Retries    *utils.RetryConfig
	PathParams RegisterSchemaPathParams
	Request    RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
}

type RegisterSchemaResponse struct {
	ContentType string
	Error       *shared.Error
	StatusCode  int64
}
