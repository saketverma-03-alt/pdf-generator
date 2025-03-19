package payslip

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func getCols(earningCmp *ComponentData, deductionCmp *ComponentData) core.Row {
	row := row.New().WithStyle(&props.Cell{
		BorderType:  border.None,
		BorderColor: &props.Color{Red: 200, Green: 200, Blue: 200},
	})
	if earningCmp != nil {
		row.Add(text.NewCol(2, earningCmp.ComponentName, props.Text{
			Size:  6,
			Style: fontstyle.Bold,
		}))
		row.Add(col.New(2).Add(text.New(fmt.Sprintf("%.2f", earningCmp.FixedAmountMonthly))))
		row.Add(col.New(2).Add(text.New(fmt.Sprintf("%.2f", earningCmp.EarningAmountMonthly))))
	} else {
		row.Add(col.New(6))
	}

	if deductionCmp != nil {
		row.Add(col.New(2).Add(text.New((deductionCmp.ComponentName))))
		row.Add(col.New(2).Add(text.New(fmt.Sprintf("%.2f", deductionCmp.FixedAmountMonthly))))
		row.Add(col.New(2).Add(text.New(fmt.Sprintf("%.2f", deductionCmp.EarningAmountMonthly))))
	} else {
		row.Add(col.New(6))
	}

	return row
}

// Add Employee Earnings and Deductions Columns
func addEmployeeEariningsAndDeductionCols(employee *EmployeePayrollDetail, m core.Maroto) error {
	earnings := []ComponentData{}
	deductions := []ComponentData{}

	for _, component := range employee.Components {
		if component.ComponentType == "earnings" {
			earnings = append(earnings, component)
		} else if component.ComponentType == "deductions" {
			deductions = append(deductions, component)
		}
	}

	headerStyles := getTableHeaderStyle()
	m.AddRow(6, text.NewCol(2, "Earnings", headerStyles), text.NewCol(2, "Fixed Amount", headerStyles), text.NewCol(2, "Earning Amount", headerStyles),
		text.NewCol(2, "Deductions", headerStyles), text.NewCol(2, "Fixed Amount", headerStyles), text.NewCol(2, "Deduction Amount", headerStyles)).WithStyle(&props.Cell{
		BorderType:  border.Bottom,
		BorderColor: &props.BlackColor,
	})
	m.AddRow(4)

	// Larger list
	var loopLength int
	earnLen := len(earnings)
	dedLen := len(deductions)
	if earnLen > dedLen {
		loopLength = earnLen
	} else {
		loopLength = dedLen
	}

	for i := 0; i < loopLength; i++ {
		var ernCmp *ComponentData = nil
		var dedCmp *ComponentData = nil
		if i < earnLen {
			ernCmp = &earnings[i]
		}
		if i < dedLen {
			dedCmp = &deductions[i]
		}
		m.AddRows(getCols(ernCmp, dedCmp))
	}

	// row is 12
	// earnings 6 | deductions 6
	// Name Fixed Calcualted | Name Fixed Calculated

	// Add Heading Row
	// m.AddRow(12, col.New(6).Add(text.New("Length Of Earnigns")), col.New(6).Add(text.New(fmt.Sprintf("%d", temp))))
	// m.AddRow(12, col.New(6).Add(text.New("Length Of Dedcution")), col.New(6).Add(text.New(fmt.Sprintf("%d", tempDed))))

	return nil
}
