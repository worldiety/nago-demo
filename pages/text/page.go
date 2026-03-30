package text

import (
	"fmt"
	"github.com/worldiety/nago-demo/layout"
	"strconv"

	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/material/round"
	"go.wdy.de/nago/presentation/ui"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"Text",
		"",
		table(wnd),
	)
}

func table(wnd core.Window) core.View {
	stateTextDefault := core.StateOf[string](wnd, "stateText")
	stateTextIconLeft := core.StateOf[string](wnd, "stateTextIconLeft")
	stateTextIconRight := core.StateOf[string](wnd, "stateTextIconRight")
	stateTextSupport := core.StateOf[string](wnd, "stateTextSupport")
	stateTextError := core.StateOf[string](wnd, "stateTextError")
	stateTextDisabled := core.StateOf[string](wnd, "stateTextDisabled")
	statePasswordDefault := core.StateOf[string](wnd, "statePasswordDefault")
	statePasswordSupport := core.StateOf[string](wnd, "statePasswordSupport")
	statePasswordError := core.StateOf[string](wnd, "statePasswordError")
	statePasswordDisabled := core.StateOf[string](wnd, "statePasswordDisabled")
	stateIntDefault := core.StateOf[int64](wnd, "stateIntDefault")
	stateFloatDefault := core.StateOf[float64](wnd, "stateFloatDefault")
	stateTextMultilineDefault := core.StateOf[string](wnd, "stateTextMultilineDefault")
	stateTextMultilineSupport := core.StateOf[string](wnd, "stateTextMultilineSupport")
	stateTextMultilineError := core.StateOf[string](wnd, "stateTextMultilineError")
	stateTextMultilineDisabled := core.StateOf[string](wnd, "stateTextMultilineDisabled")

	stateTextDisabled.Set("Хлеб")
	statePasswordDisabled.Set("Хлеб1234!")
	stateTextMultilineDisabled.Set("Хлеб и пельмени и манты и чебуреки и вареники и пиво")

	return ui.HStack(
		layout.ComponentValueTable(
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld", stateTextDefault.Get()).
					InputValue(stateTextDefault),
				Value: stateTextDefault.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld Icon links", stateTextIconLeft.Get()).
					InputValue(stateTextIconLeft).
					Leading(ui.ImageIcon(icons.People_outline).FillColor(ui.M8)),
				Value: stateTextIconLeft.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld Icon rechts", stateTextIconRight.Get()).
					InputValue(stateTextIconRight).
					Trailing(ui.ImageIcon(icons.Help_outline).FillColor(ui.M8)),
				Value: stateTextIconRight.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mit Support", stateTextSupport.Get()).
					InputValue(stateTextSupport).
					SupportingText("Hier steht ein Support-Text"),
				Value: stateTextSupport.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mit Fehler", stateTextError.Get()).
					InputValue(stateTextError).
					ErrorText("Hier steht ein Fehler-Text"),
				Value: stateTextError.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld disabled", stateTextDisabled.Get()).
					InputValue(stateTextDisabled).
					Disabled(true),
				Value: stateTextDisabled.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.PasswordField("Passwortfeld", statePasswordDefault.Get()).
					InputValue(statePasswordDefault),
				Value: statePasswordDefault.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.PasswordField("Passwortfeld mit Support", statePasswordSupport.Get()).
					InputValue(statePasswordSupport).
					SupportingText("Hier steht ein Support-Text"),
				Value: statePasswordSupport.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.PasswordField("Passwortfeld mit Fehler", statePasswordError.Get()).
					InputValue(statePasswordError).
					ErrorText("Hier steht ein Fehler-Text"),
				Value: statePasswordError.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.PasswordField("Passwortfeld disabled", statePasswordDisabled.Get()).
					InputValue(statePasswordDisabled).
					Disabled(true),
				Value: statePasswordDisabled.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.IntField("Zahlenfeld", stateIntDefault.Get(), stateIntDefault),
				Value:     strconv.FormatInt(stateIntDefault.Get(), 10),
			},
			layout.ComponentValueTableRow{
				Component: ui.FloatField("Dezimalfeld", stateFloatDefault.Get(), stateFloatDefault),
				Value:     fmt.Sprintf("%f", stateFloatDefault.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mehrzeilig", stateTextMultilineDefault.Get()).
					InputValue(stateTextMultilineDefault).
					Lines(3),
				Value: stateTextMultilineDefault.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mehrzeilig mit Support", stateTextMultilineSupport.Get()).
					InputValue(stateTextMultilineSupport).
					Lines(3).
					SupportingText("Hier steht ein Support-Text"),
				Value: stateTextMultilineSupport.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mehrzeilig mit Fehler", stateTextMultilineError.Get()).
					InputValue(stateTextMultilineError).
					Lines(3).
					ErrorText("Hier steht ein Fehler-Text"),
				Value: stateTextMultilineError.Get(),
			},
			layout.ComponentValueTableRow{
				Component: ui.TextField("Textfeld mehrzeilig disabled", stateTextMultilineDisabled.Get()).
					InputValue(stateTextMultilineDisabled).
					Lines(3).
					Disabled(true),
				Value: stateTextMultilineDisabled.Get(),
			},
		),
	)
}
