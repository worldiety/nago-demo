package main

import (
	_ "embed"

	"github.com/worldiety/nago-demo/pages/buttons"
	"github.com/worldiety/nago-demo/pages/home"
	"github.com/worldiety/nago-demo/pages/input"
	"github.com/worldiety/nago-demo/pages/typography"
	"github.com/worldiety/nago-demo/util"

	"github.com/worldiety/option"
	"go.wdy.de/nago/application"
	cfginspector "go.wdy.de/nago/application/inspector/cfg"
	cfgmigration "go.wdy.de/nago/application/migration/cfg"
	cfgrebac "go.wdy.de/nago/application/rebac/cfg"
	"go.wdy.de/nago/presentation/core"
	icons "go.wdy.de/nago/presentation/icons/flowbite/outline"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/web/vuejs"
)

//go:embed nago_logo.svg
var appIcon application.StaticBytes

func create() *application.Application {
	return application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.nago_test")
		cfg.SetName("NAGO Test")
		cfg.Serve(vuejs.Dist())

		option.MustZero(cfg.StandardSystems())
		option.Must(cfgmigration.Enable(cfg))
		option.Must(cfgrebac.Enable(cfg))

		option.Must(cfginspector.Enable(cfg))

		cfg.SetDecorator(
			cfg.NewScaffold().
				Login(false).
				Logo(ui.Image().Embed(appIcon).Frame(ui.Frame{Height: ui.L48})).
				Footer(ui.VStack(
					ui.Text("NAGO-Tester: "+util.Version()),
					ui.Text("NAGO: "+util.DependencyVersion("go.wdy.de/nago")),
				).Alignment(ui.Center).FullWidth(),
				).
				MenuEntry().Title("Buttons").Icon(icons.CirclePlus).Forward("buttons").OneOfRole().
				MenuEntry().Title("Eingabe").Icon(icons.TextSize).Forward("input").OneOfRole().
				MenuEntry().Title("Typografie").Icon(icons.TextSize).Forward("typography").OneOfRole().Decorator(),
		)

		cfg.RootViewWithDecoration(".", func(wnd core.Window) core.View {
			return home.Page(wnd)
		})

		cfg.RootViewWithDecoration("buttons", func(wnd core.Window) core.View {
			return buttons.Page(wnd)
		})
		cfg.RootViewWithDecoration("input", func(wnd core.Window) core.View {
			return input.Page(wnd)
		})
		cfg.RootViewWithDecoration("typography", func(wnd core.Window) core.View {
			return typography.Page(wnd)
		})
	})
}

// the main function of the program, which is like the java public static void main.
func main() {
	create().Run()
}
