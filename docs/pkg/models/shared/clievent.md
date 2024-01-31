# CliEvent


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `CommitHead`                                          | **string*                                             | :heavy_minus_sign:                                    | Remote commit ID.                                     |
| `CreatedAt`                                           | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Timestamp when the event was created in the database. |
| `DurationMs`                                          | **int64*                                              | :heavy_minus_sign:                                    | Duration of the event in milliseconds.                |
| `ExecutionID`                                         | *string*                                              | :heavy_check_mark:                                    | Unique identifier for each execution of the CLI.      |
| `GhActionOrganization`                                | **string*                                             | :heavy_minus_sign:                                    | GitHub organization of the action.                    |
| `GhActionRepository`                                  | **string*                                             | :heavy_minus_sign:                                    | GitHub repository of the action.                      |
| `GhActionRunLink`                                     | **string*                                             | :heavy_minus_sign:                                    | Link to the GitHub action run.                        |
| `GhActionVersion`                                     | **string*                                             | :heavy_minus_sign:                                    | Version of the GitHub action.                         |
| `GitRelativeCwd`                                      | **string*                                             | :heavy_minus_sign:                                    | Current working directory relative to the git root.   |
| `GitRemoteDefaultRepo`                                | **string*                                             | :heavy_minus_sign:                                    | Default repository name for git remote.               |
| `GitRemoteDefaultUsername`                            | **string*                                             | :heavy_minus_sign:                                    | Default username for git remote.                      |
| `GitUserEmail`                                        | **string*                                             | :heavy_minus_sign:                                    | User email from git configuration.                    |
| `GitUserName`                                         | **string*                                             | :heavy_minus_sign:                                    | User name from git configuration.                     |
| `Hostname`                                            | **string*                                             | :heavy_minus_sign:                                    | Remote hostname.                                      |
| `ID`                                                  | *string*                                              | :heavy_check_mark:                                    | Unique identifier for each event.                     |
| `InteractionType`                                     | *string*                                              | :heavy_check_mark:                                    | Type of interaction.                                  |
| `LocalCompletedAt`                                    | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Timestamp when the event completed, in local time.    |
| `LocalStartedAt`                                      | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | Timestamp when the event started, in local time.      |
| `RawCommand`                                          | **string*                                             | :heavy_minus_sign:                                    | Full CLI command.                                     |
| `RepoLabel`                                           | **string*                                             | :heavy_minus_sign:                                    | Label of the git repository.                          |
| `SpeakeasyAPIKeyID`                                   | *string*                                              | :heavy_check_mark:                                    | Identifier of the Speakeasy API key.                  |
| `SpeakeasyVersion`                                    | *string*                                              | :heavy_check_mark:                                    | Version of the Speakeasy CLI.                         |
| `Success`                                             | *bool*                                                | :heavy_check_mark:                                    | Indicates whether the event was successful.           |
| `WorkspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | Identifier of the workspace.                          |