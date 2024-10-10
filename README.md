# github.com/speakeasy-api/speakeasy-client-sdk-go

<!-- Start Summary [summary] -->
## Summary

Speakeasy API: The Speakeasy API allows teams to manage common operations with their APIs

For more information about the API: [The Speakeasy Platform Documentation](/docs)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [SDK Installation](#sdk-installation)
* [SDK Example Usage](#sdk-example-usage)
* [Available Resources and Operations](#available-resources-and-operations)
* [Retries](#retries)
* [Error Handling](#error-handling)
* [Server Selection](#server-selection)
* [Custom HTTP Client](#custom-http-client)
* [Authentication](#authentication)
* [Special Types](#special-types)
<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/speakeasy-api/speakeasy-client-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{})
	if err != nil {
		log.Fatal(err)
	}
	if res.Apis != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [APIEndpoints](docs/sdks/apiendpoints/README.md)

* [DeleteAPIEndpoint](docs/sdks/apiendpoints/README.md#deleteapiendpoint) - Delete an ApiEndpoint.
* [FindAPIEndpoint](docs/sdks/apiendpoints/README.md#findapiendpoint) - Find an ApiEndpoint via its displayName.
* [GenerateOpenAPISpecForAPIEndpoint](docs/sdks/apiendpoints/README.md#generateopenapispecforapiendpoint) - Generate an OpenAPI specification for a particular ApiEndpoint.
* [GeneratePostmanCollectionForAPIEndpoint](docs/sdks/apiendpoints/README.md#generatepostmancollectionforapiendpoint) - Generate a Postman collection for a particular ApiEndpoint.
* [GetAllAPIEndpoints](docs/sdks/apiendpoints/README.md#getallapiendpoints) - Get all Api endpoints for a particular apiID.
* [GetAllForVersionAPIEndpoints](docs/sdks/apiendpoints/README.md#getallforversionapiendpoints) - Get all ApiEndpoints for a particular apiID and versionID.
* [GetAPIEndpoint](docs/sdks/apiendpoints/README.md#getapiendpoint) - Get an ApiEndpoint.
* [UpsertAPIEndpoint](docs/sdks/apiendpoints/README.md#upsertapiendpoint) - Upsert an ApiEndpoint.

### [Apis](docs/sdks/apis/README.md)

* [DeleteAPI](docs/sdks/apis/README.md#deleteapi) - Delete an Api.
* [GenerateOpenAPISpec](docs/sdks/apis/README.md#generateopenapispec) - Generate an OpenAPI specification for a particular Api.
* [GeneratePostmanCollection](docs/sdks/apis/README.md#generatepostmancollection) - Generate a Postman collection for a particular Api.
* [GetAllAPIVersions](docs/sdks/apis/README.md#getallapiversions) - Get all Api versions for a particular ApiEndpoint.
* [GetApis](docs/sdks/apis/README.md#getapis) - Get a list of Apis for a given workspace
* [UpsertAPI](docs/sdks/apis/README.md#upsertapi) - Upsert an Api

### [Artifacts](docs/sdks/artifacts/README.md)

* [GetBlob](docs/sdks/artifacts/README.md#getblob) - Get blob for a particular digest
* [GetManifest](docs/sdks/artifacts/README.md#getmanifest) - Get manifest for a particular reference
* [GetNamespaces](docs/sdks/artifacts/README.md#getnamespaces) - Each namespace contains many revisions.
* [GetRevisions](docs/sdks/artifacts/README.md#getrevisions)
* [GetTags](docs/sdks/artifacts/README.md#gettags)
* [PostTags](docs/sdks/artifacts/README.md#posttags) - Add tags to an existing revision
* [Preflight](docs/sdks/artifacts/README.md#preflight) - Get access token for communicating with OCI distribution endpoints

### [Auth](docs/sdks/auth/README.md)

* [GetAccess](docs/sdks/auth/README.md#getaccess) - Get access allowances for a particular workspace
* [GetAccessToken](docs/sdks/auth/README.md#getaccesstoken) - Get or refresh an access token for the current workspace.
* [GetUser](docs/sdks/auth/README.md#getuser) - Get information about the current user.
* [ValidateAPIKey](docs/sdks/auth/README.md#validateapikey) - Validate the current api key.

### [Embeds](docs/sdks/embeds/README.md)

* [GetEmbedAccessToken](docs/sdks/embeds/README.md#getembedaccesstoken) - Get an embed access token for the current workspace.
* [GetValidEmbedAccessTokens](docs/sdks/embeds/README.md#getvalidembedaccesstokens) - Get all valid embed access tokens for the current workspace.
* [RevokeEmbedAccessToken](docs/sdks/embeds/README.md#revokeembedaccesstoken) - Revoke an embed access EmbedToken.

### [Events](docs/sdks/events/README.md)

* [GetEventsByTarget](docs/sdks/events/README.md#geteventsbytarget) - Load recent events for a particular workspace
* [GetTargets](docs/sdks/events/README.md#gettargets) - Load targets for a particular workspace
* [GetTargetsDeprecated](docs/sdks/events/README.md#gettargetsdeprecated) - Load targets for a particular workspace
* [Post](docs/sdks/events/README.md#post) - Post events for a specific workspace
* [Search](docs/sdks/events/README.md#search) - Search events for a particular workspace by any field

### [Github](docs/sdks/github/README.md)

* [CheckAccess](docs/sdks/github/README.md#checkaccess)
* [CheckPublishingPRs](docs/sdks/github/README.md#checkpublishingprs)
* [CheckPublishingSecrets](docs/sdks/github/README.md#checkpublishingsecrets)
* [ConfigureCodeSamples](docs/sdks/github/README.md#configurecodesamples)
* [ConfigureMintlifyRepo](docs/sdks/github/README.md#configuremintlifyrepo)
* [ConfigureTarget](docs/sdks/github/README.md#configuretarget)
* [GetAction](docs/sdks/github/README.md#getaction)
* [StorePublishingSecrets](docs/sdks/github/README.md#storepublishingsecrets)
* [TriggerAction](docs/sdks/github/README.md#triggeraction)

### [Metadata](docs/sdks/metadata/README.md)

* [DeleteVersionMetadata](docs/sdks/metadata/README.md#deleteversionmetadata) - Delete metadata for a particular apiID and versionID.
* [GetVersionMetadata](docs/sdks/metadata/README.md#getversionmetadata) - Get all metadata for a particular apiID and versionID.
* [InsertVersionMetadata](docs/sdks/metadata/README.md#insertversionmetadata) - Insert metadata for a particular apiID and versionID.

### [Organizations](docs/sdks/organizations/README.md)

* [Create](docs/sdks/organizations/README.md#create) - Create an organization
* [CreateFreeTrial](docs/sdks/organizations/README.md#createfreetrial) - Create a free trial for an organization
* [Get](docs/sdks/organizations/README.md#get) - Get organization
* [GetAll](docs/sdks/organizations/README.md#getall) - Get organizations for a user
* [GetUsage](docs/sdks/organizations/README.md#getusage) - Get billing usage summary for a particular organization

### [Reports](docs/sdks/reports/README.md)

* [GetChangesReportSignedURL](docs/sdks/reports/README.md#getchangesreportsignedurl) - Get the signed access url for the change reports for a particular document.
* [GetLintingReportSignedURL](docs/sdks/reports/README.md#getlintingreportsignedurl) - Get the signed access url for the linting reports for a particular document.
* [UploadReport](docs/sdks/reports/README.md#uploadreport) - Upload a report.

### [Requests](docs/sdks/requests/README.md)

* [GenerateRequestPostmanCollection](docs/sdks/requests/README.md#generaterequestpostmancollection) - Generate a Postman collection for a particular request.
* [GetRequestFromEventLog](docs/sdks/requests/README.md#getrequestfromeventlog) - Get information about a particular request.
* [QueryEventLog](docs/sdks/requests/README.md#queryeventlog) - Query the event log to retrieve a list of requests.

### [Schemas](docs/sdks/schemas/README.md)

* [DeleteSchema](docs/sdks/schemas/README.md#deleteschema) - Delete a particular schema revision for an Api.
* [DownloadSchema](docs/sdks/schemas/README.md#downloadschema) - Download the latest schema for a particular apiID.
* [DownloadSchemaRevision](docs/sdks/schemas/README.md#downloadschemarevision) - Download a particular schema revision for an Api.
* [GetSchema](docs/sdks/schemas/README.md#getschema) - Get information about the latest schema.
* [GetSchemaDiff](docs/sdks/schemas/README.md#getschemadiff) - Get a diff of two schema revisions for an Api.
* [GetSchemaRevision](docs/sdks/schemas/README.md#getschemarevision) - Get information about a particular schema revision for an Api.
* [GetSchemas](docs/sdks/schemas/README.md#getschemas) - Get information about all schemas associated with a particular apiID.
* [RegisterSchema](docs/sdks/schemas/README.md#registerschema) - Register a schema.

### [ShortURLs](docs/sdks/shorturls/README.md)

* [Create](docs/sdks/shorturls/README.md#create) - Shorten a URL.


### [Suggest](docs/sdks/suggest/README.md)

* [Suggest](docs/sdks/suggest/README.md#suggest) - Generate suggestions for improving an OpenAPI document.
* [SuggestItems](docs/sdks/suggest/README.md#suggestitems) - Generate generic suggestions for a list of items.
* [SuggestOpenAPI](docs/sdks/suggest/README.md#suggestopenapi) - (DEPRECATED) Generate suggestions for improving an OpenAPI document.
* [SuggestOpenAPIRegistry](docs/sdks/suggest/README.md#suggestopenapiregistry) - Generate suggestions for improving an OpenAPI document stored in the registry.

### [Workspaces](docs/sdks/workspaces/README.md)

* [Create](docs/sdks/workspaces/README.md#create) - Create a workspace
* [CreateToken](docs/sdks/workspaces/README.md#createtoken) - Create a token for a particular workspace
* [DeleteToken](docs/sdks/workspaces/README.md#deletetoken) - Delete a token for a particular workspace
* [Get](docs/sdks/workspaces/README.md#get) - Get workspace by context
* [GetAll](docs/sdks/workspaces/README.md#getall) - Get workspaces for a user
* [GetByID](docs/sdks/workspaces/README.md#getbyid) - Get workspace
* [GetFeatureFlags](docs/sdks/workspaces/README.md#getfeatureflags) - Get workspace feature flags
* [GetSettings](docs/sdks/workspaces/README.md#getsettings) - Get workspace settings
* [GetTeam](docs/sdks/workspaces/README.md#getteam) - Get team members for a particular workspace
* [GetTokens](docs/sdks/workspaces/README.md#gettokens) - Get tokens for a particular workspace
* [GrantAccess](docs/sdks/workspaces/README.md#grantaccess) - Grant a user access to a particular workspace
* [RevokeAccess](docs/sdks/workspaces/README.md#revokeaccess) - Revoke a user's access to a particular workspace
* [Update](docs/sdks/workspaces/README.md#update) - Update workspace details
* [UpdateSettings](docs/sdks/workspaces/README.md#updatesettings) - Update workspace settings

</details>
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `sdkerrors.SDKError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `DeleteAPI` function may return the following errors:

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

### Example

```go
package main

import (
	"context"
	"errors"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/sdkerrors"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	})
	if err != nil {

		var e *sdkerrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Name

You can override the default server globally using the `WithServer` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the names associated with the available servers:

| Name | Server | Variables |
| ----- | ------ | --------- |
| `prod` | `https://api.prod.speakeasyapi.dev` | None |

#### Example

```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServer("prod"),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL("https://api.prod.speakeasyapi.dev"),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                  | Type                  | Scheme                |
| --------------------- | --------------------- | --------------------- |
| `APIKey`              | apiKey                | API key               |
| `Bearer`              | http                  | HTTP Bearer           |
| `WorkspaceIdentifier` | apiKey                | API key               |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/retry"
	"log"
	"pkg/models/operations"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/retry"
	"log"
)

func main() {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.Apis.DeleteAPI(ctx, operations.DeleteAPIRequest{
		APIID:     "<value>",
		VersionID: "<value>",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
