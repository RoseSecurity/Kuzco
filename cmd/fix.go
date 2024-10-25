package cmd

import (
	"fmt"
	"os"

	"github.com/RoseSecurity/kuzco/internal"
	u "github.com/RoseSecurity/kuzco/pkg/utils"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:     "fix",
	Short:   "Diagnose configuration errors",
	Long:    `This command analyzes and diagnoses Terraform configuration errors`,
	Example: "kuzco fix",
	Run:     Diagnose,
}

func init() {
	fixCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the Terraform and OpenTofu file (required)")
	fixCmd.Flags().StringVarP(&tool, "tool", "t", "terraform", "Specifies the configuration tooling for configurations. Valid values include: `terraform` and `opentofu`")
	fixCmd.Flags().StringVarP(&model, "model", "m", "llama3.2", "LLM model to use for generating recommendations")
	fixCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
}

func Diagnose(cmd *cobra.Command, args []string) {
	// Validate that the specified model exists in Ollama
	if err := internal.ValidateModel(model, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Model validation error: %v\n", err)
		os.Exit(1)
	}

	config, err := os.ReadFile(filePath)
	if err != nil {
		u.LogErrorAndExit(err)
		return
	}
	recommendations, err := internal.GetRecommendations(model, tool, prompt, addr)
}
