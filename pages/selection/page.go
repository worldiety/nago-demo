package selection

import (
	"github.com/worldiety/nago-demo/layout"
	"strconv"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"Auswahl",
		"",
		table(wnd),
	)
}

var options = []string{"Joe Schmoe", "Joe Bro", "Joe Dough", "William Daschmoe", "Jøe Schmøe"}

func table(wnd core.Window) core.View {
	stateGroupRadioButtonDefault := ui.AutoRadioStateGroup(wnd, "stateGroupRadioDefault", len(options))
	stateGroupRadioButtonFieldDefault := ui.AutoRadioStateGroup(wnd, "stateGroupRadioButtonFieldDefault", len(options))

	return ui.HStack(
		layout.ComponentValueTable(
			layout.ComponentValueTableRow{
				Component: ui.VStack(
					ui.Each2(stateGroupRadioButtonDefault.All(), func(idx int, checked *core.State[bool]) core.View {
						return ui.RadioButton(checked.Get()).InputChecked(checked).Disabled(idx == 2)
					})...,
				).Gap(ui.L2).Alignment(ui.Leading),
				Value: strconv.Itoa(stateGroupRadioButtonDefault.SelectedIndex()),
			},
			layout.ComponentValueTableRow{
				Component: ui.VStack(
					ui.Each2(stateGroupRadioButtonFieldDefault.All(), func(idx int, checked *core.State[bool]) core.View {
						return ui.RadioButtonField(options[idx], &stateGroupRadioButtonFieldDefault, idx).Disabled(idx == 0)
					})...,
				).Gap(ui.L2).Alignment(ui.Leading),
				Value: getOptionLabel(stateGroupRadioButtonFieldDefault.SelectedIndex()),
			},
		),
	)
}

func getOptionLabel(index int) string {
	if index < 0 || index >= len(options) {
		return ""
	}

	return options[index]
}
