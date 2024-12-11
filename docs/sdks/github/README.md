# Github
(*Github*)

## Overview

REST APIs for managing the github integration

### Available Operations

* [CheckAccess](#checkaccess)
* [CheckPublishingPRs](#checkpublishingprs)
* [CheckPublishingSecrets](#checkpublishingsecrets)
* [ConfigureCodeSamples](#configurecodesamples)
* [ConfigureMintlifyRepo](#configuremintlifyrepo)
* [ConfigureTarget](#configuretarget)
* [GetAction](#getaction)
* [GetSetup](#getsetup)
* [LinkGithub](#linkgithub)
* [StorePublishingSecrets](#storepublishingsecrets)
* [TriggerAction](#triggeraction)

## CheckAccess

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

    res, err := s.Github.CheckAccess(ctx, operations.CheckGithubAccessRequest{
        Org: "<value>",
        Repo: "<value>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CheckGithubAccessRequest](../../pkg/models/operations/checkgithubaccessrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CheckGithubAccessResponse](../../pkg/models/operations/checkgithubaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## CheckPublishingPRs

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

    res, err := s.Github.CheckPublishingPRs(ctx, operations.GithubCheckPublishingPRsRequest{
        GenerateGenLockID: "<id>",
        Org: "<value>",
        Repo: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GithubPublishingPRResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GithubCheckPublishingPRsRequest](../../pkg/models/operations/githubcheckpublishingprsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GithubCheckPublishingPRsResponse](../../pkg/models/operations/githubcheckpublishingprsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## CheckPublishingSecrets

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

    res, err := s.Github.CheckPublishingSecrets(ctx, operations.GithubCheckPublishingSecretsRequest{
        GenerateGenLockID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GithubMissingPublishingSecretsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.GithubCheckPublishingSecretsRequest](../../pkg/models/operations/githubcheckpublishingsecretsrequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.GithubCheckPublishingSecretsResponse](../../pkg/models/operations/githubcheckpublishingsecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## ConfigureCodeSamples

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Github.ConfigureCodeSamples(ctx, shared.GithubConfigureCodeSamplesRequest{
        Org: "<value>",
        Repo: "<value>",
        TargetName: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GithubConfigureCodeSamplesResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [shared.GithubConfigureCodeSamplesRequest](../../pkg/models/shared/githubconfigurecodesamplesrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../pkg/models/operations/option.md)                                             | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GithubConfigureCodeSamplesResponse](../../pkg/models/operations/githubconfigurecodesamplesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## ConfigureMintlifyRepo

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Github.ConfigureMintlifyRepo(ctx, shared.GithubConfigureMintlifyRepoRequest{
        Input: "<value>",
        Org: "<value>",
        Overlays: []string{
            "<value>",
        },
        Repo: "<value>",
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
| `request`                                                                                                  | [shared.GithubConfigureMintlifyRepoRequest](../../pkg/models/shared/githubconfiguremintlifyreporequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                                               | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.GithubConfigureMintlifyRepoResponse](../../pkg/models/operations/githubconfiguremintlifyreporesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## ConfigureTarget

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Github.ConfigureTarget(ctx, shared.GithubConfigureTargetRequest{
        Org: "<value>",
        RepoName: "<value>",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [shared.GithubConfigureTargetRequest](../../pkg/models/shared/githubconfiguretargetrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                   | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GithubConfigureTargetResponse](../../pkg/models/operations/githubconfiguretargetresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetAction

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

    res, err := s.Github.GetAction(ctx, operations.GetGitHubActionRequest{
        Org: "<value>",
        Repo: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GithubGetActionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetGitHubActionRequest](../../pkg/models/operations/getgithubactionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GetGitHubActionResponse](../../pkg/models/operations/getgithubactionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## GetSetup

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

    res, err := s.Github.GetSetup(ctx, operations.GetGithubSetupStateRequest{
        GenerateGenLockID: "<id>",
        Org: "<value>",
        Repo: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GithubSetupStateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.GetGithubSetupStateRequest](../../pkg/models/operations/getgithubsetupstaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../pkg/models/operations/option.md)                                       | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.GetGithubSetupStateResponse](../../pkg/models/operations/getgithubsetupstateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## LinkGithub

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

    res, err := s.Github.LinkGithub(ctx, operations.LinkGithubAccessRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.LinkGithubAccessRequest](../../pkg/models/operations/linkgithubaccessrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.LinkGithubAccessResponse](../../pkg/models/operations/linkgithubaccessresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## StorePublishingSecrets

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Github.StorePublishingSecrets(ctx, shared.GithubStorePublishingSecretsRequest{
        GenerateGenLockID: "<value>",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [shared.GithubStorePublishingSecretsRequest](../../pkg/models/shared/githubstorepublishingsecretsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.GithubStorePublishingSecretsResponse](../../pkg/models/operations/githubstorepublishingsecretsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |

## TriggerAction

### Example Usage

```go
package main

import(
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := speakeasyclientsdkgo.New(
        speakeasyclientsdkgo.WithSecurity(shared.Security{
            APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
        }),
    )

    res, err := s.Github.TriggerAction(ctx, shared.GithubTriggerActionRequest{
        GenLockID: "<id>",
        Org: "<value>",
        RepoName: "<value>",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [shared.GithubTriggerActionRequest](../../pkg/models/shared/githubtriggeractionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../pkg/models/operations/option.md)                               | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.GithubTriggerActionResponse](../../pkg/models/operations/githubtriggeractionresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.Error    | 4XX                | application/json   |
| sdkerrors.SDKError | 5XX                | \*/\*              |