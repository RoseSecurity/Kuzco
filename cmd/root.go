package cmd

import (
	"fmt"
	"os"

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
	Run:   Banner,
}

func init() {
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(recommendCmd)
	rootCmd.AddCommand(fixCmd)
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
		fmt.Fprintf(os.Stderr, "Model validation error: %v\n", err)
		os.Exit(1)
	}
	options.FontColor = []figlet4go.Color{
		color, // Magenta
	}
	banner, err := ascii.Render("Kuzco")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating ASCII banner: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(banner)
	cmd.Help() // Print help to explain the required flags
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
