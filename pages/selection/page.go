package selection

import (
	"strconv"
	"strings"

	"github.com/worldiety/nago-demo/layout"
	icons "go.wdy.de/nago/presentation/icons/hero/outline"
	"go.wdy.de/nago/presentation/ui/dropdown"
	"go.wdy.de/nago/presentation/ui/picker"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

type ID string

type Schmoe struct {
	ID   ID
	Name string
}

var names = []string{"Joe Schmoe", "Joe Bro", "Joe Dough", "William Daschmoe", "Jøe Schmøe"}

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"Auswahl",
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
		ui.GridCell(tableRadio(wnd)),
		ui.GridCell(tableDropdown(wnd)),
		ui.GridCell(tablePicker(wnd)),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tableRadio(wnd core.Window) core.View {
	stateGroupRadioButtonDefault := ui.AutoRadioStateGroup(wnd, "stateGroupRadioDefault", len(names))
	stateGroupRadioButtonFieldDefault := ui.AutoRadioStateGroup(wnd, "stateGroupRadioButtonFieldDefault", len(names))

	return table("Radio",
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
					return ui.RadioButtonField(names[idx], &stateGroupRadioButtonFieldDefault, idx).Disabled(idx == 0)
				})...,
			).Gap(ui.L2).Alignment(ui.Leading),
			Value: getOptionLabelByIndex(stateGroupRadioButtonFieldDefault.SelectedIndex()),
		},
	)
}

func tableDropdown(wnd core.Window) core.View {
	selectOptions := make([]dropdown.Option[ID], 0, len(names))
	for i, name := range names {
		selectOptions = append(selectOptions, dropdown.Option[ID]{
			Label:    name,
			Value:    ID(name),
			Disabled: i == 0,
		})
	}

	stateDropdownDefault := core.StateOf[ID](wnd, "stateDropdownDefault")
	stateDropdownIconLeft := core.StateOf[ID](wnd, "stateDropdownIconLeft")
	stateDropdownSupport := core.StateOf[ID](wnd, "stateDropdownSupport")
	stateDropdownError := core.StateOf[ID](wnd, "stateDropdownError")
	stateDropdownDisabled := core.StateOf[ID](wnd, "stateDropdownDisabled")

	stateDropdownDisabled.Set(ID(names[1]))

	return table("Dropdown",
		layout.ComponentValueTableRow{
			Component: dropdown.Dropdown("Standard", selectOptions, stateDropdownDefault.Get()).
				InputValue(stateDropdownDefault),
			Value: getOptionLabelFromSelectOptions(selectOptions, stateDropdownDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: dropdown.Dropdown("Icon links", selectOptions, stateDropdownIconLeft.Get()).
				InputValue(stateDropdownIconLeft).
				Leading(ui.ImageIcon(icons.FaceSmile)),
			Value: getOptionLabelFromSelectOptions(selectOptions, stateDropdownIconLeft.Get()),
		},
		layout.ComponentValueTableRow{
			Component: dropdown.Dropdown("Support", selectOptions, stateDropdownSupport.Get()).
				InputValue(stateDropdownSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: getOptionLabelFromSelectOptions(selectOptions, stateDropdownSupport.Get()),
		},
		layout.ComponentValueTableRow{
			Component: dropdown.Dropdown("Fehler", selectOptions, stateDropdownError.Get()).
				InputValue(stateDropdownError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: getOptionLabelFromSelectOptions(selectOptions, stateDropdownError.Get()),
		},
		layout.ComponentValueTableRow{
			Component: dropdown.Dropdown("Disabled", selectOptions, stateDropdownDisabled.Get()).
				InputValue(stateDropdownDisabled).
				Disabled(true),
			Value: getOptionLabelFromSelectOptions(selectOptions, stateDropdownDisabled.Get()),
		},
	)
}

func tablePicker(wnd core.Window) core.View {
	statePickerDefault := core.StateOf[[]string](wnd, "statePickerDefault")
	statePickerMulti := core.StateOf[[]string](wnd, "statePickerMulti")
	statePickerSupport := core.StateOf[[]string](wnd, "statePickerSupport")
	statePickerError := core.StateOf[[]string](wnd, "statePickerError")
	statePickerDisabled := core.StateOf[[]string](wnd, "statePickerDisabled")

	statePickerDisabled.Set([]string{names[2]})

	return table("Picker",
		layout.ComponentValueTableRow{
			Component: picker.Picker[string]("Standard", names, statePickerDefault),
			Value:     strings.Join(statePickerDefault.Get(), ", "),
		},
		layout.ComponentValueTableRow{
			Component: picker.Picker[string]("Mehrfachauswahl", names, statePickerMulti).
				MultiSelect(true),
			Value: strings.Join(statePickerMulti.Get(), ", "),
		},
		layout.ComponentValueTableRow{
			Component: picker.Picker[string]("Support", names, statePickerSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: strings.Join(statePickerSupport.Get(), ", "),
		},
		layout.ComponentValueTableRow{
			Component: picker.Picker[string]("Fehler", names, statePickerError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: strings.Join(statePickerError.Get(), ", "),
		},
		layout.ComponentValueTableRow{
			Component: picker.Picker[string]("Disabled", names, statePickerDisabled).
				Disabled(true),
			Value: strings.Join(statePickerDisabled.Get(), ", "),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}

func getOptionLabelByIndex(index int) string {
	if index < 0 || index >= len(names) {
		return ""
	}

	return names[index]
}

func getOptionLabelFromSelectOptions(options []dropdown.Option[ID], selected ID) string {
	for _, option := range options {
		if option.Value == selected {
			return option.Label
		}
	}

	return ""
}
