package payslip

import "time"

// Type for payslip config recived from backend
type PayslipConfig struct {
	LogoPlacment        string `json:"logo_placment"`
	LogoUrl             string `json:"logo_url"`
	Address             string `json:"address"`
	PasswordProtection  string // NOTE: can be PAN,ADDHR,ETC,EMPCODE , should decide how meny digits it needs to be
	IsPasswordProtected bool   `json:"is_password_protected"`

	EmployeePayrollDetail []EmployeePayrollDetail `json:"employee_payroll_detail"`

	EmployeeDetailsSequence   string //TODO: add type ,Need to be deloped in backend. clocked by above data
	RemoveNonApplicableFields bool
	AddComulativeEarningCol   bool
	// TODO: Enquiry show the actial earning ammount per component,
	// is the optonal or it should be eithor or with Comulative Earinging or Actula Earning
	FootNoteOne string `json:"foot_note_one"`
	FootNoteTwo string `json:"foot_note_two"`
}

type Address struct {
	AddressLineOne string `json:"address_line_1"`
	AddressLineTwo string `json:"address_line_2"`
}

type EmployeePayrollDetail struct {
	Month      int             `json:"month"`
	Year       int             `json:"year"`
	Payroll    PayrollInfo     `json:"payroll"`
	Components []ComponentData `json:"component_data"`
}

type ComponentData struct {
	ID                   int       `json:"id"`
	PayrollComponentID   int       `json:"payroll_component_id"`
	ShortName            string    `json:"short_name"`
	ComponentType        string    `json:"component_type"`
	FixedAmountMonthly   float64   `json:"fixed_amount_monthly"`
	EarningAmountMonthly float64   `json:"earning_amount_monthly"`
	CreatedOn            time.Time `json:"created_on"`
	ComponentName        string    `json:"component_name"`
	EmployeePayrollID    int       `json:"employee_payroll_id"`
	ModifiedOn           time.Time `json:"modified_on"`
	IncludeInGross       bool      `json:"include_in_gross"`
	FixedAmountYearly    float64   `json:"fixed_amount_yearly"`
	EarningAmountYearly  float64   `json:"earning_amount_yearly"`
	IsActive             bool      `json:"is_active"`
}

type PayrollInfo struct {
	ID                    int       `json:"id"`
	EmployeeID            int       `json:"employee_id"`
	CompanyID             int       `json:"company_id"`
	PayGroupID            int       `json:"pay_group_id"`
	Year                  int       `json:"year"`
	Month                 int       `json:"month"`
	PaymentCycleStartDate time.Time `json:"payment_cycle_start_date"`
	PaymentCycleEndDate   time.Time `json:"payment_cycle_end_date"`
	SalaryStructureType   *string   `json:"salary_structure_type"` // Changed to pointer to allow null
	SalaryPayOffType      *string   `json:"salary_pay_off_type"`   // Changed to pointer to allow null
	IsPaid                bool      `json:"is_paid"`

	// Days & OT
	TotalLeaveDays     int    `json:"total_leave_days"`
	LopDays            int    `json:"lop_days"`
	MonthDays          int    `json:"month_days"`
	PaidDays           int    `json:"paid_days"`
	OTDuration         string `json:"ot_duration"`
	TotalWeeklyOffDays int    `json:"total_weekly_off_days"`
	TotalHolidayDays   int    `json:"total_holiday_days"`
	TotalWorkingDays   int    `json:"total_working_days"`

	// Calculated Data
	CalculatedGrossMonthly      float64 `json:"calculated_gross_monthly"`
	CalculatedDeductionMonthly  float64 `json:"calculated_deduction_monthly"`
	CalculatedNetPayableMonthly float64 `json:"calculated_net_payable_monthly"`

	// fixed
	FixedAmountNetPayableMonthly float64 `json:"fixed_amount_net_payable_monthly"`
	FixedAmountDeductionMonthly  float64 `json:"fixed_amount_deduction_monthly"`
	FixedAmountGrossMonthly      float64 `json:"fixed_amount_gross_monthly"`
	FixedAmountNetPayableYearly  float64 `json:"fixed_amount_net_payable_yearly"`
	FixedAmountDeductionYearly   float64 `json:"fixed_amount_deduction_yearly"`
	FixedAmountGrossYearly       float64 `json:"fixed_amount_gross_yearly"`

	// meta data
	CreatedBy      string    `json:"created_by"`
	CreatedOn      time.Time `json:"created_on"`
	ModifiedOn     time.Time `json:"modified_on"`
	LastModifiedBy *string   `json:"last_modified_by"` // Changed to pointer to allow null
	IsActive       bool      `json:"is_active"`
}
