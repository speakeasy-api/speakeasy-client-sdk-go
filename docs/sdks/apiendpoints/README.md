# APIEndpoints
(*APIEndpoints*)

## Overview

REST APIs for managing ApiEndpoint entities

### Available Operations

* [DeleteAPIEndpoint](#deleteapiendpoint) - Delete an ApiEndpoint.
* [FindAPIEndpoint](#findapiendpoint) - Find an ApiEndpoint via its displayName.
* [GenerateOpenAPISpecForAPIEndpoint](#generateopenapispecforapiendpoint) - Generate an OpenAPI specification for a particular ApiEndpoint.
* [GeneratePostmanCollectionForAPIEndpoint](#generatepostmancollectionforapiendpoint) - Generate a Postman collection for a particular ApiEndpoint.
* [GetAllAPIEndpoints](#getallapiendpoints) - Get all Api endpoints for a particular apiID.
* [GetAllForVersionAPIEndpoints](#getallforversionapiendpoints) - Get all ApiEndpoints for a particular apiID and versionID.
* [GetAPIEndpoint](#getapiendpoint) - Get an ApiEndpoint.
* [UpsertAPIEndpoint](#upsertapiendpoint) - Upsert an ApiEndpoint.

## DeleteAPIEndpoint

Delete an ApiEndpoint. This will also delete all associated Request Logs (if using a Postgres datastore).

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.DeleteAPIEndpoint(ctx, operations.DeleteAPIEndpointRequest{
        APIEndpointID: "<value>",
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.DeleteAPIEndpointRequest](../../pkg/models/operations/deleteapiendpointrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.DeleteAPIEndpointResponse](../../pkg/models/operations/deleteapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## FindAPIEndpoint

Find an ApiEndpoint via its displayName (set by operationId from a registered OpenAPI schema).
This is useful for finding the ID of an ApiEndpoint to use in the /v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID} endpoints.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.FindAPIEndpoint(ctx, operations.FindAPIEndpointRequest{
        APIID: "<value>",
        DisplayName: "Don.Strosin",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.FindAPIEndpointRequest](../../pkg/models/operations/findapiendpointrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.FindAPIEndpointResponse](../../pkg/models/operations/findapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GenerateOpenAPISpecForAPIEndpoint

This endpoint will generate a new operation in any registered OpenAPI document if the operation does not already exist in the document.
Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GenerateOpenAPISpecForAPIEndpoint(ctx, operations.GenerateOpenAPISpecForAPIEndpointRequest{
        APIEndpointID: "<value>",
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GenerateOpenAPISpecDiff != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.GenerateOpenAPISpecForAPIEndpointRequest](../../pkg/models/operations/generateopenapispecforapiendpointrequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.GenerateOpenAPISpecForAPIEndpointResponse](../../pkg/models/operations/generateopenapispecforapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GeneratePostmanCollectionForAPIEndpoint

Generates a postman collection that allows the endpoint to be called from postman variables produced for any path/query/header parameters included in the OpenAPI document.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GeneratePostmanCollectionForAPIEndpoint(ctx, operations.GeneratePostmanCollectionForAPIEndpointRequest{
        APIEndpointID: "<value>",
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PostmanCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                  | Type                                                                                                                                       | Required                                                                                                                                   | Description                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                                      | :heavy_check_mark:                                                                                                                         | The context to use for the request.                                                                                                        |
| `request`                                                                                                                                  | [operations.GeneratePostmanCollectionForAPIEndpointRequest](../../pkg/models/operations/generatepostmancollectionforapiendpointrequest.md) | :heavy_check_mark:                                                                                                                         | The request object to use for the request.                                                                                                 |


### Response

**[*operations.GeneratePostmanCollectionForAPIEndpointResponse](../../pkg/models/operations/generatepostmancollectionforapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllAPIEndpoints

Get all Api endpoints for a particular apiID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAllAPIEndpoints(ctx, operations.GetAllAPIEndpointsRequest{
        APIID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoints != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetAllAPIEndpointsRequest](../../pkg/models/operations/getallapiendpointsrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetAllAPIEndpointsResponse](../../pkg/models/operations/getallapiendpointsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAllForVersionAPIEndpoints

Get all ApiEndpoints for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAllForVersionAPIEndpoints(ctx, operations.GetAllForVersionAPIEndpointsRequest{
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoints != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GetAllForVersionAPIEndpointsRequest](../../pkg/models/operations/getallforversionapiendpointsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |


### Response

**[*operations.GetAllForVersionAPIEndpointsResponse](../../pkg/models/operations/getallforversionapiendpointsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetAPIEndpoint

Get an ApiEndpoint.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAPIEndpoint(ctx, operations.GetAPIEndpointRequest{
        APIEndpointID: "<value>",
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetAPIEndpointRequest](../../pkg/models/operations/getapiendpointrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetAPIEndpointResponse](../../pkg/models/operations/getapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpsertAPIEndpoint

Upsert an ApiEndpoint. If the ApiEndpoint does not exist it will be created, otherwise it will be updated.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.UpsertAPIEndpoint(ctx, operations.UpsertAPIEndpointRequest{
        APIEndpoint: shared.APIEndpointInput{
            APIEndpointID: "<value>",
            Description: "Public-key systematic attitude",
            DisplayName: "Camille.Schaefer11",
            Method: "<value>",
            Path: "/Library",
            VersionID: "<value>",
        },
        APIEndpointID: "<value>",
        APIID: "<value>",
        VersionID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIEndpoint != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.UpsertAPIEndpointRequest](../../pkg/models/operations/upsertapiendpointrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.UpsertAPIEndpointResponse](../../pkg/models/operations/upsertapiendpointresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
