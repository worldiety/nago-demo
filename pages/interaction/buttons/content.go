package buttons

import (
	"fmt"
	"time"

	"github.com/worldiety/nago-demo/layout"
	icons "go.wdy.de/nago/presentation/icons/flowbite/outline"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(wnd core.Window) core.View {
	cols := 2
	if wnd.Info().SizeClass < core.SizeClassLarge {
		cols = 1
	}

	return ui.Grid(
		ui.GridCell(tablePrimary(wnd)),
		ui.GridCell(tablePrimaryLink()),
		ui.GridCell(tableSecondary(wnd)),
		ui.GridCell(tableSecondaryLink()),
		ui.GridCell(tableTertiary(wnd)),
		ui.GridCell(tableTertiaryLink()),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tablePrimary(wnd core.Window) core.View {
	statePrimaryDefault := core.StateOf[time.Time](wnd, "statePrimaryDefault")
	statePrimaryIconLeft := core.StateOf[time.Time](wnd, "statePrimaryIconLeft")
	statePrimaryIconRight := core.StateOf[time.Time](wnd, "statePrimaryIconRight")
	statePrimaryIcon := core.StateOf[time.Time](wnd, "statePrimaryIcon")
	statePrimaryDisabled := core.StateOf[time.Time](wnd, "statePrimaryDisabled")

	return table("Primär",
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(func() { statePrimaryDefault.Set(time.Now()) }).Title("Standard"),
			Value:     valueLastPressed(statePrimaryDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(func() { statePrimaryIconLeft.Set(time.Now()) }).Title("Icon links").PreIcon(icons.PaperClip),
			Value:     valueLastPressed(statePrimaryIconLeft.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(func() { statePrimaryIconRight.Set(time.Now()) }).Title("Icon rechts").PostIcon(icons.Fire),
			Value:     valueLastPressed(statePrimaryIconRight.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(func() { statePrimaryIcon.Set(time.Now()) }).PreIcon(icons.List),
			Value:     valueLastPressed(statePrimaryIcon.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(func() { statePrimaryDisabled.Set(time.Now()) }).Title("Disabled").Disabled(true),
			Value:     valueLastPressed(statePrimaryDisabled.Get()),
		},
	)
}

func tablePrimaryLink() core.View {
	url := core.URI("https://worldiety.de")

	return table("Primär (Link)",
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Standard"),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon links").
				PreIcon(icons.PaperClip),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon rechts").
				PostIcon(icons.Fire),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(nil).
				HRef(url).
				Target("_blank").
				PreIcon(icons.List),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.PrimaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Disabled").
				Disabled(true),
			Value: string(url),
		},
	)
}

func tableSecondary(wnd core.Window) core.View {
	stateSecondaryDefault := core.StateOf[time.Time](wnd, "stateSecondaryDefault")
	stateSecondaryIconLeft := core.StateOf[time.Time](wnd, "stateSecondaryIconLeft")
	stateSecondaryIconRight := core.StateOf[time.Time](wnd, "stateSecondaryIconRight")
	stateSecondaryIcon := core.StateOf[time.Time](wnd, "stateSecondaryIcon")
	stateSecondaryDisabled := core.StateOf[time.Time](wnd, "stateSecondaryDisabled")

	return table("Sekundär",
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(func() { stateSecondaryDefault.Set(time.Now()) }).Title("Standard"),
			Value:     valueLastPressed(stateSecondaryDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(func() { stateSecondaryIconLeft.Set(time.Now()) }).Title("Icon links").PreIcon(icons.PaperClip),
			Value:     valueLastPressed(stateSecondaryIconLeft.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(func() { stateSecondaryIconRight.Set(time.Now()) }).Title("Icon rechts").PostIcon(icons.Fire),
			Value:     valueLastPressed(stateSecondaryIconRight.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(func() { stateSecondaryIcon.Set(time.Now()) }).PreIcon(icons.List),
			Value:     valueLastPressed(stateSecondaryIcon.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(func() { stateSecondaryDisabled.Set(time.Now()) }).Title("Disabled").Disabled(true),
			Value:     valueLastPressed(stateSecondaryDisabled.Get()),
		},
	)
}

func tableSecondaryLink() core.View {
	url := core.URI("https://worldiety.de")

	return table("Sekundär (Link)",
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Standard"),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon links").
				PreIcon(icons.PaperClip),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon rechts").
				PostIcon(icons.Fire),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(nil).
				HRef(url).
				Target("_blank").
				PreIcon(icons.List),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.SecondaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Disabled").
				Disabled(true),
			Value: string(url),
		},
	)
}

func tableTertiary(wnd core.Window) core.View {
	stateTertiaryDefault := core.StateOf[time.Time](wnd, "stateTertiaryDefault")
	stateTertiaryIconLeft := core.StateOf[time.Time](wnd, "stateTertiaryIconLeft")
	stateTertiaryIconRight := core.StateOf[time.Time](wnd, "stateTertiaryIconRight")
	stateTertiaryIcon := core.StateOf[time.Time](wnd, "stateTertiaryIcon")
	stateTertiaryDisabled := core.StateOf[time.Time](wnd, "stateTertiaryDisabled")

	return table("Tertiär",
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(func() { stateTertiaryDefault.Set(time.Now()) }).Title("Standard"),
			Value:     valueLastPressed(stateTertiaryDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(func() { stateTertiaryIconLeft.Set(time.Now()) }).Title("Icon links").PreIcon(icons.PaperClip),
			Value:     valueLastPressed(stateTertiaryIconLeft.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(func() { stateTertiaryIconRight.Set(time.Now()) }).Title("Icon rechts").PostIcon(icons.Fire),
			Value:     valueLastPressed(stateTertiaryIconRight.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(func() { stateTertiaryIcon.Set(time.Now()) }).PreIcon(icons.List),
			Value:     valueLastPressed(stateTertiaryIcon.Get()),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(func() { stateTertiaryDisabled.Set(time.Now()) }).Title("Disabled").Disabled(true),
			Value:     valueLastPressed(stateTertiaryDisabled.Get()),
		},
	)
}

func tableTertiaryLink() core.View {
	url := core.URI("https://worldiety.de")

	return table("Tertiär (Link)",
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Standard"),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon links").
				PreIcon(icons.PaperClip),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Icon rechts").
				PostIcon(icons.Fire),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(nil).
				HRef(url).
				Target("_blank").
				PreIcon(icons.List),
			Value: string(url),
		},
		layout.ComponentValueTableRow{
			Component: ui.TertiaryButton(nil).
				HRef(url).
				Target("_blank").
				Title("Disabled").
				Disabled(true),
			Value: string(url),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}

func valueLastPressed(t time.Time) string {
	return fmt.Sprintf("Zuletzt gedrückt: %s", t.Format(time.TimeOnly))
}
