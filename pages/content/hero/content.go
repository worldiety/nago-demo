package hero

import (
	_ "embed"

	"go.wdy.de/nago/application"
	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/hero/outline"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/hero"
)

type ContentHeroContext struct {
	BgURI core.URI
}

//go:embed blob.png
var blob application.StaticBytes

const (
	title       = "Überschrift NAGO Text und so"
	subtitle    = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua."
	actionLabel = "Let's start"
)

func Content(wnd core.Window, ctx ContentHeroContext) core.View {
	wnd.Context()

	return ui.VStack(
		hero.Hero(title).
			Subtitle(subtitle).
			Actions(ui.PrimaryButton(nil).Title(actionLabel).PreIcon(icons.ArrowRight)).
			BackgroundImage(ctx.BgURI).
			Alignment(ui.BottomLeading).
			Frame(ui.Frame{Width: ui.L1600, MaxWidth: ui.Full}),
		hero.Hero(title).
			Subtitle(subtitle).
			Actions(ui.PrimaryButton(nil).Title(actionLabel).PreIcon(icons.ArrowRight)).
			BackgroundImage(ctx.BgURI).
			Alignment(ui.Center).
			Frame(ui.Frame{Width: ui.L1600, MaxWidth: ui.Full}),
		hero.Hero(title).
			Subtitle(subtitle).
			Actions(ui.PrimaryButton(nil).Title(actionLabel).PreIcon(icons.ArrowRight)).
			BackgroundImage(ctx.BgURI).
			Alignment(ui.BottomTrailing).
			Frame(ui.Frame{Width: ui.L1600, MaxWidth: ui.Full}),
		hero.Hero(title).
			Subtitle(subtitle).
			Actions(ui.PrimaryButton(nil).Title(actionLabel).PreIcon(icons.ArrowRight)).
			SideImage(ui.Image().Embed(blob)).
			Alignment(ui.Leading).
			Frame(ui.Frame{Width: ui.L1600, MaxWidth: ui.Full}),
	).Gap(ui.L32).FullWidth()
}
