package admin

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"
	"mekarTechnicalTest/utils/helper"
)

//AdmUseCase is a struct for depedency injection
type AdmUseCase struct {
	db            *sql.DB
	admRepository AdmRepositoryInterface
}

//AdmUseCaseInterface is an interface from admin use case
type AdmUseCaseInterface interface {
	AdminLogin(adminLoginResult *models.Response, adminData *AdmModel)
}

//InitAdmUseCase is a function for inject admin use case
func InitAdmUseCase(db *sql.DB) *AdmUseCase {
	return &AdmUseCase{db: db, admRepository: InitAdmRepository(db)}
}

//AdminLogin in admin use case is a function for check some validation or put some logic
func (useCase AdmUseCase) AdminLogin(adminLoginResult *models.Response, adminData *AdmModel) {
	if len(adminData.Username) == 0 {
		adminLoginResult.Meta.Status = "Fail"
		adminLoginResult.Meta.Code = 404
		adminLoginResult.Meta.Message = "username is empty"
		adminLoginResult.Meta.Records = 0
		adminLoginResult.Data = nil
		return
	} else if len(adminData.Password) == 0 {
		adminLoginResult.Meta.Status = "Fail"
		adminLoginResult.Meta.Code = 404
		adminLoginResult.Meta.Message = "password is empty"
		adminLoginResult.Meta.Records = 0
		adminLoginResult.Data = nil
		return
	}
	useCase.admRepository.AdminLogin(adminLoginResult, adminData)
	if adminLoginResult.Data != nil {
		adminResultData := adminLoginResult.Data.(AdmModel)
		if adminData.Username != adminResultData.Username || adminData.Password != adminResultData.Password {
			adminLoginResult.Meta.Status = "Fail"
			adminLoginResult.Meta.Code = 404
			adminLoginResult.Meta.Message = "invalid username or password"
			adminLoginResult.Meta.Records = 0
			adminLoginResult.Data = nil
			return
		}
		tokenExpiredTime := int64(constants.TOKEN_EXPIRE)
		tokenKey := helper.GenerateToken(adminResultData.Username, tokenExpiredTime)
		adminToken := models.TokenSecurityModel{
			Username: adminResultData.Username,
			Token:    tokenKey,
		}
		adminLoginResult.Meta.Status = "Success"
		adminLoginResult.Meta.Code = 202
		adminLoginResult.Meta.Message = "Admin success login"
		adminLoginResult.Meta.Records = 1
		adminLoginResult.Data = adminToken
		return
	}
}
