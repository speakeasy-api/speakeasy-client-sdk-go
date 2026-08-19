# GetWorkspaceAccessRequest


## Fields

| Field                                        | Type                                         | Required                                     | Description                                  |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| `GenLockID`                                  | `*string`                                    | :heavy_minus_sign:                           | Unique identifier of the generation target.  |
| `Passive`                                    | `*bool`                                      | :heavy_minus_sign:                           | Skip side-effects like incrementing metrics. |
| `TargetType`                                 | `*string`                                    | :heavy_minus_sign:                           | The type of the generated target.            |