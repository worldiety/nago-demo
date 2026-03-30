package layout

import (
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

type ComponentValueTableRow struct {
	Component core.View
	Value     string
}

func ComponentValueTable(rows ...ComponentValueTableRow) core.View {
	return ui.Table(
		ui.TableColumn(ui.Text("Komponente")),
		ui.TableColumn(ui.Text("Wert")),
	).Rows(
		ui.ForEach(rows, func(row ComponentValueTableRow) ui.TTableRow {
			return ui.TableRow(
				inputCell(row.Component),
				valueCell(row.Value),
			)
		})...,
	)
}

func inputCell(view core.View) ui.TTableCell {
	return ui.TableCell(
		view,
	)
}

func valueCell(value string) ui.TTableCell {
	return ui.TableCell(
		ui.VStack(
			ui.Text(value),
		).Alignment(ui.Leading).Frame(ui.Frame{MinWidth: ui.L160, MaxWidth: ui.L320}),
	)
}
