package project7_8

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
		"/",
		Index,
	},

	Route{
		"DeclarationsGet",
		"GET",
		"/declarations",
		DeclarationsGet,
	},

	Route{
		"DeclarationsIdDelete",
		"DELETE",
		"/declarations/{id}",
		DeclarationsIdDelete,
	},

	Route{
		"DeclarationsIdGet",
		"GET",
		"/declarations/{id}",
		DeclarationsIdGet,
	},

	Route{
		"DeclarationsIdPatch",
		"PATCH",
		"/declarations/{id}",
		DeclarationsIdPatch,
	},

	Route{
		"DeclarationsPost",
		"POST",
		"/declarations",
		DeclarationsPost,
	},

	Route{
		"ReceiptIdGet",
		"GET",
		"/receipt/{id}",
		ReceiptIdGet,
	},

	Route{
		"ReceiptIdImageGet",
		"GET",
		"/receipt/{id}/image",
		ReceiptIdImageGet,
	},

	Route{
		"ReceiptPost",
		"POST",
		"/receipt",
		ReceiptPost,
	},

	Route{
		"UserGet",
		"GET",
		"/user",
		UserGet,
	},

	Route{
		"UserPost",
		"POST",
		"/user",
		UserPost,
	},

	Route{
		"UserAuthGet",
		"GET",
		"/user/auth",
		UserAuthGet,
	},

	Route{
		"UserProjectsGet",
		"GET",
		"/user/projects",
		UserProjectsGet,
	},
}
