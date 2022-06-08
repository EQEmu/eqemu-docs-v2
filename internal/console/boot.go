package console

import "github.com/spf13/cobra"

func Run() error {
	cmd := &cobra.Command{Use: "eqemu-docs"}
	cmd.SetHelpFunc(helpFunc)

	// eventually move this to google wire
	commands := []*cobra.Command{
		NewTestCommand().Command(),
		NewSyncDbSchemaConfigFromFilesCommand().Command(),
		NewQuestApiDocGeneratorCommand().Command(),
		NewDbGenerateDocsCommand().Command(),
		NewGMCommandsDocsGenerateCommand().Command(),
		NewServerRulesDocsGenerateCommand().Command(),
	}

	cmd.AddCommand(commands...)

	return cmd.Execute()
}
