package work

import (
	"database/sql"
	"encoding/json"
	"log"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/helper"
	"net/http"
)

//WrkController is a struct for depedency injection
type WrkController struct {
	workUseCase WrkUseCaseInterface
}

//InitWrkController is a function for inject database
func InitWrkController(db *sql.DB) *WrkController {
	return &WrkController{workUseCase: InitWrkUseCase(db)}
}

//Works is a function for get all works list
func (controller WrkController) Works(writer http.ResponseWriter, request *http.Request) {
	var workResult models.Response
	controller.workUseCase.Works(&workResult)
	response := models.Response{
		Meta: workResult.Meta,
		Data: workResult.Data,
	}
	byteOfResponse, _ := json.Marshal(response)
	writer.WriteHeader(http.StatusOK)
	writer.Write(byteOfResponse)
	log.Printf(" | %v %v", workResult.Meta.Status, workResult.Meta.Message)
	helper.LogApp(workResult.Meta.Status + " " + workResult.Meta.Message)
}
