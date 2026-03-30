package bool

import (
	"github.com/worldiety/nago-demo/layout"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"Bool'sche Werte",
		"",
		table(wnd),
	)
}

func table(wnd core.Window) core.View {
	stateCheckboxDefault := core.StateOf[bool](wnd, "stateCheckboxDefault")
	stateCheckboxDisabled := core.StateOf[bool](wnd, "stateCheckboxDisabled")
	stateCheckboxFieldDefault := core.StateOf[bool](wnd, "stateCheckboxFieldDefault")
	stateCheckboxFieldSupport := core.StateOf[bool](wnd, "stateCheckboxFieldSupport")
	stateCheckboxFieldError := core.StateOf[bool](wnd, "stateCheckboxFieldError")
	stateCheckboxFieldDisabled := core.StateOf[bool](wnd, "stateCheckboxFieldDisabled")
	stateToggleDefault := core.StateOf[bool](wnd, "stateToggleDefault")
	stateToggleDisabled := core.StateOf[bool](wnd, "stateToggleDisabled")
	stateToggleFieldDefault := core.StateOf[bool](wnd, "stateToggleFieldDefault")
	stateToggleFieldSupport := core.StateOf[bool](wnd, "stateToggleFieldSupport")
	stateToggleFieldError := core.StateOf[bool](wnd, "stateToggleFieldError")

	stateCheckboxDisabled.Set(true)
	stateToggleDisabled.Set(true)

	return ui.HStack(
		layout.ComponentValueTable(
			layout.ComponentValueTableRow{
				Component: ui.Checkbox(stateCheckboxDefault.Get()).
					InputValue(stateCheckboxDefault),
				Value: boolToString(stateCheckboxDefault.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.Checkbox(stateCheckboxDisabled.Get()).
					InputValue(stateCheckboxDisabled).
					Disabled(true),
				Value: boolToString(stateCheckboxDisabled.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.CheckboxField("Checkbox", stateCheckboxFieldDefault.Get()).
					InputValue(stateCheckboxFieldDefault),
				Value: boolToString(stateCheckboxFieldDefault.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.CheckboxField("Checkbox mit Support", stateCheckboxFieldSupport.Get()).
					InputValue(stateCheckboxFieldSupport).
					SupportingText("Hier steht ein Support-Text"),
				Value: boolToString(stateCheckboxFieldSupport.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.CheckboxField("Checkbox mit Fehler", stateCheckboxFieldError.Get()).
					InputValue(stateCheckboxFieldError).
					ErrorText("Hier steht ein Fehler-Text"),
				Value: boolToString(stateCheckboxFieldError.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.CheckboxField("Checkbox disabled", stateCheckboxFieldDisabled.Get()).
					InputValue(stateCheckboxFieldDisabled).
					Disabled(true),
				Value: boolToString(stateCheckboxFieldDisabled.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.Toggle(stateToggleDefault.Get()).
					InputChecked(stateToggleDefault),
				Value: boolToString(stateToggleDefault.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.Toggle(stateToggleDisabled.Get()).
					InputChecked(stateToggleDisabled).
					Disabled(true),
				Value: boolToString(stateToggleDisabled.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.ToggleField("Schalter", stateToggleFieldDefault.Get()).
					InputValue(stateToggleFieldDefault),
				Value: boolToString(stateToggleFieldDefault.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.ToggleField("Schalter mit Support", stateToggleFieldSupport.Get()).
					InputValue(stateToggleFieldSupport).
					SupportingText("Hier steht ein Support-Text"),
				Value: boolToString(stateToggleFieldSupport.Get()),
			},
			layout.ComponentValueTableRow{
				Component: ui.ToggleField("Schalter mit Fehler", stateToggleFieldError.Get()).
					InputValue(stateToggleFieldError).
					ErrorText("Hier steht ein Fehler-Text"),
				Value: boolToString(stateToggleFieldError.Get()),
			},
		),
	)
}

func boolToString(value bool) string {
	if value {
		return "true"
	}

	return "false"
}
