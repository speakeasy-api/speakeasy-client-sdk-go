# CodeSamples
(*CodeSamples*)

## Overview

REST APIs for retrieving Code Samples

### Available Operations

* [GenerateCodeSamplePreview](#generatecodesamplepreview) - Generate Code Sample previews from a file and configuration parameters.
* [GenerateCodeSamplePreviewAsync](#generatecodesamplepreviewasync) - Initiate asynchronous Code Sample preview generation from a file and configuration parameters, receiving an async JobID response for polling.
* [Get](#get) - Retrieve usage snippets
* [GetCodeSamplePreviewAsync](#getcodesamplepreviewasync) - Poll for the result of an asynchronous Code Sample preview generation.

## GenerateCodeSamplePreview

This endpoint generates Code Sample previews from a file and configuration parameters.

### Example Usage

<!-- UsageSnippet language="go" operationID="generateCodeSamplePreview" method="post" path="/v1/code_sample/preview" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.CodeSamples.GenerateCodeSamplePreview(ctx, shared.CodeSampleSchemaInput{
        Language: "<value>",
        SchemaFile: shared.SchemaFile{
            Content: example,
            FileName: "example.file",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageSnippets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CodeSampleSchemaInput](../../pkg/models/shared/codesampleschemainput.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GenerateCodeSamplePreviewResponse](../../pkg/models/operations/generatecodesamplepreviewresponse.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| sdkerrors.Error  | 4XX              | application/json |
| sdkerrors.Error  | 5XX              | application/json |

## GenerateCodeSamplePreviewAsync

This endpoint generates Code Sample previews from a file and configuration parameters, receiving an async JobID response for polling.

### Example Usage

<!-- UsageSnippet language="go" operationID="generateCodeSamplePreviewAsync" method="post" path="/v1/code_sample/preview/async" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.CodeSamples.GenerateCodeSamplePreviewAsync(ctx, shared.CodeSampleSchemaInput{
        Language: "<value>",
        SchemaFile: shared.SchemaFile{
            Content: example,
            FileName: "example.file",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.CodeSampleSchemaInput](../../pkg/models/shared/codesampleschemainput.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                     | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.GenerateCodeSamplePreviewAsyncResponse](../../pkg/models/operations/generatecodesamplepreviewasyncresponse.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| sdkerrors.Error  | 4XX              | application/json |
| sdkerrors.Error  | 5XX              | application/json |

## Get

Retrieve usage snippets from an OpenAPI document stored in the registry. Supports filtering by language and operation ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCodeSamples" method="get" path="/v1/code_sample" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.CodeSamples.Get(ctx, operations.GetCodeSamplesRequest{
        Languages: []string{
            "python",
            "javascript",
        },
        MethodPaths: []operations.MethodPaths{
            operations.MethodPaths{
                Method: shared.HTTPMethodGet,
                Path: "/pets",
            },
        },
        OperationIds: []string{
            "g",
            "e",
            "t",
            "P",
            "e",
            "t",
            "s",
        },
        RegistryURL: "https://spec.speakeasy.com/my-org/my-workspace/my-source",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UsageSnippets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetCodeSamplesRequest](../../pkg/models/operations/getcodesamplesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetCodeSamplesResponse](../../pkg/models/operations/getcodesamplesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetCodeSamplePreviewAsync

Poll for the result of an asynchronous Code Sample preview generation.

### Example Usage

<!-- UsageSnippet language="go" operationID="getCodeSamplePreviewAsync" method="get" path="/v1/code_sample/preview/async/{jobID}" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.CodeSamples.GetCodeSamplePreviewAsync(ctx, operations.GetCodeSamplePreviewAsyncRequest{
        JobID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetCodeSamplePreviewAsyncRequest](../../pkg/models/operations/getcodesamplepreviewasyncrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetCodeSamplePreviewAsyncResponse](../../pkg/models/operations/getcodesamplepreviewasyncresponse.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| sdkerrors.Error  | 4XX              | application/json |
| sdkerrors.Error  | 5XX              | application/json |