package repository


type SindicateRepositoryInterface interface {
	GetAllSindicates() ([]models.GetAllSindicates, error)
	CreateSindicate(input models.CreateSindicate) (models.SindicateModel, error)
}