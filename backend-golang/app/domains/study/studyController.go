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
	response := models.Response{
		Meta: studyResult.Meta,
		Data: studyResult.Data,
	}
	byteOfResponse, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteOfResponse)
	log.Printf(" | %v %v", studyResult.Meta.Status, studyResult.Meta.Message)
	helper.LogApp(studyResult.Meta.Status + " " + studyResult.Meta.Message)
}
