package datetime

import (
	"fmt"
	"time"

	"github.com/worldiety/nago-demo/layout"
	"go.wdy.de/nago/pkg/xtime"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/timepicker"
)

func Content(wnd core.Window) core.View {
	cols := 2
	if wnd.Info().SizeClass < core.SizeClassXL {
		cols = 1
	}

	return ui.Grid(
		ui.GridCell(tableDate(wnd)),
		ui.GridCell(tableDateRange(wnd)),
		ui.GridCell(tableDuration(wnd)),
	).Columns(cols).Gap(ui.L32).FullWidth().Heights("auto")
}

func tableDate(wnd core.Window) core.View {
	worldietyGbR := xtime.Date{Day: 16, Month: 8, Year: 2010}

	stateDateDefault := core.StateOf[xtime.Date](wnd, "stateDateDefault").Init(func() xtime.Date { return worldietyGbR })
	stateDateDouble := core.StateOf[xtime.Date](wnd, "stateDateDouble").Init(func() xtime.Date { return worldietyGbR })
	stateDateSupport := core.StateOf[xtime.Date](wnd, "stateDateSupport").Init(func() xtime.Date { return worldietyGbR })
	stateDateError := core.StateOf[xtime.Date](wnd, "stateDateError").Init(func() xtime.Date { return worldietyGbR })
	stateDateDisabled := core.StateOf[xtime.Date](wnd, "stateDateDisabled").Init(func() xtime.Date { return worldietyGbR })

	return table("Datum",
		layout.ComponentValueTableRow{
			Component: ui.SingleDatePicker("Standard", stateDateDefault.Get(), stateDateDefault),
			Value:     stateDateDefault.Get().Format(xtime.GermanDate),
		},
		layout.ComponentValueTableRow{
			Component: ui.SingleDatePicker("Doppelt", stateDateDouble.Get(), stateDateDouble).
				DoubleMode(true),
			Value: stateDateDefault.Get().Format(xtime.GermanDate),
		},
		layout.ComponentValueTableRow{
			Component: ui.SingleDatePicker("Support", stateDateSupport.Get(), stateDateSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: stateDateSupport.Get().Format(xtime.GermanDate),
		},
		layout.ComponentValueTableRow{
			Component: ui.SingleDatePicker("Fehler", stateDateError.Get(), stateDateError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: stateDateError.Get().Format(xtime.GermanDate),
		},
		layout.ComponentValueTableRow{
			Component: ui.SingleDatePicker("Disabled", stateDateDisabled.Get(), stateDateDisabled).
				Disabled(true),
			Value: stateDateDisabled.Get().Format(xtime.GermanDate),
		},
	)
}

func tableDateRange(wnd core.Window) core.View {
	worldietyGbR2 := xtime.Date{Day: 19, Month: 8, Year: 2010}
	worldietyGbR3 := xtime.Date{Day: 24, Month: 8, Year: 2010}

	stateDateStartDefault := core.StateOf[xtime.Date](wnd, "stateDateStartDefault").Init(func() xtime.Date { return worldietyGbR2 })
	stateDateEndDefault := core.StateOf[xtime.Date](wnd, "stateDateEndDefault").Init(func() xtime.Date { return worldietyGbR3 })
	stateDateStartDouble := core.StateOf[xtime.Date](wnd, "stateDateStartDouble").Init(func() xtime.Date { return worldietyGbR2 })
	stateDateEndDouble := core.StateOf[xtime.Date](wnd, "stateDateEndDouble").Init(func() xtime.Date { return worldietyGbR3 })
	stateDateStartSupport := core.StateOf[xtime.Date](wnd, "stateDateStartSupport").Init(func() xtime.Date { return worldietyGbR2 })
	stateDateEndSupport := core.StateOf[xtime.Date](wnd, "stateDateEndSupport").Init(func() xtime.Date { return worldietyGbR3 })
	stateDateStartError := core.StateOf[xtime.Date](wnd, "stateDateStartError").Init(func() xtime.Date { return worldietyGbR2 })
	stateDateEndError := core.StateOf[xtime.Date](wnd, "stateDateEndError").Init(func() xtime.Date { return worldietyGbR3 })
	stateDateStartDisabled := core.StateOf[xtime.Date](wnd, "stateDateStartDisabled").Init(func() xtime.Date { return worldietyGbR2 })
	stateDateEndDisabled := core.StateOf[xtime.Date](wnd, "stateDateEndDisabled").Init(func() xtime.Date { return worldietyGbR3 })

	return table("Datum (Zeitraum)",
		layout.ComponentValueTableRow{
			Component: ui.RangeDatePicker("Standard", stateDateStartDefault.Get(), stateDateStartDefault, stateDateEndDefault.Get(), stateDateEndDefault),
			Value:     fmt.Sprintf("%s - %s", stateDateStartDefault.Get().Format(xtime.GermanDate), stateDateEndDefault.Get().Format(xtime.GermanDate)),
		},
		layout.ComponentValueTableRow{
			Component: ui.RangeDatePicker("Doppelt", stateDateStartDouble.Get(), stateDateStartDouble, stateDateEndDouble.Get(), stateDateEndDouble).
				DoubleMode(true),
			Value: fmt.Sprintf("%s - %s", stateDateStartDouble.Get().Format(xtime.GermanDate), stateDateEndDouble.Get().Format(xtime.GermanDate)),
		},
		layout.ComponentValueTableRow{
			Component: ui.RangeDatePicker("Support", stateDateStartSupport.Get(), stateDateStartSupport, stateDateEndSupport.Get(), stateDateEndSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: fmt.Sprintf("%s - %s", stateDateStartSupport.Get().Format(xtime.GermanDate), stateDateEndSupport.Get().Format(xtime.GermanDate)),
		},
		layout.ComponentValueTableRow{
			Component: ui.RangeDatePicker("Fehler", stateDateStartError.Get(), stateDateStartError, stateDateEndError.Get(), stateDateEndError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: fmt.Sprintf("%s - %s", stateDateStartError.Get().Format(xtime.GermanDate), stateDateEndError.Get().Format(xtime.GermanDate)),
		},
		layout.ComponentValueTableRow{
			Component: ui.RangeDatePicker("Disabled", stateDateStartDisabled.Get(), stateDateStartDisabled, stateDateEndDisabled.Get(), stateDateEndDisabled).
				Disabled(true),
			Value: fmt.Sprintf("%s - %s", stateDateStartDisabled.Get().Format(xtime.GermanDate), stateDateEndDisabled.Get().Format(xtime.GermanDate)),
		},
	)
}

func tableDuration(wnd core.Window) core.View {
	stateDurationDefault := core.StateOf[time.Duration](wnd, "stateDurationDefault")
	stateDurationAll := core.StateOf[time.Duration](wnd, "stateDurationAll")
	stateDurationFormatted := core.StateOf[time.Duration](wnd, "stateDurationFormatted")
	stateDurationSupport := core.StateOf[time.Duration](wnd, "stateDurationSupport")
	stateDurationError := core.StateOf[time.Duration](wnd, "stateDurationError")
	stateDurationDisabled := core.StateOf[time.Duration](wnd, "stateDurationDisabled")

	return table("Zeit/Dauer",
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Standard", stateDurationDefault),
			Value:     stateDurationDefault.Get().String(),
		},
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Alle Felder", stateDurationAll).
				Days(true).
				Hours(true).
				Minutes(true).
				Seconds(true),
			Value: stateDurationAll.Get().String(),
		},
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Formatiert", stateDurationFormatted).
				Format(timepicker.DecomposedFormat),
			Value: stateDurationFormatted.Get().String(),
		},
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Support", stateDurationSupport).
				SupportingText("Hier steht ein Support-Text"),
			Value: stateDurationSupport.Get().String(),
		},
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Fehler", stateDurationError).
				ErrorText("Hier steht ein Fehler-Text"),
			Value: stateDurationError.Get().String(),
		},
		layout.ComponentValueTableRow{
			Component: timepicker.Picker("Disabled", stateDurationDisabled).
				Disabled(true),
			Value: stateDurationDisabled.Get().String(),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}
