package layout

import (
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Page(_ core.Window, title, subtitle string, view core.View) core.View {
	return ui.VStack(
		ui.If(len(title) > 0 || len(subtitle) > 0, ui.VStack(
			ui.If(len(title) > 0, ui.Text(title).Font(ui.HeadlineLarge)),
			ui.If(len(subtitle) > 0, ui.Text(subtitle).Font(ui.HeadlineSmall)),
		).Gap(ui.L8)),
		ui.VStack(
			view,
		),
	).Gap(ui.L32).Padding(ui.Padding{}.Vertical(ui.L32))
}
