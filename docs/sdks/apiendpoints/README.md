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
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.DeleteAPIEndpoint(ctx, operations.DeleteAPIEndpointRequest{
        APIEndpointID: "bus Gasoline Southwest",
        APIID: "candela",
        VersionID: "indexing maxime bypassing",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteAPIEndpointRequest](../../models/operations/deleteapiendpointrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteAPIEndpointResponse](../../models/operations/deleteapiendpointresponse.md), error**


## FindAPIEndpoint

Find an ApiEndpoint via its displayName (set by operationId from a registered OpenAPI schema).
This is useful for finding the ID of an ApiEndpoint to use in the /v1/apis/{apiID}/version/{versionID}/api_endpoints/{apiEndpointID} endpoints.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.FindAPIEndpoint(ctx, operations.FindAPIEndpointRequest{
        APIID: "insignificant",
        DisplayName: "payment",
        VersionID: "Highland North Awesome",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.FindAPIEndpointRequest](../../models/operations/findapiendpointrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.FindAPIEndpointResponse](../../models/operations/findapiendpointresponse.md), error**


## GenerateOpenAPISpecForAPIEndpoint

This endpoint will generate a new operation in any registered OpenAPI document if the operation does not already exist in the document.
Returns the original document and the newly generated document allowing a diff to be performed to see what has changed.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GenerateOpenAPISpecForAPIEndpoint(ctx, operations.GenerateOpenAPISpecForAPIEndpointRequest{
        APIEndpointID: "Tactics ew explicabo",
        APIID: "gray act",
        VersionID: "bah",
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

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.GenerateOpenAPISpecForAPIEndpointRequest](../../models/operations/generateopenapispecforapiendpointrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.GenerateOpenAPISpecForAPIEndpointResponse](../../models/operations/generateopenapispecforapiendpointresponse.md), error**


## GeneratePostmanCollectionForAPIEndpoint

Generates a postman collection that allows the endpoint to be called from postman variables produced for any path/query/header parameters included in the OpenAPI document.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GeneratePostmanCollectionForAPIEndpoint(ctx, operations.GeneratePostmanCollectionForAPIEndpointRequest{
        APIEndpointID: "silver Bronze Outdoors",
        APIID: "Virtual admirable Southeast",
        VersionID: "Northeast Bicycle mark",
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

| Parameter                                                                                                                              | Type                                                                                                                                   | Required                                                                                                                               | Description                                                                                                                            |
| -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                  | :heavy_check_mark:                                                                                                                     | The context to use for the request.                                                                                                    |
| `request`                                                                                                                              | [operations.GeneratePostmanCollectionForAPIEndpointRequest](../../models/operations/generatepostmancollectionforapiendpointrequest.md) | :heavy_check_mark:                                                                                                                     | The request object to use for the request.                                                                                             |


### Response

**[*operations.GeneratePostmanCollectionForAPIEndpointResponse](../../models/operations/generatepostmancollectionforapiendpointresponse.md), error**


## GetAllAPIEndpoints

Get all Api endpoints for a particular apiID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAllAPIEndpoints(ctx, operations.GetAllAPIEndpointsRequest{
        APIID: "Director JSON candela",
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetAllAPIEndpointsRequest](../../models/operations/getallapiendpointsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.GetAllAPIEndpointsResponse](../../models/operations/getallapiendpointsresponse.md), error**


## GetAllForVersionAPIEndpoints

Get all ApiEndpoints for a particular apiID and versionID.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAllForVersionAPIEndpoints(ctx, operations.GetAllForVersionAPIEndpointsRequest{
        APIID: "THX Tuna sparkling",
        VersionID: "Assurance",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetAllForVersionAPIEndpointsRequest](../../models/operations/getallforversionapiendpointsrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetAllForVersionAPIEndpointsResponse](../../models/operations/getallforversionapiendpointsresponse.md), error**


## GetAPIEndpoint

Get an ApiEndpoint.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.GetAPIEndpoint(ctx, operations.GetAPIEndpointRequest{
        APIEndpointID: "Optimized male SSD",
        APIID: "Soul Land",
        VersionID: "Androgyne Hydrogen",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetAPIEndpointRequest](../../models/operations/getapiendpointrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetAPIEndpointResponse](../../models/operations/getapiendpointresponse.md), error**


## UpsertAPIEndpoint

Upsert an ApiEndpoint. If the ApiEndpoint does not exist it will be created, otherwise it will be updated.

### Example Usage

```go
package main

import(
	"context"
	"log"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/pkg/models/operations"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: "",
        }),
    )

    ctx := context.Background()
    res, err := s.APIEndpoints.UpsertAPIEndpoint(ctx, operations.UpsertAPIEndpointRequest{
        APIEndpointInput: shared.APIEndpointInput{
            APIEndpointID: "Convertible Turner Cotton",
            Description: "Triple-buffered optimal methodology",
            DisplayName: "Langosh Northwest Mabelle",
            Method: "South",
            Path: "/rescue",
            VersionID: "East Land viral",
        },
        APIEndpointID: "Electric Operations systems",
        APIID: "Operations East Unbranded",
        VersionID: "parsing Convertible",
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
| `request`                                                                                  | [operations.UpsertAPIEndpointRequest](../../models/operations/upsertapiendpointrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.UpsertAPIEndpointResponse](../../models/operations/upsertapiendpointresponse.md), error**

