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
	"os"
	"testing"
)

func TestSpeakeasy_GenerateCodeSamplePreview(t *testing.T) {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("generateCodeSamplePreview")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	content, fileErr := os.Open("../.speakeasy/testfiles/example.file")
	require.NoError(t, fileErr)

	res, err := s.CodeSamples.GenerateCodeSamplePreview(ctx, shared.CodeSampleSchemaInput{
		Language: "<value>",
		SchemaFile: shared.SchemaFile{
			Content:  content,
			FileName: "example.file",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestSpeakeasy_GenerateCodeSamplePreviewAsync(t *testing.T) {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("generateCodeSamplePreviewAsync")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	content, fileErr := os.Open("../.speakeasy/testfiles/example.file")
	require.NoError(t, fileErr)

	res, err := s.CodeSamples.GenerateCodeSamplePreviewAsync(ctx, shared.CodeSampleSchemaInput{
		Language: "<value>",
		SchemaFile: shared.SchemaFile{
			Content:  content,
			FileName: "example.file",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 202, res.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GenerateCodeSamplePreviewAsyncResponseBody{
		JobID:  "<id>",
		Status: shared.CodeSamplesJobStatusRunning,
	}, res.Object)

}

func TestSpeakeasy_GetCodeSamplePreviewAsync(t *testing.T) {
	ctx := context.Background()

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("getCodeSamplePreviewAsync")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.CodeSamples.GetCodeSamplePreviewAsync(ctx, operations.GetCodeSamplePreviewAsyncRequest{
		JobID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 202, res.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GetCodeSamplePreviewAsyncResponseBody{
		Status: shared.CodeSamplesJobStatusPending,
	}, res.Object)

}
