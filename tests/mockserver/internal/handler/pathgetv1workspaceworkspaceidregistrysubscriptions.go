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

func pathGetV1WorkspaceWorkspaceIDRegistrySubscriptions(dir *logging.HTTPFileDirectory) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")

		switch test {
		case "listRegistrySubscriptions":
			dir.HandlerFunc("listRegistrySubscriptions", testListRegistrySubscriptionsListRegistrySubscriptions)(w, req)
		default:
			http.Error(w, "Unknown test: "+test, http.StatusBadRequest)
		}
	}
}

func testListRegistrySubscriptionsListRegistrySubscriptions(w http.ResponseWriter, req *http.Request) {
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
	respBody := []components.RegistrySubscription{
		components.RegistrySubscription{
			CreatedAt:            types.MustTimeFromString("2022-12-26T20:06:41.993Z"),
			EventType:            components.EventTypeUpdate,
			ID:                   "<id>",
			NamespaceName:        "<value>",
			SubscriptionSettings: "<value>",
			SubscriptionType:     components.SubscriptionTypeCli,
			UpdatedAt:            types.MustTimeFromString("2023-10-12T18:58:27.338Z"),
			WorkspaceID:          "<id>",
		},
		components.RegistrySubscription{
			CreatedAt:            types.MustTimeFromString("2022-07-07T11:51:20.859Z"),
			EventType:            components.EventTypeUpdate,
			ID:                   "<id>",
			NamespaceName:        "<value>",
			SubscriptionSettings: "<value>",
			SubscriptionType:     components.SubscriptionTypeCli,
			UpdatedAt:            types.MustTimeFromString("2022-03-26T04:16:26.649Z"),
			WorkspaceID:          "<id>",
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
