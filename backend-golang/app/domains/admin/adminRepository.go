package admin

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"
)

//AdmRepository is a struct for depedency injection
type AdmRepository struct {
	db *sql.DB
}

//AdmRepositoryInterface is an interface from admin repository
type AdmRepositoryInterface interface {
	AdminLogin(adminLoginResult *models.Response, adminData *AdmModel)
}

//InitAdmRepository is a function for inject admin repository database
func InitAdmRepository(db *sql.DB) *AdmRepository {
	return &AdmRepository{db: db}
}

//AdminLogin is a function for admin to log in
func (repository AdmRepository) AdminLogin(adminLoginResult *models.Response, adminData *AdmModel) {
	var admin AdmModel
	repository.db.QueryRow(constants.CHECK_ADMIN_USERNAME, adminData.Username, "A").Scan(&admin.AdminID, &admin.Username, &admin.Password)
	if len(admin.Username) == 0 {
		adminLoginResult.Meta.Status = "Fail"
		adminLoginResult.Meta.Code = 404
		adminLoginResult.Meta.Message = "Username not found"
		adminLoginResult.Meta.Records = 0
		adminLoginResult.Data = nil
		return
	}
	adminLoginResult.Meta.Status = "Success"
	adminLoginResult.Meta.Code = 202
	adminLoginResult.Meta.Message = "Admin success login"
	adminLoginResult.Meta.Records = 1
	adminLoginResult.Data = admin
}
