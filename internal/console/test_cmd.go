package console

import (
	"fmt"
	"github.com/EQEmu/eqemu-docs-v2/config"
	"github.com/k0kubun/pp"
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

	pp.Println(config.GetMkDocsDbSchemaNavTables())

	//fmt.Println(config.GetMkDocsDatabaseConfig())
}

func (c *TestCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
