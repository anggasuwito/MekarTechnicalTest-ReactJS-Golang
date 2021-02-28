package work

import (
	"mekarTechnicalTest/app/middlewares"
	"mekarTechnicalTest/app/models"
)

//InitWorkRoute is a function to initiate work route path and methods
func InitWorkRoute(appConfig *models.AppConfigModel, mainRoute string) {
	workController := InitWrkController(appConfig.Database)
	work := appConfig.Route.PathPrefix(mainRoute).Subrouter()
	work.Use(middlewares.APILog, middlewares.WorkJSONSetting, middlewares.AuthorizationToken)
	work.HandleFunc("/works", workController.Works).Methods("GET")
}
