package switcher

import (
	icons "go.wdy.de/nago/presentation/icons/hero/outline"
	"go.wdy.de/nago/presentation/ui/switcher"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(wnd core.Window) core.View {
	pages := getSwitcherPages()
	stateSwitcherDefault := core.StateOf[string](wnd, "stateSwitcherDefault")
	stateSwitcherObjectFit := core.StateOf[string](wnd, "stateSwitcherObjectFit")
	stateSwitcherDynamic := core.StateOf[string](wnd, "stateSwitcherDynamic")

	return ui.Grid(
		ui.GridCell(
			ui.VStack(
				ui.Text("Standard").Font(ui.HeadlineSmall),
				switcher.Switcher(pages, stateSwitcherDefault).
					FullWidth(),
			).Alignment(ui.Top).Gap(ui.L4).FullWidth(),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Image Object-Fit").Font(ui.HeadlineSmall),
				switcher.Switcher(pages, stateSwitcherObjectFit).
					FullWidth().
					ImageObjectFit(ui.FitCover),
			).Alignment(ui.Top).Gap(ui.L4).FullWidth(),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Dynamische Höhe").Font(ui.HeadlineSmall),
				switcher.Switcher(pages, stateSwitcherDynamic).
					FullWidth().
					DynamicHeight(),
			).Alignment(ui.Top).Gap(ui.L4).FullWidth(),
		),
	).Columns(1).Gap(ui.L32).FullWidth().Heights("auto")
}

func getSwitcherPages() []switcher.TSwitcherPage {
	return []switcher.TSwitcherPage{
		switcher.SwitcherPage(
			"switcher_page_1",
			"Switcher Seite 1",
			icons.Banknotes,
			ui.VStack(
				ui.Text("Lorem Ipsum").Font(ui.HeadlineMedium),
				ui.Text("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum.\nLorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum."),
				ui.HStack(
					ui.Button(ui.ButtonStylePrimary, func() {}).Title("Dummy Button").FullWidth(),
				).FullWidth().Padding(ui.Padding{Top: ui.L8}),
			).Alignment(ui.BottomLeading),
		).Img("https://picsum.photos/300/600"),
		switcher.SwitcherPage(
			"switcher_page_2",
			"Switcher Seite 2",
			icons.Heart,
			ui.VStack(
				ui.Text("Lorem Ipsum").Font(ui.HeadlineMedium),
				ui.Text("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."),
				ui.HStack(
					ui.Button(ui.ButtonStylePrimary, func() {}).Title("Dummy Button").FullWidth(),
				).FullWidth().Padding(ui.Padding{Top: ui.L8}),
			).Alignment(ui.BottomLeading),
		).Img("https://picsum.photos/1200/900"),
		switcher.SwitcherPage(
			"switcher_page_3",
			"Switcher Seite 3",
			icons.DocumentText,
			ui.VStack(
				ui.Text("Lorem Ipsum").Font(ui.HeadlineMedium),
				ui.Text("Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."),
				ui.HStack(
					ui.Button(ui.ButtonStylePrimary, func() {}).Title("Dummy Button"),
				).FullWidth().Alignment(ui.Center).Padding(ui.Padding{Top: ui.L8}),
			).Alignment(ui.BottomLeading),
		),
	}
}
