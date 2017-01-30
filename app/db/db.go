// Package db connects to MySQL with simple database/sql
// TODO: Include configuration instead of taking it from the params
package db

// Using MySQL driver: https://github.com/go-sql-driver/mysql/
import _ "github.com/go-sql-driver/mysql"
import "database/sql"

var db *sql.DB
var err error

// Connect to Database
func Connect(user string, password string, database string) (*sql.DB, error) {
	// db, err = sql.Open("mysql", "root:<password>@/<dbname>")
	db, err = sql.Open("mysql", user+":"+password+"@tcp(localhost:6603)/"+database)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic(err.Error())
	}

	return db, err
}
