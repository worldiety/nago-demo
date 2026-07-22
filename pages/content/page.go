package content

import (
	"fmt"

	"github.com/worldiety/nago-demo/layout"
	"github.com/worldiety/nago-demo/pages"
	"github.com/worldiety/nago-demo/pages/content/accordion"
	"github.com/worldiety/nago-demo/pages/content/alert"
	"github.com/worldiety/nago-demo/pages/content/chart"
	"github.com/worldiety/nago-demo/pages/content/dataview"
	"github.com/worldiety/nago-demo/pages/content/dialog"
	"github.com/worldiety/nago-demo/pages/content/flowchart"
	"github.com/worldiety/nago-demo/pages/content/hero"
	"github.com/worldiety/nago-demo/pages/content/pdf"
	"github.com/worldiety/nago-demo/pages/content/stepper"
	"github.com/worldiety/nago-demo/pages/content/switcher"
	"go.wdy.de/nago/presentation/ui/dropdown"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

type ContentContext struct {
	Hero hero.ContentHeroContext
}

var categoryOptions = []dropdown.Option[string]{
	{
		Value: "accordion",
		Label: "Akkordeon",
	},
	{
		Value: "alert",
		Label: "Alert",
	},
	{
		Value: "hero",
		Label: "Banner",
	},
	{
		Value: "dataview",
		Label: "Daten-Tabelle",
	},
	{
		Value: "chart",
		Label: "Diagramm",
	},
	{
		Value: "dialog",
		Label: "Dialog",
	},
	{
		Value: "flowchart",
		Label: "Flussdiagramm",
	},
	{
		Value: "pdf",
		Label: "PDF",
	},
	{
		Value: "stepper",
		Label: "Stepper",
	},
	{
		Value: "switcher",
		Label: "Switcher",
	},
}

func Page(wnd core.Window, ctx ContentContext) core.View {
	p, wide := page(wnd, ctx)

	return layout.Page(wnd,
		wide,
		"",
		"",
		p,
	)
}

func page(wnd core.Window, ctx ContentContext) (core.View, bool) {
	stateCategory := core.StateOf[string](wnd, "stateCategory").Init(func() string {
		fromQuery, ok := getCategoryQuery(wnd)
		if ok {
			return fromQuery
		}

		return categoryOptions[0].Value
	})

	stateCategory.Observe(func(cat string) {
		setCategoryQuery(wnd, cat)
	})

	return ui.VStack(
		pages.HeaderWithSelectFilter("Inhalt", "Kategorie", categoryOptions, stateCategory),
		pageContent(wnd, stateCategory.Get(), ctx),
	).Gap(ui.L32).FullWidth(), stateCategory.Get() == "hero"
}

func pageContent(wnd core.Window, page string, ctx ContentContext) core.View {
	switch page {
	case "accordion":
		return accordion.Content(wnd)
	case "alert":
		return alert.Content(wnd)
	case "chart":
		return chart.Content(wnd)
	case "dataview":
		return dataview.Content(wnd)
	case "dialog":
		return dialog.Content(wnd)
	case "flowchart":
		return flowchart.Content(wnd)
	case "hero":
		return hero.Content(wnd, ctx.Hero)
	case "pdf":
		return pdf.Content(wnd)
	case "stepper":
		return stepper.Content(wnd)
	case "switcher":
		return switcher.Content(wnd)
	}

	return ui.HStack(ui.Text(fmt.Sprintf("unknown category '%s'", page)))
}

func getDefaultCategory() string {
	return categoryOptions[0].Value
}

func getCategoryQuery(wnd core.Window) (string, bool) {
	val, ok := wnd.Values()["category"]
	return val, ok
}

func setCategoryQuery(wnd core.Window, page string) {
	val, ok := getCategoryQuery(wnd)
	if (!ok && page != getDefaultCategory()) || (ok && page != val) {
		values := wnd.Values()
		values = values.Put("category", page)
		wnd.Navigation().Replace(wnd.Path(), values)
	}
}
