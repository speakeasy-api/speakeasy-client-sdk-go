# InsertVersionMetadataRequest


## Fields

| Field                                                                             | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `VersionMetadata`                                                                 | [shared.VersionMetadataInput](../../../pkg/models/shared/versionmetadatainput.md) | :heavy_check_mark:                                                                | A JSON representation of the metadata to insert.                                  |
| `APIID`                                                                           | *string*                                                                          | :heavy_check_mark:                                                                | The ID of the Api to insert metadata for.                                         |
| `VersionID`                                                                       | *string*                                                                          | :heavy_check_mark:                                                                | The version ID of the Api to insert metadata for.                                 |