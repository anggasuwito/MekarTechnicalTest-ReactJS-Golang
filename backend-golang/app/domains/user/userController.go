package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/helper"
	"net/http"
)

//UsrController is a struct for depedency injection
type UsrController struct {
	userUseCase UsrUseCaseInterface
}

//InitUsrController is a function for inject database
func InitUsrController(db *sql.DB) *UsrController {
	return &UsrController{userUseCase: InitUsrUseCase(db)}
}

//Users is a function for get all users list
func (controller UsrController) Users(writer http.ResponseWriter, request *http.Request) {
	var userResult models.Response
	controller.userUseCase.Users(&userResult)
	response := models.Response{
		Meta: userResult.Meta,
		Data: userResult.Data,
	}
	byteOfResponse, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteOfResponse)
	log.Printf(" | %v %v", userResult.Meta.Status, userResult.Meta.Message)
	helper.LogApp(userResult.Meta.Status + " " + userResult.Meta.Message)
}

//UserByID is a function for get user by id
func (controller UsrController) UserByID(writer http.ResponseWriter, request *http.Request) {
	var userByIDResult models.Response
	userID := helper.GomuxPathVariable("userId", request)
	controller.userUseCase.UserByID(&userByIDResult, userID)
	response := models.Response{
		Meta: userByIDResult.Meta,
		Data: userByIDResult.Data,
	}
	byteOfResponse, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteOfResponse)
	log.Printf(" | %v %v", userByIDResult.Meta.Status, userByIDResult.Meta.Message)
	helper.LogApp(userByIDResult.Meta.Status + " " + userByIDResult.Meta.Message)
}

//CreateUser is a function for create new user
func (controller UsrController) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var createUserResult models.Response
	var userDataBody UsrBodyModel
	json.NewDecoder(request.Body).Decode(&userDataBody)
	controller.userUseCase.CreateUser(&createUserResult, &userDataBody)
	response := models.Response{
		Meta: createUserResult.Meta,
		Data: createUserResult.Data,
	}
	byteOfResponse, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteOfResponse)
	log.Printf(" | %v %v", createUserResult.Meta.Status, createUserResult.Meta.Message)
	helper.LogApp(createUserResult.Meta.Status + " " + createUserResult.Meta.Message)
}
