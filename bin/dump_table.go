package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Velocidex/ordereddict"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

var (
	dumpCommand = app.Command(
		"dump", "Dump table.")

	dumpCommandFileArg = dumpCommand.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))

	dumpCommandTableName = dumpCommand.Arg(
		"table", "The name of the table to dump").
		Required().String()
)

func doDump() {
	eseCtx, err := parser.NewESEContext(*dumpCommandFileArg)
	kingpin.FatalIfError(err, "Unable to open ese file")

	catalog, err := parser.ReadCatalog(eseCtx)
	kingpin.FatalIfError(err, "Unable to open ese file")

	err = catalog.DumpTable(*dumpCommandTableName, func(row *ordereddict.Dict) error {
		serialized, err := json.Marshal(row)
		if err != nil {
			return err
		}
		fmt.Printf("%v\n", string(serialized))

		return nil
	})
	kingpin.FatalIfError(err, "Unable to open ese file")
}

func init() {
	commandHandlers = append(commandHandlers, func(command string) bool {
		switch command {
		case dumpCommand.FullCommand():
			doDump()
		default:
			return false
		}
		return true
	})
}
