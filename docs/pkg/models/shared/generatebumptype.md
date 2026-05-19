# GenerateBumpType

Bump type of the lock file (calculated semver delta, custom change (manual release), or prerelease/graduate)

## Example Usage

```go
import (
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
)

value := shared.GenerateBumpTypeMajor
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GenerateBumpTypeMajor`      | major                        |
| `GenerateBumpTypeMinor`      | minor                        |
| `GenerateBumpTypePatch`      | patch                        |
| `GenerateBumpTypeCustom`     | custom                       |
| `GenerateBumpTypeGraduate`   | graduate                     |
| `GenerateBumpTypePrerelease` | prerelease                   |
| `GenerateBumpTypeNone`       | none                         |