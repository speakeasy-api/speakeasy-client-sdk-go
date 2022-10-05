package operations

import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
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
	PathParams RegisterSchemaPathParams
	Request    RegisterSchemaRequestBody `request:"mediaType=multipart/form-data"`
}

type RegisterSchemaResponses struct {
	Error *shared.Error
}

type RegisterSchemaResponse struct {
	ContentType string
	Responses   map[int64]map[string]RegisterSchemaResponses
	StatusCode  int64
}