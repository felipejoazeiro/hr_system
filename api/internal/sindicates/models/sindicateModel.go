package sindicate

type SindicateModel struct {
	ID 							int				`json:"id" db:"id"`
	NroRegister					string			`json:"nro_register" db:"nro_register"`
	DateStart					date			`json:"date_start" db:"date_start"`
	DateFinish					date			`json:"date_finish" db:"date_finish"`
	FkSindicatePercents			int				`json:"fk_sindicate_percents" db:"fk_sindicate_percentes"`
	FkSindicateValues 			int				`json:"fk_sindicate_values" db:"fk_sindicate_values"`
	FkSindicateAuthorization	int				`json:"fk_sindicate_authorization" db:"fk_sindicate_authorization"`
	FkSindicateBreakfest		int				`json:"fk_sindicate_breakfest" db:"fk_sindicate_breakfest"`
	FkSindicateCarePackage		int				`json:"fk_sindicate_care_package db:"fk_sindicate_care_package`
	FkSindicateVoucher			int				`json:"fk_sindicate_voucher db:"fk_sindicate_voucher`
	IsActive					bool 			`json:"is_activate" db:"is_activate"`
}

type CreateSindicateReq struct {
	NroRegister					string							`json:"nro_register" db:"nro_register"`
	DateStart					date							`json:"date_start" db:"date_start"`
	DateFinish					date							`json:"date_finish" db:"date_finish`
	SindicatePercents			CreateSindicatePercents			`json:"sindicate_percents"`
	SindicateValues				CreateSindicateValues			`json:"sindicate_values"`
	SindicateAuthorization		CreateSindicateAuthorization	`json:"sindicate_authorization"`
	SindicateBreakfest			CreateSindicateBreakfest		`json:"sindicate_breakfest"`
	SindicateCarePackage		CreateSindicateCarePackage		`json:"sindicate_care_package"`
	SindicateVoucher			CreateSindicateVaucher			`json:"sindicate_voucher"`
}

type GetSindicatesModel struct {
	ID 							int				`json:"id" db:"id"`
	NroRegister					string			`json:"nro_register" db:"nro_register"`
	DateStart					date			`json:"date_start" db:"date_start"`
	DateFinish					date			`json:"date_finish" db:"date_finish"`
}

type CreateSindicate struct {
	NroRegister					string			`json:"nro_register" db:"nro_register"`
	DateStart					date			`json:"date_start" db:"date_start"`
	DateFinish					date			`json:"date_finish" db:"date_finish"`
	FkSindicatePercents			int				`json:"fk_sindicate_percents" db:"fk_sindicate_percentes"`
	FkSindicateValues 			int				`json:"fk_sindicate_values" db:"fk_sindicate_values"`
	FkSindicateAuthorization	int				`json:"fk_sindicate_authorization" db:"fk_sindicate_authorization"`
	FkSindicateBreakfest		int				`json:"fk_sindicate_breakfest" db:"fk_sindicate_breakfest"`
	FkSindicateCarePackage		int				`json:"fk_sindicate_care_package db:"fk_sindicate_care_package`
	FkSindicateVoucher			int				`json:"fk_sindicate_voucher db:"fk_sindicate_voucher`
}

type EditSindicate struct {
	NroRegister					*string			`json:"nro_register" db:"nro_register"`
	DateStart					*date			`json:"date_start" db:"date_start"`
	DateFinish					*date			`json:"date_finish" db:"date_finish"`
}

type SindicatePercents struct {
	ID 							int				`json:"id" db:"id"`
	ExtraHourWeek				int				`json:"extra_hour_week" db:"extra_hour_week"`
	ExtraHourSaturday			int				`json:"extra_hour_saturday" db:"extra_hour_saturday"`
	ExtraHourSundayHolidays		int				`json:"extra_hour_sunday_holidays" db:"extra_hour_sunday_holidays"`
	SindicateContr				int				`json:"sindicate_contr" db:"sindicate_contr"`
	Seconci						int				`json:"seconci" db:"seconci"`
	SalaryAdvancePercents		int				`json:"salary_advance_percents" db:"salary_advance_percents"`
	IncentiveQualify			int				`json:"incentive_qualify" db:"incentive_qualify"`
	ComplemRetirement			int				`json:"complement_retirement" db:"complement_retirement"`
	NightTime					int				`json:"night_time" db:"night_time"`	
}

type CreateSindicatePercents struct {
	ExtraHourWeek				int				`json:"extra_hour_week" db:"extra_hour_week"`
	ExtraHourSaturday			int				`json:"extra_hour_saturday" db:"extra_hour_saturday"`
	ExtraHourSundayHolidays		int				`json:"extra_hour_sunday_holidays" db:"extra_hour_sunday_holidays"`
	SindicateContr				int				`json:"sindicate_contr" db:"sindicate_contr"`
	Seconci						int				`json:"seconci" db:"seconci"`
	SalaryAdvancePercents		int				`json:"salary_advance_percents" db:"salary_advance_percents"`
	IncentiveQualify			int				`json:"incentive_qualify" db:"incentive_qualify"`
	ComplemRetirement			int				`json:"complement_retirement" db:"complement_retirement"`
	NightTime					int				`json:"night_time" db:"night_time"`	
}

type EditSindicatePercents struct {
	ExtraHourWeek				*int			`json:"extra_hour_week" db:"extra_hour_week"`
	ExtraHourSaturday			*int			`json:"extra_hour_saturday" db:"extra_hour_saturday"`
	ExtraHourSundayHolidays		*int			`json:"extra_hour_sunday_holidays" db:"extra_hour_sunday_holidays"`
	SindicateContr				*int			`json:"sindicate_contr" db:"sindicate_contr"`
	Seconci						*int			`json:"seconci" db:"seconci"`
	SalaryAdvancePercents		*int			`json:"salary_advance_percents" db:"salary_advance_percents"`
	IncentiveQualify			*int			`json:"incentive_qualify" db:"incentive_qualify"`
	ComplemRetirement			*int			`json:"complement_retirement" db:"complement_retirement"`
	NightTime					*int			`json:"night_time" db:"night_time"`	
}

type SindicateValues struct {
	ID							int				`json:"id" db:"id"`
	Attendance					int				`json:"attendance" db:"attendance"`
	BaseSalary					int				`json:"base_salary" db:"base_salary"`
	HealthInsurance				int				`json:"health_insurance" db:"health_insurance"`
	ChristmasBonus				int				`json:"christmas_bonus" db:"chrismas_bonus"`
	AssociativeContribution		int				`json:"associative_contribution" db:"associative_contribution"`
	ChildcareAssistance			int				`json:"childcare_assistance" db:"chilcare_assistance"`
	EducationAid				int				`json:"education_aid" db:"education_aid"`
	Unhealthness				int				`json:"unhealthness" db:"unhealthness"`
	Dangerousness				int				`json:"dangerousness" db:"dangerousness"`
	NightTimeSupplement			int				`json:"nighttime_supplement" db:"nighttime_supplement"`
}

type CreateSindicateValues struct {
	Attendance					int				`json:"attendance" db:"attendance"`
	BaseSalary					int				`json:"base_salary" db:"base_salary"`
	HealthInsurance				int				`json:"health_insurance" db:"health_insurance"`
	ChristmasBonus				int				`json:"christmas_bonus" db:"chrismas_bonus"`
	AssociativeContribution		int				`json:"associative_contribution" db:"associative_contribution"`
	ChildcareAssistance			int				`json:"childcare_assistance" db:"chilcare_assistance"`
	EducationAid				int				`json:"education_aid" db:"education_aid"`
	Unhealthness				int				`json:"unhealthness" db:"unhealthness"`
	Dangerousness				int				`json:"dangerousness" db:"dangerousness"`
	NightTimeSupplement			int				`json:"nighttime_supplement" db:"nighttime_supplement"`
}

type EditSindicateValues struct {
	Attendance					*int			`json:"attendance" db:"attendance"`
	BaseSalary					*int			`json:"base_salary" db:"base_salary"`
	HealthInsurance				*int			`json:"health_insurance" db:"health_insurance"`
	ChristmasBonus				*int			`json:"christmas_bonus" db:"chrismas_bonus"`
	AssociativeContribution		*int			`json:"associative_contribution" db:"associative_contribution"`
	ChildcareAssistance			*int			`json:"childcare_assistance" db:"chilcare_assistance"`
	EducationAid				*int			`json:"education_aid" db:"education_aid"`
	Unhealthness				*int			`json:"unhealthness" db:"unhealthness"`
	Dangerousness				*int			`json:"dangerousness" db:"dangerousness"`
	NightTimeSupplement			*int			`json:"nighttime_supplement" db:"nighttime_supplement"`
}

type SindicateVoucher struct {
	ID								int				`json:"id" db:"id"`
	AbscenceDiscountVAVR			bool 			`json:"abscence_discount_va_vr" db:"abscence_discount_va"`
	AbscenceDiscountVT				bool			`json:"abscence_discount_vt" db:"abscence_discount_vr"`
	MedCertificateDiscountVAVR		bool			`json:"med_certificate_dicount_va_vr" db:"med_certificate_discount_va"`
	MedCertificateDiscountVt		bool			`json:"med_certificate_discount_vt" db:"med_certificate_discount_vt"`
	VacationDiscountVAVR			bool			`json:"vacation_discount_va_vr" db:"vacation_discount_va_vr"`
	VacationDiscountVT				bool			`json:"vacation_discount_vt" db:"vacation_discount_vt"`
	SickLeavetDiscountVAVRINSS		bool			`json:"sick_leave_dicount_va_vr_inss" db:"sick_leave_discount_va_vr_inss"`
	SickLeaveDiscountVTINSS			bool			`json:"sick_leave_discount_vt_inss" db:"sick_leave_discount_vt_inss"`
	ValueAlimDay					int				`json:"value_alim_day" db:"value_alim_day"`
	DeductLimitByAllowanceonPayroll	int				`json:"deduct_limit_by_allowance_on_payroll" db:"deduc_limit_by_allowance_on_payroll"`
	PercentualDeduct				int				`json:"percentual_deduct" by:"percentual_deduct"`
	DeductLimitByVAVRonPayroll		int				`json:"deduct_limit_by_va_vr_on_payroll" db:"deduct_limit_by_va_vr_on_payroll"`
	DiscountVtPerc					int				`json:"discount_vt_perc" db:"discount_vt_perc"`
}

type CreateSindicateVaucher struct {
	AbscenceDiscountVAVR			bool 			`json:"abscence_discount_va_vr" db:"abscence_discount_va"`
	AbscenceDiscountVT				bool			`json:"abscence_discount_vt" db:"abscence_discount_vr"`
	MedCertificateDiscountVAVR		bool			`json:"med_certificate_dicount_va_vr" db:"med_certificate_discount_va"`
	MedCertificateDiscountVt		bool			`json:"med_certificate_discount_vt" db:"med_certificate_discount_vt"`
	VacationDiscountVAVR			bool			`json:"vacation_discount_va_vr" db:"vacation_discount_va_vr"`
	VacationDiscountVT				bool			`json:"vacation_discount_vt" db:"vacation_discount_vt"`
	SickLeavetDiscountVAVRINSS		bool			`json:"sick_leave_dicount_va_vr_inss" db:"sick_leave_discount_va_vr_inss"`
	SickLeaveDiscountVTINSS			bool			`json:"sick_leave_discount_vt_inss" db:"sick_leave_discount_vt_inss"`
	ValueAlimDay					int				`json:"value_alim_day" db:"value_alim_day"`
	DeductLimitByAllowanceonPayroll	int				`json:"deduct_limit_by_allowance_on_payroll" db:"deduc_limit_by_allowance_on_payroll"`
	PercentualDeduct				int				`json:"percentual_deduct" by:"percentual_deduct"`
	DeductLimitByVAVRonPayroll		int				`json:"deduct_limit_by_va_vr_on_payroll" db:"deduct_limit_by_va_vr_on_payroll"`
	DiscountVtPerc					int				`json:"discount_vt_perc" db:"discount_vt_perc"`
}

type EditSindicateVaucher struct {
	AbscenceDiscountVAVR			*bool 			`json:"abscence_discount_va_vr" db:"abscence_discount_va"`
	AbscenceDiscountVT				*bool			`json:"abscence_discount_vt" db:"abscence_discount_vr"`
	MedCertificateDiscountVAVR		*bool			`json:"med_certificate_dicount_va_vr" db:"med_certificate_discount_va"`
	MedCertificateDiscountVt		*bool			`json:"med_certificate_discount_vt" db:"med_certificate_discount_vt"`
	VacationDiscountVAVR			*bool			`json:"vacation_discount_va_vr" db:"vacation_discount_va_vr"`
	VacationDiscountVT				*bool			`json:"vacation_discount_vt" db:"vacation_discount_vt"`
	SickLeaveDiscountVAVRINSS		*bool			`json:"sick_leave_dicount_va_vr_inss" db:"sick_leave_discount_va_vr_inss"`
	SickLeaveDiscountVTINSS			*bool			`json:"sick_leave_discount_vt_inss" db:"sick_leave_discount_vt_inss"`
	ValueAlimDay					*int			`json:"value_alim_day" db:"value_alim_day"`
	DeductLimitByAllowanceonPayroll	*int			`json:"deduct_limit_by_allowance_on_payroll" db:"deduc_limit_by_allowance_on_payroll"`
	PercentualDeduct				*int			`json:"percentual_deduct" by:"percentual_deduct"`
	DeductLimitByVAVRonPayroll		*int			`json:"deduct_limit_by_va_vr_on_payroll" db:"deduct_limit_by_va_vr_on_payroll"`
	DiscountVtPerc					*int			`json:"discount_vt_perc" db:"discount_vt_perc"`
}

type SindicateBreakfest struct {
	ID								int				`json:"id" db:"id"`
	MedCertificationDisc			bool			`json:"med_certification_disc" bd:"med_certification_disc"`
	MedCertificationDisProp			bool			`json:"med_certification_disc_prop" bd:"med_certification_dic_prop"`
	AbscenceDiscount				bool			`json:"abscence_discount" bd:"abscence_discount"`
	AbscenceDiscountProp			bool			`json:"abscence_discount_prop" bd:"abscence_discount_prop"`
	VacationDiscount				bool			`json:"vacation_discount" bd:"vacation_discount"`
	VacationDiscountProp			bool			`json:"vacation_discount_prop" bd:"vacation_discount_prop"`
	LeaveDiscount					bool			`json:"leave_discount" bd:"leave_discount"`
	LeaveDiscountProp				bool			`json:"leave_discount_prop" bd:"leave_discount_prop"`
	BreakNature						bool			`json:"break_nature" bd:"break_nature"`
	LimitDescPercent				int				`json:"limit_desc_percent" bd:"limit_desc_percent"`
	ValueDay						int				`json:"value_day" bd:"value_day"`
	ValueMonth						int				`json:"value_month" bd:"value_month"`
	LimitDesc						int				`json:"limit_desc" bd:"limit_desc"`
}

type CreateSindicateBreakfest struct {
	MedCertificationDisc			bool			`json:"med_certification_disc" bd:"med_certification_disc"`
	MedCertificationDisProp			bool			`json:"med_certification_disc_prop" bd:"med_certification_dic_prop"`
	AbscenceDiscount				bool			`json:"abscence_discount" bd:"abscence_discount"`
	AbscenceDiscountProp			bool			`json:"abscence_discount_prop" bd:"abscence_discount_prop"`
	VacationDiscount				bool			`json:"vacation_discount" bd:"vacation_discount"`
	VacationDiscountProp			bool			`json:"vacation_discount_prop" bd:"vacation_discount_prop"`
	LeaveDiscount					bool			`json:"leave_discount" bd:"leave_discount"`
	LeaveDiscountProp				bool			`json:"leave_discount_prop" bd:"leave_discount_prop"`
	BreakNature						bool			`json:"break_nature" bd:"break_nature"`
	LimitDescPercent				int				`json:"limit_desc_percent" bd:"limit_desc_percent"`
	ValueDay						int				`json:"value_day" bd:"value_day"`
	ValueMonth						int				`json:"value_month" bd:"value_month"`
	LimitDesc						int				`json:"limit_desc" bd:"limit_desc"`
}

type EditSindicateBreakfest struct {
	MedCertificationDisc			*bool			`json:"med_certification_disc" bd:"med_certification_disc"`
	MedCertificationDiscProp		*bool			`json:"med_certification_disc_prop" bd:"med_certification_dic_prop"`
	AbscenceDiscount				*bool			`json:"abscence_discount" bd:"abscence_discount"`
	AbscenceDiscountProp			*bool			`json:"abscence_discount_prop" bd:"abscence_discount_prop"`
	VacationDiscount				*bool			`json:"vacation_discount" bd:"vacation_discount"`
	VacationDiscountProp			*bool			`json:"vacation_discount_prop" bd:"vacation_discount_prop"`
	LeaveDiscount					*bool			`json:"leave_discount" bd:"leave_discount"`
	LeaveDiscountProp				*bool			`json:"leave_discount_prop" bd:"leave_discount_prop"`
	BreakNature						*bool			`json:"break_nature" bd:"break_nature"`
	LimitDescPercent				*int			`json:"limit_desc_percent" bd:"limit_desc_percent"`
	ValueDay						*int			`json:"value_day" bd:"value_day"`
	ValueMonth						*int			`json:"value_month" bd:"value_month"`
	LimitDesc						*int			`json:"limit_desc" bd:"limit_desc"`
}

type SindicateAuthorization struct {
	ID						int					`json:"id" bd:"id"`
	NormalHour				boolean				`json:"normal_hour" bd:"normal_hour"`
	SalaryAdvance			boolean				`json:"salary_advance" bd:"salary_advance"`
	BankHours				boolean				`json:"bank_hours" bd:"bank_hours"`
	ShiftWork				boolean				`json:"shift_work" bd:"shift_word"`
	SindPayBeneithEmpl		boolean				`json:"sind_pay_beneith_emplo" bd:"sind_pay_beneith_emplo"`
	MobilePointAppApproval	boolean				`json:"mobile_point_app_approval" bd:"mobile_point_app_approval"`
}

type CreateSindicateAuthorization struct {
	NormalHour				boolean				`json:"normal_hour" bd:"normal_hour"`
	SalaryAdvance			boolean				`json:"salary_advance" bd:"salary_advance"`
	BankHours				boolean				`json:"bank_hours" bd:"bank_hours"`
	ShiftWork				boolean				`json:"shift_work" bd:"shift_word"`
	SindPayBeneithEmpl		boolean				`json:"sind_pay_beneith_emplo" bd:"sind_pay_beneith_emplo"`
	MobilePointAppApproval	boolean				`json:"mobile_point_app_approval" bd:"mobile_point_app_approval"`
}

type EditSindicateAuthorization struct {
	NormalHour				*boolean			`json:"normal_hour" bd:"normal_hour"`
	SalaryAdvance			*boolean			`json:"salary_advance" bd:"salary_advance"`
	BankHours				*boolean			`json:"bank_hours" bd:"bank_hours"`
	ShiftWork				*boolean			`json:"shift_work" bd:"shift_word"`
	SindPayBeneithEmpl		*boolean			`json:"sind_pay_beneith_emplo" bd:"sind_pay_beneith_emplo"`
	MobilePointAppApproval	*boolean			`json:"mobile_point_app_approval" bd:"mobile_point_app_approval"`
}

type SindicateCarePackage struct {
	ID 						int				`json:"id" db:"id"`
	DiscountBasic			int				`json:"discount_basic" db:"discount_basic"`
	ValueBasic				int				`json:"value_basic" db:"value_basic"`
	LimitDiscBasic			int				`json:"limit_disc_basic" db:"limit_disc_basic"`
	AbscenceDiscBasic		bool			`json:"abscence_disc_basic" db:"abscence_disc_basic"`
	LeaveDiscBasic			bool			`json:"leave_disc_basic" db:"leave_disc_basic"`
	VacationDisc			bool			`json:"vacation_disc" db:"vacation_disc"`
	VacationDiscountProp	bool			`json:"vacation_discount_prop" db:"vacation_discount_prop"`	
	GetBasicWithOneDay		bool			`json:"get_basic_with_one_day" db:"get_basic_with_one_day"`
	MagneticCard			bool			`json:"magnetic_card" db:"magnetic_card"`
	MissesDayPropDisc		bool			`json:"misses_day_prop_disc" db:"misses_day_prop_disc"`
	LeavePropDisc			bool			`json:"leave_prop_disc" db:"leave_prop_disc"`
	MedCertificateDisc		bool			`json:"med_certificate_disc" db:"med_certificate_disc"`
	MedCertificatePropDisc	bool			`json:"med_certificate_prop_disc" db:"med_certificate_prop_disc"`
	BasicNature				bool			`json:"basic_nature" db:"basic_nature"`
}

type CreateSindicateCarePackage struct {
	DiscountBasic			int				`json:"discount_basic" db:"discount_basic"`
	ValueBasic				int				`json:"value_basic" db:"value_basic"`
	LimitDiscBasic			int				`json:"limit_disc_basic" db:"limit_disc_basic"`
	AbscenceDiscBasic		bool			`json:"abscence_disc_basic" db:"abscence_disc_basic"`
	LeaveDiscBasic			bool			`json:"leave_disc_basic" db:"leave_disc_basic"`
	VacationDisc			bool			`json:"vacation_disc" db:"vacation_disc"`
	VacationDiscountProp	bool			`json:"vacation_discount_prop" db:"vacation_discount_prop"`	
	GetBasicWithOneDay		bool			`json:"get_basic_with_one_day" db:"get_basic_with_one_day"`
	MagneticCard			bool			`json:"magnetic_card" db:"magnetic_card"`
	MissesDayPropDisc		bool			`json:"misses_day_prop_disc" db:"misses_day_prop_disc"`
	LeavePropDisc			bool			`json:"leave_prop_disc" db:"leave_prop_disc"`
	MedCertificateDisc		bool			`json:"med_certificate_disc" db:"med_certificate_disc"`
	MedCertificatePropDisc	bool			`json:"med_certificate_prop_disc" db:"med_certificate_prop_disc"`
	BasicNature				bool			`json:"basic_nature" db:"basic_nature"`
}

type EditSindicateCarePackage struct {
	DiscountBasic			*int			`json:"discount_basic" db:"discount_basic"`
	ValueBasic				*int			`json:"value_basic" db:"value_basic"`
	LimitDiscBasic			*int			`json:"limit_disc_basic" db:"limit_disc_basic"`
	AbscenceDiscBasic		*bool			`json:"abscence_disc_basic" db:"abscence_disc_basic"`
	LeaveDiscBasic			*bool			`json:"leave_disc_basic" db:"leave_disc_basic"`
	VacationDisc			*bool			`json:"vacation_disc" db:"vacation_disc"`
	VacationDiscountProp	*bool			`json:"vacation_discount_prop" db:"vacation_discount_prop"`	
	GetBasicWithOneDay		*bool			`json:"get_basic_with_one_day" db:"get_basic_with_one_day"`
	MagneticCard			*bool			`json:"magnetic_card" db:"magnetic_card"`
	MissesDayPropDisc		*bool			`json:"misses_day_prop_disc" db:"misses_day_prop_disc"`
	LeavePropDisc			*bool			`json:"leave_prop_disc" db:"leave_prop_disc"`
	MedCertificateDisc		*bool			`json:"med_certificate_disc" db:"med_certificate_disc"`
	MedCertificatePropDisc	*bool			`json:"med_certificate_prop_disc" db:"med_certificate_prop_disc"`
	BasicNature				*bool			`json:"basic_nature" db:"basic_nature"`

}