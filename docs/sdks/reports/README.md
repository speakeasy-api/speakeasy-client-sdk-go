# Reports
(*Reports*)

## Overview

REST APIs for managing reports (lint reports, change reports, etc)

### Available Operations

* [GetChangesReportSignedURL](#getchangesreportsignedurl) - Get the signed access url for the change reports for a particular document.
* [GetLintingReportSignedURL](#getlintingreportsignedurl) - Get the signed access url for the linting reports for a particular document.
* [UploadReport](#uploadreport) - Upload a report.

## GetChangesReportSignedURL

Get the signed access url for the change reports for a particular document.

### Example Usage

<!-- UsageSnippet language="go" operationID="getChangesReportSignedUrl" method="get" path="/v1/reports/changes/{documentChecksum}" -->
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

    res, err := s.Reports.GetChangesReportSignedURL(ctx, operations.GetChangesReportSignedURLRequest{
        DocumentChecksum: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SignedAccess != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetChangesReportSignedURLRequest](../../pkg/models/operations/getchangesreportsignedurlrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetChangesReportSignedURLResponse](../../pkg/models/operations/getchangesreportsignedurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## GetLintingReportSignedURL

Get the signed access url for the linting reports for a particular document.

### Example Usage

<!-- UsageSnippet language="go" operationID="getLintingReportSignedUrl" method="get" path="/v1/reports/linting/{documentChecksum}" -->
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

    res, err := s.Reports.GetLintingReportSignedURL(ctx, operations.GetLintingReportSignedURLRequest{
        DocumentChecksum: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.SignedAccess != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetLintingReportSignedURLRequest](../../pkg/models/operations/getlintingreportsignedurlrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../pkg/models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.GetLintingReportSignedURLResponse](../../pkg/models/operations/getlintingreportsignedurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |

## UploadReport

Upload a report.

### Example Usage

<!-- UsageSnippet language="go" operationID="uploadReport" method="post" path="/v1/reports" -->
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

    res, err := s.Reports.UploadReport(ctx, operations.UploadReportRequestBody{
        Data: shared.Report{},
        File: operations.File{
            Content: example,
            FileName: "example.file",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UploadedReport != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UploadReportRequestBody](../../pkg/models/operations/uploadreportrequestbody.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../pkg/models/operations/option.md)                                 | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.UploadReportResponse](../../pkg/models/operations/uploadreportresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4XX, 5XX           | \*/\*              |