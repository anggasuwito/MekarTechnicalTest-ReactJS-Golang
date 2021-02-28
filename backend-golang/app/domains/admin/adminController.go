package admin

import (
	"database/sql"
	"encoding/json"
	"log"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/helper"
	"net/http"
)

//AdmController is a struct for depedency injection
type AdmController struct {
	admUseCase AdmUseCaseInterface
}

//InitAdmController is a function for inject database
func InitAdmController(db *sql.DB) *AdmController {
	return &AdmController{admUseCase: InitAdmUseCase(db)}
}

//AdminLogin is a function for log in an admin
func (controller AdmController) AdminLogin(writer http.ResponseWriter, request *http.Request) {
	var adminLoginResult models.Response
	var adminData AdmModel
	json.NewDecoder(request.Body).Decode(&adminData)
	controller.admUseCase.AdminLogin(&adminLoginResult, &adminData)
	if adminLoginResult.Meta.Code == 404 {
		response := models.Response{
			Meta: adminLoginResult.Meta,
			Data: adminLoginResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(adminLoginResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Printf(" | %v", adminLoginResult.Meta.Message)
		helper.LogApp(adminLoginResult.Meta.Message)
	} else {
		response := models.Response{
			Meta: adminLoginResult.Meta,
			Data: adminLoginResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(adminLoginResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Println(" | Success admin login")
		helper.LogApp("Success admin login")
	}
}
