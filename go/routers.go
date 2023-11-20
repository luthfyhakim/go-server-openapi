/*
 * Todolist RESTful API
 *
 * OpenAPI for Todolist RESTful API
 *
 * API version: 1
 * Contact: luthfyhakim250404@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
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

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v1/",
		Index,
	},

	Route{
		"TodolistGet",
		strings.ToUpper("Get"),
		"/api/v1/todolist",
		TodolistGet,
	},

	Route{
		"TodolistPost",
		strings.ToUpper("Post"),
		"/api/v1/todolist",
		TodolistPost,
	},

	Route{
		"TodolistTodolistIdDelete",
		strings.ToUpper("Delete"),
		"/api/v1/todolist/{todolistId}",
		TodolistTodolistIdDelete,
	},

	Route{
		"TodolistTodolistIdPut",
		strings.ToUpper("Put"),
		"/api/v1/todolist/{todolistId}",
		TodolistTodolistIdPut,
	},
}