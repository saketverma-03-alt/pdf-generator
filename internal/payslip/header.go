package payslip

import (
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

/**
 * Header can have three layout, it contans LOGO ,ADDRESS and Heading
 * Logo can be placed at L,R or M, and accordingly address can be placed at R,L,M
 * L is left, R is right, M is middle
 */
func getHeaderRows(pc *PayslipConfig) []core.Row {
	// |logo 6||Address 6|
	var address string
	if pc.Address != "" {
		address = pc.Address
	}
	var rows []core.Row

	switch pc.LogoPlacment {

	case "left":
		newRow := row.New(30)
		newRow.Add(
			image.NewFromFileCol(4, pc.LogoUrl),
			text.NewCol(8, address, props.Text{
				Align: align.Right,
			}),
		)
		rows = append(rows, newRow)

	case "right":
		newRow := row.New(30)
		newRow.Add(
			text.NewCol(8, address),
			image.NewFromFileCol(4, pc.LogoUrl),
		)
		rows = append(rows, newRow)

	case "middle":
		imageRow := row.New(15).Add(
			image.NewFromFileCol(4, pc.LogoUrl),
			text.NewCol(8, address),
		)
		rows = append(rows, imageRow)
		addressRow := row.New(15).Add(
			text.NewCol(8, address, props.Text{
				Align: align.Center,
			}),
		)
		rows = append(rows, addressRow)
	}
	// TODO add default case

	// Add row for any additonal Headings
	// ..

	return rows
}
