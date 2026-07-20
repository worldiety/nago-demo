package layout

import (
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

type ComponentValueTableRow struct {
	Component core.View
	Value     string
	ValueView core.View
}

func ComponentValueTable(rows ...ComponentValueTableRow) core.View {
	return ui.Table(
		ui.TableColumn(ui.Text("Komponente")),
		ui.TableColumn(ui.Text("Wert")),
	).Rows(
		ui.ForEach(rows, func(row ComponentValueTableRow) ui.TTableRow {
			return ui.TableRow(
				inputCell(row.Component),
				ui.TableCell(
					ui.IfElse(len(row.Value) > 0, valueCell(row.Value), valueViewCell(row.ValueView)),
				),
			)
		})...,
	)
}

func inputCell(view core.View) ui.TTableCell {
	return ui.TableCell(
		ui.HStack(
			view,
		),
	)
}

func valueCell(value string) core.View {
	return ui.VStack(
		ui.Text(value),
	).Alignment(ui.Leading).Frame(ui.Frame{MinWidth: ui.L160, MaxWidth: ui.L320})
}

func valueViewCell(view core.View) core.View {
	return ui.VStack(
		view,
	).Alignment(ui.Leading).Frame(ui.Frame{MinWidth: ui.L160, MaxWidth: ui.L320})
}
