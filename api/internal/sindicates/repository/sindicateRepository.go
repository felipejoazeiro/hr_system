package repository

type SindicateRepository struct {
	db *sql.DB
}

func NewSindicateRepository(db *sql.DB) &SindicateRepository {
	return &SindicateRepository{db: db}
}

func (r *SindicateRepository) GetAllSindicate() ([]models.GetSindicatesModel, error){
	query := `SELECT s.id, s.nro_register, s.data_start, s._finish FROM sindicates s;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sindicates []models.GetSindicatesModel

	for rows.Next{
		var sind models.GetSindicatesModel
		err:=rows.Scan(
			&sind.ID,
			&sind.NroRegister,
			&sind.DateStart,
			&sind.DateFinish,
		)
		if err != nil {
			return nil, err
		}
		sindicates = append(sindicates, sind)
	}
	if err := rows.Err(); err != nil{
		return nil, err
	}

	return sindicates, nil
}

func (r *SindicateRepository) getSindicateAuthorizationId(id int) (int, error){
	query := `SELECT s.fk_sindicate_authorization FROM sindicate s WHERE s.id = $1;`
	authorizationId, err := r.db.Exec(query, id)
	return authorizationId, err
}

func (r *SindicateRepository) getSindicateBreakfestId(id int) (int, error) {
	query := `SELECT s.fk_sindicate_breakfest FROM sindicate s WHERE s.id = $1;`
	breakfestId, err := r.db.Exec(query, id)
	return breakfestId, err
}

func (r *SindicateRepository) getSindicateCarePackageId(id int) (int, error) {
	query := `SELECT s.fk_sindicate_care_package FROM sindicate s WHERE s.id = $1;`
	carePackageId, err := r.db.Exec(query, id)
	return carePackageId, err
}

func (r *SindicateRepository) getSindicatePercentsId(id int) (int, error) {
	query := `SELECT s.fk_sindicate_percents FROM sindicate s WHERE s.id = $1;`
	PercentsId, err := r.db.Exec(query, id)
	return PercentsId, err
}

func (r *SindicateRepository) getSindicateValuesId(id int) (int, error) {
	query := `SELECT s.fk_sindicate_values FROM sindicate s WHERE s.id = $1;`
	valuesId, err := r.db.Exec(query, id)
	return valuesId, err
}

func (r *SindicateRepository) getSindicateVoucherId(id int) (int, error) {
	query := `SELECT s.fk_sindicate_voucher FROM sindicate s WHERE s.id = $1;`
	voucherId, err := r.db.Exec(query, id)
	return voucherId, err
}

func (r *SindicateRepository) EditSindicateAuthorization(id int, d EditSindicateAuthorization) (int, error){
	setParts := []string{}{}
	args := []interface{}{}
	argPos := 1

	if d.NormalHour != nil {
		setParts = append(setParts, fmt.Sprintf("normal_hour=$%d", argPos))
		args = append(args, *d.NormalHour)
		argPos++
	}

	if d.SalaryAdvance != nil {
		setParts = append(setParts, fmt.Sprintf("salary_advance=$%d", argPos))
		args = append(args, *d.SalaryAdvance)
		argPos++
	}

	if d.BankHours != nil {
		setParts = append(setParts, fmt.Sprintf("bank_hours=$%d", argPos))
		args = append(args, *d.BankHours)
		argPos++
	}

	if d.ShiftWork != nil {
		setParts = append(setParts, fmt.Sprintf("shift_work=$%d", argPos))
		args = append(args, *d.ShiftWork)
		argPos++
	}

	if d.SindPayBeneithEmpl != nil {
		setParts = append(setParts, fmt.Sprintf("sind_pay_beneith_empl=$%d", argPos))
		args = append(args, *d.SindPayBeneithEmpl)
		argPos++
	}

	if d.MobilePointAppApproval != nil {
		setParts = append(setParts, fmt.Sprintf("mobile_point_app_approval=$%d", argPos))
		args = append(arsg, *d.MobilePointAppApproval)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_authorization SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID.ID)

	return updatedID, err
}

func (r *SindicateRepository) EditSindicateBreakfest(id int, d EditSindicateBreakfest) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if d.MedCertificationDisc != nil {
		setParts = append(setParts, fmt.Sprintf("med_certification_disc=$%d"))
		args = append(args, *d.MedCertificationDisc)
		argPos++
	}

	if d.MedCertificationDiscProp != nil {
		setParts = append(setParts, fmt.Sprintf("med_certification_disc_prop=$%d"))
		args = append(args, *d.MedCertificationDiscProp)
		argPos++
	}
}

