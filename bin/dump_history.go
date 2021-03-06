package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Velocidex/ordereddict"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

var (
	dumpHistCmd = app.Command(
		"dump_hist", "Dump dump history.")

	dumpHistCmdFileArg = dumpHistCmd.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))
)

func findHistContainers() {
	eseCtx, err := parser.NewESEContext(*dumpHistCmdFileArg)
	kingpin.FatalIfError(err, "Unable to open ese file")

	catalog, err := parser.ReadCatalog(eseCtx)
	kingpin.FatalIfError(err, "Unable to open ese file")

	err = catalog.DumpTable("Containers", func(row *ordereddict.Dict) error {
		dirVal, pres := row.Get("Directory")
		containerVal, pres := row.Get("Name")
		containerName := fmt.Sprintf("%v", containerVal)
		// Remove Null Unicode
		containerName = strings.Replace(containerName, "\u0000", "", -1)
		if pres {
			dir := fmt.Sprintf("%v", dirVal)
			res := strings.Contains(dir, "History.IE5")
			fmt.Printf("%v\n", res)
			if res {
				containerINT, present := row.Get("ContainerId")
				containerID := fmt.Sprintf("Container_%v", containerINT)
				if present == true {
					err = catalog.DumpTable(containerID, func(row *ordereddict.Dict) error {
						row.Set("Directory", dirVal)
						row.Set("ContainerName", containerName)
						serialized, err := json.Marshal(row)
						if err != nil {
							return nil
						}
						fmt.Printf("%v\n", string(serialized))
						return nil
					})
				}
			}
		}
		return nil
	})
	kingpin.FatalIfError(err, "Unable to find History.IE5")
}

func init() {
	commandHandlers = append(commandHandlers, func(command string) bool {
		switch command {
		case dumpHistCmd.FullCommand():
			findHistContainers()
		default:
			return false
		}
		return true
	})
}
