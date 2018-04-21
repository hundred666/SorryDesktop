package main

import (
	"github.com/asticode/go-astilectron"
	"flag"
	"github.com/asticode/go-astilog"
	"github.com/asticode/go-astilectron-bootstrap"
)

var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

func main() {
	flag.Parse()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug:          true,
		Homepage:       "index.html",
		MessageHandler: handleMessages,
		RestoreAssets:  RestoreAssets,
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(600),
			Width:           astilectron.PtrInt(600),
		},
	}); err != nil {
		astilog.Fatal(err)
	}
}