package user

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
	"mekarTechnicalTest/utils/constants"
	"mekarTechnicalTest/utils/helper"
)

//UsrRepository is a struct for depedency injection
type UsrRepository struct {
	db *sql.DB
}

//UsrRepositoryInterface is an interface from user repository
type UsrRepositoryInterface interface {
	Users(userResult *models.Response)
	UserByID(userByIDResult *models.Response, userID string)
	CreateUser(createUserResult *models.Response, userDataBody *UsrBodyModel)
	UpdateUserByID(updateUserResult *models.Response, userDataBody *UsrBodyModel, userID string)
	DeleteUserByID(deleteUserResult *models.Response, userID string)
}

//InitUsrRepository is a function for inject user repository database
func InitUsrRepository(db *sql.DB) *UsrRepository {
	return &UsrRepository{db: db}
}

//Users is a function for get all users from database
func (repository UsrRepository) Users(userResult *models.Response) {
	var users []*UsrModel
	data, err := repository.db.Query(constants.GET_ALL_USERS, "A")
	getUserErr := helper.ErrorNotNil(err)
	if getUserErr {
		userResult.Meta.Status = "Fail"
		userResult.Meta.Code = 404
		userResult.Meta.Message = "get all users error"
		userResult.Meta.Records = 0
		userResult.Data = nil
		return
	}

	for data.Next() {
		user := UsrModel{}
		errUser := data.Scan(&user.UserID, &user.Name, &user.BirthDate, &user.NumberIDCard, &user.Work.WorkID, &user.Work.Name, &user.Study.StudyID, &user.Study.Name)
		loopingUserErr := helper.ErrorNotNil(errUser)
		if loopingUserErr {
			userResult.Meta.Status = "Fail"
			userResult.Meta.Code = 404
			userResult.Meta.Message = "looping user error"
			userResult.Meta.Records = 0
			userResult.Data = nil
			return
		}
		users = append(users, &user)
	}
	countUsers := len(users)
	userResult.Meta.Status = "Success"
	userResult.Meta.Code = 202
	userResult.Meta.Message = "Success get all users"
	userResult.Meta.Records = countUsers
	userResult.Data = users
}

//UserByID is a function for get user by id from database
func (repository UsrRepository) UserByID(userByIDResult *models.Response, userID string) {
	var user UsrModel
	repository.db.QueryRow(constants.GET_USER_BY_ID, userID, "A").Scan(&user.UserID, &user.Name, &user.BirthDate, &user.NumberIDCard, &user.Work.WorkID, &user.Work.Name, &user.Study.StudyID, &user.Study.Name)
	if user.Name == "" {
		userByIDResult.Meta.Status = "Fail"
		userByIDResult.Meta.Code = 404
		userByIDResult.Meta.Message = "user not found"
		userByIDResult.Meta.Records = 0
		userByIDResult.Data = nil
		return
	}
	userByIDResult.Meta.Status = "Success"
	userByIDResult.Meta.Code = 202
	userByIDResult.Meta.Message = "Success get user by id"
	userByIDResult.Meta.Records = 1
	userByIDResult.Data = user
}

//CreateUser is a function for insert new user to database
func (repository UsrRepository) CreateUser(createUserResult *models.Response, userDataBody *UsrBodyModel) {
	createNewUser, createNewUserErr := repository.db.Prepare(constants.CREATE_NEW_USER)
	newUserErr := helper.ErrorNotNil(createNewUserErr)
	if newUserErr {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "prepare create new user query error"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	}
	_, execCreateUserErr := createNewUser.Exec(userDataBody.Name, userDataBody.BirthDate, userDataBody.NumberIDCard, userDataBody.WorkID, userDataBody.StudyID)
	execnewUserErr := helper.ErrorNotNil(execCreateUserErr)
	if execnewUserErr {
		createUserResult.Meta.Status = "Fail"
		createUserResult.Meta.Code = 404
		createUserResult.Meta.Message = "exec create new user query error"
		createUserResult.Meta.Records = 0
		createUserResult.Data = nil
		return
	}
	createUserResult.Meta.Status = "Success"
	createUserResult.Meta.Code = 202
	createUserResult.Meta.Message = "Success create new user"
	createUserResult.Meta.Records = 0
	createUserResult.Data = nil
}

//UpdateUserByID is a function for update user data from database
func (repository UsrRepository) UpdateUserByID(updateUserResult *models.Response, userDataBody *UsrBodyModel, userID string) {
	var user UsrModel
	repository.db.QueryRow(constants.GET_USER_BY_ID, userID, "A").Scan(&user.UserID, &user.Name, &user.BirthDate, &user.NumberIDCard, &user.Work.WorkID, &user.Work.Name, &user.Study.StudyID, &user.Study.Name)
	if user.Name == "" {
		updateUserResult.Meta.Status = "Fail"
		updateUserResult.Meta.Code = 404
		updateUserResult.Meta.Message = "user id not found"
		updateUserResult.Meta.Records = 0
		updateUserResult.Data = nil
		return
	}

	updateUser, updateUserErr := repository.db.Prepare(constants.UPDATE_USER)
	prepUpdateUserErr := helper.ErrorNotNil(updateUserErr)
	if prepUpdateUserErr {
		updateUserResult.Meta.Status = "Fail"
		updateUserResult.Meta.Code = 404
		updateUserResult.Meta.Message = "prepare update user query error"
		updateUserResult.Meta.Records = 0
		updateUserResult.Data = nil
		return
	}

	_, execUpdateUserBodyErr := updateUser.Exec(userDataBody.Name, userDataBody.BirthDate, userDataBody.NumberIDCard, userDataBody.WorkID, userDataBody.StudyID, userID, "A")
	execUpdateUserErr := helper.ErrorNotNil(execUpdateUserBodyErr)
	if execUpdateUserErr {
		updateUserResult.Meta.Status = "Fail"
		updateUserResult.Meta.Code = 404
		updateUserResult.Meta.Message = "exec update user query error"
		updateUserResult.Meta.Records = 0
		updateUserResult.Data = nil
		return
	}
	updateUserResult.Meta.Status = "Success"
	updateUserResult.Meta.Code = 202
	updateUserResult.Meta.Message = "Success update user"
	updateUserResult.Meta.Records = 0
	updateUserResult.Data = nil
}

//DeleteUserByID is a function for soft delete user by id from database
func (repository UsrRepository) DeleteUserByID(deleteUserResult *models.Response, userID string) {
	var user UsrModel
	repository.db.QueryRow(constants.GET_USER_BY_ID, userID, "A").Scan(&user.UserID, &user.Name, &user.BirthDate, &user.NumberIDCard, &user.Work.WorkID, &user.Work.Name, &user.Study.StudyID, &user.Study.Name)
	if user.Name == "" {
		deleteUserResult.Meta.Status = "Fail"
		deleteUserResult.Meta.Code = 404
		deleteUserResult.Meta.Message = "user id not found"
		deleteUserResult.Meta.Records = 0
		deleteUserResult.Data = nil
		return
	}

	deleteUser, deleteUserErr := repository.db.Prepare(constants.DELETE_USER)
	prepDeleteUserErr := helper.ErrorNotNil(deleteUserErr)
	if prepDeleteUserErr {
		deleteUserResult.Meta.Status = "Fail"
		deleteUserResult.Meta.Code = 404
		deleteUserResult.Meta.Message = "prepare delete user query error"
		deleteUserResult.Meta.Records = 0
		deleteUserResult.Data = nil
		return
	}

	_, execDeleteUserDataErr := deleteUser.Exec("I", userID, "A")
	execDeleteUserErr := helper.ErrorNotNil(execDeleteUserDataErr)
	if execDeleteUserErr {
		deleteUserResult.Meta.Status = "Fail"
		deleteUserResult.Meta.Code = 404
		deleteUserResult.Meta.Message = "exec delete user query error"
		deleteUserResult.Meta.Records = 0
		deleteUserResult.Data = nil
		return
	}
	deleteUserResult.Meta.Status = "Success"
	deleteUserResult.Meta.Code = 202
	deleteUserResult.Meta.Message = "Success delete user"
	deleteUserResult.Meta.Records = 0
	deleteUserResult.Data = nil
}
