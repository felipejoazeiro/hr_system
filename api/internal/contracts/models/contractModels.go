package models

type GetAllContracts struct {
	ID				int		`json:"id" db:"id"`
	Name			string	`json:"name" db:"name"`
	Code			string	`json:"code" db:"code"`
	InitialDate		time	`json:"date_initial" db:"date_initial"`	
}

type Contract struct {
	ID		int			`json:"id" db:"id"`
	Name				string		`json:"name" db:"name"`
	Code				string		`json:"code" db:"code"`
	Research			bool		`json:"research" db:"research"`
	UsesCpf				bool		`json:"uses_cpf" db:"uses_cpf"`
	IsActive 			bool		`json:"is_active" db:"is_active"`
	FkContractContact	int			`json:"fk_contract_contact" db:"fk_contract_contact"`
	FkContractInfo		int			`json:"fk_contract_info" db:"fk_contract_info"`
	FkContractDates		int			`json:"fk_contract_dates" db:"fk_contract_dates"`
	FkContractValues	int			`json:"fk_contract_values" db:"fk_contract_values"`
	FkContractDiscount	int			`json:"fk_contract_discount" db:"fk_contract_discount"`
	FkContractRh		int			`json:"fk_contract_rh" db:"fk_contract_rh"`
}

type ContractModel struct {
	ID		int			`json:"id" db:"id"`
	Name	string		`json:"name" db:"name"`
	Code	string		`json:"code" db:"code"`
}

type GetContract struct {
	ID				int			`json:"id" db:"id"`
	Name			string		`json:"name" db:"name"`
	Code			string		`json:"code" db:"code"`
	Employee		string		`json:"employee" db:"employee"`
	Research		bool		`json"research" db:"research"`
	IsActive		bool		`json:"is_active" db:"is_active"`
	Cpf 			string		`json:"cpf" db:"cpf"`
	Address			string		`json:"address" db:"address"`
	NRO				string		`json:"nro" db:"nro"`
	State			string		`json:"state" db:"state"`
	City 			string 		`json:"city" db:"city"`	
	InitialDate		string		`json:"date_initial" db:"date_initial"`
	LimitDate		string		`json:"limit_date" db:"limit_date"`
	Email			string		`json:"email" db:"email"`
	Contact			string		`json:"contact" db:"contact"`
	Cellphone		string		`json:"cellphone" db:"cellphone"`
}

type CreateContract struct {
	Name				string		`json:"name" db:"name"`
	Code				string		`json:"code" db:"code"`
	Research			bool		`json:"research" db:"research"`
	UsesCpf				bool		`json:"uses_cpf" db:"uses_cpf"`
	IsActive 			bool		`json:"is_active" db:"is_active"`
	FkContractInfo		int			`json:"fk_contract_info" db:"fk_contract_info"`
	FkContractDates		int			`json:"fk_contract_dates" db:"fk_contract_dates"`
	FkContractValues	int			`json:"fk_contract_values" db:"fk_contract_values"`
	FkContractDiscount	int			`json:"fk_contract_discount" db:"fk_contract_discount"`
	FkContractRh		int			`json:"fk_contract_rh" db:"fk_contract_rh"`
}

type EditContract struct {
	Name				*string		`json:"name" db:"name"`
	Code				*string		`json:"code" db:"code"`
	Research			*bool		`json:"research" db:"research"`
	UsesCpf				*bool		`json:"uses_cpf" db:"uses_cpf"`
	IsActive 			*bool		`json:"is_active" db:"is_active"`
	FkContractContact	*int			`json:"fk_contract_contact" db:"fk_contract_contact"`
	FkContractInfo		*int			`json:"fk_contract_info" db:"fk_contract_info"`
	FkContractDates		*int			`json:"fk_contract_dates" db:"fk_contract_dates"`
	FkContractValues	*int			`json:"fk_contract_values" db:"fk_contract_values"`
	FkContractDiscount	*int			`json:"fk_contract_discount" db:"fk_contract_discount"`
	FkContractRh		*int			`json:"fk_contract_rh" db:"fk_contract_rh"`
}

type ContractDatesModel struct {
	ID				int			`json:"id" db:"id"`
	DateInitial		date 		`json:"date_initial" db:"date_initial"`
	DateLimit		date		`json:"date_limit" db:"date_limit"`
	DateGuarantee	date		`json:"date_guarantee" db:"date_guarantee"`
	DateProposal	date		`json:"date_proposal" db:"date_proposal"`
	DateBudget		date		`json:"date_budget" db:"date_budget"`
	DateTables		date		`json:"date_tables" db:"date_tables"`
}

type CreateContractDates struct{
	DateInitial		date 		`json:"date_initial" db:"date_initial"`
	DateLimit		date		`json:"date_limit" db:"date_limit"`
	DateGuarantee	date		`json:"date_guarantee" db:"date_guarantee"`
	DateProposal	date		`json:"date_proposal" db:"date_proposal"`
	DateBudget		date		`json:"date_budget" db:"date_budget"`
	DateTables		date		`json:"date_tables" db:"date_tables"`
}

type EditContractDates struct {
	DateInitial		*date 		`json:"date_initial" db:"date_initial"`
	DateLimit		*date		`json:"date_limit" db:"date_limit"`
	DateGuarantee	*date		`json:"date_guarantee" db:"date_guarantee"`
	DateProposal	*date		`json:"date_proposal" db:"date_proposal"`
	DateBudget		*date		`json:"date_budget" db:"date_budget"`
	DateTables		*date		`json:"date_tables" db:"date_tables"`
}

type ContractRhInfoModel struct {
	ID					int			`json:"id" db:"id"`
	HourLimit			date		`json:"hour_limit" db:"hour_limit"`
	MinutesLimit		date		`json:"minutes_limit" db:"minutes_limit"`
	DaysFirstExp		int			`json:"days_first_exp" db:"days_first_exp"`
	DaysSecondExp		int			`json:"days_second_exp" db:"days_second_exp"`
	DataInit			date		`json:"data_init" db:"data_init"`
	PayExtraHour		bool		`json:"pay_extra_hour" db:"pay_extra_hour"`
	ManualStitch		bool		`json:"manual_stitch" db:"manual_stitch"`
	PaysBreakfast		bool		`json:"pays_breakfast" db:"pays_breakfast"`
}

type CreateContractRhInfo struct {
	HourLimit			date		`json:"hour_limit" db:"hour_limit"`
	MinutesLimit		date		`json:"minutes_limit" db:"minutes_limit"`
	DaysFirstExp		int			`json:"days_first_exp" db:"days_first_exp"`
	DaysSecondExp		int			`json:"days_second_exp" db:"days_second_exp"`
	DataInit			date		`json:"data_init" db:"data_init"`
	PayExtraHour		bool		`json:"pay_extra_hour" db:"pay_extra_hour"`
	ManualStitch		bool		`json:"manual_stitch" db:"manual_stitch"`
	PaysBreakfast		bool		`json:"pays_breakfast" db:"pays_breakfast"`
}

type EditContractRhInfo	struct {
	HourLimit			*date		`json:"hour_limit" db:"hour_limit"`
	MinutesLimit		*date		`json:"minutes_limit" db:"minutes_limit"`
	DaysFirstExp		*int		`json:"days_first_exp" db:"days_first_exp"`
	DaysSecondExp		*int		`json:"days_second_exp" db:"days_second_exp"`
	DataInit			*date		`json:"data_init" db:"data_init"`
	PayExtraHour		*bool		`json:"pay_extra_hour" db:"pay_extra_hour"`
	ManualStitch		*bool		`json:"manual_stitch" db:"manual_stitch"`
	PaysBreakfast		*bool		`json:"pays_breakfast" db:"pays_breakfast"`
}

type ContractDiscountModel struct {
	ID					int 		`json:"id" db:"id"`
	DiscIdentifier		int			`json:"disc_identifier" db:"dics_identifier"`
	DiscService			int			`json:"disc_service" db:"disc_service"`
	DiscTransport		int			`json:"disc_transport" db:"disc_transport"`
	DiscTranpEmployee	int			`json:"disc_tranp_employee" db:"disc_tranp_employee"`
	DiscLabor			int			`json:"disc_labor" db:"disc_labor"`
	DiscMaterial		int			`json:"disc_material" db:"disc_material"`
	DiscField			int			`json:"disc_field" db:"disc_field"`
}

type CreateContractDiscount struct {
	DiscIdentifier		int			`json:"disc_identifier" db:"dics_identifier"`
	DiscService			int			`json:"disc_service" db:"disc_service"`
	DiscTransport		int			`json:"disc_transport" db:"disc_transport"`
	DiscTranpEmployee	int			`json:"disc_tranp_employee" db:"disc_tranp_employee"`
	DiscLabor			int			`json:"disc_labor" db:"disc_labor"`
	DiscMaterial		int			`json:"disc_material" db:"disc_material"`
	DiscField			int			`json:"disc_field" db:"disc_field"`
}

type EditContractDiscount struct {
	DiscIdentifier		*int		`json:"disc_identifier" db:"dics_identifier"`
	DiscService			*int		`json:"disc_service" db:"disc_service"`
	DiscTransport		*int		`json:"disc_transport" db:"disc_transport"`
	DiscTranpEmployee	*int		`json:"disc_tranp_employee" db:"disc_tranp_employee"`
	DiscLabor			*int		`json:"disc_labor" db:"disc_labor"`
	DiscMaterial		*int		`json:"disc_material" db:"disc_material"`
	DiscField			*int		`json:"disc_field" db:"disc_field"`
}

type ContractInfoModel struct {
	ID				int			`json:"id" db:"id"`
	Construction	string		`json:"construction" db:"construction"`
	CAP 			string		`json:"cap" db:"cap"`
	Process			string		`json:"process" db:"process"`
	InfoPmco		string		`json:"info_pmco" db:"info_pmco"`
	MaxEmployee		int			`json:"max_employee" db:"max_employee"`
	Address			string		`json:"address" db:"address"`
	NRO				int			`json:"nro" db:"nro"`
	Complement		string		`json:"complement" db:"complement"`
	Phone			int			`json:"phone" db:"phone"`
	State			string		`json:"state" db:"state"`
	City			string		`json:"city" db:"city"`
	CEP 			string		`json:"cep" db:"cep"`
	Email			string		`json:"email" db:"email"`
	Contact			string		`json:"contact" db:"contact"`
	FkEmployee		int			`json:"fk_employee" db:"fk_employee"`
}

type CreateContractInfo struct {
	Construction	string		`json:"construction" db:"construction"`
	CAP 			string		`json:"cap" db:"cap"`
	Process			string		`json:"process" db:"process"`
	InfoPmco		string		`json:"info_pmco" db:"info_pmco"`
	MaxEmployee		int			`json:"max_employee" db:"max_employee"`
	Address			string		`json:"address" db:"address"`
	NRO				int			`json:"nro" db:"nro"`
	Complement		string		`json:"complement" db:"complement"`
	Phone			int			`json:"phone" db:"phone"`
	State			string		`json:"state" db:"state"`
	City			string		`json:"city" db:"city"`
	CEP 			string		`json:"cep" db:"cep"`
	Email			string		`json:"email" db:"email"`
	Contact			string		`json:"contact" db:"contact"`
	FkEmployee		int			`json:"fk_employee" db:"fk_employee"`
}

type EditContractInfo struct {
	Construction	*string		`json:"construction" db:"construction"`
	CAP 			*string		`json:"cap" db:"cap"`
	Process			*string		`json:"process" db:"process"`
	InfoPmco		*string		`json:"info_pmco" db:"info_pmco"`
	MaxEmployee		*int			`json:"max_employee" db:"max_employee"`
	Address			*string		`json:"address" db:"address"`
	NRO				*int		`json:"nro" db:"nro"`
	Complement		*string		`json:"complement" db:"complement"`
	Phone			*int		`json:"phone" db:"phone"`
	State			*string		`json:"state" db:"state"`
	City			*string		`json:"city" db:"city"`
	CEP 			*string		`json:"cep" db:"cep"`
	Email			*string		`json:"email" db:"email"`
	Contact			*string		`json:"contact" db:"contact"`
	FkEmployee		*int		`json:"fk_employee" db:"fk_employee"`
}

type ContractValuesModel struct {
	ID					int				`json:"id" db:"id"`
	AcronymValues		string			`json:"acronym_values" db:"acronym_values"`
	BdiService			int				`json:"bdi_service" db:"bdi_service"`
	BdiMaterial			int				`json:"bdi_material" db:"bdi_material"`
	BdiLabor			int				`json:"bdi_labor" db:"bid_labor"`
	EntryTable			bool 			`json:"entry_table" db:"entry_table"`
	SendEmail			bool			`json:"send_email" db:"send_email"`
}

type CreateContractValues struct {
	AcronymValues		string			`json:"acronym_values" db:"acronym_values"`
	BdiService			int				`json:"bdi_service" db:"bdi_service"`
	BdiMaterial			int				`json:"bdi_material" db:"bdi_material"`
	BdiLabor			int				`json:"bdi_labor" db:"bid_labor"`
	EntryTable			bool 			`json:"entry_table" db:"entry_table"`
	SendEmail			bool			`json:"send_email" db:"send_email"`
}

type EditContractValues struct {
	AcronymValues		*string			`json:"acronym_values" db:"acronym_values"`
	BdiService			*int			`json:"bdi_service" db:"bdi_service"`
	BdiMaterial			*int			`json:"bdi_material" db:"bdi_material"`
	BdiLabor			*int			`json:"bdi_labor" db:"bid_labor"`
	EntryTable			*bool 			`json:"entry_table" db:"entry_table"`
	SendEmail			*bool			`json:"send_email" db:"send_email"`
}

type CreateFullContractRequest struct {
	Name 			string					`json:"name"`
	Code			string					`json:"code"`
	Research		bool					`json:"research"`
	UsesCpf			bool					`json:"uses_cpf"`
	IsActive		bool					`json:"is_active"`

	Info 			CreateContractInfo		`json:"info"`
	Dates 			CreateContractDates		`json:"dates"`
	Values 			CreateContractValues	`json:"values"`
	Discount		CreateContractDiscount	`json:"discount"`
	RhInfo 			CreateContractRhInfo	`json:"rh_info"`
}