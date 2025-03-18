package payslip

// Type for payslip config recived from backend
type PayslipConfig struct {
	LogoPlacment              string `json:"logo_placment"`
	LogoUrl                   string `json:"logo_url"`
	Address                   string `json:"address"`
	EmployeeDetails           string //TODO: add type ,Clerify from backend
	EmployeeDetailsSequence   string //TODO: add type ,Need to be deloped in backend. clocked by above data
	EmployeeEarnings          string //TODO: add type, []{com_name:string,earned:0,fixed:0, cumulative?:0}
	EmployeeDeductions        string //TODO: add type, []{com_name:string,deduction:0,fixed:0}
	PasswordProtection        string // NOTE: can be PAN,ADDHR,ETC,EMPCODE , should decide how meny digits it needs to be
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
