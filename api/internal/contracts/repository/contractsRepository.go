package repository

import (
	"app/internal/contracts/models"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type ContractsRepository struct {
	db *sql.DB
}

func NewContractsRepository(db *sql.DB) *ContractsRepository {
	return &ContractsRepository{db: db}
}

func (r *ContractsRepository) GetAllContracts() ([]models.GetAllContracts, error) {
	query := `SELECT c.id, c.name, c.code, d.date_initial FROM contracts c LEFT JOIN contracts_date d ON c.fk_contract_dates = d.id;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contracts []models.GetAllContracts

	for rows.Next() {
		var contract models.GetAllContracts
		err := rows.Scan(
			&contract.ID,
			&contract.Name,
			&contract.Code,
			&contract.InitialDate,
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

func (r *ContractsRepository) CreateContract(tx *sql.Tx, input models.CreateContract) (int, error) {
	query := `INSERT INTO contracts (name, code, research, uses_cpf, is_active, fk_contract_info, fk_contract_dates, fk_contract_values, fk_contract_discount, fk_contract_rh, is_active) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, TRUE) RETURNING id;`

	var contractID int

	err := tx.QueryRow(
		query,
		input.Name,
		input.Code,
		input.Research,
		input.UsesCpf,
		input.IsActive,
		input.FkContractInfo,
		input.FkContractDates,
		input.FkContractValues,
		input.FkContractDiscount,
		input.FkContractRh,
	).Scan(
		&contractID,
	)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return contractID, err
}

func (r *ContractsRepository) CreateContractInfo(tx *sql.Tx, input models.CreateContractInfo) (int, error) {
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

func (r *ContractsRepository) CreateContractDates(tx *sql.Tx, input models.CreateContractDates) (int, error) {
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

func (r *ContractsRepository) CreateContractValues(tx *sql.Tx, input models.CreateContractValues) (int, error) {
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

	err := tx.QueryRow(query,
		input.AcronymValues,
		input.BdiService,
		input.BdiMaterial,
		input.BdiLabor,
		input.EntryTable,
		input.SendEmail,
	).Scan(&valuesId)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return valuesId, nil
}

func (r *ContractsRepository) CreateContractDiscount(tx *sql.Tx, input models.CreateContractDiscount) (int, error) {
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

func (r *ContractsRepository) CreateContractRhInfo(tx *sql.Tx, input models.CreateContractRhInfo) (int, error) {
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

func (r *ContractsRepository) EditContract(id string, c models.EditContract) (int, error) {
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

	if c.UsesCpf != nil {
		setParts = append(setParts, fmt.Sprintf("uses_cpf=$%d", argPos))
		args = append(args, *c.UsesCpf)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.New("Missing fields for update")
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE contracts SET %s WHERE id=$%d RETURNING id;", strings.Join(setParts, ", "), argPos)

	var contractId int
	err := r.db.QueryRow(query, args...).Scan(
		&contractId,
	)

	return contractId, err

}

func (r *ContractsRepository) EditContractValues(id string, c models.EditContractValues) (int, error) {
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
		return 0, errors.New("Missing fields for update")
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
		&updated.SendEmail,
	)

	return updated.ID, err
}

func (r *ContractsRepository) EditContractInfo(id string, c models.EditContractInfo) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.Construction != nil {
		setParts = append(setParts, fmt.Sprintf("construction=$%d", argPos))
		args = append(args, *c.Construction)
		argPos++
	}

	if c.CAP != nil {
		setParts = append(setParts, fmt.Sprintf("cap=$%d", argPos))
		args = append(args, *c.CAP)
		argPos++
	}

	if c.Process != nil {
		setParts = append(setParts, fmt.Sprintf("process=$%d", argPos))
		args = append(args, *c.Process)
		argPos++
	}
	if c.InfoPmco != nil {
		setParts = append(setParts, fmt.Sprintf("info_pmco=$%d", argPos))
		args = append(args, *c.InfoPmco)
		argPos++
	}
	if c.MaxEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("max_employee=$%d", argPos))
		args = append(args, *c.MaxEmployee)
		argPos++
	}
	if c.Address != nil {
		setParts = append(setParts, fmt.Sprintf("address=$%d", argPos))
		args = append(args, *c.Address)
		argPos++
	}
	if c.NRO != nil {
		setParts = append(setParts, fmt.Sprintf("nro=$%d", argPos))
		args = append(args, *c.NRO)
		argPos++
	}
	if c.Complement != nil {
		setParts = append(setParts, fmt.Sprintf("complement=$%d", argPos))
		args = append(args, *c.Complement)
		argPos++
	}
	if c.Phone != nil {
		setParts = append(setParts, fmt.Sprintf("phone=$%d", argPos))
		args = append(args, *c.Phone)
		argPos++
	}
	if c.State != nil {
		setParts = append(setParts, fmt.Sprintf("state=$%d", argPos))
		args = append(args, *c.State)
		argPos++
	}
	if c.City != nil {
		setParts = append(setParts, fmt.Sprintf("city=$%d", argPos))
		args = append(args, *c.City)
		argPos++
	}
	if c.CEP != nil {
		setParts = append(setParts, fmt.Sprintf("cep=$%d", argPos))
		args = append(args, *c.CEP)
		argPos++
	}
	if c.Email != nil {
		setParts = append(setParts, fmt.Sprintf("email=$%d", argPos))
		args = append(args, *c.Email)
		argPos++
	}
	if c.Contact != nil {
		setParts = append(setParts, fmt.Sprintf("contact=$%d", argPos))
		args = append(args, *c.Contact)
		argPos++
	}
	if c.FkEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("fk_employee=$%d", argPos))
		args = append(args, *c.FkEmployee)
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

func (r *ContractsRepository) EditContractDates(id string, c models.EditContractDates) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.DateInitial != nil {
		setParts = append(setParts, fmt.Sprintf("date_initial=$%d", argPos))
		args = append(args, *c.DateInitial)
		argPos++
	}

	if c.DateLimit != nil {
		setParts = append(setParts, fmt.Sprintf("date_limit=$%d", argPos))
		args = append(args, *c.DateLimit)
		argPos++
	}

	if c.DateGuarantee != nil {
		setParts = append(setParts, fmt.Sprintf("date_guarantee=$%d", argPos))
		args = append(argPos, *c.DateProposal)
		argPos++
	}

	if c.DateProposal != nil {
		setParts = append(setParts, fmt.Sprintf("date_proposal=%d", argPos))
		args = append(argPos, *c.DateProposal)
		argPos++
	}

	if c.DateBudget != nil {
		setParts = append(setParts, fmt.Sprintf("date_budget=$%d", argPos))
		args = append(argPos, *c.DateBudget)
		argPos++
	}

	if c.DateTables != nil {
		setParts = append(setParts, fmt.Sprintf("date_tables=$%d", argPos))
		args = append(argPos, *c.DateTables)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.New("nenhum campo para atualizar")
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE contract_dates SET %s WHERE id=$%d RETURNING id;", strings.Join(setParts, ", "), argPos)

	var valuesId int

	err := r.db.QueryRow(query, args...).Scan(&valuesId)
	return valuesId, err
}

func (r *ContractsRepository) EditContractDiscount(id string, c models.EditContractDiscount) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if c.DiscIdentifier != nil {
		setParts = append(setParts, fmt.Sprintf("disc_identifier=$%d", argPos))
		args = append(args, *c.DiscIdentifier)
		argPos++
	}
	if c.DiscService != nil {
		setParts = append(setParts, fmt.Sprintf("disc_service=$%d", argPos))
		args = append(args, *c.DiscService)
		argPos++
	}
	if c.DiscTransport != nil {
		setParts = append(setParts, fmt.Sprintf("disc_transport=$%d", argPos))
		args = append(args, *c.DiscTransport)
		argPos++
	}
	if c.DiscTranpEmployee != nil {
		setParts = append(setParts, fmt.Sprintf("disc_tranp_employee=$%d", argPos))
		args = append(args, *c.DiscTranpEmployee)
		argPos++
	}
	if c.DiscLabor != nil {
		setParts = append(setParts, fmt.Sprintf("disc_labor=$%d", argPos))
		args = append(args, *c.DiscLabor)
		argPos++
	}
	if c.DiscMaterial != nil {
		setParts = append(setParts, fmt.Sprintf("disc_material=$%d", argPos))
		args = append(args, *c.DiscMaterial)
		argPos++
	}
	if c.DiscField != nil {
		setParts = append(setParts, fmt.Sprintf("disc_field=$%d", argPos))
		args = append(args, *c.DiscField)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, errors.New("Missing fields for update")
	}

	var discountId int

	query := fmt.Sprintf("")

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
		&contractRhId,
	)

	return contractRhId, err
}

func (r *ContractsRepository) DeactivateContract(id string) {
	query := `UPDATE contract SET is_active = false WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *ContractsRepository) ReactivateDepartment(id string) {
	query := `UPDATE contract SET is_active = true WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
