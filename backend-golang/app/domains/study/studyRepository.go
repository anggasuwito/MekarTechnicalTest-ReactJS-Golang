package study

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"
	"mekarTechnicalTest/utils/helper"
)

//StdRepository is a struct for depedency injection
type StdRepository struct {
	db *sql.DB
}

//StdRepositoryInterface is an interface from study repository
type StdRepositoryInterface interface {
	Studies(studyResult *models.Response)
}

//InitStdRepository is a function for inject study repository database
func InitStdRepository(db *sql.DB) *StdRepository {
	return &StdRepository{db: db}
}

//Studies is a function for get all studies from database
func (repository StdRepository) Studies(studyResult *models.Response) {
	var studies []*StdModel
	data, err := repository.db.Query(constants.GET_ALL_STUDIES, "A")
	getStudyErr := helper.ErrorNotNil(err)
	if getStudyErr {
		studyResult.Meta.Status = "Fail"
		studyResult.Meta.Code = 404
		studyResult.Meta.Message = "get all studies error"
		studyResult.Meta.Records = 0
		studyResult.Data = nil
		return
	}

	for data.Next() {
		study := StdModel{}
		errStudy := data.Scan(&study.StudyID, &study.Name)
		loopingStudyErr := helper.ErrorNotNil(errStudy)
		if loopingStudyErr {
			studyResult.Meta.Status = "Fail"
			studyResult.Meta.Code = 404
			studyResult.Meta.Message = "looping study error"
			studyResult.Meta.Records = 0
			studyResult.Data = nil
			return
		}
		studies = append(studies, &study)
	}
	countStudies := len(studies)
	studyResult.Meta.Status = "Success"
	studyResult.Meta.Code = 202
	studyResult.Meta.Message = "Success get all studies"
	studyResult.Meta.Records = countStudies
	studyResult.Data = studies
}
