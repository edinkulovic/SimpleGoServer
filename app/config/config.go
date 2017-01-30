package config

import "os"

// Config stores the application-wide configurations
var appConfig Config

// Config structure for Retrieving Configuration from the Env files
type Config struct {
	ServerPort      int
	IsDevelopment   bool
	DatabaseUser    string
	DatabasePass    string
	DatabaseName    string
	DatabasePort    int
	DatabaseAddress string
}

// Load the files from environment
func Load() (Config, error) {
	var appConfig Config
	var err error
	//appConfig.ServerPort, err = strconv.ParseInt(os.Getenv("SERVER_PORT"), 64, 32)
	//appConfig.DatabasePort, err = strconv.ParseInt(os.Getenv("DATABASE_PORT"), 64, 32)

	appConfig.ServerPort = 8002
	appConfig.DatabasePort = 6603

	appConfig.DatabaseUser = os.Getenv("DATABASE_USER")
	appConfig.DatabasePass = os.Getenv("DATABASE_PASS")
	appConfig.DatabaseName = os.Getenv("DATABASE_NAME")
	appConfig.DatabaseAddress = os.Getenv("DATABASE_ADDRESS")

	return appConfig, err
}
