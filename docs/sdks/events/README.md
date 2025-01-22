# Events
(*Events*)

## Overview

REST APIs for managing events captured by a speakeasy binary (CLI, GitHub Action etc)

### Available Operations

* [GetEventsByTarget](#geteventsbytarget) - Load recent events for a particular workspace
* [GetTargets](#gettargets) - Load targets for a particular workspace
* [GetTargetsDeprecated](#gettargetsdeprecated) - Load targets for a particular workspace
* [Post](#post) - Post events for a specific workspace
* [Search](#search) - Search events for a particular workspace by any field

## GetEventsByTarget

Load recent events for a particular workspace

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Events.GetEventsByTarget(ctx, operations.GetWorkspaceEventsByTargetRequest{
        TargetID: "<id>",
        WorkspaceID: speakeasyclientsdkgo.String("<id>"),
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
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetWorkspaceEventsByTargetResponse](../../pkg/models/operations/getworkspaceeventsbytargetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## GetTargets

Load targets for a particular workspace

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Events.GetTargets(ctx, operations.GetWorkspaceTargetsRequest{})
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
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetWorkspaceTargetsResponse](../../pkg/models/operations/getworkspacetargetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## GetTargetsDeprecated

Load targets for a particular workspace

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Events.GetTargetsDeprecated(ctx, operations.GetWorkspaceTargetsDeprecatedRequest{
        WorkspaceID: speakeasyclientsdkgo.String("<id>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TargetSDKList != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.GetWorkspaceTargetsDeprecatedRequest](../../pkg/models/operations/getworkspacetargetsdeprecatedrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.GetWorkspaceTargetsDeprecatedResponse](../../pkg/models/operations/getworkspacetargetsdeprecatedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## Post

Sends an array of events to be stored for a particular workspace.

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Events.Post(ctx, operations.PostWorkspaceEventsRequest{
        RequestBody: []shared.CliEvent{
            shared.CliEvent{
                CreatedAt: types.MustTimeFromString("2025-03-02T10:07:28.113Z"),
                ExecutionID: "<id>",
                ID: "<id>",
                InteractionType: shared.InteractionTypeAuthenticate,
                LocalStartedAt: types.MustTimeFromString("2025-08-12T17:54:17.538Z"),
                SpeakeasyAPIKeyName: "<value>",
                SpeakeasyVersion: "<value>",
                Success: true,
                WorkspaceID: "<id>",
            },
        },
        WorkspaceID: speakeasyclientsdkgo.String("<id>"),
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

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## Search

Search events for a particular workspace by any field

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Events.Search(ctx, operations.SearchWorkspaceEventsRequest{
        WorkspaceID: speakeasyclientsdkgo.String("<id>"),
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

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.SearchWorkspaceEventsRequest](../../pkg/models/operations/searchworkspaceeventsrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../pkg/models/operations/option.md)                                           | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.SearchWorkspaceEventsResponse](../../pkg/models/operations/searchworkspaceeventsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |