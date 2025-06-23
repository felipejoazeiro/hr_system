package repository

func (r *EmployeeRepository) GetEmployeeByID(id string) (*models.detailsEmployeeResponse, error){
	query := `
        SELECT 
            e.id,
            e.register,
            e.name_employee,
            e.email,
            p.name as position_name,
            d.name as department_name,
            e.active_employee,
            e.entry_date,
            e.contract_date,
            e.photo,
            e.birthdate,
            e.nro_title,
            e.electoral_zone,
            e.electoral_section,
            e.nro_rg,
            e.state_rg,
            e.nro_work_card,
            e.series_work_card,
            e.cpf,
            e.phone,
            e.address,
            e.district,
            e.city,
            e.states,
            e.uf,
            e.cep,
            e.identification,
            e.code_operator,
            e.code_line,
            e.card_number,
            e.qtt_daily_ticker,
            e.ticket_value,
            e.fk_ticket_type,
            e.updated_at
        FROM employees e
        LEFT JOIN positions p ON e.fk_cargo = p.id
        LEFT JOIN departments d ON e.department_id = d.id
        WHERE e.id = ?
    `

	var employee models.EmployeeDetailReponse
	err := r.db.QueryRow(query, id).Scan( 
		&employee.ID,
        &employee.Register,
        &employee.Name,
        &employee.Email,
        &employee.Position,
        &employee.Department,
        &employee.Active,
        &employee.EntryDate,
        &employee.ContractDate,
        &employee.PhotoURL,
        &employee.BirthDate,
        &employee.NroTitle,
        &employee.ElectoralZone,
        &employee.ElectoralSection,
        &employee.NroRG,
        &employee.StateRG,
        &employee.NroWorkCard,
        &employee.SeriesWorkCard,
        &employee.CPF,
        &employee.Phone,
        &employee.Address,
        &employee.District,
        &employee.City,
        &employee.States,
        &employee.UF,
        &employee.CEP,
        &employee.Identification,
        &employee.CodeOperator,
        &employee.CodeLine,
        &employee.CardNumber,
        &employee.QttDailyTicket,
        &employee.TicketValue,
        &employee.FkTicketType,
        &employee.LastUpdate,
	)

	if err != nil {
		return nil, err
	}
	return &employee, nil
}