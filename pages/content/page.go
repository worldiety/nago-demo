package content

import (
	"github.com/worldiety/nago-demo/layout"
	"github.com/worldiety/nago-demo/pages/content/accordion"
	"github.com/worldiety/nago-demo/pages/content/dialog"
	"github.com/worldiety/nago-demo/pages/content/switcher"
	"go.wdy.de/nago/presentation/ui/dropdown"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

var categoryOptions = []dropdown.Option[string]{
	{
		Value: "accordion",
		Label: "Akkordeon",
	},
	{
		Value: "dialog",
		Label: "Dialog",
	},
	{
		Value: "switcher",
		Label: "Switcher",
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
		ui.Stack(
			ui.Text("Inhalt").Font(ui.HeadlineLarge),
			ui.Spacer(),
			dropdown.Dropdown("", categoryOptions, stateCategory.Get()).
				InputValue(stateCategory),
		).Gap(ui.L8).FullWidth(),
		pageContent(wnd, stateCategory.Get()),
	).Gap(ui.L32).FullWidth()
}

func pageContent(wnd core.Window, page string) core.View {
	switch page {
	case "accordion":
		return accordion.Content(wnd)
	case "dialog":
		return dialog.Content(wnd)
	case "switcher":
		return switcher.Content(wnd)
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
