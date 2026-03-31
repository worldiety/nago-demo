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
			ui.VStack(
				ui.Text("Bevor Du loslegst...").Font(ui.HeadlineMedium),
				ui.RichText("<ul>"+
					"<li>Eingabekomponenten (bspw. Textfelder) sind in Tabellen dargestellt, welche Auswirkungen auf deren Inhalt haben. Sollte etwas nicht korrekt dargestellt werden, besprich dies bitte mit dem EntwicklerInnen-Team. Gemeinsam entscheiden wir dann, ob es sich um einen Bug der Komponente handelt, oder ob die Darstellung im NAGO-Tester angepasst werden muss.</li>"+
					"<li>Du findest die Version des NAGO-Testers und von NAGO am unteren Rand der Anwendung. Bitte gib in Bug-Tickets immer <b>beide</b> Versionen an.</li>"+
					"<li>Der NAGO-Tester kann nicht alle möglichen Szenarien und Kombinationen von Eigenschaften seiner Komponenten darstellen, versucht aber möglichst viel abzudecken. Sollte Dir doch mal etwas fehlen, melde dich bitte beim EntwicklerInnen-Team.</li>"+
					"</ul>"),
			).Alignment(ui.Leading).Frame(ui.Frame{MaxWidth: ui.L1200}),
		).Gap(ui.L32),
	)
}
