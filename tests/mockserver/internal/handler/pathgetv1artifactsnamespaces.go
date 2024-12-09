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

func pathGetV1ArtifactsNamespaces(dir *logging.HTTPFileDirectory) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")

		switch test {
		case "getNamespaces":
			dir.HandlerFunc("getNamespaces", testGetNamespacesGetNamespaces)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testGetNamespacesGetNamespaces(w http.ResponseWriter, req *http.Request) {
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
	respBody := &components.GetNamespacesResponse{
		Items: []components.Namespace{
			components.Namespace{
				CreatedAt: types.MustTimeFromString("2024-01-20T10:11:46.428Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2022-01-07T04:43:36.873Z"),
			},
			components.Namespace{
				CreatedAt: types.MustTimeFromString("2023-07-06T07:26:30.970Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2024-01-31T11:34:19.590Z"),
			},
			components.Namespace{
				CreatedAt: types.MustTimeFromString("2022-08-10T13:43:30.983Z"),
				ID:        "<id>",
				Name:      "<value>",
				UpdatedAt: types.MustTimeFromString("2022-07-15T08:41:36.371Z"),
			},
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
