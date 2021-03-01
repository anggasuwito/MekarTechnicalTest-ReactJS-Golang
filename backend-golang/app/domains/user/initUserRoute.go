package user

import (
	"mekarTechnicalTest/app/middlewares"
	"mekarTechnicalTest/app/models"
)

//InitUserRoute is a function to initiate user route path and methods
func InitUserRoute(appConfig *models.AppConfigModel, mainRoute string) {
	userController := InitUsrController(appConfig.Database)
	user := appConfig.Route.PathPrefix(mainRoute).Subrouter()
	user.Use(middlewares.APILog, middlewares.UserJSONSetting, middlewares.AuthorizationToken)
	user.HandleFunc("/users", userController.Users).Methods("GET").Queries("keywords", "{keywords}", "page", "{page}", "limit", "{limit}")
	user.HandleFunc("/{userId}", userController.UserByID).Methods("GET")
	user.HandleFunc("/createuser", userController.CreateUser).Methods("POST")
	user.HandleFunc("/updateuser/{userId}", userController.UpdateUserByID).Methods("PUT")
	user.HandleFunc("/deleteuser/{userId}", userController.DeleteUserByID).Methods("DELETE")
}
