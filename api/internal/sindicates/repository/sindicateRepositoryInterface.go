package repository


type SindicateRepositoryInterface interface {
	GetAllSindicates() ([]models.GetAllSindicates, error)
	GetSindicateByID(id int) (models.SindicateModel, error)
	CreateSindicate(input models.CreateSindicateReq) (models.SindicateModel, error)
	CreateSindicateAuthorization(input models.CreateSindicateAuthorization) (models.SindicateAuthorization, error)
	CreateSindicateBreakfest(input models.CreateSindicateBreakfest) (models.SindicateBreakfest, error)
	CreateSindicateCarePackage(input models.CreateSindicateCarePackage) (models.SindicateCarePackage, error)
	CreateSindicatePercents(input models.CreateSindicatePercents) (models.SindicatePercents, error)
	CreateSindicateValues(input models.CreateSindicateValues) (models.SindicateValues, error)
	CreateSindicateVoucher(input models.CreateSindicateVoucher) (models.SindicateVoucher, error)
	EditSindicate(id int, input models.EditSindicate) (models.SindicateModel, error)
	EditSindicateAuthorization(id int, input models.EditSindicateAuthorization) (models.SindicateAuthorization, error)
	EditSindicateBreakfest(id int, input models.EditSindicateBreakfest) (models.SindicateBreakfest, error)
	EditSindicateCarePackage(id int, input models.EditSindicateCarePackage) (models.SindicateCarePackage, error)
	EditSindicatePercents(id int, input models.EditSindicatePercents) (models.SindicatePercents, error)
	EditSindicateValues(id int, input models.EditSindicateValues) (models.SindicateValues, error)
	EditSindicateVaucher(id int, input models.EditSindicateVaucher) (models.SindicateVoucher, error)
	DeactivateSindicate(id int) (models.SindicateModel, error)
	ReactivateSindicate(id int) (models.SindicateModel, error)
}