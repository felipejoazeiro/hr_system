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

func (r *ContractsRepository) CreateContract(tx *sql.Tx, input CreateContract) (Contract, error){
	query := `INSERT INTO contracts (name, code, research, uses_cpf, is_active, fk_contract_info, fk_contract_dates, fk_contract_values, fk_contract_discount, fk_contract_rh, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, TRUE) RETURNING id;`

	var contractID int

	err := tx.QueryRow(
		query, 
		input.name, 
		input.code, 
		input.research, 
		input.uses_cpf, 
		input.is_active, 
		input.fk_contract_info, 
		input.fk_contract_dates, 
		input.fk_contract_values, 
		input.fk_contract_discount, 
		input.fk_contract_rh, 
	).Scan(
		&contractID
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return contractID, err
} 

func (r *ContractsRepository) CreateContractInfo(tx *sql.Tx, input CreateContractInfo) (int, error) {
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
		tx.Rollback()
		return 0, err
	}
	return infoId, nil
}

func (r *ContractsRepository) CreateContractDates(tx *sql.Tx, input CreateContractDates) (int, error) {
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
		tx.Rollback()
		return 0, err
	}

	return datesId, nil
}

func (r *ContractsRepository) CreateContractValues(tx *sql.Tx, input CreateContractValues) (int, error) {
	query := `
		INSERT INTO contract_values (
			acronym_value, 
			bdi_service,
			bdi_material,
			bdi_labor,
			entry_table,
			send_email
		) VALUES (
			$1, $2, $3, $4, $5, $6 
		) RETURNING id
	`
	var valuesId int

	err  := tx.QueryRow(query,
		input.AcronymValues,
		input.BdiService,
		input.BdiMaterial,
		input.BdiLabor,
		input.EntryTable,
		input.Email,
	).Scan(&valuesId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return valuesId, nil
}

func (r *ContractsRepository) CreateContractDiscount(tx *sql.Tx, input CreateContractDiscount) (int, error) {
	query := `
		INSERT INTO contract_discount (
			disc_identifier,
			disc_service,
			disc_transport,
			disc_tranp_employee,
			disc_labor,
			disc_material,
			disc_field
		) VALUES (
		 $1, $2, $3, $4, $5, $6, $7 
		) RETURNING id
	`

	var discountId int

	err := tx.QueryRow(
		query,
		input.DiscIdentifier,
		input.DiscService,
		input.DiscTransport,
		input.DiscTranpEmployee,
		input.DiscLabor,
		input.DiscMaterial,
		input.DiscField,
	).Scan(&discountId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return discountId, nil
}

func (r *ContractsRepository) CreateContractRhInfo(tx *sql.Tx, input CreateContractRhInfo) (int, error) {
	query := `
		INSERT INTO contract_rh (
			hour_limit,
			minutes_limit,
			days_first_exp,
			days_second_exp,
			data_init,
			pay_extra_hour,
			manual_stitch,
			pays_breakfast
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8 
		) RETURNING id
	`

	var rhId int
	err := tx.QueryRow(
		query,
		input.HourLimit,
		input.MinutesLimit,
		input.DaysFirstExp,
		input.DaysSecondExp,
		input.DataInit,
		input.PayExtraHour,
		input.ManualStitch,
		input.PaysBreakfast,
	).Scan(&rhId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return rhId, nil
}

func(r *ContractsRepository) EditContract(id string, c models.EditContract) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.Name != nil {
		setParts = append(setParts, fmt.Sprintf("name=$%d", argPos))
		args = append(args, *c.Name)
		argPos++
	}

	if c.Code != nil {
		setParts = append(setParts, fmt.Sprintf("code=$%d", argPos))
		args = append(args, *c.Code)
		argPos++
	}

	if c.Research != nil {
		setParts = append(setParts, fmt.Sprintf("research=$%d", argPos))
		args = append(args, *c.Research)
		argPos++
	}

	if  c.UsesCpf != nil {
		setParts = append(setParts, fmt.Sprintf("uses_cpf=$%d", argPos))
		args = append(args, *c.UsesCpf)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.new("Missing fields for update")
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE contracts SET %s WHERE id=$%d RETURNING id;", strings.Join(setParts, ", "), argPos)

	var contractId int
	err := r.db.QueryRow(query, args...).Scan(
		&contractId
	)

	return contractId, err

}

func (r *ContractsRepository) EditContractValues(id string, c models.EditContractValues) (int, error){
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.AcronymValues != nil {
		setParts = append(setParts, fmt.Sprintf("acronym_values=$%d", argPos))
		args = append(args, *c.AcronymValues)
		argPos++
	}
	if c.BdiService != nil {
		setParts = append(setParts, fmt.Sprintf("bdi_service=$%d", argPos))
		args = append(args, c.BdiService)
		argPos++
	}
	if c.BdiMaterial != nil {
		setParts = append(setParts, fmt.Sprintf("bdi_material=$%d", argPos))
		args = append(args, c.BdiMaterial)
		argPos++
	}
	if c.BdiLabor != nil {
		setParts = append(setParts, fmt.Sprintf("bdi_labor=#%d", argPos))
		args = append(args, c.BdiLabor)
		argPos++
	}
	if c.EntryTable != nil {
		setParts = append(setParts, fmt.Sprintf("entry_table=$%d", argPos))
		args = append(args, c.EntryTable)
		argPos++
	}
	if c.SendEmail != nil {
		setParts = append(setParts, fmt.Sprintf("send_email=$%d", argPos))
		args = append(args, c.SendEmail)
		argPos++
	}

	if len(setParts) == 0 {
		return models.ContractValuesModel{}, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE contract_value SET %s WHERE id=$%d RETURNING id, acronym_values, bdi_service, bdi_material, bdi_labor, entry_table, send_email; ", strings.Join(setParts, ", "), argPos)
	var updated models.ContractValuesModel
	err := r.db.QueryRow(query, args...).Scan(
		&updated.ID,
		&updated.AcronymValues,
		&updated.BdiService,
		&updated.BdiMaterial,
		&updated.BdiLabor,
		&updated.EntryTable,
		&updated.SendEmail
	)

	return updated, err
}

func (r *ContractsRepository) EditContractInfo(id string, c models.EditContractInfo) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.Construction != nil {
		setParts = append(setParts, fmt.Sprintf("construction=$%d", argPos))
		args = append(args, *d.Construction)
		argPos++
	}

	if c.CAP != nil {
		setParts = append(setParts, fmt.Sprintf("cap=$%d", argPos))
		args = append(args, *d.CAP)
		argPos++
	}

	if d.Process != nil {
		setParts = append(setParts, fmt.Sprintf("process=$%d", argPos))
		args = append(args, *d.Process)
		argPos++
	}
	if d.InfoPmco != nil {
		setParts = append(setParts, fmt.Sprintf("info_pmco=$%d", argPos))
		args = append(args, *d.InfoPmco)
		argPos++
	}
	if d.MaxEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("max_employee=$%d", argPos))
		args = append(args, *d.MaxEmployee)
		argPos++
	}
	if d.Address != nil {
		setParts = append(setParts, fmt.Sprintf("address=$%d", argPos))
		args = append(args, *d.Address)
		argPos++
	}
	if d.NRO != nil {
		setParts = append(setParts, fmt.Sprintf("nro=$%d", argPos))
		args = append(args, *d.NRO)
		argPos++
	}
	if d.Complement != nil {
		setParts = append(setParts, fmt.Sprintf("complement=$%d", argPos))
		args = append(args, *d.Complement)
		argPos++
	}
	if d.Phone != nil {
		setParts = append(setParts, fmt.Sprintf("phone=$%d", argPos))
		args = append(args, *d.Phone)
		argPos++
	}
	if d.State != nil {
		setParts = append(setParts, fmt.Sprintf("state=$%d", argPos))
		args = append(args, *d.State)
		argPos++
	}
	if d.City != nil {
		setParts = append(setParts, fmt.Sprintf("city=$%d", argPos))
		args = append(args, *d.City)
		argPos++
	}
	if d.CEP != nil {
		setParts = append(setParts, fmt.Sprintf("cep=$%d", argPos))
		args = append(args, *d.CEP)
		argPos++
	}
	if d.Email != nil {
		setParts = append(setParts, fmt.Sprintf("email=$%d", argPos))
		args = append(args, *d.Email)
		argPos++
	}
	if d.Contact != nil {
		setParts = append(setParts, fmt.Sprintf("contact=$%d", argPos))
		args = append(args, *d.Contact)
		argPos++
	}
	if d.FkEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("fk_employee=$%d", argPos))
		args = append(args, *d.FkEmployee)
		argPos++
	} 

	if len(setParts) == 0 {
    	return 0, errors.New("nenhum campo para atualizar")
	}		

	args = append(args, id) 
	query := fmt.Sprintf("UPDATE contract_info SET %s WHERE id=$%d RETURNING id;", strings.Join(setParts, ", "), argPos)

	var infoId int
	err := r.db.QueryRow(query, args...).Scan(
		&infoId,
	)
	return infoId, err
}

func (r *ContractsRepository) EditContractDates(id string, c models.EditContractDates) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.DateInitial := nil {
		setParts = append(setParts, fmt.Sprintf("date_initial=$%d", argPos))
		args = append(args, *c.DateInitial)
		argPos++
	}

	if c.DateLimit := nil {
		setParts = append(setParts, fmt.Sprintf("date_limit=$%d", argPos))
		args = append(args, *c.DateLimit)
		argPos++
	}

	if c.DateGuarantee := nil {
		setParts = append(setParts, fmt.Sprintf("date_guarantee=$%d", argPos))
		args = append(argPos, *c.DateProposal)
		argPos++
	}

	if c.DateProposal := nil {
		setParts = append(setParts, fmt.Sprintf("date_proposal=%d", argPos))
		args = append(argPos, *c.DateProposal)
		argPos++
	}

	if c.DateBudget := nil {
		setParts = append(setParts, fmt.Sprintf("date_budget=$%d", argPos))
		args = append(argPos, *c.DateBudget)
		argPos++
	}

	if c.DateTables := nil {
		setParts = append(setParts, fmt.Sprintf("date_tables=$%d", argPos))
		args = append(argPos, *c.DateTables)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.New("nenhum campo para atualizar")
	}

	args = append(args,id)
	query = fmt.Sprintf("UPDATE contract_dates SET %s WHERE id=$%d RETURNING id;", strings.Join(setParts, ", "), argPos)

	var valuesId int

	err := r.db.QueryRow(query, args...).Scan(&valuesId)
	return valuesId, err
}

func (r *ContractsRepository) EditContractDiscount(id string, c models.EditContractDiscount) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if input.DiscIdentifier != nil {
		setParts = append(setParts, fmt.Sprintf("disc_identifier=$%d", argPos))
		args = append(args, *input.DiscIdentifier)
		argPos++
	}
	if input.DiscService != nil {
		setParts = append(setParts, fmt.Sprintf("disc_service=$%d", argPos))
		args = append(args, *input.DiscService)
		argPos++
	}
	if input.DiscTransport != nil {
		setParts = append(setParts, fmt.Sprintf("disc_transport=$%d", argPos))
		args = append(args, *input.DiscTransport)
		argPos++
	}
	if input.DiscTranpEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("disc_tranp_employee=$%d", argPos))
		args = append(args, *input.DiscTranpEmployee)
		argPos++
	}
	if input.DiscLabor != nil {
		setParts = append(setParts, fmt.Sprintf("disc_labor=$%d", argPos))
		args = append(args, *input.DiscLabor)
		argPos++
	}
	if input.DiscMaterial != nil {
		setParts = append(setParts, fmt.Sprintf("disc_material=$%d", argPos))
		args = append(args, *input.DiscMaterial)
		argPos++
	}
	if input.DiscField != nil {
		setParts = append(setParts, fmt.Sprintf("disc_field=$%d", argPos))
		args = append(args, *input.DiscField)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.new("Missing fields for update")
	}

	var discountId int

	err := r.db.QueryRow(query, args...).Scan(&discountId)

	return discountId, err
}

func (r *ContractsRepository) EditContractRhInfo(id string, c models.EditContractRhInfo) {
	setParts := []string
	args := []interface{}{}
	argPos := 1

	if input.HourLimit != nil {
		setParts = append(setParts, fmt.Sprintf("hour_limit=$%d", argPos))
		args = append(args, *input.HourLimit)
		argPos++
	}
	if input.MinutesLimit != nil {
		setParts = append(setParts, fmt.Sprintf("minutes_limit=$%d", argPos))
		args = append(args, *input.MinutesLimit)
		argPos++
	}
	if input.DaysFirstExp != nil {
		setParts = append(setParts, fmt.Sprintf("days_first_exp=$%d", argPos))
		args = append(args, *input.DaysFirstExp)
		argPos++
	}
	if input.DaysSecondExp != nil {
		setParts = append(setParts, fmt.Sprintf("days_second_exp=$%d", argPos))
		args = append(args, *input.DaysSecondExp)
		argPos++
	}
	if input.DataInit != nil {
		setParts = append(setParts, fmt.Sprintf("data_init=$%d", argPos))
		args = append(args, *input.DataInit)
		argPos++
	}
	if input.PayExtraHour != nil {
		setParts = append(setParts, fmt.Sprintf("pay_extra_hour=$%d", argPos))
		args = append(args, *input.PayExtraHour)
		argPos++
	}
	if input.ManualStitch != nil {
		setParts = append(setParts, fmt.Sprintf("manual_stitch=$%d", argPos))
		args = append(args, *input.ManualStitch)
		argPos++
	}
	if input.PaysBreakfast != nil {
		setParts = append(setParts, fmt.Sprintf("pays_breakfast=$%d", argPos))
		args = append(args, *input.PaysBreakfast)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.new("Missing fields for update")
	}

	query := fmt.Sprintf("UPDATE contract_rh_info SET %s WHERE id=$%d RETURNING id", string.Join(setParts, ", "), argPos)

	var contractRhId int
	err := r.db.QueryRow(query, args...).Scan(
		&contractRhId
	)

	return contractRhId, err
}