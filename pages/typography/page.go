package typography

import (
	"github.com/worldiety/nago-demo/layout"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"",
		"",
		grid(wnd),
	)
}

func grid(wnd core.Window) core.View {
	cols := 2
	if wnd.Info().SizeClass < core.SizeClassLarge {
		cols = 1
	}

	return ui.Grid(
		ui.GridCell(
			ui.VStack(
				ui.Text("Display Large").Font(ui.DisplayLarge),
				ui.Text("Display Medium").Font(ui.DisplayMedium),
				ui.Text("Display Small").Font(ui.DisplaySmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Headline Large").Font(ui.HeadlineLarge),
				ui.Text("Headline Medium").Font(ui.HeadlineMedium),
				ui.Text("Headline Small").Font(ui.HeadlineSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Title Large").Font(ui.TitleLarge),
				ui.Text("Title Medium").Font(ui.TitleMedium),
				ui.Text("Title Small").Font(ui.TitleSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Label Large").Font(ui.LabelLarge),
				ui.Text("Label Medium").Font(ui.LabelMedium),
				ui.Text("Label Small").Font(ui.LabelSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Body Large: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.BodyLarge),
				ui.Text("Body Medium: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.BodyMedium),
				ui.Text("Body Small: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.BodySmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Mono Large: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoLarge),
				ui.Text("Mono Medium: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoMedium),
				ui.Text("Mono Small: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Mono Bold Large: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoBoldLarge),
				ui.Text("Mono Bold Medium: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoBoldMedium),
				ui.Text("Mono Bold Small: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoBoldSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Mono Italic Large: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoItalicLarge),
				ui.Text("Mono Italic Medium: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoItalicMedium),
				ui.Text("Mono Italic Small: Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua.").Font(ui.MonoItalicSmall),
			).Gap(ui.L8).Alignment(ui.TopLeading),
		),
	).Columns(cols).Gap(ui.L64).Heights("auto").FullWidth()
}
