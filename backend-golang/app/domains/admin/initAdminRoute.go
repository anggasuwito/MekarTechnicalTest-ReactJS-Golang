package admin

import (
	"mekarTechnicalTest/app/middlewares"
	"mekarTechnicalTest/app/models"
)

//InitAdminRoute is a function to initiate admin route path and methods
func InitAdminRoute(appConfig *models.AppConfigModel, mainRoute string) {
	adminController := InitAdmController(appConfig.Database)
	admin := appConfig.Route.PathPrefix(mainRoute).Subrouter()
	admin.Use(middlewares.APILog, middlewares.AdminJSONSetting)
	admin.HandleFunc("/login", adminController.AdminLogin).Methods("POST")
}
