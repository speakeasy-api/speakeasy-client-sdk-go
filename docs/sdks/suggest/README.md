# Suggest
(*Suggest*)

## Overview

REST APIs for managing LLM OAS suggestions

### Available Operations

* [Suggest](#suggest) - Generate suggestions for improving an OpenAPI document.
* [SuggestItems](#suggestitems) - Generate generic suggestions for a list of items.
* [SuggestOpenAPI](#suggestopenapi) - (DEPRECATED) Generate suggestions for improving an OpenAPI document.
* [SuggestOpenAPIRegistry](#suggestopenapiregistry) - Generate suggestions for improving an OpenAPI document stored in the registry.

## Suggest

Get suggestions from an LLM model for improving an OpenAPI document.

### Example Usage

<!-- UsageSnippet language="go" operationID="suggest" method="post" path="/v1/suggest/openapi_from_summary" -->
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

    res, err := s.Suggest.Suggest(ctx, operations.SuggestRequest{
        SuggestRequestBody: shared.SuggestRequestBody{
            Diagnostics: []shared.Diagnostic{
                shared.Diagnostic{
                    Message: "<value>",
                    Path: []string{
                        "/usr/src",
                    },
                    Type: "<value>",
                },
            },
            OasSummary: shared.OASSummary{
                Info: shared.OASInfo{
                    Description: "prioritize bell vainly",
                    License: shared.License{},
                    Summary: "<value>",
                    Title: "<value>",
                    Version: "<value>",
                },
                Operations: []shared.OASOperation{
                    shared.OASOperation{
                        Description: "though since instead accurate safe unnaturally charming",
                        Method: "<value>",
                        OperationID: "<id>",
                        Path: "/usr/local/bin",
                        Tags: []string{
                            "<value 1>",
                            "<value 2>",
                        },
                    },
                },
            },
            SuggestionType: shared.SuggestRequestBodySuggestionTypeMethodNames,
        },
        XSessionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Schema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.SuggestRequest](../../pkg/models/operations/suggestrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |
| `opts`                                                                     | [][operations.Option](../../pkg/models/operations/option.md)               | :heavy_minus_sign:                                                         | The options for this request.                                              |

### Response

**[*operations.SuggestResponse](../../pkg/models/operations/suggestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SuggestItems

Generate generic suggestions for a list of items.

### Example Usage

<!-- UsageSnippet language="go" operationID="suggestItems" method="post" path="/v1/suggest/items" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Suggest.SuggestItems(ctx, shared.SuggestItemsRequestBody{
        Items: []string{
            "<value 1>",
        },
        Prompt: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Strings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [shared.SuggestItemsRequestBody](../../pkg/models/shared/suggestitemsrequestbody.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.SuggestItemsResponse](../../pkg/models/operations/suggestitemsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SuggestOpenAPI

Get suggestions from an LLM model for improving an OpenAPI document.

### Example Usage

<!-- UsageSnippet language="go" operationID="suggestOpenAPI" method="post" path="/v1/suggest/openapi" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"os"
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

    example, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }

    res, err := s.Suggest.SuggestOpenAPI(ctx, operations.SuggestOpenAPIRequest{
        RequestBody: operations.SuggestOpenAPIRequestBody{
            Schema: operations.Schema{
                Content: example,
                FileName: "example.file",
            },
        },
        XSessionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Schema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.SuggestOpenAPIRequest](../../pkg/models/operations/suggestopenapirequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                             | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.SuggestOpenAPIResponse](../../pkg/models/operations/suggestopenapiresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## SuggestOpenAPIRegistry

Get suggestions from an LLM model for improving an OpenAPI document stored in the registry.

### Example Usage

<!-- UsageSnippet language="go" operationID="suggestOpenAPIRegistry" method="post" path="/v1/suggest/openapi/{namespace_name}/{revision_reference}" -->
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

    res, err := s.Suggest.SuggestOpenAPIRegistry(ctx, operations.SuggestOpenAPIRegistryRequest{
        NamespaceName: "<value>",
        RevisionReference: "<value>",
        XSessionID: "<id>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Schema != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.SuggestOpenAPIRegistryRequest](../../pkg/models/operations/suggestopenapiregistryrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.SuggestOpenAPIRegistryResponse](../../pkg/models/operations/suggestopenapiregistryresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |