package signature

import (
	_ "embed"

	"github.com/worldiety/nago-demo/layout"
	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
)

//go:embed signature.svg
var signature string

func Content(wnd core.Window) core.View {
	return ui.Stack(
		tableDefault(wnd),
	).FullWidth()
}

func tableDefault(wnd core.Window) core.View {
	stateDefault := core.StateOf[ui.Signature](wnd, "stateDefault")
	stateSupport := core.StateOf[ui.Signature](wnd, "stateSupport")
	stateError := core.StateOf[ui.Signature](wnd, "stateError")
	stateDisabled := core.StateOf[ui.Signature](wnd, "stateDisabled").Init(func() ui.Signature {
		return ui.Signature{
			SVG: signature,
		}
	})

	return table("Unterschrift",
		layout.ComponentValueTableRow{
			Component: ui.SignatureField("Standard", stateDefault),
			ValueView: ui.If(len(stateDefault.Get().SVG) > 0, ui.Image().Embed([]byte(stateDefault.Get().SVG)).Frame(ui.Frame{Width: ui.L160, Height: ui.L80})),
		},
		layout.ComponentValueTableRow{
			Component: ui.SignatureField("Support", stateSupport).SupportingText("Hier steht ein Support-Text"),
			ValueView: ui.If(len(stateSupport.Get().SVG) > 0, ui.Image().Embed([]byte(stateSupport.Get().SVG)).Frame(ui.Frame{Width: ui.L160, Height: ui.L80})),
		},
		layout.ComponentValueTableRow{
			Component: ui.SignatureField("Fehler", stateError).ErrorText("Hier steht ein Fehler-Text"),
			ValueView: ui.If(len(stateError.Get().SVG) > 0, ui.Image().Embed([]byte(stateError.Get().SVG)).Frame(ui.Frame{Width: ui.L160, Height: ui.L80})),
		},
		layout.ComponentValueTableRow{
			Component: ui.SignatureField("Disabled", stateDisabled).Disabled(true),
			ValueView: ui.If(len(stateDisabled.Get().SVG) > 0, ui.Image().Embed([]byte(stateDisabled.Get().SVG)).Frame(ui.Frame{Width: ui.L160, Height: ui.L80})),
		},
	)
}

func table(title string, rows ...layout.ComponentValueTableRow) core.View {
	return ui.VStack(
		ui.Text(title).Font(ui.HeadlineSmall),
		layout.ComponentValueTable(rows...),
	).Alignment(ui.Top).Gap(ui.L4)
}
