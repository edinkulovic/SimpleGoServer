package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/edinkulovic/SimpleGoServer/models"
)

/** Health check method for validating if server is live. **/
func healtCheck(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "plain/text")
	writer.WriteHeader(200)
	writer.Write([]byte("The beast is live"))
}

/** Retrieve Test User. **/
func getUser(writer http.ResponseWriter, request *http.Request) {
	// Creating Test User
	testUser := models.User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		UserName:  "john.doe",
		Email:     "john.doe@johndoe.it",
		Password:  "JohnDoe",
	}
	// NOTE: Second parameter error will not be handled for now
	userMarshalObject, _ := json.Marshal(testUser)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(200)
	writer.Write([]byte(userMarshalObject))

}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &mainServerHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = healtCheck
	mux["/user"] = getUser

	server.ListenAndServe()
}

type mainServerHandler struct{}

func (*mainServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
