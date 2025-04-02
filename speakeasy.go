// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package speakeasyclientsdkgo

import (
	"context"
	"fmt"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/internal/globals"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/internal/hooks"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/models/shared"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/retry"
	"github.com/speakeasy-api/speakeasy-client-sdk-go/v3/pkg/utils"
	"net/http"
	"time"
)

const (
	ServerProd string = "prod"
)

// ServerList contains the list of servers available to the SDK
var ServerList = map[string]string{
	ServerProd: "https://api.prod.speakeasyapi.dev",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

// Pointer provides a helper function to return a pointer to a type
func Pointer[T any](v T) *T { return &v }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	Server            string
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	Globals           globals.Globals
	RetryConfig       *retry.Config
	Hooks             *hooks.Hooks
	Timeout           *time.Duration
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	if c.Server == "" {
		c.Server = "prod"
	}

	return ServerList[c.Server], nil
}

// Speakeasy API: The Subscriptions API manages subscriptions for CLI and registry events
//
// /docs - The Speakeasy Platform Documentation
type Speakeasy struct {
	// REST APIs for working with Registry artifacts
	Artifacts *Artifacts
	// REST APIs for managing Authentication
	Auth *Auth
	// REST APIs for retrieving Code Samples
	CodeSamples *CodeSamples
	// REST APIs for managing the github integration
	Github *Github
	// REST APIs for managing Organizations (speakeasy L1 Tenancy construct)
	Organizations *Organizations
	// REST APIs for managing reports (lint reports, change reports, etc)
	Reports     *Reports
	SchemaStore *SchemaStore
	// REST APIs for managing short URLs
	ShortURLs *ShortURLs
	// REST APIs for managing subscriptions
	Subscriptions *Subscriptions
	// REST APIs for managing LLM OAS suggestions
	Suggest *Suggest
	// REST APIs for managing Workspaces (speakeasy tenancy)
	Workspaces *Workspaces
	// REST APIs for managing events captured by a speakeasy binary (CLI, GitHub Action etc)
	Events *Events

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Speakeasy)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Speakeasy) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServer allows the overriding of the default server by name
func WithServer(server string) SDKOption {
	return func(sdk *Speakeasy) {
		_, ok := ServerList[server]
		if !ok {
			panic(fmt.Errorf("server %s not found", server))
		}

		sdk.sdkConfiguration.Server = server
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.Client = client
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.Security = utils.AsSecuritySource(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

// WithWorkspaceID allows setting the WorkspaceID parameter for all supported operations
func WithWorkspaceID(workspaceID string) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.Globals.WorkspaceID = &workspaceID
	}
}

func WithRetryConfig(retryConfig retry.Config) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// WithTimeout Optional request timeout applied to each operation
func WithTimeout(timeout time.Duration) SDKOption {
	return func(sdk *Speakeasy) {
		sdk.sdkConfiguration.Timeout = &timeout
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Speakeasy {
	sdk := &Speakeasy{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.4.0",
			SDKVersion:        "3.26.1",
			GenVersion:        "2.565.1",
			UserAgent:         "speakeasy-sdk/go 3.26.1 2.565.1 0.4.0 github.com/speakeasy-api/speakeasy-client-sdk-go",
			Globals:           globals.Globals{},
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Artifacts = newArtifacts(sdk.sdkConfiguration)

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.CodeSamples = newCodeSamples(sdk.sdkConfiguration)

	sdk.Github = newGithub(sdk.sdkConfiguration)

	sdk.Organizations = newOrganizations(sdk.sdkConfiguration)

	sdk.Reports = newReports(sdk.sdkConfiguration)

	sdk.SchemaStore = newSchemaStore(sdk.sdkConfiguration)

	sdk.ShortURLs = newShortURLs(sdk.sdkConfiguration)

	sdk.Subscriptions = newSubscriptions(sdk.sdkConfiguration)

	sdk.Suggest = newSuggest(sdk.sdkConfiguration)

	sdk.Workspaces = newWorkspaces(sdk.sdkConfiguration)

	sdk.Events = newEvents(sdk.sdkConfiguration)

	return sdk
}
