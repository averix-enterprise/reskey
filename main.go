package main

import (
	"embed"
	"fmt"
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

	var data interface{} = &Test{}
	backend.LoadOrCreateFile("test.json", &data, Test{Name: "Maga"})
	test := data.(*Test)
	fmt.Println(test.Name)

	backend.Initialize()

	err := wails.Run(&options.App{
		Title:  "Resolution Key",
		Width:  1024,
		Height: 768,
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
