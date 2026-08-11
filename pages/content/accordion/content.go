package accordion

import (
	_ "embed"
	"fmt"

	"go.wdy.de/nago/presentation/ui/accordion"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

//go:embed accordion-content.gohtml
var accordionContent string

//go:embed accordion-content-small.gohtml
var accordionContentSmall string

func Content(wnd core.Window) core.View {
	accordions := make([]core.View, 0)
	for i := range 8 {
		accordions = append(accordions, accordion.Accordion(
			ui.Text(fmt.Sprintf("Accordion %d", i+1)),
			ui.RichText(fmt.Sprintf("Content %d: %s", i+1, accordionContent)),
			core.StateOf[bool](wnd, fmt.Sprintf("accordion_state_%d", i)),
		).FullWidth())
	}

	accordionsSmall := make([]core.View, 0)
	for i := range 8 {
		accordionsSmall = append(accordionsSmall, accordion.Accordion(
			ui.Text(fmt.Sprintf("Accordion (small) %d", i+1)).Font(ui.TitleMedium),
			ui.RichText(fmt.Sprintf("Content %d: %s", i+1, accordionContentSmall)),
			core.StateOf[bool](wnd, fmt.Sprintf("accordion_small_state_%d", i)),
		).Small().HideSeparator().Frame(ui.Frame{MaxWidth: ui.L256}))
	}

	return ui.Grid(
		ui.GridCell(
			ui.VStack(
				accordions...,
			),
		),
		ui.GridCell(
			ui.VStack(
				accordionsSmall...,
			).FullWidth(),
		),
	).Columns(1).Gap(ui.L32).FullWidth().Heights("auto")
}
