package project7_8

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//type action func(http.ResponseWriter, *http.Request, *gorm.DB)
type action func(http.ResponseWriter, *http.Request, Utils) interface{}

type Utils struct {
	db *gorm.DB
}

func (action action) ToHandlerFunc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseBody := action(w, r, Utils{db: db})
		enc := json.NewEncoder(w)
		err := enc.Encode(&responseBody)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
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
	router := mux.NewRouter().StrictSlash(false)
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

func Index(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	return "Declaration API"
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
