package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/worldiety/nago-demo/pages/content"
	"github.com/worldiety/nago-demo/pages/home"
	"github.com/worldiety/nago-demo/pages/input"
	"github.com/worldiety/nago-demo/pages/interaction"
	"github.com/worldiety/nago-demo/pages/typography"
	"github.com/worldiety/nago-demo/util"
	"github.com/worldiety/option"
	"go.wdy.de/nago/application"
	cfginspector "go.wdy.de/nago/application/inspector/cfg"
	cfgmigration "go.wdy.de/nago/application/migration/cfg"
	cfgrebac "go.wdy.de/nago/application/rebac/cfg"
	"go.wdy.de/nago/presentation/core"
	flowbiteOutline "go.wdy.de/nago/presentation/icons/flowbite/outline"
	heroOutline "go.wdy.de/nago/presentation/icons/hero/outline"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
	"go.wdy.de/nago/web/vuejs"
)

//go:embed nago_logo.svg
var appIcon application.StaticBytes

func create() *application.Application {
	return application.Configure(func(cfg *application.Configurator) {
		cfg.SetApplicationID("de.worldiety.nago_test")
		cfg.SetName("NAGO-Tester")
		cfg.Serve(vuejs.Dist())

		option.MustZero(cfg.StandardSystems())
		option.Must(cfgmigration.Enable(cfg))
		option.Must(cfgrebac.Enable(cfg))

		option.Must(cfginspector.Enable(cfg))

		cfg.SetDecorator(
			cfg.NewScaffold().
				Login(false).
				Logo(ui.Image().Embed(appIcon).Frame(ui.Frame{Height: ui.L48})).
				Footer(
					ui.HStack(
						ui.Stack(
							ui.Image().Embed(appIcon).Frame(ui.Frame{Height: ui.L48}),
							versionsView(),
						).
							Gap(ui.L32).
							Alignment(ui.Leading).
							Padding(ui.Padding{}.Horizontal(ui.L16)).
							Frame(ui.Frame{Width: ui.Full, MaxWidth: ui.L1600}),
					).
						Alignment(ui.Center).
						FullWidth().
						Padding(ui.Padding{}.Vertical(ui.L32)).
						Border(ui.Border{TopWidth: ui.L1, TopColor: ui.M5}),
				).
				MenuEntry().Title("Eingabe").Icon(flowbiteOutline.Keyboard).Forward("input").OneOfRole().
				MenuEntry().Title("Inhalt").Icon(flowbiteOutline.Image).Forward("content").OneOfRole().
				MenuEntry().Title("Interaktion").Icon(heroOutline.CursorArrowRays).Forward("interaction").OneOfRole().
				MenuEntry().Title("Typografie").Icon(flowbiteOutline.TextSize).Forward("typography").OneOfRole().
				MenuEntry().Dynamic(darkModeToggleMenuEntry()).Decorator(),
		)

		cfg.RootViewWithDecoration(".", func(wnd core.Window) core.View {
			return home.Page(wnd)
		})

		cfg.RootViewWithDecoration("input", func(wnd core.Window) core.View {
			return input.Page(wnd)
		})
		cfg.RootViewWithDecoration("content", func(wnd core.Window) core.View {
			return content.Page(wnd)
		})
		cfg.RootViewWithDecoration("interaction", func(wnd core.Window) core.View {
			return interaction.Page(wnd)
		})
		cfg.RootViewWithDecoration("typography", func(wnd core.Window) core.View {
			return typography.Page(wnd)
		})
	})
}

func darkModeToggleMenuEntry() func(wnd core.Window, entry *application.MenuEntryBuilder) {
	return func(wnd core.Window, entry *application.MenuEntryBuilder) {
		if wnd.Info().ColorScheme == core.Light {
			entry.Title("Dark Mode")
			entry.Icon(flowbiteOutline.Moon)
			entry.Action(func(wnd core.Window) {
				wnd.SetColorScheme(core.Dark)
			})
		} else {
			entry.Title("Light Mode")
			entry.Icon(flowbiteOutline.Sun)
			entry.Action(func(wnd core.Window) {
				wnd.SetColorScheme(core.Light)
			})
		}
	}
}

func versionsView() core.View {
	versionText := fmt.Sprintf("NAGO-Tester: %s\nNAGO: %s", util.Version(), util.DependencyVersion("go.wdy.de/nago"))

	return ui.VStack(
		ui.HStack(
			ui.Text("Versionen:").
				Font(ui.LabelSmall).
				Color(ui.ColorText.WithTransparency(50)).
				Padding(ui.Padding{Left: ui.L4}),
		).Position(ui.Position{Type: ui.PositionAbsolute, Left: ui.L0, Bottom: ui.Full}),
		ui.Text(versionText).Font(ui.MonoMedium),
	).
		NoClip(true).
		Position(ui.Position{Type: ui.PositionRelative}).
		Responsive(func(wnd core.Window, stack ui.TStack) ui.TStack {
			copyVersionsState := core.StateOf[bool](wnd, "copyVersionsState")

			stack = stack.Append(
				ui.If(copyVersionsState.Get(),
					ui.HStack(
						ui.Text("Kopiert!").
							Font(ui.LabelSmall).
							Color(ui.ColorText.WithTransparency(50)).Padding(ui.Padding{Left: ui.L4}),
					).Position(ui.Position{Type: ui.PositionAbsolute, Right: ui.L0, Bottom: ui.Full}),
				),
			)

			return stack.Action(func() {
				if err := wnd.Clipboard().SetText(versionText); err != nil {
					alert.ShowBannerError(wnd, err)
					copyVersionsState.Set(false)
					return
				}

				if !copyVersionsState.Get() {
					copyVersionsState.Set(true)
					go func() {
						time.Sleep(2 * time.Second)
						copyVersionsState.Set(false)
					}()
				}
			})
		}).HoveredBackgroundColor(ui.ColorText.WithTransparency(85)).Padding(ui.Padding{}.Vertical(ui.L2).Horizontal(ui.L4))
}

// the main function of the program, which is like the java public static void main.
func main() {
	create().Run()
}
