package models

import (
	"database/sql"

	"github.com/gorilla/mux"
)

//AppConfigModel is a configurations for running application blueprint or struct
type AppConfigModel struct {
	Database *sql.DB
	Route    *mux.Router
}
