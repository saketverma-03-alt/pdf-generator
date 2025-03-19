package payslip

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
)

type EmployeePayslip struct {
	MartorInstance core.Maroto
}

// Generate payslip of single employee
func (ep *EmployeePayslip) GenerateEmployeePayslip(employee *EmployeePayrollDetail) error {
	ep.configMaroto() // configure maroto
	// Add headers
	addEmployeeEariningsAndDeductionCols(employee, ep.MartorInstance) // Add Earnings & Dedctions
	// Add Totals
	// Add foot notes
	// Handle Passowrd Protection
	doc, err := ep.MartorInstance.Generate()
	if err != nil {
		panic(err)
	}
	doc.Save("./pdf/test.pdf")
	return nil
}

func (ep *EmployeePayslip) configMaroto() {
	// refrence: https://maroto.io/#/v1/documentation?id=documentation
	cfg := config.NewBuilder().WithPageSize(pagesize.A4).
		WithBottomMargin(15).WithLeftMargin(6).WithRightMargin(6).WithBottomMargin(4).WithCompression(true).
		Build()

	cfg.DefaultFont.Size = 8

	mrt := maroto.New(cfg)
	ep.MartorInstance = maroto.NewMetricsDecorator(mrt)
}
