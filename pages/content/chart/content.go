package chart

import (
	_ "embed"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/barchart"
	"go.wdy.de/nago/presentation/ui/chart"
	"go.wdy.de/nago/presentation/ui/linechart"
	"go.wdy.de/nago/presentation/ui/piechart"
)

func Content(_ core.Window) core.View {
	return ui.VStack(
		gridBar(),
		gridLine(),
		gridPie(),
	).Gap(ui.L48).FullWidth().NoClip(true)
}

func gridBar() core.View {
	chartFrame := ui.Frame{Width: ui.Full, MinHeight: ui.L256}

	series := []chart.Series{
		{
			Label: "Beliebtheit",
			DataPoints: []chart.DataPoint{
				{
					X: "Burger",
					Y: 200,
				},
				{
					X: "Pizza",
					Y: 310,
				},
				{
					X: "Pasta",
					Y: 280,
				},
				{
					X: "Palak Paneer",
					Y: 1000,
				},
			},
		},
		{
			Label: "Triedel-Score",
			DataPoints: []chart.DataPoint{
				{
					X: "Burger",
					Y: 700,
				},
				{
					X: "Pizza",
					Y: 700,
				},
				{
					X: "Pasta",
					Y: 750,
				},
				{
					X: "Palak Paneer",
					Y: 400,
				},
			},
		},
	}

	markers := []barchart.Marker{
		{
			Label:          "Marker ohne Bedeutung",
			SeriesIndex:    1,
			DataPointIndex: 0,
			Value:          "250",
		},
		{
			Label:          "Runder Marker mit benutzerdefinierter Farbe",
			SeriesIndex:    1,
			DataPointIndex: 2,
			Value:          "550",
			Color:          ui.SE0,
			Width:          10,
			Height:         10,
			Round:          true,
		},
	}

	return grid("Balken",
		ui.GridCell(
			chartVariant("Standard",
				barchart.BarChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).Markers(markers),
			),
		),
		ui.GridCell(
			chartVariant("Horizontal",
				barchart.BarChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).Horizontal(true).Markers(markers),
			),
		),
		ui.GridCell(
			chartVariant("Keine Daten",
				barchart.BarChart(
					chart.Chart{
						Frame:         chartFrame,
						NoDataMessage: "Optionaler Keine-Daten-Text",
					},
				),
			),
		),
	)
}

func gridLine() core.View {
	chartFrame := ui.Frame{Width: ui.Full, MinHeight: ui.L256}

	series := []chart.Series{
		{
			Label: "Hofmanns",
			DataPoints: []chart.DataPoint{
				{
					X: "01/2026",
					Y: 180,
				},
				{
					X: "02/2026",
					Y: 210,
				},
				{
					X: "03/2026",
					Y: 224,
				},
				{
					X: "04/2026",
					Y: 13,
				},
			},
		},
		{
			Label: "Pommes",
			DataPoints: []chart.DataPoint{
				{
					X: "01/2026",
					Y: 24,
				},
				{
					X: "02/2026",
					Y: 42,
				},
				{
					X: "03/2026",
					Y: 67,
				},
				{
					X: "04/2026",
					Y: 160,
				},
			},
		},
	}

	markers := linechart.Markers{
		Size: 5,
	}

	return grid("Linien",
		ui.GridCell(
			chartVariant("Standard",
				linechart.LineChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).Markers(markers),
			),
		),
		ui.GridCell(
			chartVariant("Smooth",
				linechart.LineChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).Curve(linechart.CurveSmooth).Markers(markers),
			),
		),
		ui.GridCell(
			chartVariant("Keine Daten",
				linechart.LineChart(
					chart.Chart{
						Frame:         chartFrame,
						NoDataMessage: "Optionaler Keine-Daten-Text",
					},
				),
			),
		),
	)
}

func gridPie() core.View {
	chartFrame := ui.Frame{Width: ui.Full, MinHeight: ui.L256}

	series := []chart.Series{
		{
			Label: "Beliebtheit",
			DataPoints: []chart.DataPoint{
				{
					X: "Harzer Roller",
					Y: 20,
				},
				{
					X: "Burger",
					Y: 200,
				},
				{
					X: "Pizza",
					Y: 310,
				},
				{
					X: "Pasta",
					Y: 280,
				},
				{
					X: "Palak Paneer",
					Y: 1000,
				},
			},
		},
	}

	return grid("Kuchen",
		ui.GridCell(
			chartVariant("Standard",
				piechart.PieChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series),
			),
		),
		ui.GridCell(
			chartVariant("Absolute Werte",
				piechart.PieChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).ShowAbsoluteValues(true),
			),
		),
		ui.GridCell(
			chartVariant("Gerundete Werte",
				piechart.PieChart(
					chart.Chart{
						Frame:         chartFrame,
						LabelRounding: chart.RoundingRound,
					},
				).Series(series),
			),
		),
		ui.GridCell(
			chartVariant("Donut",
				piechart.PieChart(
					chart.Chart{
						Frame: chartFrame,
					},
				).Series(series).ShowAsDonut(true),
			),
		),
		ui.GridCell(
			chartVariant("Keine Daten",
				piechart.PieChart(
					chart.Chart{
						Frame:         chartFrame,
						NoDataMessage: "Optionaler Keine-Daten-Text",
					},
				),
			),
		),
	)
}

func chartVariant(title string, barChart core.View) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.TitleMedium),
		barChart,
	).Gap(ui.L4).NoClip(true)
}

func grid(title string, cells ...ui.TGridCell) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		ui.Grid(cells...).Columns(3).Gap(ui.L32).FullWidth(),
	).Alignment(ui.Top).Gap(ui.L16).NoClip(true).FullWidth()
}
