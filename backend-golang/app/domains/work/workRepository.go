package work

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"
	"mekarTechnicalTest/utils/helper"
)

//WrkRepository is a struct for depedency injection
type WrkRepository struct {
	db *sql.DB
}

//WrkRepositoryInterface is an interface from work repository
type WrkRepositoryInterface interface {
	Works(workResult *models.Response)
}

//InitWrkRepository is a function for inject work repository database
func InitWrkRepository(db *sql.DB) *WrkRepository {
	return &WrkRepository{db: db}
}

//Works is a function for get all works from database
func (repository WrkRepository) Works(workResult *models.Response) {
	var works []*WrkModel
	data, err := repository.db.Query(constants.GET_ALL_WORKS, "A")
	getWorkErr := helper.ErrorNotNil(err)
	if getWorkErr {
		workResult.Meta.Status = "Fail"
		workResult.Meta.Code = 404
		workResult.Meta.Message = "get all works error"
		workResult.Meta.Records = 0
		workResult.Data = nil
		return
	}

	for data.Next() {
		work := WrkModel{}
		errWork := data.Scan(&work.WorkID, &work.Name)
		loopingWorkErr := helper.ErrorNotNil(errWork)
		if loopingWorkErr {
			workResult.Meta.Status = "Fail"
			workResult.Meta.Code = 404
			workResult.Meta.Message = "looping work error"
			workResult.Meta.Records = 0
			workResult.Data = nil
			return
		}
		works = append(works, &work)
	}
	countWorks := len(works)
	workResult.Meta.Status = "Success"
	workResult.Meta.Code = 202
	workResult.Meta.Message = "Success get all works"
	workResult.Meta.Records = countWorks
	workResult.Data = works
}
