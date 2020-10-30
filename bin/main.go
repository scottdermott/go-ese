package main

import (
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"www.velocidex.com/golang/go-ese/parser"
)

// CommandHandler : This is a cmd handler.. honestly linter
type CommandHandler func(command string) bool

var (
	app = kingpin.New("esewebcache",
		"A tool for inspecting ese files.")

	debug = app.Flag("debug", "Enable debug messages").Bool()

	commandHandlers []CommandHandler
)

func main() {
	app.HelpFlag.Short('h')
	app.UsageTemplate(kingpin.CompactUsageTemplate)
	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	if *debug {
		parser.Debug = true
		parser.DebugWalk = true
	}

	for _, commandHandler := range commandHandlers {
		if commandHandler(command) {
			break
		}
	}
}
