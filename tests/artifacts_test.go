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
	"testing"
)

func TestArtifacts_CreateRemoteSource(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("createRemoteSource")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.CreateRemoteSource(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestArtifacts_GetBlob(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getBlob")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.GetBlob(ctx, operations.GetBlobRequest{
		Digest:           "<value>",
		NamespaceName:    "<value>",
		OrganizationSlug: "<value>",
		WorkspaceSlug:    "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestArtifacts_GetManifest(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getManifest")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.GetManifest(ctx, operations.GetManifestRequest{
		NamespaceName:     "<value>",
		OrganizationSlug:  "<value>",
		RevisionReference: "<value>",
		WorkspaceSlug:     "<value>",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.Manifest)
	assert.Equal(t, &shared.Manifest{
		Layers: []shared.V2Descriptor{
			shared.V2Descriptor{
				Digest:    speakeasyclientsdkgo.String("sha256:5d20c808ce198565ff70b3ed23a991dd49afac45dece63474b27ce6ed036adc6"),
				MediaType: speakeasyclientsdkgo.String("application/vnd.docker.image.rootfs.diff.tar.gzip"),
				Size:      speakeasyclientsdkgo.Int64(2107098),
			},
		},
		MediaType:     speakeasyclientsdkgo.String("application/vnd.docker.distribution.manifest.v2+json"),
		SchemaVersion: speakeasyclientsdkgo.Int64(2),
	}, res.Manifest)

}

func TestArtifacts_GetNamespaces(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getNamespaces")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.GetNamespaces(ctx)
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.GetNamespacesResponse)
	assert.Equal(t, &shared.GetNamespacesResponse{
		Items: []shared.Namespace{
			shared.Namespace{
				CreatedAt: types.MustTimeFromString("2024-01-20T10:11:46.428Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2022-01-07T04:43:36.873Z"),
			},
			shared.Namespace{
				CreatedAt: types.MustTimeFromString("2023-07-06T07:26:30.970Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2024-01-31T11:34:19.590Z"),
			},
			shared.Namespace{
				CreatedAt: types.MustTimeFromString("2022-08-10T13:43:30.983Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2022-07-15T08:41:36.371Z"),
			},
		},
	}, res.GetNamespacesResponse)

}

func TestArtifacts_GetRevisions(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getRevisions")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.GetRevisions(ctx, operations.GetRevisionsRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.GetRevisionsResponse)
	assert.Equal(t, &shared.GetRevisionsResponse{
		Items: []shared.Revision{
			shared.Revision{
				CreatedAt:     types.MustTimeFromString("2024-09-16T19:17:04.361Z"),
				Digest:        "sha256:6d1ef012b5674ad8a127ecfa9b5e6f5178d171b90ee462846974177fd9bdd39f",
				ID:            "<id>",
				NamespaceName: "<value>",
				Tags: []string{
					"<value>",
				},
				UpdatedAt: types.MustTimeFromString("2023-12-06T05:58:38.953Z"),
			},
		},
		NextPageToken: "<value>",
	}, res.GetRevisionsResponse)

}

func TestArtifacts_GetTags(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getTags")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.GetTags(ctx, operations.GetTagsRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.GetTagsResponse)
	assert.Equal(t, &shared.GetTagsResponse{
		Items: []shared.Tag{},
	}, res.GetTagsResponse)

}

func TestArtifacts_ListRemoteSources(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("listRemoteSources")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.ListRemoteSources(ctx, operations.ListRemoteSourcesRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.RemoteSource)
	assert.Equal(t, &shared.RemoteSource{
		Inputs: []shared.RemoteDocument{
			shared.RemoteDocument{
				RegistryURL: "https://well-lit-cap.net",
			},
			shared.RemoteDocument{
				RegistryURL: "https://vibrant-labourer.net",
			},
		},
		Output: shared.RemoteDocument{
			RegistryURL: "https://unfit-minor.biz",
		},
	}, res.RemoteSource)

}

func TestArtifacts_PostTags(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("postTags")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.PostTags(ctx, operations.PostTagsRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestArtifacts_SetVisibility(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("setVisibility")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.SetVisibility(ctx, operations.SetVisibilityRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestArtifacts_ArchiveNamespace(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("archiveNamespace")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Artifacts.SetArchived(ctx, operations.ArchiveNamespaceRequest{
		NamespaceName: "<value>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}
