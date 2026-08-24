# OrganizationUsageResponse

A billing summary of organization usage


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `AllowedLanguages`                                                            | []`string`                                                                    | :heavy_check_mark:                                                            | List of allowed languages                                                     |
| `FreeTrialExpiry`                                                             | [*time.Time](https://pkg.go.dev/time#Time)                                    | :heavy_minus_sign:                                                            | Expiry date of the free trial, will be null if no trial                       |
| `TotalAllowedLanguages`                                                       | `int64`                                                                       | :heavy_check_mark:                                                            | Total number of allowed languages, -1 if unlimited                            |
| `Usage`                                                                       | [][shared.OrganizationUsage](../../../pkg/models/shared/organizationusage.md) | :heavy_check_mark:                                                            | N/A                                                                           |