# OrganizationUsage


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Accessible`                                              | *bool*                                                    | :heavy_check_mark:                                        | Indicates if the features are accessible                  |
| `AccessibleFeatures`                                      | []*string*                                                | :heavy_check_mark:                                        | Features that are accessible                              |
| `GenLockIds`                                              | []*string*                                                | :heavy_check_mark:                                        | List of generation lock IDs                               |
| `Language`                                                | *string*                                                  | :heavy_check_mark:                                        | The programming language used                             |
| `MaxOperations`                                           | *int64*                                                   | :heavy_check_mark:                                        | Maximum Number of operations per SDK specific in contract |
| `NumberOfOperations`                                      | *int64*                                                   | :heavy_check_mark:                                        | Number of operations performed                            |
| `UsedFeatures`                                            | []*string*                                                | :heavy_check_mark:                                        | Features that have been used                              |
| `Workspaces`                                              | []*string*                                                | :heavy_check_mark:                                        | List of workspace IDs                                     |