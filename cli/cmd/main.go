package main

import (
	"fmt"
	"os"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func testaction(c *kingpin.ParseContext) error {
	fmt.Printf("Would have run command %v.\n", c.String())
	return nil
}

func addSubCommand(app *kingpin.Application, name string, description string) {
	app.Command(name, description).Action(testaction)
}

func main() {
	app := kingpin.New("my-app", "My Sample Kingpin App!")
	app.Flag("flag-1", "").String()
	app.Flag("flag-2", "").HintOptions("opt1", "opt2").String()

	// Add some additional top level commands
	addSubCommand(app, "ls", "Additional top level command to show command completion")
	addSubCommand(app, "ping", "Additional top level command to show command completion")
	addSubCommand(app, "nmap", "Additional top level command to show command completion")

	kingpin.MustParse(app.Parse(os.Args[1:]))
}
