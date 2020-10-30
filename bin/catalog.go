package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

var (
	catalogCommand = app.Command(
		"catalog", "Dump the catalog")
	catalogCommandFileArg = catalogCommand.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))
)

func doCatalog() {
	eseCtx, err := parser.NewESEContext(*catalogCommandFileArg)
	kingpin.FatalIfError(err, "Unable to open ese file")

	catalog, err := parser.ReadCatalog(eseCtx)
	kingpin.FatalIfError(err, "Unable to open ese file")
	fmt.Printf(catalog.Dump())
}

func init() {
	commandHandlers = append(commandHandlers, func(command string) bool {
		switch command {
		case catalogCommand.FullCommand():
			doCatalog()

		default:
			return false
		}
		return true
	})
}
