package text

import (
	"fmt"
	"strconv"

	"github.com/worldiety/nago-demo/layout"

	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/material/round"
	"go.wdy.de/nago/presentation/ui"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"Text",
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
		ui.GridCell(tableSingleLine(wnd)),
		ui.GridCell(tableMultiLine(wnd)),
		ui.GridCell(tableOther(wnd)),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tableSingleLine(wnd core.Window) core.View {
	stateTextDefault := core.StateOf[string](wnd, "stateText")
	stateTextIconLeft := core.StateOf[string](wnd, "stateTextIconLeft")
	stateTextIconRight := core.StateOf[string](wnd, "stateTextIconRight")
	stateTextSupport := core.StateOf[string](wnd, "stateTextSupport")
	stateTextError := core.StateOf[string](wnd, "stateTextError")
	stateTextDisabled := core.StateOf[string](wnd, "stateTextDisabled")

	stateTextDisabled.Set("Хлеб")

	return table("Einzeilig",
		layout.ComponentValueTableRow{
			Component: ui.TextField("Standard", stateTextDefault.Get()).
				InputValue(stateTextDefault),
			Value: stateTextDefault.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Icon links", stateTextIconLeft.Get()).
				InputValue(stateTextIconLeft).
				Leading(ui.ImageIcon(icons.People_outline)),
			Value: stateTextIconLeft.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Icon rechts", stateTextIconRight.Get()).
				InputValue(stateTextIconRight).
				Trailing(ui.ImageIcon(icons.Help_outline)),
			Value: stateTextIconRight.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Support", stateTextSupport.Get()).
				InputValue(stateTextSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: stateTextSupport.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Fehler", stateTextError.Get()).
				InputValue(stateTextError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: stateTextError.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Disabled", stateTextDisabled.Get()).
				InputValue(stateTextDisabled).
				Disabled(true),
			Value: stateTextDisabled.Get(),
		},
	)
}

func tableMultiLine(wnd core.Window) core.View {
	stateTextMultilineDefault := core.StateOf[string](wnd, "stateTextMultilineDefault")
	stateTextMultilineSupport := core.StateOf[string](wnd, "stateTextMultilineSupport")
	stateTextMultilineError := core.StateOf[string](wnd, "stateTextMultilineError")
	stateTextMultilineDisabled := core.StateOf[string](wnd, "stateTextMultilineDisabled")

	stateTextMultilineDisabled.Set("Хлеб и пельмени и манты и чебуреки и вареники и пиво")

	return table("Mehrzeilig",
		layout.ComponentValueTableRow{
			Component: ui.TextField("Standard", stateTextMultilineDefault.Get()).
				InputValue(stateTextMultilineDefault).
				Lines(3),
			Value: stateTextMultilineDefault.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Support", stateTextMultilineSupport.Get()).
				InputValue(stateTextMultilineSupport).
				Lines(3).
				SupportingText("Hier steht ein Support-Text"),
			Value: stateTextMultilineSupport.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Fehler", stateTextMultilineError.Get()).
				InputValue(stateTextMultilineError).
				Lines(3).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: stateTextMultilineError.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.TextField("Disabled", stateTextMultilineDisabled.Get()).
				InputValue(stateTextMultilineDisabled).
				Lines(3).
				Disabled(true),
			Value: stateTextMultilineDisabled.Get(),
		},
	)
}

func tableOther(wnd core.Window) core.View {
	statePasswordDefault := core.StateOf[string](wnd, "statePasswordDefault")
	statePasswordSupport := core.StateOf[string](wnd, "statePasswordSupport")
	statePasswordError := core.StateOf[string](wnd, "statePasswordError")
	statePasswordDisabled := core.StateOf[string](wnd, "statePasswordDisabled")
	stateIntDefault := core.StateOf[int64](wnd, "stateIntDefault")
	stateFloatDefault := core.StateOf[float64](wnd, "stateFloatDefault")

	statePasswordDisabled.Set("Хлеб1234!")

	return table("Weitere",
		layout.ComponentValueTableRow{
			Component: ui.PasswordField("Passwort", statePasswordDefault.Get()).
				InputValue(statePasswordDefault),
			Value: statePasswordDefault.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.PasswordField("Passwort mit Support", statePasswordSupport.Get()).
				InputValue(statePasswordSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: statePasswordSupport.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.PasswordField("Passwort mit Fehler", statePasswordError.Get()).
				InputValue(statePasswordError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: statePasswordError.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.PasswordField("Passwort disabled", statePasswordDisabled.Get()).
				InputValue(statePasswordDisabled).
				Disabled(true),
			Value: statePasswordDisabled.Get(),
		},
		layout.ComponentValueTableRow{
			Component: ui.IntField("Zahlen", stateIntDefault.Get(), stateIntDefault),
			Value:     strconv.FormatInt(stateIntDefault.Get(), 10),
		},
		layout.ComponentValueTableRow{
			Component: ui.FloatField("Dezimal", stateFloatDefault.Get(), stateFloatDefault),
			Value:     fmt.Sprintf("%f", stateFloatDefault.Get()),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}
