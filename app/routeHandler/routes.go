/*
 Routes are used to add all routes which API supports.
 There is lot of things to be done in here like extracting Bussiness and Database Logic out of this file, better error handling etc ...
*/

package routeHandler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"errors"

	"github.com/edinkulovic/SimpleGoServer/models"
)

// SQLDb Property Shared Database Instance
var SQLDb *sql.DB
var err error
var mux map[string]func(http.ResponseWriter, *http.Request) (int, error)

// Init Method initialize routes and Database
func Init(database *sql.DB) map[string]func(http.ResponseWriter, *http.Request) (int, error) {
	SQLDb = database

	mux = make(map[string]func(http.ResponseWriter, *http.Request) (int, error))
	mux["/"] = HealtCheck
	// Users
	mux["/user"] = GetUser

	return mux
}

// HealtCheck method is used to check if server is live.
func HealtCheck(writer http.ResponseWriter, request *http.Request) (int, error) {
	return http.StatusOK, nil
}

// GetUser Method retrieves user by Username
func GetUser(writer http.ResponseWriter, request *http.Request) (int, error) {
	if request.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	var username = request.URL.Query().Get("username")
	var testUser models.User

	err := SQLDb.QueryRow("SELECT ID, Username, FirstName, LastName FROM Users WHERE Username=?", username).Scan(
		&testUser.ID, &testUser.UserName, &testUser.FirstName, &testUser.LastName)

	if err != nil {
		return http.StatusBadRequest, err
	}

	// NOTE: Second parameter error will not be handled for now
	userMarshalObject, _ := json.Marshal(testUser)
	writer.Write([]byte(userMarshalObject))

	return http.StatusOK, nil
}
