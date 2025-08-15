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

func (r *DepartmentRepository) EditSindicate(id int, d models.EditSindicate) (int, error){
	setParts := []string{}{}
	args := []interface{}{}
	argPos := 1
	if d.NroRegister != nil {
		setParts = append(setParts, fmt.Sprintf("nro_register=$%d", argPos))
		args = append(args, *d.NroRegister)
		argPos++
	}
	if d.DateStart != nil {
		setParts = append(setParts, fmt.Sprintf("date_start=$%d", argPos))
		args = append(args, *d.DateStart)
		argPos++
	}
	if d.DateFinish != nil {
		setParts = append(setParts, fmt.Sprintf("date_finish=$%d", argPos))
		args = append(args, *d.DateFinish)
		argPos++
	}
	if len(setParts) == 0 {
		return 0, nil
	}
	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)
	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID)
	return updatedID, err

}

func (r *SindicateRepository) EditSindicateAuthorization(id int, d models.EditSindicateAuthorization) (int, error){
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

func (r *SindicateRepository) EditSindicateBreakfest(id int, d models.EditSindicateBreakfest) (int, error) {
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
	if d.AbscenceDiscount != nil {
		setParts = append(setParts, fmt.Sprintf("abscence_discount=$%d"))
		args = append(args, *d.AbscenceDiscount)
		argPos++
	}
	if d.AbscenceDiscountProp != nil {
		setParts = append(setParts, fmt.Sprintf("abscence_discount_prop=$%d"))
		args = append(args, *d.AbscenceDiscountProp)
		argPos++
	}
	if d.VacationDiscount != nil {
		setParts = append(setParts, fmt.Sprintf("vacation_discount=$%d"))
		args = append(args, *d.VacationDiscount)
		argPos++
	}
	if d.VacationDiscountProp != nil {
		setParts = append(setParts, fmt.Sprintf("vacation_discount_prop=$%d"))
		args = append(args, *d.VacationDiscountProp)
		argPos++
	}
	if d.LeaveDiscount != nil {
		setParts = append(setParts, fmt.Sprintf("leave_discount=$%d"))
		args = append(args, *d.LeaveDiscount)
		argPos++
	}
	if d.LeaveDiscountProp != nil {
		setParts = append(setParts, fmt.Sprintf("leave_discount_prop=$%d"))
		args = append(args, *d.LeaveDiscountProp)
		argPos++
	}
	if d.BreakNature != nil {
		setParts = append(setParts, fmt.Sprintf("break_nature=$%d"))
		args = append(args, *d.BreakNature)
		argPos++
	}
	if d.LimitDescPercent != nil {
		setParts = append(setParts, fmt.Sprintf("limit_desc_percent=$%d"))
		args = append(args, *d.LimitDescPercent)
		argPos++
	}
	if d.ValueDay != nil {
		setParts = append(setParts, fmt.Sprintf("value_day=$%d"))
		args = append(args, *d.ValueDay)
		argPos++
	}
	if d.ValueMonth != nil {
		setParts = append(setParts, fmt.Sprintf("value_month=$%d"))
		args = append(args, *d.ValueMonth)
		argPos++
	}
	if d.LimitDesc != nil {
		setParts = append(setParts, fmt.Sprintf("limit_desc=$%d"))
		args = append(args, *d.LimitDesc)
		argPos++
	}
	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_breakfest SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID) 
	return updatedID, err
}

func (r *SindicateRepository) EditSindicatePercents(id int, d models.EditSindicatePercents) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if d.ExtraHourWeek != nil {
		setParts = append(setParts, fmt.Sprintf("extra_hour_week=$%d", argPos))
		args = append(args, *d.ExtraHourWeek)
		argPos++
	}
	if d.ExtraHourSaturday != nil {
		setParts = append(setParts, fmt.Sprintf("extra_hour_saturday=$%d", argPos))
		args = append(args, *d.ExtraHourSaturday)
		argPos++
	}
	if d.ExtraHourSundayHolidays != nil {
		setParts = append(setParts, fmt.Sprintf("extra_hour_sunday_holidays=$%d", argPos))
		args = append(args, *d.ExtraHourSundayHolidays)
		argPos++
	}
	if d.SindicateContr != nil {
		setParts = append(setParts, fmt.Sprintf("sindicate_contr=$%d", argPos))
		args = append(args, *d.SindicateContr)
		argPos++
	}
	if d.Seconci != nil {
		setParts = append(setParts, fmt.Sprintf("seconci=$%d", argPos))
		args = append(args, *d.Seconci)
		argPos++
	}
	if d.SalaryAdvancePercents != nil {
		setParts = append(setParts, fmt.Sprintf("salary_advance_percents=$%d", argPos))
		args = append(args, *d.SalaryAdvancePercents)
		argPos++
	}
	if d.IncentiveQualify != nil {
		setParts = append(setParts, fmt.Sprintf("incentive_qualify=$%d", argPos))
		args = append(args, *d.IncentiveQualify)
		argPos++
	}
	if d.ComplemRetirement != nil {
		setParts = append(setParts, fmt.Sprintf("complem_retirement=$%d", argPos))
		args = append(args, *d.ComplemRetirement)
		argPos++
	}
	if d.NightTime != nil {
		setParts = append(setParts, fmt.Sprintf("night_time=$%d", argPos))
		args = append(args, *d.NightTime)
		argPos++
	}
	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_percents SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID)
	return updatedID, err
}

func (r *SindicateRepository) EditSindicateValues(id int, d models.EditSindicateValues) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if d.Attendance != nil {
		setParts = append(setParts, fmt.Sprintf("attendance=$%d", argPos))
		args = append(args, *d.Attendance)
		argPos++
	}
	if d.BaseSalary != nil {
		setParts = append(setParts, fmt.Sprintf("base_salary=$%d", argPos))
		args = append(args, *d.BaseSalary)
		argPos++
	}
	if d.HealthInsurance != nil {
		setParts = append(setParts, fmt.Sprintf("health_insurance=$%d", argPos))
		args = append(args, *d.HealthInsurance)
		argPos++
	}
	if d.ChristmasBonus != nil {
		setParts = append(setParts, fmt.Sprintf("christmas_bonus=$%d", argPos))
		args = append(args, *d.ChristmasBonus)
		argPos++
	}
	if d.AssociativeContribution != nil {
		setParts = append(setParts, fmt.Sprintf("associative_contribution=$%d", argPos))
		args = append(args, *d.AssociativeContribution)
		argPos++
	}
	if d.ChildcareAssistance != nil {
		setParts = append(setParts, fmt.Sprintf("childcare_assistance=$%d", argPos))
		args = append(args, *d.ChildcareAssistance)
		argPos++
	}
	if d.EducationAid != nil {
		setParts = append(setParts, fmt.Sprintf("education_aid=$%d", argPos))
		args = append(args, *d.EducationAid)
		argPos++
	}
	if d.UnhealthnessAid != nil {
		setParts = append(setParts, fmt.Sprintf("unhealthness_aid=$%d", argPos))
		args = append(args, *d.UnhealthnessAid)
		argPos++
	}
	if.Dangerousness != nil {
		setParts = append(setParts, fmt.Sprintf("dangerousness=$%d", argPos))
		args = append(args, *d.Dangerousness)
		argPos++
	}
	if d.NightTimeSupplement != nil {
		setParts = append(setParts, fmt.Sprintf("night_time_supplement=$%d", argPos))
		args = append(args, *d.NightTimeSupplement)
		argPos++
	}
	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_values SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID)
	return updatedID, err
}

func (r *SindicateRepository) EditSindicateVoucher(id int, d models.EditSindicateVoucher) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if d.AbscenceDiscountVAVR != nil {
		setParts = append(setParts, fmt.Sprintf("abscence_discount_vavr=$%d", argPos))
		args = append(args, *d.AbscenceDiscountVAVR)
		argPos++
	}
	if d.AbscenceDiscountVT != nil {
		setParts = append(setParts, fmt.Sprintf("abscence_discount_vt=$%d", argPos))
		args = append(args, *d.AbscenceDiscountVT)
		argPos++
	}
	if d.MedCertificateDiscountVAVR != nil {
		setParts = append(setParts, fmt.Sprintf("med_certificate_discount_vavr=$%d", argPos))
		args = append(args, *d.MedCertificateDiscountVAVR)
		argPos++
	}
	if d.MedCertificateDiscountVT != nil {
		setParts = append(setParts, fmt.Sprintf("med_certificate_discount_vt=$%d", argPos))
		args = append(args, *d.MedCertificateDiscountVT)
		argPos++
	}
	if d.VacationDiscountVAVR != nil {	
		setParts = append(setParts, fmt.Sprintf("vacation_discount_vavr=$%d", argPos))
		args = append(args, *d.VacationDiscountVAVR)
		argPos++
	}
	if d.VacationDiscountVT != nil {
		setParts = append(setParts, fmt.Sprintf("vacation_discount_vt=$%d", argPos))
		args = append(args, *d.VacationDiscountVT)
		argPos++
	}
	if d.SickLeaveDiscountVAVRINSS != nil {
		setParts = append(setParts, fmt.Sprintf("sick_leave_discount_vavr_inss=$%d", argPos))
		args = append(args, *d.SickLeaveDiscountVAVRINSS)
		argPos++
	}
	if d.SickLeaveDiscountVTINSS != nil {
		setParts = append(setParts, fmt.Sprintf("sick_leave_discount_vt_inss=$%d", argPos))
		args = append(args, *d.SickLeaveDiscountVTINSS)
		argPos++
	}
	if d.ValueAlimDay != nil {
		setParts = append(setParts, fmt.Sprintf("value_alim_day=$%d", argPos))
		args = append(args, *d.ValueAlimDay)
		argPos++
	}
	if d.DeductLimitByAllowanceonPayroll != nil {
		setParts = append(setParts, fmt.Sprintf("deduct_limit_by_allowanceon_payroll=$%d", argPos))
		args = append(args, *d.DeductLimitByAllowanceonPayroll)
		argPos++
	}
	if d.PercentualDeduct != nil {
		setParts = append(setParts, fmt.Sprintf("percentual_deduct=$%d", argPos))
		args = append(args, *d.PercentualDeduct)
		argPos++
	}
	if d.DeductLimitByVAVRonPayroll != nil {
		setParts = append(setParts, fmt.Sprintf("deduct_limit_by_vavr_on_payroll=$%d", argPos))
		args = append(args, *d.DeductLimitByVAVRonPayroll)
		argPos++
	}
	if d.DiscountVtPerc != nil {
		setParts = append(setParts, fmt.Sprintf("discount_vt_perc=$%d", argPos))
		args = append(args, *d.DiscountVtPerc)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_voucher SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID)
	return updatedID, err
}

func (r *SindicateRepository) EditSindicateCarePackage(id int, d models.EditSindicateCarePackage) (int, error) {
	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if d.DiscountBasic != nil {
		setParts = append(setParts, fmt.Sprintf("discount_basic=$%d", argPos))
		args = append(args, *d.DiscountBasic)
		argPos++
	}
	if d.ValueBasic != nil {
		setParts = append(setParts, fmt.Sprintf("value_basic=$%d", argPos))
		args = append(args, *d.ValueBasic)	
		argPos++
	}
	if d.LimitDiscBasic != nil {
		setParts = append(setParts, fmt.Sprintf("limit_disc_basic=$%d", argPos))
		args = append(args, *d.LimitDiscBasic)
		argPos++
	}
	if d.AbscenceDiscBasic != nil {
		setParts = append(setParts, fmt.Sprintf("abscence_disc_basic=$%d", argPos))
		args = append(args, *d.AbscenceDiscBasic)
		argPos++
	}
	if d.LeaveDiscBasic != nil {
		setParts = append(setParts, fmt.Sprintf("leave_disc_basic=$%d", argPos))
		args = append(args, *d.LeaveDiscBasic)
		argPos++
	}
	if d.VacationDisc != nil {
		setParts = append(setParts, fmt.Sprintf("vacation_disc=$%d", argPos))
		args = append(args, *d.VacationDisc)
		argPos++
	}
	if d.VacationDiscountProp != nil {
		setParts = append(setParts, fmt.Sprintf("vacation_discount_prop=$%d", argPos))
		args = append(args, *d.VacationDiscountProp)
		argPos++
	}
	if d.GetBasicWithOneDay != nil {
		setParts = append(setParts, fmt.Sprintf("get_basic_with_one_day=$%d", argPos))
		args = append(args, *d.GetBasicWithOneDay)
		argPos++
	}
	if d.MagneticCard != nil {
		setParts = append(setParts, fmt.Sprintf("magnetic_card=$%d", argPos))
		args = append(args, *d.MagneticCard)
		argPos++
	}
	if d.MissesDayPropDiscount != nil {
		setParts = append(setParts, fmt.Sprintf("misses_day_prop_discount=$%d", argPos))
		args = append(args, *d.MissesDayPropDiscount)
		argPos++
	}
	if d.LeavePropDisc != nil {
		setParts = append(setParts, fmt.Sprintf("leave_prop_disc=$%d", argPos))
		args = append(args, *d.LeavePropDisc)
		argPos++
	}
	if d.MedCertificateDisc != nil {
		setParts = append(setParts, fmt.Sprintf("med_certificate_disc=$%d", argPos))
		args = append(args, *d.MedCertificateDisc)
		argPos++
	}
	if d.MedCertificatePropDisc != nil {
		setParts = append(setParts, fmt.Sprintf("med_certificate_prop_disc=$%d", argPos))
		args = append(args, *d.MedCertificatePropDisc)
		argPos++
	}
	if d.BasicNature != nil {
		setParts = append(setParts, fmt.Sprintf("basic_nature=$%d", argPos))
		args = append(args, *d.BasicNature)
		argPos++
	}

	if len(setParts) == 0 {
		return 0, nil
	}

	args = append(args, id)
	query := fmt.Sprintf("UPDATE sindicate_care_package SET %s WHERE id=$%d returning id", strings.Join(setParts, ", "), argPos)

	var updatedID int
	err := r.db.QueryRow(query, args...).Scan(&updatedID)
	return updatedID, err
}

func (r *SindicateRepository) DeactivateSindicate(id int) error {
	query := `UPDATE sindicate SET is_active = false WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err	
}

func (r *SindicateRepository) ReactivateSindicate(id int) error {
	query := `UPDATE sindicate SET is_active = true WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *SindicateRepository) CreateSindicatePercents(d models.SindicatePercents) (int, error) {
	query := `INSERT INTO sindicate_percents (extra_hour_week, extra_hour_saturday, extra_hour_sunday_holidays, sindicate_contr, seconci, salary_advance_percents, incentive_qualify, complem_retirement, night_time) 
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.ExtraHourWeek,
		d.ExtraHourSaturday,
		d.ExtraHourSundayHolidays,
		d.SindicateContr,
		d.Seconci,
		d.SalaryAdvancePercents,
		d.IncentiveQualify,
		d.ComplemRetirement,
		d.NightTime).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicateValues(d models.CreateSindicateValues) (int, error) {
	query := `INSERT INTO sindicate_values (attendance, base_salary, health_insurance, christmas_bonus, associative_contribution, childcare_assistance, education_aid, unhealthness_aid, dangerousness, night_time_supplement)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.Attendance,
		d.BaseSalary,
		d.HealthInsurance,
		d.ChristmasBonus,
		d.AssociativeContribution,
		d.ChildcareAssistance,
		d.EducationAid,
		d.UnhealthnessAid,
		d.Dangerousness,
		d.NightTimeSupplement).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicateVoucher(d models.CreateSindicateVoucher) (int, error) {
	query := `INSERT INTO sindicate_voucher (abscence_discount_vavr, abscence_discount_vt, med_certificate_discount_vavr, med_certificate_discount_vt, vacation_discount_vavr, vacation_discount_vt, sick_leave_discount_vavr_inss, sick_leave_discount_vt_inss, value_alim_day, deduct_limit_by_allowanceon_payroll, percentual_deduct, deduct_limit_by_vavr_on_payroll, discount_vt_perc)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.AbscenceDiscountVAVR,
		d.AbscenceDiscountVT,
		d.MedCertificateDiscountVAVR,
		d.MedCertificateDiscountVT,
		d.VacationDiscountVAVR,
		d.VacationDiscountVT,
		d.SickLeaveDiscountVAVRINSS,
		d.SickLeaveDiscountVTINSS,
		d.ValueAlimDay,
		d.DeductLimitByAllowanceonPayroll,
		d.PercentualDeduct,
		d.DeductLimitByVAVRonPayroll,
		d.DiscountVtPerc).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicateCarePackage(d models.CreateSindicateCarePackage) (int, error) {
	query := `INSERT INTO sindicate_care_package (discount_basic, value_basic, limit_disc_basic, abscence_disc_basic, leave_disc_basic, vacation_disc, vacation_discount_prop, get_basic_with_one_day, magnetic_card, misses_day_prop_discount, leave_prop_disc, med_certificate_disc, med_certificate_prop_disc, basic_nature)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.DiscountBasic,
		d.ValueBasic,
		d.LimitDiscBasic,
		d.AbscenceDiscBasic,
		d.LeaveDiscBasic,
		d.VacationDisc,
		d.VacationDiscountProp,
		d.GetBasicWithOneDay,
		d.MagneticCard,
		d.MissesDayPropDiscount,
		d.LeavePropDisc,
		d.MedCertificateDisc,
		d.MedCertificatePropDisc,
		d.BasicNature).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicateAuthorization(d models.CreateSindicateAuthorization) (int, error) {
	query := `INSERT INTO sindicate_authorization (normal_hour, salary_advance, bank_hours, shift_work, sind_pay_beneith_empl, mobile_point_app_approval)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.NormalHour,
		d.SalaryAdvance,
		d.BankHours,
		d.ShiftWork,
		d.SindPayBeneithEmpl,
		d.MobilePointAppApproval).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicateBreakfest(d models.CreateSindicateBreakfest) (int, error) {
	query := `INSERT INTO sindicate_breakfest (med_certification_disc, med_certification_disc_prop, abscence_discount, abscence_discount_prop, vacation_discount, vacation_discount_prop, leave_discount, leave_discount_prop, break_nature, limit_desc_percent, value_day, value_month, limit_desc)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	var id int
	err := r.db.QueryRow(query,
		d.MedCertificationDisc,
		d.MedCertificationDiscProp,
		d.AbscenceDiscount,
		d.AbscenceDiscountProp,
		d.VacationDiscount,
		d.VacationDiscountProp,
		d.LeaveDiscount,
		d.LeaveDiscountProp,
		d.BreakNature,
		d.LimitDescPercent,
		d.ValueDay,
		d.ValueMonth,
		d.LimitDesc).Scan(&id)
	return id, err
}

func (r *SindicateRepository) CreateSindicate(d models.CreateSindicateReq) (int, error) {
	
	percentsId, err := r.CreateSindicatePercents(d.SindicatePercents)
	if err != nil {	
		return 0, err
	}

	valuesId, err := r.CreateSindicateValues(d.SindicateValues)
	if err != nil {
		return 0, err
	}

	authorizationId, err := r.CreateSindicateAuthorization(d.SindicateAuthorization)
	if err != nil {
		return 0, err
	}

	breakfestId, err := r.CreateSindicateBreakfest(d.SindicateBreakfest)
	if err != nil {
		return 0, err
	}

	carePackageId, err := r.CreateSindicateCarePackage(d.SindicateCarePackage)
	if err != nil {	
		return 0, err
	}

	voucherId, err := r.CreateSindicateVoucher(d.SindicateVoucher)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO sindicate (nro_register, date_start, date_finish, fk_sindicate_authorization, fk_sindicate_breakfest, fk_sindicate_care_package, fk_sindicate_percents, fk_sindicate_values, fk_sindicate_voucher)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	err := r.db.QueryRow(query,
		d.NroRegister,
		d.DateStart,
		d.DateFinish,
		authorizationId,
		breakfestId,
		carePackageId,
		percentsId,
		valuesId,
		voucherId).Scan(&id)
	return id, err
}

