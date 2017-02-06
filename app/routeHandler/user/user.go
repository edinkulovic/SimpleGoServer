package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"strconv"

	"github.com/edinkulovic/SimpleGoServer/models"
)

type (
	UserRoutes struct{}
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
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

	var loginRequest LoginRequest

	if request.Body == nil {
		return http.StatusBadRequest, errors.New("Body Empty")
	}

	err := json.NewDecoder(request.Body).Decode(&loginRequest)

	if err != nil {
		fmt.Println(err)
		return http.StatusBadRequest, errors.New("Incorrec Body Format")
	}

	// Empty Struct does not take any memory models.User{}
	user, err := models.User{}.Login(loginRequest.Username, loginRequest.Password)

	if err != nil {
		fmt.Println(err)
		return http.StatusUnauthorized, errors.New("Username or Password is Incorrect")
	}

	token, expiration, err := models.Claims{}.CreateToken(user.UserName)

	if err != nil {
		fmt.Println(err)
		return http.StatusUnauthorized, err
	}

	// TODO: Move this into Redis DB
	writer.Header().Set("Token", token)
	writer.Header().Set("TokenExpiration", strconv.FormatInt(int64(expiration), 10))

	// NOTE: Second parameter error will not be handled for now
	userMarshalObject, _ := json.Marshal(user)
	writer.Write([]byte(userMarshalObject))

	return http.StatusOK, nil
}
