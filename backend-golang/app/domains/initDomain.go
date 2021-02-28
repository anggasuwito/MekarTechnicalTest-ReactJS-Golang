package domains

import (
	"database/sql"
	"mekarTechnicalTest/app/domains/admin"
	"mekarTechnicalTest/app/domains/study"
	"mekarTechnicalTest/app/domains/user"
	"mekarTechnicalTest/app/domains/work"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"

	"github.com/gorilla/mux"
)

//InitDomain is a function to initiate domain service
func InitDomain(database *sql.DB, route *mux.Router) {
	appConfig := &models.AppConfigModel{
		Database: database,
		Route:    route,
	}
	admin.InitAdminRoute(appConfig, constants.MAIN_ROUTE_ADMIN)
	study.InitStudyRoute(appConfig, constants.MAIN_ROUTE_STUDY)
	work.InitWorkRoute(appConfig, constants.MAIN_ROUTE_WORK)
	user.InitUserRoute(appConfig, constants.MAIN_ROUTE_USER)
}
