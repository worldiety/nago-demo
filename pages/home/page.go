package home

import (
	"github.com/worldiety/nago-demo/layout"

	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/flowbite/outline"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/hero"
)

func Page(wnd core.Window) core.View {
	return layout.Page(wnd, "", "",
		ui.VStack(
			hero.Hero("Willkommen beim NAGO-Tester").
				Subtitle("Teil der NAGO-Qualitätsoffensive").Alignment(ui.Leading).
				SideSVG(icons.ShieldCheck),
		),
	)
}
