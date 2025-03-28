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

func TestSubscriptions_CreateSubscription(t *testing.T) {
	t.Skip("incomplete test found please make sure to address the following errors: [`workflow step createSubscription.test referencing operation createSubscription not found in document`]")
}

func TestSubscriptions_ListRegistrySubscriptions(t *testing.T) {
	t.Skip("incomplete test found please make sure to address the following errors: [`workflow step listRegistrySubscriptions.test referencing operation listRegistrySubscriptions not found in document`]")
}

func TestSubscriptions_ActivateSubscriptionNamespace(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("activateSubscriptionNamespace")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Subscriptions.ActivateSubscriptionNamespace(ctx, operations.ActivateSubscriptionNamespaceRequest{
		NamespaceName:  "<value>",
		SubscriptionID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}

func TestSubscriptions_IgnoreSubscriptionNamespace(t *testing.T) {
	ctx := context.Background()

	testHTTPClient := createTestHTTPClient("ignoreSubscriptionNamespace")

	s := speakeasyclientsdkgo.New(
		speakeasyclientsdkgo.WithServerURL(utils.GetEnv("TEST_SERVER_URL", "http://localhost:18080")),
		speakeasyclientsdkgo.WithClient(testHTTPClient),
		speakeasyclientsdkgo.WithSecurity(shared.Security{
			APIKey: speakeasyclientsdkgo.String("<YOUR_API_KEY_HERE>"),
		}),
	)

	res, err := s.Subscriptions.IgnoreSubscriptionNamespace(ctx, operations.IgnoreSubscriptionNamespaceRequest{
		NamespaceName:  "<value>",
		SubscriptionID: "<id>",
	})
	require.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

}
