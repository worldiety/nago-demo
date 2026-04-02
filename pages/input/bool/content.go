package bool

import (
	"github.com/worldiety/nago-demo/layout"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(wnd core.Window) core.View {
	cols := 2
	if wnd.Info().SizeClass < core.SizeClassLarge {
		cols = 1
	}

	return ui.Grid(
		ui.GridCell(tableCheckbox(wnd)),
		ui.GridCell(tableToggle(wnd)),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tableCheckbox(wnd core.Window) core.View {
	stateCheckboxDefault := core.StateOf[bool](wnd, "stateCheckboxDefault")
	stateCheckboxDisabled := core.StateOf[bool](wnd, "stateCheckboxDisabled")
	stateCheckboxFieldDefault := core.StateOf[bool](wnd, "stateCheckboxFieldDefault")
	stateCheckboxFieldSupport := core.StateOf[bool](wnd, "stateCheckboxFieldSupport")
	stateCheckboxFieldError := core.StateOf[bool](wnd, "stateCheckboxFieldError")
	stateCheckboxFieldDisabled := core.StateOf[bool](wnd, "stateCheckboxFieldDisabled")

	stateCheckboxDisabled.Set(true)

	return table("Checkbox",
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
			Component: ui.CheckboxField("Standard", stateCheckboxFieldDefault.Get()).
				InputValue(stateCheckboxFieldDefault),
			Value: boolToString(stateCheckboxFieldDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.CheckboxField("Support", stateCheckboxFieldSupport.Get()).
				InputValue(stateCheckboxFieldSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: boolToString(stateCheckboxFieldSupport.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.CheckboxField("Fehler", stateCheckboxFieldError.Get()).
				InputValue(stateCheckboxFieldError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: boolToString(stateCheckboxFieldError.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.CheckboxField("Disabled", stateCheckboxFieldDisabled.Get()).
				InputValue(stateCheckboxFieldDisabled).
				Disabled(true),
			Value: boolToString(stateCheckboxFieldDisabled.Get()),
		},
	)
}

func tableToggle(wnd core.Window) core.View {
	stateToggleDefault := core.StateOf[bool](wnd, "stateToggleDefault")
	stateToggleDisabled := core.StateOf[bool](wnd, "stateToggleDisabled")
	stateToggleFieldDefault := core.StateOf[bool](wnd, "stateToggleFieldDefault")
	stateToggleFieldSupport := core.StateOf[bool](wnd, "stateToggleFieldSupport")
	stateToggleFieldError := core.StateOf[bool](wnd, "stateToggleFieldError")
	stateToggleFieldDisabled := core.StateOf[bool](wnd, "stateToggleFieldDisabled")

	stateToggleDisabled.Set(true)
	stateToggleFieldDisabled.Set(false)

	return table("Toggle",
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
			Component: ui.ToggleField("Standard", stateToggleFieldDefault.Get()).
				InputValue(stateToggleFieldDefault),
			Value: boolToString(stateToggleFieldDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.ToggleField("Support", stateToggleFieldSupport.Get()).
				InputValue(stateToggleFieldSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: boolToString(stateToggleFieldSupport.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.ToggleField("Fehler", stateToggleFieldError.Get()).
				InputValue(stateToggleFieldError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: boolToString(stateToggleFieldError.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.ToggleField("Disabled", stateToggleFieldDisabled.Get()).
				InputValue(stateToggleFieldDisabled).
				Disabled(true),
			Value: boolToString(stateToggleFieldDisabled.Get()),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}

func boolToString(value bool) string {
	if value {
		return "true"
	}

	return "false"
}
