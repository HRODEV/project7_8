package project7_8

import (
	"encoding/json"
	"net/http"

	"fmt"
	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"strings"
)

//type action func(http.ResponseWriter, *http.Request, *gorm.DB)
type action func(http.ResponseWriter, *http.Request, Utils) interface{}

type Utils struct {
	db          *gorm.DB
	currentUser *models.User
}

func (action action) ToHandlerFunc(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responseBody := action(w, r, Utils{db: db})

		if responseBody != nil {
			enc := json.NewEncoder(w)
			err := enc.Encode(&responseBody)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
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

func requireAuthentication(handler action) action {
	return func(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
		authorizationHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

		if len(authorizationHeader) != 2 || authorizationHeader[0] != "Bearer" {
			http.Error(w, "Authorization header not found or empty", http.StatusBadRequest)
			return nil
		}

		authService := services.AuthService{}
		tokenString := authService.DecodeAuthorzationHeader(authorizationHeader[1])

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte("gXV6aSf4NJ7Ah@S!DrE5$Pm!"), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return nil
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			utils.currentUser = &models.User{}
			dbActions.GetUserByID(uint(claims["sub"].(float64)), utils.currentUser, utils.db)

			return handler(w, r, utils)
		} else {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return nil
		}
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
		requireAuthentication(DeclarationsGet),
	},

	Route{
		"DeclarationsIdGet",
		"GET",
		"/declarations/{id}",
		requireAuthentication(DeclarationsIdGet),
	},

	Route{
		"DeclarationsIdGetImage",
		"GET",
		"/declarations/{id}/image",
		requireAuthentication(DeclarationsIdGetImage),
	},

	Route{
		"DeclarationsIdDelete",
		"DELETE",
		"/declarations/{id}",
		requireAuthentication(DeclarationsIdDelete),
	},

	Route{
		"DeclarationsIdPatch",
		"PATCH",
		"/declarations/{id}",
		requireAuthentication(DeclarationsIdPatch),
	},

	Route{
		"DeclarationsPost",
		"POST",
		"/declarations",
		requireAuthentication(DeclarationsPost),
	},

	Route{
		"ReceiptIdGet",
		"GET",
		"/receipt/{id}",
		requireAuthentication(ReceiptIdGet),
	},

	Route{
		"ReceiptIdImageGet",
		"GET",
		"/receipt/{id}/image",
		requireAuthentication(ReceiptIdImageGet),
	},

	Route{
		"ReceiptPost",
		"POST",
		"/receipt",
		requireAuthentication(ReceiptPost),
	},

	Route{
		"UserGet",
		"GET",
		"/user",
		requireAuthentication(UserGet),
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
		requireAuthentication(UserProjectsGet),
	},
}
