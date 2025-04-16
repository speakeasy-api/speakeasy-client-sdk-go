# UpdatePublishingTokenExpirationRequestBody

The publishing token to update


## Fields

| Field                                             | Type                                              | Required                                          | Description                                       |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| `TokenName`                                       | **string*                                         | :heavy_minus_sign:                                | The new name for the publishing token.            |
| `ValidUntil`                                      | [time.Time](https://pkg.go.dev/time#Time)         | :heavy_check_mark:                                | The new expiration date for the publishing token. |