# GetWorkspaceTargetsRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AfterLastEventCreatedAt`                                              | [*time.Time](https://pkg.go.dev/time#Time)                             | :heavy_minus_sign:                                                     | Filter to only return targets with events created after this timestamp |
| `WorkspaceID`                                                          | **string*                                                              | :heavy_minus_sign:                                                     | Unique identifier of the workspace.                                    |