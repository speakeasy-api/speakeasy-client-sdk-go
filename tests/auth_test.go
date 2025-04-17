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

func TestAuth_GetWorkspaceAccess(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getWorkspaceAccess")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Auth.GetAccess(ctx, operations.GetWorkspaceAccessRequest{})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.AccessDetails)
	assert.Equal(t, &shared.AccessDetails{
		GenerationAllowed: true,
		Message:           "<value>",
	}, res.AccessDetails)

}

func TestAuth_GetAccessToken(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getAccessToken")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
	)

	res, err := s.Auth.GetAccessToken(ctx, operations.GetAccessTokenRequest{
		WorkspaceID: "<value>",
	})
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.AccessToken)
	assert.Equal(t, &shared.AccessToken{
		AccessToken: "<value>",
		Claims:      shared.Claims{},
		User:        shared.AccessTokenUser{},
	}, res.AccessToken)

}

func TestAuth_GetUser(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("getUser")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Auth.GetUser(ctx)
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.User)
	assert.Equal(t, &shared.User{
		Admin:         false,
		Confirmed:     true,
		CreatedAt:     types.MustTimeFromString("2024-09-05T08:59:40.988Z"),
		DisplayName:   "Tianna_Prohaska",
		Email:         "Morton82@gmail.com",
		EmailVerified: false,
		ID:            "<id>",
		UpdatedAt:     types.MustTimeFromString("2022-08-28T03:26:52.335Z"),
		Whitelisted:   true,
	}, res.User)

}

func TestAuth_ValidateAPIKey(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("validateApiKey")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Auth.ValidateAPIKey(ctx)
	require.NoError(t, err)
	assert.Contains(t, []any{200, 201, 202, 203, 204, 205, 206, 207, 208, 226}, res.StatusCode)
	assert.NotNil(t, res.APIKeyDetails)
	assert.Equal(t, &shared.APIKeyDetails{
		AccountTypeV2: shared.AccountTypeEnterprise,
		BillingAddOns: []shared.BillingAddOn{
			shared.BillingAddOnSnippetAi,
		},
		EnabledFeatures: []string{
			"<value>",
			"<value>",
			"<value>",
		},
		OrgSlug:            "<value>",
		TelemetryDisabled:  true,
		WorkspaceCreatedAt: types.MustTimeFromString("2024-04-24T00:30:38.626Z"),
		WorkspaceID:        "<id>",
		WorkspaceSlug:      "<value>",
	}, res.APIKeyDetails)

}
