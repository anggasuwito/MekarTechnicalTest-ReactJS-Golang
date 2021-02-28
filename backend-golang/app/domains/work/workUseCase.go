package work

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
)

//WrkUseCase is a struct for depedency injection
type WrkUseCase struct {
	db             *sql.DB
	workRepository WrkRepositoryInterface
}

//WrkUseCaseInterface is an interface from work use case
type WrkUseCaseInterface interface {
	Works(workResult *models.Response)
}

//InitWrkUseCase is a function for inject work use case
func InitWrkUseCase(db *sql.DB) *WrkUseCase {
	return &WrkUseCase{db: db, workRepository: InitWrkRepository(db)}
}

//Works in work use case is a function for check some validation or put some logic
func (useCase WrkUseCase) Works(workResult *models.Response) {
	useCase.workRepository.Works(workResult)
}
