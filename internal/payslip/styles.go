package payslip

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func getTableHeaderStyle() props.Text {
	return props.Text{
		Top:   1,
		Size:  8,
		Style: fontstyle.Bold,
	}
}
