package dialog

import (
	_ "embed"

	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/flowbite/outline"
	"go.wdy.de/nago/presentation/ui"
)

//go:embed dialog-content.gohtml
var dialogContent string

func Content(wnd core.Window) core.View {
	stateDialogDefault := core.StateOf[bool](wnd, "stateDialogDefault")
	stateDialogFooter := core.StateOf[bool](wnd, "stateDialogFooter")

	closeFunc := func(state *core.State[bool]) func() {
		return func() {
			state.Set(false)
		}
	}

	return ui.Grid(
		ui.GridCell(
			ui.VStack(
				ui.Text("Standard").Font(ui.HeadlineSmall),
				ui.PrimaryButton(func() { stateDialogDefault.Set(true) }).Title("Klicken zum Öffnen"),
				ui.If(stateDialogDefault.Get(),
					ui.Modal(
						ui.Dialog(ui.RichText(dialogContent)).
							Title(ui.Text("Standard")).
							TitleX(ui.TertiaryButton(closeFunc(stateDialogDefault)).PreIcon(icons.Close)),
					).OnDismissRequest(closeFunc(stateDialogDefault)),
				),
			).Alignment(ui.Top).Gap(ui.L4).FullWidth(),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Mit Footer").Font(ui.HeadlineSmall),
				ui.PrimaryButton(func() { stateDialogFooter.Set(true) }).Title("Klicken zum Öffnen"),
				ui.If(stateDialogFooter.Get(),
					ui.Modal(
						ui.Dialog(ui.RichText(dialogContent)).
							Title(ui.Text("Mit Footer")).
							TitleX(ui.TertiaryButton(closeFunc(stateDialogFooter)).PreIcon(icons.Close)).
							Footer(
								ui.HStack(
									ui.TertiaryButton(closeFunc(stateDialogFooter)).Title("Abbrechen"),
									ui.Spacer(),
									ui.PrimaryButton(closeFunc(stateDialogFooter)).Title("Speichern"),
								).Gap(ui.L8),
							),
					).OnDismissRequest(closeFunc(stateDialogFooter)),
				),
			).Alignment(ui.Top).Gap(ui.L4).FullWidth(),
		),
	).Columns(1).Gap(ui.L32).FullWidth().Heights("auto")
}
