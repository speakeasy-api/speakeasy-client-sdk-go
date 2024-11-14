# Speakeasy SDK

## Overview

Speakeasy API: The Subscriptions API manages subscriptions for CLI and registry events

The Speakeasy Platform Documentation
</docs>

### Available Operations

* [GenerateCodeSamplePreview](#generatecodesamplepreview) - Generate Code Sample previews from a file and configuration parameters.
* [GenerateCodeSamplePreviewAsync](#generatecodesamplepreviewasync) - Initiate asynchronous Code Sample preview generation from a file and configuration parameters, receiving an async JobID response for polling.
* [GetCodeSamplePreviewAsync](#getcodesamplepreviewasync) - Poll for the result of an asynchronous Code Sample preview generation.

## GenerateCodeSamplePreview

This endpoint generates Code Sample previews from a file and configuration parameters.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"os"
	"context"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.GenerateCodeSamplePreview(ctx, shared.CodeSampleSchemaInput{
        Languages: []string{
            "<value>",
            "<value>",
        },
        SchemaFile: shared.SchemaFile{
            Content: content,
            FileName: "example.file",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONResponseStream != nil {
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
| sdkerrors.Error  | 4XX, 5XX         | application/json |

## GenerateCodeSamplePreviewAsync

This endpoint generates Code Sample previews from a file and configuration parameters, receiving an async JobID response for polling.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"os"
	"context"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    ctx := context.Background()
    res, err := s.GenerateCodeSamplePreviewAsync(ctx, shared.CodeSampleSchemaInput{
        Languages: []string{
            "<value>",
            "<value>",
            "<value>",
        },
        SchemaFile: shared.SchemaFile{
            Content: content,
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
| sdkerrors.Error  | 4XX, 5XX         | application/json |

## GetCodeSamplePreviewAsync

Poll for the result of an asynchronous Code Sample preview generation.

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
    )

    ctx := context.Background()
    res, err := s.GetCodeSamplePreviewAsync(ctx, operations.GetCodeSamplePreviewAsyncRequest{
        JobID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONResponseStream != nil {
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
| sdkerrors.Error  | 4XX, 5XX         | application/json |