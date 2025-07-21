package repository


type SindicateRepositoryInterface interface {
	GetAllSindicates() ([]models.GetAllSindicates, error)
	CreateSindicate(input models.CreateSindicateReq) (models.SindicateModel, error)
	EditSindicateAuthorization(id int, input models.EditSindicateAuthorization) (models.SindicateAuthorization, error)
	EditSindicateBreakfest(id int, input models.EditSindicateBreakfest) (models.SindicateBreakfest, error)
	EditSindicateCarePackage(id int, input models.EditSindicateCarePackage) (models.SindicateCarePackage, error)
	EditSindicatePercents(id int, input models.EditSindicatePercents) (models.SindicatePercents, error)
	EditSindicateValues(id int, input models.EditSindicateValues) (models.SindicateValues, error)
	EditSindicateVaucher(id int, input models.EditSindicateVaucher) (models.SindicateVoucher, error)
	DeactivateSindicate(id int) (models.SindicateModel, error)
	ReactivateSindicate(id int) (models.SindicateModel, error)
}