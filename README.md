# speakeasy-client-sdk-go

This is the Speakeasy API Client SDK for Go. It is generated from our OpenAPI spec found at https://docs.speakeasyapi.dev/openapi.yaml and used for interacting with the [Speakeasy API](https://docs.speakeasyapi.dev/docs/speakeasy-api/speakeasy-api).

This SDK was generated using Speakeasy's SDK Generator. For more information on how to use the generator to generate your own SDKs, please see the [Speakeasy Client SDK Generator Docs](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks).

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-api/speakeasy-client-sdk-go
```
<!-- End SDK Installation -->

## Example usage
```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
	ctx := context.Background()

	opts := []sdk.SDKOption{
		sdk.WithSecurity(shared.Security{
			APIKey: shared.SchemeAPIKey{
				APIKey: "YOUR_API_KEY", // Replace with your API key from your Speakeasy Workspace
			},
		}),
	}

	s := sdk.New(opts...)

	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
		QueryParams: operations.GetApisQueryParams{
			Metadata: map[string][]string{
				"label": {"1"},
			},
			Op: &operations.GetApisOp{
				And: true,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println(res.Error)
	} else {
		fmt.Println(res.Apis)
	}
}
```

<!-- Start SDK Available Operations -->
## SDK Available Operations

### SDK SDK

* `ValidateAPIKey` - Validate the current api key.

### ApiEndpoints

* `DeleteAPIEndpoint` - Delete an ApiEndpoint.
* `FindAPIEndpoint` - Find an ApiEndpoint via its displayName.
* `GenerateOpenAPISpecForAPIEndpoint` - Generate an OpenAPI specification for a particular ApiEndpoint.
* `GeneratePostmanCollectionForAPIEndpoint` - Generate a Postman collection for a particular ApiEndpoint.
* `GetAllAPIEndpoints` - Get all Api endpoints for a particular apiID.
* `GetAllForVersionAPIEndpoints` - Get all ApiEndpoints for a particular apiID and versionID.
* `GetAPIEndpoint` - Get an ApiEndpoint.
* `UpsertAPIEndpoint` - Upsert an ApiEndpoint.

### Apis

* `DeleteAPI` - Delete an Api.
* `GenerateOpenAPISpec` - Generate an OpenAPI specification for a particular Api.
* `GeneratePostmanCollection` - Generate a Postman collection for a particular Api.
* `GetAllAPIVersions` - Get all Api versions for a particular ApiEndpoint.
* `GetApis` - Get a list of Apis for a given workspace
* `UpsertAPI` - Upsert an Api

### Embeds

* `GetEmbedAccessToken` - Get an embed access token for the current workspace.
* `GetValidEmbedAccessTokens` - Get all valid embed access tokens for the current workspace.
* `RevokeEmbedAccessToken` - Revoke an embed access EmbedToken.

### Metadata

* `DeleteVersionMetadata` - Delete metadata for a particular apiID and versionID.
* `GetVersionMetadata` - Get all metadata for a particular apiID and versionID.
* `InsertVersionMetadata` - Insert metadata for a particular apiID and versionID.

### Plugins

* `GetPlugins` - Get all plugins for the current workspace.
* `RunPlugin` - Run a plugin
* `UpsertPlugin` - Upsert a plugin

### Requests

* `GenerateRequestPostmanCollection` - Generate a Postman collection for a particular request.
* `GetRequestFromEventLog` - Get information about a particular request.
* `QueryEventLog` - Query the event log to retrieve a list of requests.

### Schemas

* `DeleteSchema` - Delete a particular schema revision for an Api.
* `DownloadSchema` - Download the latest schema for a particular apiID.
* `DownloadSchemaRevision` - Download a particular schema revision for an Api.
* `GetSchema` - Get information about the latest schema.
* `GetSchemaDiff` - Get a diff of two schema revisions for an Api.
* `GetSchemaRevision` - Get information about a particular schema revision for an Api.
* `GetSchemas` - Get information about all schemas associated with a particular apiID.
* `RegisterSchema` - Register a schema.

<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
