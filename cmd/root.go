package cmd

import (
	"fmt"

	"github.com/RoseSecurity/kuzco/internal"
	tuiUtils "github.com/RoseSecurity/kuzco/internal/tui/utils"
	u "github.com/RoseSecurity/kuzco/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	filePath string
	tool     string
	model    string
	prompt   string
	addr     string
)

var rootCmd = &cobra.Command{
	Use:   "kuzco",
	Short: "Intelligently analyze your Terraform and OpenTofu configurations",
	Long:  `Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations and fixes for boosting efficiency, security, and performance.`,
	Run:   runAnalyzer,
}

func init() {
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(recommendCmd)
	rootCmd.AddCommand(fixCmd)

	// Disable auto generated string from documentation so that documentation is cleanly built and updated
	rootCmd.DisableAutoGenTag = true
}

func runAnalyzer(cmd *cobra.Command, args []string) {
	var err error

	// Check if the required flag is set
	if filePath == "" {
		fmt.Println()
		err = tuiUtils.PrintStyledText("KUZCO")
		if err != nil {
			u.LogErrorAndExit(err)
		}
		cmd.Help() // Print help to explain the required flags
		return
	}

	// Validate that the specified model exists in Ollama
	if err := internal.ValidateModel(model, addr); err != nil {
		u.LogErrorAndExit(err)
	}
	cmd.Help()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		u.LogErrorAndExit(err)
	}
}
