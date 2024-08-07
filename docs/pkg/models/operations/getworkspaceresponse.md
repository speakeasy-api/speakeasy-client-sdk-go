# GetWorkspaceResponse


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ContentType`                                                | *string*                                                     | :heavy_check_mark:                                           | HTTP response content type for this operation                |
| `Error`                                                      | **sdkerrors.Error*                                           | :heavy_minus_sign:                                           | Default error response                                       |
| `StatusCode`                                                 | *int*                                                        | :heavy_check_mark:                                           | HTTP response status code for this operation                 |
| `RawResponse`                                                | [*http.Response](https://pkg.go.dev/net/http#Response)       | :heavy_check_mark:                                           | Raw HTTP response; suitable for custom response parsing      |
| `Workspace`                                                  | [*shared.Workspace](../../../pkg/models/shared/workspace.md) | :heavy_minus_sign:                                           | OK                                                           |