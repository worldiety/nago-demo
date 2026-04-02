package interaction

import (
	"github.com/worldiety/nago-demo/layout"
	"github.com/worldiety/nago-demo/pages"
	"github.com/worldiety/nago-demo/pages/interaction/buttons"
	"go.wdy.de/nago/presentation/ui/dropdown"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

var categoryOptions = []dropdown.Option[string]{
	{
		Value: "buttons",
		Label: "Buttons",
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
	stateCategory := core.StateOf[string](wnd, "stateCategory")

	fromQuery, ok := getCategoryQuery(wnd)
	if ok {
		stateCategory.Set(fromQuery)
	} else {
		stateCategory.Set(categoryOptions[0].Value)
	}

	stateCategory.Observe(func(cat string) {
		setCategoryQuery(wnd, cat)
	})

	return ui.VStack(
		pages.HeaderWithSelectFilter("Interaktion", "Kategorie", categoryOptions, stateCategory),
		pageContent(wnd, stateCategory.Get()),
	).Gap(ui.L32).FullWidth()
}

func pageContent(wnd core.Window, page string) core.View {
	switch page {
	case "buttons":
		return buttons.Content(wnd)
	}

	return ui.HStack()
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
		wnd.Navigation().ForwardTo(wnd.Path(), values)
	}
}
