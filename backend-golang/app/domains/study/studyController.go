package study

import (
	"database/sql"
	"encoding/json"
	"log"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/helper"
	"net/http"
)

//StdController is a struct for depedency injection
type StdController struct {
	studyUseCase StdUseCaseInterface
}

//InitStdController is a function for inject database
func InitStdController(db *sql.DB) *StdController {
	return &StdController{studyUseCase: InitStdUseCase(db)}
}

//Studies is a function for get all studies list
func (controller StdController) Studies(writer http.ResponseWriter, request *http.Request) {
	var studyResult models.Response
	controller.studyUseCase.Studies(&studyResult)
	if studyResult.Meta.Code == 404 {
		response := models.Response{
			Meta: studyResult.Meta,
			Data: studyResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(studyResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Printf(" | %v", studyResult.Meta.Message)
		helper.LogApp(studyResult.Meta.Message)
	} else {
		response := models.Response{
			Meta: studyResult.Meta,
			Data: studyResult.Data,
		}
		byteOfResponse, _ := json.Marshal(response)
		writer.WriteHeader(studyResult.Meta.Code)
		writer.Write(byteOfResponse)
		log.Println(" | Success get all studies")
		helper.LogApp("Success get all studies")
	}
}
