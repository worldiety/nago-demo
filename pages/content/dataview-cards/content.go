package dataview_card

import (
	_ "embed"

	"github.com/worldiety/nago-demo/pages/content/dataview-table"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui/dataview"
)

func Content(wnd core.Window) core.View {
	return dataview_table.Dataview(wnd, dataview.Card)
}
