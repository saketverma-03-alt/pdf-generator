package payslip

import (
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func getTableHeaderStyle() props.Text {
	return props.Text{
		Style: fontstyle.Bold,
	}
}
