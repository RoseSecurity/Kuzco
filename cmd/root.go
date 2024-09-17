package cmd

import (
	"fmt"
	"os"

	"github.com/RoseSecurity/kuzco/internal"
	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

var (
	filePath string
	model    string
)

var rootCmd = &cobra.Command{
	Use:   "kuzco",
	Short: "Intelligently analyze your Terraform configurations",
	Long:  `Intelligently analyze your Terraform configurations to receive personalized recommendations for boosting efficiency, security, and performance.`,
	Run:   runAnalyzer,
}

func init() {
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the Terraform file (required)")
	rootCmd.Flags().StringVarP(&model, "model", "m", "llama3.1", "LLM model to use for generating recommendations")
}

func runAnalyzer(cmd *cobra.Command, args []string) {
	// Check if the required flag is set
	if filePath == "" {
		ascii := figlet4go.NewAsciiRender()
		options := figlet4go.NewRenderOptions()
		color, err := figlet4go.NewTrueColorFromHexString("FF00FF")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating color: %v\n", err)
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
		return
	}

	// Proceed with the main logic if all required flags are set
	if err := internal.Run(filePath, model); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
