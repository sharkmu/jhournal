package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	list := tview.NewList().
		AddItem("Add", "Record a new entry", 'a', func() {
			addEntry()
		}).
		AddItem("View entries", "See all your previous journal entries", 'v', func() {
			listEntries()
		}).
		AddItem("Settings", "Open settings, to configure your configurations", 's', func() {
			openSettings()
		}).
		AddItem("Quit", "Press to exit", 'q', func() {
			app.Stop()
		})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func addEntry() {
	println("TODO")
}

func listEntries() {
	println("TODO")
}

func openSettings() {
	println("TODO")
}
