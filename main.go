package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

type TestCommand struct {
}

func (c TestCommand) Help() string {
	return "Help for the test command"
}

func (c TestCommand) Run(args []string) int {
	log.Print("Command Test successfuly executed")
	return 0
}

func (c TestCommand) Synopsis() string {
	return "This is the test command"
}

func TestCommandFactory() (cli.Command, error) {
	return TestCommand{}, nil
}

func main() {
	c := cli.NewCLI("app", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"test": TestCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
