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
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("generateCodeSamplePreview")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	content, fileErr := os.Open("../.speakeasy/testfiles/example.file")
	require.NoError(t, fileErr)

	ctx := context.Background()
	res, err := s.GenerateCodeSamplePreview(ctx, shared.CodeSampleSchemaInput{
		Languages: []string{
			"<value>",
			"<value>",
		},
		SchemaFile: shared.SchemaFile{
			Content:  content,
			FileName: "example.file",
		},
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
}

func TestSpeakeasy_GenerateCodeSamplePreviewAsync(t *testing.T) {
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("generateCodeSamplePreviewAsync")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	content, fileErr := os.Open("../.speakeasy/testfiles/example.file")
	require.NoError(t, fileErr)

	ctx := context.Background()
	res, err := s.GenerateCodeSamplePreviewAsync(ctx, shared.CodeSampleSchemaInput{
		Languages: []string{
			"<value>",
			"<value>",
			"<value>",
		},
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
	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(createTestHTTPClient("getCodeSamplePreviewAsync")),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	ctx := context.Background()
	res, err := s.GetCodeSamplePreviewAsync(ctx, operations.GetCodeSamplePreviewAsyncRequest{
		JobID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 202, res.StatusCode)
	assert.NotNil(t, res.TwoHundredAndTwoApplicationJSONObject)
	assert.Equal(t, &operations.GetCodeSamplePreviewAsyncResponseBody{
		Status: shared.CodeSamplesJobStatusPending,
	}, res.TwoHundredAndTwoApplicationJSONObject)
}
