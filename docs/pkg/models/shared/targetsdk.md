# TargetSDK


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `CommitHead`                                           | **string*                                              | :heavy_minus_sign:                                     | Remote commit ID.                                      |
| `CreatedAt`                                            | [time.Time](https://pkg.go.dev/time#Time)              | :heavy_check_mark:                                     | Timestamp when the event was created in the database.  |
| `GenerateConfigPostVersion`                            | **string*                                              | :heavy_minus_sign:                                     | Version of the generated target (post generation)      |
| `GenerateGenLockID`                                    | **string*                                              | :heavy_minus_sign:                                     | gen.lock ID (expected to be a uuid).                   |
| `GeneratePublished`                                    | **bool*                                                | :heavy_minus_sign:                                     | Indicates whether the target was considered published. |
| `GenerateTarget`                                       | *string*                                               | :heavy_check_mark:                                     | The target of the event.                               |
| `GenerateTargetVersion`                                | **string*                                              | :heavy_minus_sign:                                     | The version of the target.                             |
| `GitRelativeCwd`                                       | **string*                                              | :heavy_minus_sign:                                     | Current working directory relative to the git root.    |
| `GitRemoteDefaultOwner`                                | **string*                                              | :heavy_minus_sign:                                     | Default owner for git remote.                          |
| `GitRemoteDefaultRepo`                                 | **string*                                              | :heavy_minus_sign:                                     | Default repository name for git remote.                |
| `ID`                                                   | *string*                                               | :heavy_check_mark:                                     | Unique identifier for each event.                      |
| `Success`                                              | **bool*                                                | :heavy_minus_sign:                                     | Indicates whether the event was successful.            |