# Workspaces
(*Workspaces*)

## Overview

REST APIs for managing Workspaces (speakeasy tenancy)

### Available Operations

* [Create](#create) - Create a workspace
* [CreateToken](#createtoken) - Create a token for a particular workspace
* [DeleteToken](#deletetoken) - Delete a token for a particular workspace
* [Get](#get) - Get workspace by context
* [GetAll](#getall) - Get workspaces for a user
* [GetByID](#getbyid) - Get workspace
* [GetFeatureFlags](#getfeatureflags) - Get workspace feature flags
* [GetSettings](#getsettings) - Get workspace settings
* [GetTeam](#getteam) - Get team members for a particular workspace
* [GetTokens](#gettokens) - Get tokens for a particular workspace
* [GrantAccess](#grantaccess) - Grant a user access to a particular workspace
* [RevokeAccess](#revokeaccess) - Revoke a user's access to a particular workspace
* [SetFeatureFlags](#setfeatureflags) - Set workspace feature flags
* [Update](#update) - Update workspace details
* [UpdateSettings](#updatesettings) - Update workspace settings

## Create

Creates a workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="createWorkspace" method="post" path="/v1/workspace" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.Create(ctx, shared.Workspace{
        CreatedAt: types.MustTimeFromString("2023-11-18T13:41:10.525Z"),
        ID: "<id>",
        Name: "<value>",
        OrganizationID: "<id>",
        Slug: "<value>",
        UpdatedAt: types.MustTimeFromString("2024-11-21T08:36:32.740Z"),
        Verified: true,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.Workspace](../../pkg/models/shared/workspace.md)     | :heavy_check_mark:                                           | The request object to use for the request.                   |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.CreateWorkspaceResponse](../../pkg/models/operations/createworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## CreateToken

Create a token for a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="createWorkspaceToken" method="post" path="/v1/workspace/{workspace_id}/tokens" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.CreateToken(ctx, operations.CreateWorkspaceTokenRequest{
        WorkspaceToken: shared.WorkspaceToken{
            Alg: "<value>",
            CreatedAt: types.MustTimeFromString("2024-10-04T10:23:04.522Z"),
            ID: "<id>",
            Key: "<key>",
            Name: "<value>",
            WorkspaceID: "<id>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateWorkspaceTokenRequest](../../pkg/models/operations/createworkspacetokenrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.CreateWorkspaceTokenResponse](../../pkg/models/operations/createworkspacetokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## DeleteToken

Delete a token for a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteWorkspaceToken" method="delete" path="/v1/workspace/{workspace_id}/tokens/{tokenID}" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.DeleteToken(ctx, operations.DeleteWorkspaceTokenRequest{
        TokenID: "<id>",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.DeleteWorkspaceTokenRequest](../../pkg/models/operations/deleteworkspacetokenrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.DeleteWorkspaceTokenResponse](../../pkg/models/operations/deleteworkspacetokenresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## Get

Get information about a particular workspace by context.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaceByContext" method="get" path="/v1/workspace" -->
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

    res, err := s.Workspaces.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceAndOrganization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetWorkspaceByContextResponse](../../pkg/models/operations/getworkspacebycontextresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetAll

Returns a list of workspaces a user has access too

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaces" method="get" path="/v1/workspaces" -->
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

    res, err := s.Workspaces.GetAll(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `opts`                                                       | [][operations.Option](../../pkg/models/operations/option.md) | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetWorkspacesResponse](../../pkg/models/operations/getworkspacesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetByID

Get information about a particular workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspace" method="get" path="/v1/workspace/{workspace_id}" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GetByID(ctx, operations.GetWorkspaceRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetWorkspaceRequest](../../pkg/models/operations/getworkspacerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                         | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.GetWorkspaceResponse](../../pkg/models/operations/getworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetFeatureFlags

Get workspace feature flags

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaceFeatureFlags" method="get" path="/v1/workspace/{workspace_id}/feature_flags" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GetFeatureFlags(ctx, operations.GetWorkspaceFeatureFlagsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceFeatureFlagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetWorkspaceFeatureFlagsRequest](../../pkg/models/operations/getworkspacefeatureflagsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GetWorkspaceFeatureFlagsResponse](../../pkg/models/operations/getworkspacefeatureflagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## GetSettings

Get settings about a particular workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaceSettings" method="get" path="/v1/workspace/{workspace_id}/settings" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GetSettings(ctx, operations.GetWorkspaceSettingsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceSettings != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.GetWorkspaceSettingsRequest](../../pkg/models/operations/getworkspacesettingsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.GetWorkspaceSettingsResponse](../../pkg/models/operations/getworkspacesettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetTeam

Get team members for a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaceTeam" method="get" path="/v1/workspace/{workspace_id}/team" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GetTeam(ctx, operations.GetWorkspaceTeamRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceTeamResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.GetWorkspaceTeamRequest](../../pkg/models/operations/getworkspaceteamrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.GetWorkspaceTeamResponse](../../pkg/models/operations/getworkspaceteamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetTokens

Get tokens for a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="getWorkspaceTokens" method="get" path="/v1/workspace/{workspace_id}/tokens" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GetTokens(ctx, operations.GetWorkspaceTokensRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetWorkspaceTokensRequest](../../pkg/models/operations/getworkspacetokensrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                     | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.GetWorkspaceTokensResponse](../../pkg/models/operations/getworkspacetokensresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GrantAccess

Grant a user access to a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="grantUserAccessToWorkspace" method="put" path="/v1/workspace/{workspace_id}/team/email/{email}" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.GrantAccess(ctx, operations.GrantUserAccessToWorkspaceRequest{
        Email: "Idella24@gmail.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceInviteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GrantUserAccessToWorkspaceRequest](../../pkg/models/operations/grantuseraccesstoworkspacerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../pkg/models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GrantUserAccessToWorkspaceResponse](../../pkg/models/operations/grantuseraccesstoworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## RevokeAccess

Revoke a user's access to a particular workspace

### Example Usage

<!-- UsageSnippet language="go" operationID="revokeUserAccessToWorkspace" method="delete" path="/v1/workspace/{workspace_id}/team/{userId}" -->
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
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.RevokeAccess(ctx, operations.RevokeUserAccessToWorkspaceRequest{
        UserID: "<id>",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.RevokeUserAccessToWorkspaceRequest](../../pkg/models/operations/revokeuseraccesstoworkspacerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.RevokeUserAccessToWorkspaceResponse](../../pkg/models/operations/revokeuseraccesstoworkspaceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## SetFeatureFlags

Set workspace feature flags

### Example Usage

<!-- UsageSnippet language="go" operationID="setWorkspaceFeatureFlags" method="post" path="/v1/workspace/feature_flags" -->
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

    res, err := s.Workspaces.SetFeatureFlags(ctx, shared.WorkspaceFeatureFlagRequest{
        FeatureFlags: []shared.WorkspaceFeatureFlag{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.WorkspaceFeatureFlagResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [shared.WorkspaceFeatureFlagRequest](../../pkg/models/shared/workspacefeatureflagrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.SetWorkspaceFeatureFlagsResponse](../../pkg/models/operations/setworkspacefeatureflagsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 5XX                | application/json   |
| sdkerrors.SDKError | 4XX                | \*/\*              |

## Update

Update information about a particular workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWorkspaceDetails" method="post" path="/v1/workspace/{workspace_id}/details" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.Update(ctx, operations.UpdateWorkspaceDetailsRequest{
        Workspace: shared.Workspace{
            CreatedAt: types.MustTimeFromString("2023-08-02T22:30:24.264Z"),
            ID: "<id>",
            Name: "<value>",
            OrganizationID: "<id>",
            Slug: "<value>",
            UpdatedAt: types.MustTimeFromString("2025-01-24T03:53:13.581Z"),
            Verified: true,
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.UpdateWorkspaceDetailsRequest](../../pkg/models/operations/updateworkspacedetailsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.UpdateWorkspaceDetailsResponse](../../pkg/models/operations/updateworkspacedetailsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## UpdateSettings

Update settings about a particular workspace.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateWorkspaceSettings" method="put" path="/v1/workspace/{workspace_id}/settings" -->
```go
package main

import(
	"context"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := v3.New(
        v3.WithWorkspaceID("<id>"),
        v3.WithSecurity(shared.Security{
            APIKey: v3.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Workspaces.UpdateSettings(ctx, operations.UpdateWorkspaceSettingsRequest{
        WorkspaceSettings: shared.WorkspaceSettings{
            CreatedAt: types.MustTimeFromString("2025-03-09T15:48:09.330Z"),
            UpdatedAt: types.MustTimeFromString("2025-11-24T16:37:53.492Z"),
            WebhookURL: "https://wobbly-lid.org",
            WorkspaceID: "<id>",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.UpdateWorkspaceSettingsRequest](../../pkg/models/operations/updateworkspacesettingsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.UpdateWorkspaceSettingsResponse](../../pkg/models/operations/updateworkspacesettingsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |