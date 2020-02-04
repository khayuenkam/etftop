package etftop

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (et *EtfTop) SetUpTable() *tview.Table {
	table := et.table
	table.SetSelectedStyle(tcell.ColorBlack, tcell.ColorLightBlue, tcell.AttrNone)
	table.SetSelectable(true, false)
	table.Select(1, 1).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			et.app.Stop()
		}
	})

	setHeader(table)
	et.RefreshTable()

	return table
}

func (et *EtfTop) RefreshTable() error {
	et.log("refreshTable()")
	var i = 1

	for k, v := range et.State.etfsData {
		et.table.SetCell(i, 0,
			tview.NewTableCell(k).
				SetAlign(tview.AlignCenter))

		et.table.SetCell(i, 1,
			tview.NewTableCell(fmt.Sprintf("   %vEUR   ", v.Last)).
				SetAlign(tview.AlignCenter))
		et.table.SetCell(i, 2,
			tview.NewTableCell(fmt.Sprintf("   %vEUR   ", v.Ask)).
				SetAlign(tview.AlignCenter))
		et.table.SetCell(i, 3,
			tview.NewTableCell(fmt.Sprintf("   %vEUR   ", v.Bid)).
				SetAlign(tview.AlignCenter))
		et.table.SetCell(i, 4,
			tview.NewTableCell(fmt.Sprintf("   %vEUR   ", v.High)).
				SetAlign(tview.AlignCenter))
		et.table.SetCell(i, 5,
			tview.NewTableCell(fmt.Sprintf("   %vEUR   ", v.Low)).
				SetAlign(tview.AlignCenter))
		et.table.SetCell(i, 6,
			tview.NewTableCell(fmt.Sprintf("%v", v.Stueck)).
				SetAlign(tview.AlignCenter))
		i++
	}

	return nil
}

func setHeader(table *tview.Table) {
	headers := []string{"name", "price", "ask_price", "bid_price", "high", "low", "24H_volume"}

	for i, h := range headers {
		table.SetCell(0, i,
			tview.NewTableCell(h).
				SetSelectable(false).
				SetBackgroundColor(tcell.ColorGreen).
				SetTextColor(tcell.ColorBlack).
				SetAlign(tview.AlignCenter))
	}
}
