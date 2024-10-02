# APIEndpointInput

An ApiEndpoint is a description of an Endpoint for an API.


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `APIEndpointID`                                                    | *string*                                                           | :heavy_check_mark:                                                 | The ID of this ApiEndpoint. This is a hash of the method and path. |
| `Description`                                                      | *string*                                                           | :heavy_check_mark:                                                 | A detailed description of the ApiEndpoint.                         |
| `DisplayName`                                                      | *string*                                                           | :heavy_check_mark:                                                 | A human-readable name for the ApiEndpoint.                         |
| `Method`                                                           | *string*                                                           | :heavy_check_mark:                                                 | HTTP verb.                                                         |
| `Path`                                                             | *string*                                                           | :heavy_check_mark:                                                 | Path that handles this Api.                                        |
| `VersionID`                                                        | *string*                                                           | :heavy_check_mark:                                                 | The version ID of the Api this ApiEndpoint belongs to.             |