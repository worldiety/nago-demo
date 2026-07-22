package home

import (
	_ "embed"

	"github.com/worldiety/nago-demo/layout"
	"go.wdy.de/nago/application"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/hero"
)

//go:embed hero_side.png
var heroSide application.StaticBytes

func Page(wnd core.Window, heroBgUrl core.URI) core.View {
	return layout.Page(wnd, "", "",
		ui.VStack(
			hero.Hero("Willkommen in der NAGO Demo-App").
				Subtitle("Teil der NAGO-Qualitätsoffensive").
				BackgroundImage(heroBgUrl).
				SideImage(ui.Image().Embed(heroSide)).
				Alignment(ui.Leading).
				Frame(ui.Frame{}.FullWidth()),
			ui.VStack(
				ui.Text("Bevor Du loslegst...").Font(ui.HeadlineMedium),
				ui.RichText("<ul>"+
					"<li>Eingabekomponenten (bspw. Textfelder) sind in Tabellen dargestellt, welche Auswirkungen auf deren Inhalt haben. Sollte etwas nicht korrekt dargestellt werden, besprich dies bitte mit dem EntwicklerInnen-Team. Gemeinsam entscheiden wir dann, ob es sich um einen Bug der Komponente handelt, oder ob die Darstellung in der Demo-App angepasst werden muss.</li>"+
					"<li>Du findest die Version der Demo-App und von NAGO am unteren Rand der Anwendung. Bitte gib in Bug-Tickets immer <b>beide</b> Versionen an.</li>"+
					"<li>Die Demo-App kann nicht alle möglichen Szenarien und Kombinationen von Eigenschaften seiner Komponenten darstellen, versucht aber möglichst viel abzudecken. Sollte Dir doch mal etwas fehlen, melde dich bitte beim EntwicklerInnen-Team.</li>"+
					"</ul>"),
			).Alignment(ui.Leading).Frame(ui.Frame{MaxWidth: ui.L1200}),
		).Gap(ui.L32),
	)
}
