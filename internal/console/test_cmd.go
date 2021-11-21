package console

import (
	"fmt"
	"github.com/spf13/cobra"
)

type TestCommand struct {
	command *cobra.Command
}

func (c *TestCommand) Command() *cobra.Command {
	return c.command
}

func NewTestCommand() *TestCommand {
	i := &TestCommand{
		command: &cobra.Command{
			Use:   "test:test",
			Short: "Command used for temporary non-committal code testing",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *TestCommand) Handle(_ *cobra.Command, _ []string) {
	fmt.Println("Hello world")
}

func (c *TestCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
