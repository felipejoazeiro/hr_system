package repository

type ContractsRepository struct {
	db *sql.DB
}

func NewContractsRepository(db *sql.DB) *ContractsRepository {
	return &ContractsRepository{db:db}
}

func (r *ContractsRepository) GetAllContracts() ([]models.GetAllContracts, error) {
	query := `SELECT c.id, c.name, c.code, d.date_initial FROM contracts c LEFT JOIN contracts_date d ON c.fk_contract_dates = d.id;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	} 
	defer rows.Close()

	var contracts []models.GetAllContracts

	for rows.Next(){
		var contract models.GetAllContracts
		err := rows.Scan(
			&contract.ID,
			&contract.Name,
			&contract.Code,
			&contract.InitialDate
		)

		if err != nil {
			return nil, err
		}

		contracts = append(contracts, contract)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return contracts, nil
}

func (r *ContractsRepository) CreateContract(input CreateContract) (Contract, error){
	query := `INSERT INTO contracts (name, code, research, uses_cpf, is_active, fk_contract_info, fk_contract_dates, fk_contract_values, fk_contract_discount, fk_contract_rh, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, TRUE) RETURNING id,name, code, research, uses_cpf, is_active, fk_contract_contact, fk_contract_info, fk_contract_dates, fk_contract_values, fk_contract_discount, fk_contract_rh, is_active;`

	var contract models.Contract

	err := r.db.QueryRow(query, input.name, input.code, input.research, input.uses_cpf, input.is_active, input.fk_contract_info, input.fk_contract_dates, input.fk_contract_values, input.fk_contract_discount, input.fk_contract_rh, is_active).Scan(
		&contract.ID,
		&contract.Name,
		&contract.Code,
		&contract.Research,
		&contract.UsesCpf,
		&contract.IsActive,
		&contract.FkContractDates,
		&contract.FkContractDiscount,
		&contract.FkContractInfo,
		&contract.FkContractRh,
		&contract.FkContractValues,
	)
	return contract, err
} 

func (r *ContractsRepository) createContractInfo(tx *sql.Tx, input CreateContractInfo) (int, error) {
	query := `
		INSERT INTO contract_info (
			construction, cap, process, info_pmco, max_employee, address, nro, complement, phone, state, city, cep, email, contact, fk_employee
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		)
		RETURNING id;
	`

	var infoId int

	err := tx.QueryRow(
		query,
		input.Construction,
		input.CAP,
		input.Process,
		input.InfoPmco,
		input.MaxEmployee,
		input.Address,
		input.NRO,
		input.Complement,
		input.Phone,
		input.State,
		input.City,
		input.CEP,
		input.Email,
		input.Contact,
		input.FkEmployee,
	).Scan(&infoId)

	if err != nil {
		return 0, err
	}
	return infoId, nil
}

func (r *ContractsRepository) createContractDates(tx *sql.Tx, input CreateContractDates) (int, error) {
	query := `
		INSERT INTO contract_dates (
			date_initial,
			date_limit,
			date_guarantee,
			date_proposal,
			date_budget,
			date_tables
		) VALUES (
			$1, $2, $3, $4, $5, $6
		) RETURNING id
	`

	var datesId int

	err := tx.QueryRow(
		query,
		input.DateInitial,
		input.DateLimit,
		input.DateGuarantee,
		input.DateProposal,
		input.DateBudget,
		input.DateTables,
	).Scan(&datesId)

	if err != nil {
		return 0, err
	}

	return datesId, nil
}