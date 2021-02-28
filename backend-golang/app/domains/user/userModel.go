package user

import (
	"mekarTechnicalTest/app/domains/study"
	"mekarTechnicalTest/app/domains/work"
)

//UsrModel is a blueprint or struct for user
type UsrModel struct {
	UserID       int            `json:"id"`
	Name         string         `json:"name"`
	BirthDate    string         `json:"birthDate"`
	NumberIDCard string         `json:"numberIdCard"`
	Work         work.WrkModel  `json:"work"`
	Study        study.StdModel `json:"study"`
}

//UsrBodyModel is a blueprint or struct for user body
type UsrBodyModel struct {
	Name         string `json:"name"`
	BirthDate    string `json:"birthDate"`
	NumberIDCard string `json:"numberIdCard"`
	WorkID       int    `json:"workId"`
	StudyID      int    `json:"studyId"`
}
