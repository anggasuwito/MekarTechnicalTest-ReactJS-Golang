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
	if userResult.Meta.Code == 404 {
		response := models.Response{
			Meta: userResult.Meta,
			Data: userResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(userResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Printf(" | %v", userResult.Meta.Message)
		helper.LogApp(userResult.Meta.Message)
	} else {
		response := models.Response{
			Meta: userResult.Meta,
			Data: userResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(userResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Println(" | Success get all users")
		helper.LogApp("Success get all users")
	}
}

//UserByID is a function for get user by id
func (controller UsrController) UserByID(writer http.ResponseWriter, request *http.Request) {
	var userByIDResult models.Response
	userID := helper.GomuxPathVariable("userId", request)
	controller.userUseCase.UserByID(&userByIDResult, userID)
	if userByIDResult.Meta.Code == 404 {
		response := models.Response{
			Meta: userByIDResult.Meta,
			Data: userByIDResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(userByIDResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Printf(" | %v", userByIDResult.Meta.Message)
		helper.LogApp(userByIDResult.Meta.Message)
	} else {
		response := models.Response{
			Meta: userByIDResult.Meta,
			Data: userByIDResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(userByIDResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Println(" | Success get user by id")
		helper.LogApp("Success get user by id")
	}
}

//CreateUser is a function for create new user
func (controller UsrController) CreateUser(writer http.ResponseWriter, request *http.Request) {
	var createUserResult models.Response
	var userDataBody UsrBodyModel
	json.NewDecoder(request.Body).Decode(&userDataBody)
	controller.userUseCase.CreateUser(&createUserResult, &userDataBody)
	if createUserResult.Meta.Code == 404 {
		response := models.Response{
			Meta: createUserResult.Meta,
			Data: createUserResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(createUserResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Printf(" | %v", createUserResult.Meta.Message)
		helper.LogApp(createUserResult.Meta.Message)
	} else {
		response := models.Response{
			Meta: createUserResult.Meta,
			Data: createUserResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(createUserResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Println(" | Success create new user ")
		helper.LogApp("Success create new user ")
	}
}
