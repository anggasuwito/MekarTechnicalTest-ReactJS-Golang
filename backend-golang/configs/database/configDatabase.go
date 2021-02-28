package database

import (
	"database/sql"
	"fmt"
	"log"
	"mekarTechnicalTest/utils/helper"

	//MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

//ConfigDatabase is a function for create a new connection to database
func ConfigDatabase() *sql.DB {
	dbUser := helper.GetKeyValue("DB_USER", "root")
	dbPass := helper.GetKeyValue("DB_PASSWORD", "password")
	dbHost := helper.GetKeyValue("DB_HOST", "localhost")
	dbPort := helper.GetKeyValue("DB_PORT", "3306")
	dbSchemaName := helper.GetKeyValue("DB_SCHEMANAME", "dbName")

	dbPath := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbSchemaName)
	dbConnection, _ := sql.Open("mysql", dbPath)
	err := dbConnection.Ping()
	helper.FatalError(err, "DB error")
	logData := fmt.Sprintf("Starting Database Connection : %v", dbSchemaName)
	log.Printf(" | %v", logData)
	helper.LogApp(logData)
	return dbConnection
}
