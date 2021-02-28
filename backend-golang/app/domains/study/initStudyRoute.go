package study

import (
	"mekarTechnicalTest/app/middlewares"
	"mekarTechnicalTest/app/models"
)

//InitStudyRoute is a function to initiate study route path and methods
func InitStudyRoute(appConfig *models.AppConfigModel, mainRoute string) {
	studyController := InitStdController(appConfig.Database)
	study := appConfig.Route.PathPrefix(mainRoute).Subrouter()
	study.Use(middlewares.APILog, middlewares.StudyJSONSetting, middlewares.AuthorizationToken)
	study.HandleFunc("/studies", studyController.Studies).Methods("GET")
}
