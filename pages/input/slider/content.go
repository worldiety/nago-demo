package slider

import (
	"fmt"

	"github.com/worldiety/nago-demo/layout"
	"go.wdy.de/nago/presentation/ui/slider"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(wnd core.Window) core.View {
	cols := 2
	if wnd.Info().SizeClass < core.SizeClassLarge {
		cols = 1
	}

	return ui.Grid(
		ui.GridCell(tableDefault(wnd)),
		ui.GridCell(tableRange(wnd)),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tableDefault(wnd core.Window) core.View {
	stateDefault := core.StateOf[float64](wnd, "stateDefault").Init(func() float64 { return 3 })
	stateMarkers := core.StateOf[float64](wnd, "stateMarkers").Init(func() float64 { return 3 })
	stateUnit := core.StateOf[float64](wnd, "stateUnit").Init(func() float64 { return 3 })
	stateSupport := core.StateOf[float64](wnd, "stateSupport").Init(func() float64 { return 3 })
	stateError := core.StateOf[float64](wnd, "stateError").Init(func() float64 { return 3 })
	stateSmallSteps := core.StateOf[float64](wnd, "stateSmallSteps").Init(func() float64 { return 3 })
	stateDisabled := core.StateOf[float64](wnd, "stateDisabled").Init(func() float64 { return 3 })

	return table("Einzelwert",
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Standard").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateDefault),
			Value:     fmt.Sprintf("%f", stateDefault.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Mit Marker").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateMarkers).ShowMarkers(true),
			Value:     fmt.Sprintf("%f", stateMarkers.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Mit Einheit").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateUnit).Unit("g"),
			Value:     fmt.Sprintf("%f", stateUnit.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Mit Support").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateSupport).SupportingText("Ich bin ein Support-Text"),
			Value:     fmt.Sprintf("%f", stateSupport.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Mit Fehler").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateError).ErrorText("Ich bin ein Fehler-Text"),
			Value:     fmt.Sprintf("%f", stateError.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Feingranular").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateSmallSteps).Step(0.01),
			Value:     fmt.Sprintf("%f", stateSmallSteps.Get()),
		},
		layout.ComponentValueTableRow{
			Component: slider.Slider(0, 10).Label("Disabled").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateDisabled).Disabled(true),
			Value:     fmt.Sprintf("%f", stateDisabled.Get()),
		},
	)
}

func tableRange(wnd core.Window) core.View {
	stateRangeDefault := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeDefault").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeMarkers := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeMarkers").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeUnit := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeUnit").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeSupport := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeSupport").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeError := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeError").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeSmallSteps := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeSmallSteps").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })
	stateRangeDisabled := core.StateOf[slider.RangeSliderValue](wnd, "stateRangeDisabled").Init(func() slider.RangeSliderValue { return slider.RangeSliderValue{From: 3, To: 7} })

	return table("Range",
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range Standard").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeDefault),
			Value:     fmt.Sprintf("%f - %f", stateRangeDefault.Get().From, stateRangeDefault.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range mit Marker").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeMarkers).ShowMarkers(true),
			Value:     fmt.Sprintf("%f - %f", stateRangeMarkers.Get().From, stateRangeMarkers.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range mit Einheit").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeUnit).Unit("g"),
			Value:     fmt.Sprintf("%f - %f", stateRangeUnit.Get().From, stateRangeUnit.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range mit Support").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeSupport).SupportingText("Ich bin ein Support-Text"),
			Value:     fmt.Sprintf("%f - %f", stateRangeSupport.Get().From, stateRangeSupport.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range mit Fehler").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeError).ErrorText("Ich bin ein Fehler-Text"),
			Value:     fmt.Sprintf("%f - %f", stateRangeError.Get().From, stateRangeError.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range feingranular").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeSmallSteps).Step(0.01),
			Value:     fmt.Sprintf("%f - %f", stateRangeSmallSteps.Get().From, stateRangeSmallSteps.Get().To),
		},
		layout.ComponentValueTableRow{
			Component: slider.RangeSlider(0, 10).Label("Range disabled").Frame(ui.Frame{MinWidth: ui.L200}).InputValue(stateRangeDisabled).Disabled(true),
			Value:     fmt.Sprintf("%f - %f", stateRangeDisabled.Get().From, stateRangeDisabled.Get().To),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}
