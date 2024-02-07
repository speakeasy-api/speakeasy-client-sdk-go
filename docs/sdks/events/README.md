# Events
(*Events*)

## Overview

REST APIs for capturing event data

### Available Operations

* [PostWorkspaceEvents](#postworkspaceevents) - Post events for a specific workspace

## PostWorkspaceEvents

Sends an array of events to be stored for a particular workspace.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity("<YOUR_API_KEY_HERE>"),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("string")),
    )

    ctx := context.Background()
    res, err := s.Events.PostWorkspaceEvents(ctx, operations.PostWorkspaceEventsRequest{
        RequestBody: []shared.CliEvent{
            shared.CliEvent{
                CreatedAt: types.MustTimeFromString("2024-11-21T06:58:42.120Z"),
                ExecutionID: "string",
                ID: "<ID>",
                InteractionType: shared.InteractionTypeCliExec,
                LocalStartedAt: types.MustTimeFromString("2024-05-07T12:35:47.182Z"),
                SpeakeasyAPIKeyName: "string",
                SpeakeasyVersion: "string",
                Success: false,
                WorkspaceID: "string",
            },
        },
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostWorkspaceEventsRequest](../../pkg/models/operations/postworkspaceeventsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostWorkspaceEventsResponse](../../pkg/models/operations/postworkspaceeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
