package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

var (
	pageCommand = app.Command(
		"page", "Dump information about a database page")

	pageCommandFileArg = pageCommand.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))

	pageCommandPageNumber = pageCommand.Arg(
		"page_number", "The page to inspect",
	).Required().Int64()
)

func doPage() {
	eseCtx, err := parser.NewESEContext(*pageCommandFileArg)
	kingpin.FatalIfError(err, "Unable to open ese file")

	parser.DumpPage(eseCtx, *pageCommandPageNumber)
}

func init() {
	commandHandlers = append(commandHandlers, func(command string) bool {
		switch command {
		case pageCommand.FullCommand():
			doPage()
		default:
			return false
		}
		return true
	})
}
