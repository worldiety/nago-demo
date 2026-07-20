package input

import (
	"fmt"

	"github.com/worldiety/nago-demo/layout"
	"github.com/worldiety/nago-demo/pages"
	boolInput "github.com/worldiety/nago-demo/pages/input/bool"
	"github.com/worldiety/nago-demo/pages/input/datetime"
	"github.com/worldiety/nago-demo/pages/input/selection"
	"github.com/worldiety/nago-demo/pages/input/signature"
	"github.com/worldiety/nago-demo/pages/input/slider"
	"github.com/worldiety/nago-demo/pages/input/text"
	"go.wdy.de/nago/presentation/ui/dropdown"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

var categoryOptions = []dropdown.Option[string]{
	{
		Value: "selection",
		Label: "Auswahl",
	},
	{
		Value: "bool",
		Label: "Bool'sche Werte",
	},
	{
		Value: "datetime",
		Label: "Datum/Zeit",
	},
	{
		Value: "slider",
		Label: "Slider",
	},
	{
		Value: "text",
		Label: "Text",
	},
	{
		Value: "signature",
		Label: "Unterschrift",
	},
}

func Page(wnd core.Window) core.View {
	return layout.Page(wnd,
		"",
		"",
		page(wnd),
	)
}

func page(wnd core.Window) core.View {
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
		pages.HeaderWithSelectFilter("Eingabe", "Kategorie", categoryOptions, stateCategory),
		pageContent(wnd, stateCategory.Get()),
	).Gap(ui.L32).FullWidth()
}

func pageContent(wnd core.Window, page string) core.View {
	switch page {
	case "bool":
		return boolInput.Content(wnd)
	case "datetime":
		return datetime.Content(wnd)
	case "selection":
		return selection.Content(wnd)
	case "signature":
		return signature.Content(wnd)
	case "slider":
		return slider.Content(wnd)
	case "text":
		return text.Content(wnd)
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
