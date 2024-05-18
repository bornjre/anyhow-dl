package code

import (
	"github.com/alecthomas/kong"
	"github.com/bornjre/anyhow-dl/code/download"
	"github.com/bornjre/anyhow-dl/code/serve"
)

var CLI struct {
	Download struct {
	} `cmd:"" help:"Download files."`

	Serve struct {
	} `cmd:"" help:"Serve files"`
}

func Run() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "download":
		download.DownloadRun()

	case "serve":
		serve.ServeRun()

	default:
		panic(ctx.Command())
	}
}
