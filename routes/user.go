package project7_8

import (
	"encoding/json"
	"github.com/HRODEV/project7_8/dbActions"
	"github.com/HRODEV/project7_8/models"
	"github.com/HRODEV/project7_8/services"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func UserAuthGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	authorizationHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(authorizationHeader) != 2 || authorizationHeader[0] != "Basic" {
		http.Error(w, "Authorization header not found or empty", http.StatusBadRequest)
		return nil
	}

	authService := services.AuthService{}
	userCredentials := strings.Split(authService.DecodeAuthorzationHeader(authorizationHeader[1]), ":")

	// Verify email and password
	user := models.User{}
	utils.db.Where("email = ? AND password = ?", userCredentials[0], userCredentials[1]).Find(&user)

	if user.ID == 0 {
		http.Error(w, "Combination of username and password cannot be found", http.StatusUnauthorized)
		return nil
	}

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// token claims
	claims["sub"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with our secret
	tokenString, err := token.SignedString([]byte("gXV6aSf4NJ7Ah@S!DrE5$Pm!"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return &models.Auth{UserId: user.ID, Token: tokenString}
}

func UserGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	var user models.User
	dbActions.GetUserByID(utils.currentUser.ID, &user, utils.db)

	return &user
}

//
func UserPost(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	// Convert request body to interface
	var user models.User

	reqBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBody, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	err = dbActions.CreateUser(&user, utils.db)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}

	return &user
}

func UserProjectsGet(w http.ResponseWriter, r *http.Request, utils Utils) interface{} {
	http.Error(w, "Not implemented yet", http.StatusNotImplemented)
	return nil
}
