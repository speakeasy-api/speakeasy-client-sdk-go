# Namespace

A namespace contains many revisions.


## Fields

| Field                                                 | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `CreatedAt`                                           | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | N/A                                                   |
| `ID`                                                  | *string*                                              | :heavy_check_mark:                                    | {organization_slug}/{workspace_slug}/{namespace_name} |
| `Name`                                                | *string*                                              | :heavy_check_mark:                                    | A human-readable name for the namespace.              |
| `Public`                                              | **bool*                                               | :heavy_minus_sign:                                    | N/A                                                   |
| `UpdatedAt`                                           | [time.Time](https://pkg.go.dev/time#Time)             | :heavy_check_mark:                                    | N/A                                                   |