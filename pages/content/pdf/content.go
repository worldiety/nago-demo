package pdf

import (
	_ "embed"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

func Content(_ core.Window) core.View {
	pdfUrl := "https://ontheline.trincoll.edu/images/bookdown/sample-local-pdf.pdf"

	return ui.Stack(
		ui.PDF(core.URI(pdfUrl)).Frame(ui.Frame{Width: ui.L880, Height: ui.L880, MaxWidth: ui.Full}),
	).Gap(ui.L32).FullWidth()
}
