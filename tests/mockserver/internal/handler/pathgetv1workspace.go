// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathGetV1Workspace(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "getWorkspaceByContext[0]":
			dir.HandlerFunc("getWorkspaceByContext", testGetWorkspaceByContextGetWorkspaceByContext0)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testGetWorkspaceByContextGetWorkspaceByContext0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.SecurityHeader(req, "x-api-key", true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.SecurityHeader(req, "x-workspace-identifier", true); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody := &components.WorkspaceAndOrganization{
		Organization: components.Organization{
			AccountType:       components.AccountTypeBusiness,
			CreatedAt:         types.MustTimeFromString("2023-09-05T11:33:52.011Z"),
			ID:                "<id>",
			Name:              "<value>",
			Slug:              "<value>",
			SsoActivated:      false,
			TelemetryDisabled: false,
			UpdatedAt:         types.MustTimeFromString("2023-07-26T06:33:15.810Z"),
		},
		Workspace: components.Workspace{
			CreatedAt:         types.MustTimeFromString("2024-11-29T01:50:48.954Z"),
			ID:                "<id>",
			Name:              "<value>",
			OrganizationID:    "<id>",
			Slug:              "<value>",
			TelemetryDisabled: types.Bool(true),
			UpdatedAt:         types.MustTimeFromString("2023-05-10T02:28:23.533Z"),
			Verified:          true,
		},
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}