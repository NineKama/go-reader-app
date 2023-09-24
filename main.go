package main

import (
	"embed"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func OpenTextFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

//func OpenPDFFile(filePath string) (string, error) {
//	// Create a PDF configuration.
//	config := model.NewDefaultConfiguration()
//
//	// Open the existing PDF file.
//	pdf, err := pdfcpu.ReadFile(filePath, config)
//	if err != nil {
//		return "", err
//	}
//
//	// Extract text content from the PDF.
//	textContent, err := api.ExtractText(pdf)
//	if err != nil {
//		return "", err
//	}
//
//	return textContent, nil
//}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "go-reader-app",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
