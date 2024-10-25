package cmd

import (
	"fmt"
	"os"

	"github.com/RoseSecurity/kuzco/internal"
	"github.com/spf13/cobra"
)

var recommendCmd = &cobra.Command{
	Use:   "recommend",
	Short: "Intelligently analyze your Terraform and OpenTofu configurations",
	Long:  `Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations for boosting efficiency, security, and performance.`,
	Run:   Analyze,
}

func init() {
	recommendCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the Terraform and OpenTofu file (required)")
	recommendCmd.Flags().StringVarP(&tool, "tool", "t", "terraform", "Specifies the configuration tooling for configurations. Valid values include: `terraform` and `opentofu`")
	recommendCmd.Flags().StringVarP(&model, "model", "m", "llama3.2", "LLM model to use for generating recommendations")
	recommendCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "User prompt for guiding the response format of the LLM model")
	recommendCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
}

func Analyze(cmd *cobra.Command, args []string) {
	// Validate that the specified model exists in Ollama
	if err := internal.ValidateModel(model, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Model validation error: %v\n", err)
		os.Exit(1)
	}

	// Proceed with the main logic if all required flags are set
	if err := internal.Run(filePath, tool, model, prompt, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
