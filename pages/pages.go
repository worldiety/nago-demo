package pages

import (
	"fmt"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/dropdown"
)

func HeaderWithSelectFilter(title, filterName string, options []dropdown.Option[string], state *core.State[string]) core.View {
	return ui.Stack(
		ui.Text(title).Font(ui.HeadlineLarge),
		ui.Spacer(),
		ui.HStack(
			ui.Text(fmt.Sprintf("%s:", filterName)).Font(ui.TitleLarge),
			dropdown.Dropdown("", options, state.Get()).
				InputValue(state),
		).
			Alignment(ui.Center).
			Gap(ui.L16).
			BackgroundColor(ui.M2).
			Padding(ui.Padding{}.All(ui.L16)).
			Border(ui.Border{}.Radius(ui.L14)),
	).Gap(ui.L8).FullWidth()
}
