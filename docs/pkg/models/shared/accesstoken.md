# AccessToken

An AccessToken is a token that can be used to authenticate with the Speakeasy API.


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AccessToken`                                                           | `string`                                                                | :heavy_check_mark:                                                      | N/A                                                                     |
| `Claims`                                                                | [shared.Claims](../../../pkg/models/shared/claims.md)                   | :heavy_check_mark:                                                      | N/A                                                                     |
| `FeatureFlags`                                                          | [][shared.FeatureFlag](../../../pkg/models/shared/featureflag.md)       | :heavy_minus_sign:                                                      | N/A                                                                     |
| `User`                                                                  | [shared.AccessTokenUser](../../../pkg/models/shared/accesstokenuser.md) | :heavy_check_mark:                                                      | N/A                                                                     |
| `Workspaces`                                                            | [][shared.Workspaces](../../../pkg/models/shared/workspaces.md)         | :heavy_minus_sign:                                                      | N/A                                                                     |