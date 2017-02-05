package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/edinkulovic/SimpleGoServer/models"
)

type (
	UserRoutes struct{}
)

// GetByUsername Method retrieves user by Username
func (ur UserRoutes) GetByUsername(writer http.ResponseWriter, request *http.Request) (int, error) {
	if request.Method != http.MethodGet {
		return http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	username := request.URL.Query().Get("username")
	// Empty Struct does not take any memory models.User{}
	user, err := models.User{}.GetByUsername(username)

	if err != nil {
		// TODO: Logging
		return http.StatusBadRequest, err
	}

	// NOTE: Second parameter error will not be handled for now
	userMarshalObject, _ := json.Marshal(user)
	writer.Write([]byte(userMarshalObject))

	return http.StatusOK, nil
}

// Login Method retrieves user data if Credentials are correct
func (ur UserRoutes) Login(writer http.ResponseWriter, request *http.Request) (int, error) {
	if request.Method != http.MethodPost && request.Method != http.MethodOptions {
		return http.StatusMethodNotAllowed, errors.New(http.StatusText(http.StatusMethodNotAllowed))
	}

	username := request.URL.Query().Get("username")
	password := request.URL.Query().Get("password")
	// Empty Struct does not take any memory models.User{}
	user, err := models.User{}.Login(username, password)

	if err != nil {
		// TODO: Logging
		// TODO: Add custom Status Codes
		return http.StatusUnauthorized, err
	}

	// NOTE: Second parameter error will not be handled for now
	userMarshalObject, _ := json.Marshal(user)
	writer.Write([]byte(userMarshalObject))

	return http.StatusOK, nil
}
