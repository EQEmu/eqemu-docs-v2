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
		NewBotCommandsDocsGenerateCommand().Command(),
		NewGMCommandsDocsGenerateCommand().Command(),
		NewServerRulesDocsGenerateCommand().Command(),
	}

	cmd.AddCommand(commands...)

	return cmd.Execute()
}

func GetStatusValue(statusLevel string) int {
	m := map[string]int{
		"Player":          0,
		"Steward":         10,
		"ApprenticeGuide": 20,
		"Guide":           50,
		"QuestTroupe":     80,
		"SeniorGuide":     81,
		"GMTester":        85,
		"EQSupport":       90,
		"GMStaff":         95,
		"GMAdmin":         100,
		"GMLeadAdmin":     150,
		"QuestMaster":     160,
		"GMAreas":         170,
		"GMCoder":         180,
		"GMMgmt":          200,
		"GMImpossible":    250,
		"Max":             255,
	}

	if _, ok := m[statusLevel]; !ok {
		return 0
	}

	return m[statusLevel]
}
