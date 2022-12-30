package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/wailsapp/wails/exp/pkg/options"

	"github.com/wailsapp/wails/exp/pkg/application"
)

func main() {

	app := application.New(options.Application{
		Name:        "Clipboard Demo",
		Description: "A demo of the clipboard API",
		Mac: options.Mac{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	// Create a custom menu
	menu := app.NewMenu()
	menu.AddRole(application.AppMenu)

	setClipboardMenu := menu.AddSubmenu("Set Clipboard")
	setClipboardMenu.Add("Set Text 'Hello'").OnClick(func(ctx *application.Context) {
		success := app.Clipboard().SetText("Hello")
		if !success {
			app.NewInfoDialog().SetMessage("Failed to set clipboard text").Show()
		}
	})
	setClipboardMenu.Add("Set Text 'World'").OnClick(func(ctx *application.Context) {
		success := app.Clipboard().SetText("World")
		if !success {
			app.NewInfoDialog().SetMessage("Failed to set clipboard text").Show()
		}
	})
	setClipboardMenu.Add("Set Text (current time)").OnClick(func(ctx *application.Context) {
		success := app.Clipboard().SetText(time.Now().String())
		if !success {
			app.NewInfoDialog().SetMessage("Failed to set clipboard text").Show()
		}
	})
	getClipboardMenu := menu.AddSubmenu("Get Clipboard")
	getClipboardMenu.Add("Get Text").OnClick(func(ctx *application.Context) {
		result := app.Clipboard().Text()
		app.NewInfoDialog().SetMessage("Got:\n\n" + result).Show()
	})

	clearClipboardMenu := menu.AddSubmenu("Clear Clipboard")
	clearClipboardMenu.Add("Clear Text").OnClick(func(ctx *application.Context) {
		success := app.Clipboard().SetText("")
		if success {
			app.NewInfoDialog().SetMessage("Clipboard text cleared").Show()
		} else {
			app.NewInfoDialog().SetMessage("Clipboard text not cleared").Show()
		}
	})

	app.SetMenu(menu)

	app.NewWindow()

	err := app.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
