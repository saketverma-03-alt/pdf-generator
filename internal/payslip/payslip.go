package payslip

import (
	"fmt"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type EmployeePayslip struct {
	MartorInstance core.Maroto
}

// Generate payslip of single employee
func (ep *EmployeePayslip) GenerateEmployeePayslip(employee *EmployeePayrollDetail) error {
	ep.configMaroto() // configure maroto
	// Add headers
	addEmployeeEariningsAndDeductionCols(employee, ep.MartorInstance) // Add Earnings & Dedctions
	addTotals(employee, ep.MartorInstance)                            // Add Totals
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

func addTotals(employee *EmployeePayrollDetail, martorInstance core.Maroto) core.Row {
	styles := getTableHeaderStyle()
	rowOne := row.New(6).Add(
		text.NewCol(2, "Total Gross Pay", styles),
		text.NewCol(2, fmt.Sprintf("%.2f", employee.Payroll.FixedAmountGrossMonthly), styles),
		text.NewCol(2, fmt.Sprintf("%.2f", employee.Payroll.CalculatedGrossMonthly), styles),
		text.NewCol(2, "Deduction Total", styles),
		text.NewCol(2, fmt.Sprintf("%.2f", employee.Payroll.FixedAmountDeductionMonthly), styles),
		text.NewCol(2, fmt.Sprintf("%.2f", employee.Payroll.CalculatedDeductionMonthly), styles),
	).WithStyle(&props.Cell{
		BorderType:      "tb",
		BorderThickness: 0.1,
		BorderColor:     &props.Color{Red: 96, Green: 96, Blue: 96},
	})
	rowTwo := row.New(6).Add(
		text.NewCol(2, "Net Pay:", styles),
		text.NewCol(2, fmt.Sprintf("%.2f", employee.Payroll.CalculatedNetPayableMonthly), styles),
	).WithStyle(&props.Cell{
		BorderType:      "tb",
		BorderThickness: 0.1,
		BorderColor:     &props.Color{Red: 96, Green: 96, Blue: 96},
	})
	martorInstance.AddRows(row.New(4), rowOne, rowTwo)
	return rowOne
}
