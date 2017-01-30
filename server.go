package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/edinkulovic/SimpleGoServer/app/config"
	"github.com/edinkulovic/SimpleGoServer/app/db"
	"github.com/edinkulovic/SimpleGoServer/app/routeHandler"
)

// AppConfig Property Stores the Configuration retrieved from environment files
var AppConfig config.Config
var mux map[string]func(http.ResponseWriter, *http.Request) (int, error)
var sqlDb *sql.DB

type appContext struct {
	db     *sql.DB
	config config.Config
}

func main() {
	// Load Application Configurations
	AppConfig, err := config.Load()
	if err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// Connect to Database Server
	sqlDb, err := db.Connect(AppConfig.DatabaseUser, AppConfig.DatabasePass, AppConfig.DatabaseName)
	if err != nil {
		panic(fmt.Errorf("Unable to connect to database: %s", err))
	}

	defer sqlDb.Close()

	// Setup Http Server
	server := http.Server{
		Addr:         ":" + strconv.Itoa(AppConfig.ServerPort),
		Handler:      &mainServerHandler{},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Initialize Route Handler
	mux = routeHandler.Init(sqlDb)

	// Start the server
	err = server.ListenAndServe()

	if err != nil {
		panic(fmt.Errorf("Unable to start server: %s", err))
	}

	fmt.Println("Starting Server at port: " + strconv.Itoa(AppConfig.ServerPort))
}

type mainServerHandler struct{}

func (*mainServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json charset=UTF-8")

	if h, ok := mux[r.URL.Path]; ok { // Check if route exists
		status, err := h(w, r)
		if err != nil {
			switch status {
			case http.StatusNotFound, http.StatusInternalServerError:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			case http.StatusBadRequest:
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			default:
				// Catch any other errors we haven't explicitly handled
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}
	} else {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
	}
}
