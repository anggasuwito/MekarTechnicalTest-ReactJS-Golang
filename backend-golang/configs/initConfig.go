package configs

import (
	"database/sql"
	"mekarTechnicalTest/configs/database"
	"mekarTechnicalTest/configs/router"

	"github.com/gorilla/mux"
)

//InitConfig is function for initiate configuration
func InitConfig() (*sql.DB, *mux.Router) {
	database := database.ConfigDatabase()
	route := router.ConfigRouter()
	return database, route
}
