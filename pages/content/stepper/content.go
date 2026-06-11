package stepper

import (
	_ "embed"

	"go.wdy.de/nago/presentation/ui/stepper"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(wnd core.Window) core.View {
	state := core.StateOf[int](wnd, "state")

	steps := []stepper.TStep{
		stepper.Step().Headline("Schritt 1").SupportingText("Einkaufsliste"),
		stepper.Step().Headline("Schritt 2").SupportingText("Einkaufen"),
		stepper.Step().Headline("Schritt 3").SupportingText("Zutaten schnibbeln"),
		stepper.Step().Headline("Schritt 4").SupportingText("Kochen"),
		stepper.Step().Headline("Schritt 5").SupportingText("Essen und freuen"),
	}

	return ui.VStack(
		controls(state, len(steps)),
		steppers(state, steps),
		controls(state, len(steps)),
	).Gap(ui.L32).FullWidth()
}

func steppers(state *core.State[int], steps []stepper.TStep) core.View {
	return ui.Grid(
		ui.GridCell(
			variant(
				"Standard",
				stepper.Stepper(steps...).InputValue(state),
			),
		).ColSpan(3),
		ui.GridCell(
			variant(
				"Standard (keine Zahlen)",
				stepper.Stepper(steps...).InputValue(state).Numbers(false),
			),
		).ColSpan(3),
		ui.GridCell(
			variant(
				"Vertikal",
				stepper.Stepper(steps...).InputValue(state).Layout(stepper.StepperLayoutVertical),
			),
		).RowSpan(2),
		ui.GridCell(
			variant(
				"Simpel",
				stepper.Stepper(steps...).InputValue(state).Layout(stepper.StepperLayoutSimple),
			),
		),
		ui.GridCell(
			variant(
				"Simpel (keine Linien)",
				stepper.Stepper(steps...).InputValue(state).Layout(stepper.StepperLayoutSimple).Lines(false),
			),
		),
		ui.GridCell(
			variant(
				"Simpel Liste",
				stepper.Stepper(steps...).InputValue(state).Layout(stepper.StepperLayoutSimpleList),
			),
		),
		ui.GridCell(
			variant(
				"Simpel Liste (keine Linien)",
				stepper.Stepper(steps...).InputValue(state).Layout(stepper.StepperLayoutSimpleList).Lines(false),
			),
		),
	).Columns(3).ColGap(ui.L32).RowGap(ui.L64).FullWidth().Heights("auto")
}

func variant(title string, content core.View) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		content,
	).Gap(ui.L16).Alignment(ui.Top).NoClip(true)
}

func controls(state *core.State[int], stepCount int) core.View {
	nextText := "Weiter >"
	if state.Get() >= stepCount {
		nextText = "Fertig >"
	}

	return ui.Stack(
		ui.PrimaryButton(func() {
			state.Set(state.Get() - 1)
		}).Title("< Zurück").Disabled(state.Get() <= 0),
		ui.PrimaryButton(func() {
			state.Set(state.Get() + 1)
		}).Title(nextText).Disabled(state.Get() > stepCount-1),
	).Gap(ui.L8)
}
