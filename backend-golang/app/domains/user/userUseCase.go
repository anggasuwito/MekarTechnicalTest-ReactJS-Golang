package user

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
)

//UsrUseCase is a struct for depedency injection
type UsrUseCase struct {
	db             *sql.DB
	userRepository UsrRepositoryInterface
}

//UsrUseCaseInterface is an interface from user use case
type UsrUseCaseInterface interface {
	Users(userResult *models.Response)
	UserByID(userByIDResult *models.Response, userID string)
	CreateUser(createUserResult *models.Response, userDataBody *UsrBodyModel)
}

//InitUsrUseCase is a function for inject user use case
func InitUsrUseCase(db *sql.DB) *UsrUseCase {
	return &UsrUseCase{db: db, userRepository: InitUsrRepository(db)}
}

//Users in user use case is a function for check some validation or put some logic
func (useCase UsrUseCase) Users(userResult *models.Response) {
	useCase.userRepository.Users(userResult)
}

//UserByID in user use case is a function for check some validation or put some logic
func (useCase UsrUseCase) UserByID(userByIDResult *models.Response, userID string) {
	useCase.userRepository.UserByID(userByIDResult, userID)
}

//CreateUser in user use case is a function for check some validation or put some logic
func (useCase UsrUseCase) CreateUser(createUserResult *models.Response, userDataBody *UsrBodyModel) {
	if len(userDataBody.Name) == 0 {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "user name is empty"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	} else if len(userDataBody.BirthDate) == 0 {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "user birth date is empty"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	} else if len(userDataBody.NumberIDCard) == 0 {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "user number id card is empty"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	} else if userDataBody.WorkID == 0 {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "user work id is empty"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	} else if userDataBody.StudyID == 0 {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "user study id is empty"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	}
	useCase.userRepository.CreateUser(createUserResult, userDataBody)
}
