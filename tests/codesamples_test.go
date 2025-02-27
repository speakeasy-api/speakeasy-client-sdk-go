// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCodesamples_GetCodeSamples(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getCodeSamples")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.CodeSamples.Get(ctx, operations.GetCodeSamplesRequest{
		RegistryURL: "https://normal-making.name",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.UsageSnippets)
	assert.Equal(t, &shared.UsageSnippets{
		Snippets: []shared.UsageSnippet{
			shared.UsageSnippet{
				Code:        "<value>",
				Language:    "<value>",
				Method:      "<value>",
				OperationID: "<id>",
				Path:        "/root",
			},
			shared.UsageSnippet{
				Code:        "<value>",
				Language:    "<value>",
				Method:      "<value>",
				OperationID: "<id>",
				Path:        "/usr/bin",
			},
			shared.UsageSnippet{
				Code:        "<value>",
				Language:    "<value>",
				Method:      "<value>",
				OperationID: "<id>",
				Path:        "/sys",
			},
		},
	}, res.UsageSnippets)

}
