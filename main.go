package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	top := row.New().Add(
		col.New(4).Add(text.New("Logo Will Be Here")),
		col.New(4),
		col.New(4).Add(text.New("Address will be here")),
	)

	m.AddRows(top)
	document, err := m.Generate()
	if err != nil {
		log.Fatal("Error Generating Pdf")
	}
	document.Save("./pdf/test.pdf")
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().WithPageSize(pagesize.A4).WithTopMargin(15).WithLeftMargin(10).WithRightMargin(10).WithBottomMargin(4).Build()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	m.RegisterFooter(GetPageFooter())

	return m
}

func GetPageFooter() core.Row {
	return row.New().WithStyle(&props.Cell{
		BorderType:  border.Top,
		BorderColor: &props.BlackColor,
	}).Add(
		col.New(12).Add(
			text.New("This is a computer generated payslip", props.Text{
				Top:   4,
				Align: align.Center,
			}),
		),
	)
}
