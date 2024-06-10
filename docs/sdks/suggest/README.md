# Suggest
(*Suggest*)

## Overview

REST APIs for managing LLM OAS suggestions

### Available Operations

* [SuggestOperationIDs](#suggestoperationids) - Generate operation ID suggestions.

## SuggestOperationIDs

Get suggestions from an LLM model for improving the operationIDs in the provided schema.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )
    request := operations.SuggestOperationIDsRequestBody{
        Schema: operations.Schema{
            Content: []byte("0xb2de88c98a"),
            FileName: "your_file_here",
        },
    }
    ctx := context.Background()
    res, err := s.Suggest.SuggestOperationIDs(ctx, request)
    if err != nil {
        log.Fatal(err)
    }
    if res.Suggestion != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.SuggestOperationIDsRequestBody](../../pkg/models/operations/suggestoperationidsrequestbody.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.SuggestOperationIDsResponse](../../pkg/models/operations/suggestoperationidsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
