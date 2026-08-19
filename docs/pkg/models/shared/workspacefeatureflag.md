# WorkspaceFeatureFlag

enum value workspace feature flag

## Example Usage

```go
import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
)

value := shared.WorkspaceFeatureFlagSchemaRegistry

// Open enum: custom values can be created with a direct type cast
custom := shared.WorkspaceFeatureFlag("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `WorkspaceFeatureFlagSchemaRegistry`     | schema_registry                          |
| `WorkspaceFeatureFlagChangesReport`      | changes_report                           |
| `WorkspaceFeatureFlagSkipSchemaRegistry` | skip_schema_registry                     |
| `WorkspaceFeatureFlagWebhooks`           | webhooks                                 |