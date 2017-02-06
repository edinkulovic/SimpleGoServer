/*
 Routes are used to add all routes which API supports.
 There is lot of things to be done in here like extracting Bussiness and Database Logic out of this file, better error handling etc ...
*/

package routeHandler

import (
	"errors"
	"net/http"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/edinkulovic/SimpleGoServer/app/routeHandler/user"
	"github.com/edinkulovic/SimpleGoServer/models"
)

var routes map[string]func(http.ResponseWriter, *http.Request) (int, error)

// Init Method initialize routes
func Init() map[string]func(http.ResponseWriter, *http.Request) (int, error) {

	routes = make(map[string]func(http.ResponseWriter, *http.Request) (int, error))
	routes["/"] = HealtCheck
	// Users
	routes["/user"] = validateJWT(user.UserRoutes{}.GetByUsername)
	routes["/user/login"] = user.UserRoutes{}.Login

	return routes
}

func validateJWT(protectedHandler func(http.ResponseWriter, *http.Request) (int, error)) func(http.ResponseWriter, *http.Request) (int, error) {
	return (func(res http.ResponseWriter, req *http.Request) (int, error))(func(res http.ResponseWriter, req *http.Request) (int, error) {

		headerToken := req.Header.Get("Token")

		if headerToken == "" {
			return http.StatusUnauthorized, errors.New("Token Missing")
		}

		// Return a Token using the cookie
		token, err := jwt.ParseWithClaims(headerToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			// Make sure token's signature wasn't changed
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte("secret"), nil
		})

		if err != nil {
			fmt.Println(err)
			return http.StatusUnauthorized, err
		}

		// Grab the tokens claims and pass it into the original request
		if _, ok := token.Claims.(*models.Claims); ok && token.Valid {
			return protectedHandler(res, req)
		}

		return http.StatusUnauthorized, err
	})
}

// HealtCheck method is used to check if server is live.
func HealtCheck(writer http.ResponseWriter, request *http.Request) (int, error) {
	return http.StatusOK, nil
}
