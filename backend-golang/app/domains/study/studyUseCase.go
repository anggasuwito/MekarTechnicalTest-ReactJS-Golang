package study

import (
	"database/sql"
	"mekarTechnicalTest/app/models"
)

//StdUseCase is a struct for depedency injection
type StdUseCase struct {
	db              *sql.DB
	studyRepository StdRepositoryInterface
}

//StdUseCaseInterface is an interface from study use case
type StdUseCaseInterface interface {
	Studies(studyResult *models.Response)
}

//InitStdUseCase is a function for inject study use case
func InitStdUseCase(db *sql.DB) *StdUseCase {
	return &StdUseCase{db: db, studyRepository: InitStdRepository(db)}
}

//Studies in study use case is a function for check some validation or put some logic
func (useCase StdUseCase) Studies(studyResult *models.Response) {
	useCase.studyRepository.Studies(studyResult)
}
