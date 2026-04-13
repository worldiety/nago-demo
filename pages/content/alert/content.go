package alert

import (
	"fmt"
	"math/rand"
	"time"

	"go.wdy.de/nago/presentation/core"
	"go.wdy.de/nago/presentation/ui"
	"go.wdy.de/nago/presentation/ui/alert"
)

type AlertContent struct {
	title   string
	message string
}

var messages = []AlertContent{
	{
		title:   "Wer hat Torbens Hut gesehen?",
		message: "Torbens Hut wurde zuletzt in der Cafeteria neben der Kaffeemaschine gesehen. Der Hut besteht aus Echtgold und hat die Aufschrift TS420.",
	},
	{
		title:   "Achtung, Explosionsgefahr!",
		message: "Alle Geräte ab MacOS Version 26.4.0 sind betroffen. Wiederholt schnelles Auf- und Zuklappen des Gerätes kann zu heftigen Explosionen führen.",
	},
	{
		title:   "Miau?",
		message: "Miau miau Leckerchen miau miau. Miau schleck miau Hunger miau? GaLiGrü Mio miau prrrrrr",
	},
	{
		title:   "Hofmanns-Superkräfte?",
		message: "In der Mikrowelle erwärmte Hofmanns-Gerichte können bei Verzehr zu Superkräften führen.",
	},
	{
		title:   "NicNac-Vorrat kritisch",
		message: "Der Vorrat an NicNacs ist auf einem kritisch niedrigen Stand. Der NicNac-Radar zeigt zudem eine starke NicNac-Zunahme in direkter Nachbarschaft des WZO.",
	},
	{
		title:   "Wichtig!",
		message: "Um die Leistungsfähigkeit deines Geräts zu gewährleisten, sind folgende Schritte durchzuführen: Setze Dich unter deinen Tisch, ziehe deinen linken Schuh aus und singe die Titelmelodie der Biene Maja.",
	},
	{
		title:   "Rezept für Suppe",
		message: "Zutaten: Suppe\nZubereitung: Erwärme die Suppe.\nGuten Appetit!",
	},
	{
		title:   "Viel Spaß beim nächsten Test.",
		message: "Ich gehe jetzt zur Oberfläche. Es ist so ein schöner Tag. Gestern habe ich ein Reh gesehen. Wenn du den nächsten Test bestehst, lasse ich dich vielleicht im Fahrstuhl bis ganz nach oben fahren und erzähle dir von dem Reh.",
	},
}

func Content(wnd core.Window) core.View {
	return ui.Grid(
		ui.GridCell(
			ui.VStack(
				ui.Text("Standard"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:   content.title,
						Message: content.message,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Okay"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:   content.title,
						Message: content.message,
						Intent:  alert.IntentOk,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Fehler"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:   content.title,
						Message: content.message,
						Intent:  alert.IntentError,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Standard (5 Minuten)"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:    content.title,
						Message:  content.message,
						Duration: 5 * time.Minute,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Okay (5 Minuten)"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:    content.title,
						Message:  content.message,
						Intent:   alert.IntentOk,
						Duration: 5 * time.Minute,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Fehler (5 Minuten)"),
				ui.PrimaryButton(func() {
					content := getRandomAlertContent()
					alert.ShowBannerMessage(wnd, alert.Message{
						Title:    content.title,
						Message:  content.message,
						Intent:   alert.IntentError,
						Duration: 5 * time.Minute,
					})
				}).Title("Alert anzeigen"),
			),
		),
		ui.GridCell(
			ui.VStack(
				ui.Text("Technischer Fehler"),
				ui.PrimaryButton(func() {
					alert.ShowBannerError(wnd, fmt.Errorf("[abc123] Ich bin ein Fehler"))
				}).Title("Alert anzeigen"),
			),
		),
	).Columns(3).Gap(ui.L16)
}

func getRandomAlertContent() AlertContent {
	index := rand.Intn(len(messages))
	msg := messages[index]
	msg.message = fmt.Sprintf("%s\n(id: %d)", msg.message, rand.Intn(999999999))
	return msg
}
