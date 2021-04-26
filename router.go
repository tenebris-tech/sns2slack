// Copyright (c) 2021 Tenebris Technologies Inc.
// https://www.tenebris.com
// Use of this source code is governed by the MIT license.
// Please see the LICENSE file for details.

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Create gorilla/mux router and load routes from route.go
func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Iterate through routes and add
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// Add routes for Not Found (404) and Method Not Allowed (405)
	router.NotFoundHandler = Logger(http.HandlerFunc(NotFoundHandler), "NotFound")
	router.MethodNotAllowedHandler = Logger(http.HandlerFunc(NotAllowedHandler), "NotAllowed")
	return router
}
