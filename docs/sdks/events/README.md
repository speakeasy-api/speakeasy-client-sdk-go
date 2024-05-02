# Events
(*Events*)

## Overview

REST APIs for capturing event data

### Available Operations

* [GetWorkspaceEventsByTarget](#getworkspaceeventsbytarget) - Load recent events for a particular workspace
* [GetWorkspaceTargets](#getworkspacetargets) - Load targets for a particular workspace
* [PostWorkspaceEvents](#postworkspaceevents) - Post events for a specific workspace
* [SearchWorkspaceEvents](#searchworkspaceevents) - Search events for a particular workspace by any field

## GetWorkspaceEventsByTarget

Load recent events for a particular workspace

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
    res, err := s.Events.GetWorkspaceEventsByTarget(ctx, operations.GetWorkspaceEventsByTargetRequest{
        TargetID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CliEventBatch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetWorkspaceEventsByTargetRequest](../../pkg/models/operations/getworkspaceeventsbytargetrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetWorkspaceEventsByTargetResponse](../../pkg/models/operations/getworkspaceeventsbytargetresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetWorkspaceTargets

Load targets for a particular workspace

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
    res, err := s.Events.GetWorkspaceTargets(ctx, operations.GetWorkspaceTargetsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.TargetSDKList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetWorkspaceTargetsRequest](../../pkg/models/operations/getworkspacetargetsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.GetWorkspaceTargetsResponse](../../pkg/models/operations/getworkspacetargetsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
)

func main() {
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
        speakeasyclientsdkgo.WithWorkspaceID(speakeasyclientsdkgo.String("<value>")),
    )

    ctx := context.Background()
    res, err := s.Events.PostWorkspaceEvents(ctx, operations.PostWorkspaceEventsRequest{
        RequestBody: []shared.CliEvent{
            shared.CliEvent{
                CreatedAt: types.MustTimeFromString("2024-11-21T06:58:42.120Z"),
                ExecutionID: "<value>",
                ID: "<id>",
                InteractionType: shared.InteractionTypeCliExec,
                LocalStartedAt: types.MustTimeFromString("2024-05-07T12:35:47.182Z"),
                SpeakeasyAPIKeyName: "<value>",
                SpeakeasyVersion: "<value>",
                Success: false,
                WorkspaceID: "<value>",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostWorkspaceEventsRequest](../../pkg/models/operations/postworkspaceeventsrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |


### Response

**[*operations.PostWorkspaceEventsResponse](../../pkg/models/operations/postworkspaceeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## SearchWorkspaceEvents

Search events for a particular workspace by any field

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
    res, err := s.Events.SearchWorkspaceEvents(ctx, operations.SearchWorkspaceEventsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.CliEventBatch != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SearchWorkspaceEventsRequest](../../pkg/models/operations/searchworkspaceeventsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.SearchWorkspaceEventsResponse](../../pkg/models/operations/searchworkspaceeventsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
