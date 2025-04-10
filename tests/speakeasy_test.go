// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	speakeasyclientsdkgo "github.com/speakeasy-api/speakeasy-client-sdk-go/v3"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/operations"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/types"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSpeakeasy_GenerateCodeSamplePreview(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("generateCodeSamplePreview")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
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

	testHTTPClient := createTestHTTPClient("generateCodeSamplePreviewAsync")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
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

	testHTTPClient := createTestHTTPClient("getCodeSamplePreviewAsync")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
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

func TestSpeakeasy_CreatePublishingToken(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("createPublishingToken")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.CreatePublishingToken(ctx, shared.PublishingToken{
		CreatedAt:      types.MustTimeFromString("2025-10-25T02:17:15.413Z"),
		ID:             "<id>",
		TargetID:       "<id>",
		TargetResource: "<value>",
		Token:          "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotNil(t, res.PublishingToken)
	assert.Equal(t, &shared.PublishingToken{
		CreatedAt:      types.MustTimeFromString("2023-12-11T20:09:47.767Z"),
		ID:             "<id>",
		TargetID:       "<id>",
		TargetResource: "<value>",
		Token:          "<value>",
	}, res.PublishingToken)

}

func TestSpeakeasy_DeletePublishingToken(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("deletePublishingToken")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.DeletePublishingToken(ctx, operations.DeletePublishingTokenRequest{
		TokenID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestSpeakeasy_GetPublishingTokenByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getPublishingTokenByID")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.GetPublishingTokenByID(ctx, operations.GetPublishingTokenByIDRequest{
		TokenID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotNil(t, res.PublishingToken)
	assert.Equal(t, &shared.PublishingToken{
		CreatedAt:      types.MustTimeFromString("2025-10-20T08:51:57.553Z"),
		ID:             "<id>",
		TargetID:       "<id>",
		TargetResource: "<value>",
		Token:          "<value>",
	}, res.PublishingToken)

}

func TestSpeakeasy_GetPublishingTokenTargetByID(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getPublishingTokenTargetByID")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.GetPublishingTokenTargetByID(ctx, operations.GetPublishingTokenTargetByIDRequest{
		TokenID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotNil(t, res.Object)
	assert.Equal(t, &operations.GetPublishingTokenTargetByIDResponseBody{}, res.Object)

}

func TestSpeakeasy_UpdatePublishingTokenExpiration(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("updatePublishingTokenExpiration")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.UpdatePublishingTokenExpiration(ctx, operations.UpdatePublishingTokenExpirationRequest{
		TokenID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}
