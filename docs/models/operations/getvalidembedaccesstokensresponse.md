# GetValidEmbedAccessTokensResponse


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ContentType`                                            | *string*                                                 | :heavy_check_mark:                                       | HTTP response content type for this operation            |
| `EmbedTokens`                                            | [][shared.EmbedToken](../../models/shared/embedtoken.md) | :heavy_minus_sign:                                       | OK                                                       |
| `Error`                                                  | [*shared.Error](../../models/shared/error.md)            | :heavy_minus_sign:                                       | Default error response                                   |
| `StatusCode`                                             | *int*                                                    | :heavy_check_mark:                                       | HTTP response status code for this operation             |
| `RawResponse`                                            | [*http.Response](https://pkg.go.dev/net/http#Response)   | :heavy_minus_sign:                                       | Raw HTTP response; suitable for custom response parsing  |