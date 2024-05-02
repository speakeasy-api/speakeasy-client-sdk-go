# AddTags

Request body of tags to add to a revision


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `RevisionDigest`                           | *string*                                   | :heavy_check_mark:                         | revision digest to add tags too sha256:... |
| `Tags`                                     | []*string*                                 | :heavy_check_mark:                         | string tags to add to the revision         |