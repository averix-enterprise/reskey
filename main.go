package main

import (
	"embed"
	"reskey/backend"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type Test struct {
	Name string `json:"name"`
}

func main() {
	app := &backend.App{}

	backend.Initialize()

	err := wails.Run(&options.App{
		Title:         "Resolution Key",
		Width:         500,
		Height:        250,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error while trying to start main window:", err.Error())
	}
}
