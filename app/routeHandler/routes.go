/*
 Routes are used to add all routes which API supports.
 There is lot of things to be done in here like extracting Bussiness and Database Logic out of this file, better error handling etc ...
*/

package routeHandler

import (
	"net/http"

	"github.com/edinkulovic/SimpleGoServer/app/routeHandler/user"
)

var routes map[string]func(http.ResponseWriter, *http.Request) (int, error)

// Init Method initialize routes
func Init() map[string]func(http.ResponseWriter, *http.Request) (int, error) {

	routes = make(map[string]func(http.ResponseWriter, *http.Request) (int, error))
	routes["/"] = HealtCheck
	// Users
	routes["/user"] = user.UserRoutes{}.GetByUsername
	routes["/user/login"] = user.UserRoutes{}.Login

	return routes
}

// HealtCheck method is used to check if server is live.
func HealtCheck(writer http.ResponseWriter, request *http.Request) (int, error) {
	return http.StatusOK, nil
}
