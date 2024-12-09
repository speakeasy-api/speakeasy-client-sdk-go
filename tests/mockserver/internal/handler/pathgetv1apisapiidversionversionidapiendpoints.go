// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"net/http"
)

func pathGetV1ApisAPIIDVersionVersionIDAPIEndpoints(dir *logging.HTTPFileDirectory) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")

		switch test {
		case "getAllForVersionApiEndpoints":
			dir.HandlerFunc("getAllForVersionApiEndpoints", testGetAllForVersionAPIEndpointsGetAllForVersionAPIEndpoints)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testGetAllForVersionAPIEndpointsGetAllForVersionAPIEndpoints(w http.ResponseWriter, req *http.Request) {
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
	respBody := []components.APIEndpoint{
		components.APIEndpoint{
			APIEndpointID: "<id>",
			APIID:         "<id>",
			CreatedAt:     types.MustTimeFromString("2024-06-17T00:07:30.468Z"),
			Description:   "cutover knottily productive thump scrabble abaft gracious pulp",
			DisplayName:   "Alek.Kunze41",
			Method:        "<value>",
			Path:          "/home/user",
			UpdatedAt:     types.MustTimeFromString("2024-07-27T06:29:44.391Z"),
			VersionID:     "<id>",
			WorkspaceID:   "<id>",
		},
		components.APIEndpoint{
			APIEndpointID: "<id>",
			APIID:         "<id>",
			CreatedAt:     types.MustTimeFromString("2022-06-22T07:41:12.922Z"),
			Description:   "sparse obligation er honorable offensively shallow",
			DisplayName:   "Jarrell82",
			Method:        "<value>",
			Path:          "/sys",
			UpdatedAt:     types.MustTimeFromString("2024-08-17T07:17:55.780Z"),
			VersionID:     "<id>",
			WorkspaceID:   "<id>",
		},
		components.APIEndpoint{
			APIEndpointID: "<id>",
			APIID:         "<id>",
			CreatedAt:     types.MustTimeFromString("2023-06-20T19:40:53.688Z"),
			Description:   "yieldingly widow blah phooey uh-huh corny fess",
			DisplayName:   "Moshe.Streich",
			Method:        "<value>",
			Path:          "/selinux",
			UpdatedAt:     types.MustTimeFromString("2023-07-07T21:35:49.283Z"),
			VersionID:     "<id>",
			WorkspaceID:   "<id>",
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
