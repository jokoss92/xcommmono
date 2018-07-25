package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {

	// connect to our database server with data source name
	// data source name configuration has the following parameters :
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]

	// example config :
	// user:password@tcp(127.0.0.1:3306)/database

	dbDriver := "mysql"
	dbUser := "root"
	dbURL := "localhost"
	dbPort := "3306"
	dbPass := "admin12345"
	dbName := "db_company"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbURL+":"+dbPort+")/"+dbName)
	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
