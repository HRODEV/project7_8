package project7_8

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

type action func(http.ResponseWriter, *http.Request, *gorm.DB)

func (action action) ToHandlerFunc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		action(w, r, db)
	}
}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc action
}

func decorateJsonHeader(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		handler(w, r)
	}
}

func decorateBasicHeaders(handler http.HandlerFunc) http.HandlerFunc {
	return decorateJsonHeader(handler)
}

type Routes []Route

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = decorateBasicHeaders(route.HandlerFunc.ToHandlerFunc(db))
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	w.Write([]byte("Declaration API"))
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
