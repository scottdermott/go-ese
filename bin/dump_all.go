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
	dumpAllCmd = app.Command(
		"dump_all", "Dump all tables.")

	dumpAllCmdFileArg = dumpAllCmd.Arg(
		"file", "The image file to inspect",
	).Required().OpenFile(os.O_RDONLY, os.FileMode(0666))
)

func findAllContainers() {
	eseCtx, err := parser.NewESEContext(*dumpAllCmdFileArg)
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
			containerINT, present := row.Get("ContainerId")
			containerID := fmt.Sprintf("Container_%v", containerINT)
			if present {
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
		return nil
	})
	kingpin.FatalIfError(err, "Unable to find Containers")
}

func init() {
	commandHandlers = append(commandHandlers, func(command string) bool {
		switch command {
		case dumpAllCmd.FullCommand():
			findAllContainers()
		default:
			return false
		}
		return true
	})
}
