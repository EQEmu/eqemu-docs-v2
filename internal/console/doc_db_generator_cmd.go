package console

import (
	"fmt"
	"github.com/spf13/cobra"
)

type DocDbSchemaGeneratorCommand struct {
	command *cobra.Command
}

func (c *DocDbSchemaGeneratorCommand) Command() *cobra.Command {
	return c.command
}

func NewDocDbSchemaGeneratorCommand() *DocDbSchemaGeneratorCommand {
	i := &DocDbSchemaGeneratorCommand{
		command: &cobra.Command{
			Use:   "doc:db-schema-gen-file",
			Short: "Generates database documentation from configuration file",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

func (c *DocDbSchemaGeneratorCommand) Handle(cmd *cobra.Command, _ []string) {
	fmt.Println("Hello world")
}

func (c *DocDbSchemaGeneratorCommand) Validate(_ *cobra.Command, _ []string) error {
	return nil
}
