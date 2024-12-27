// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

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
	Example: "kuzco fix -f path/to/config.tf -t terraform",
	Run:     Diagnose,
}

func init() {
	fixCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the Terraform and OpenTofu file (required)")
	fixCmd.Flags().StringVarP(&tool, "tool", "t", "terraform", "Specifies the configuration tooling for configurations. Valid values include: `terraform` and `opentofu`")
	fixCmd.Flags().StringVarP(&model, "model", "m", "llama3.2", "LLM model to use for generating recommendations")
	fixCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
}

func Diagnose(cmd *cobra.Command, args []string) {
	// Ensure file path is provided
	if filePath == "" {
		fmt.Fprintf(os.Stderr, "Error: file path is required. Use the -f flag to specify the configuration file.\n")
		os.Exit(1)
	}

	// Validate that the specified model exists in Ollama
	if err := internal.ValidateModel(model, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Model validation error: %v\n", err)
		os.Exit(1)
	}

	// Read the configuration file content
	config, err := os.ReadFile(filePath)
	if err != nil {
		u.LogErrorAndExit(err)
	}

	if len(config) == 0 {
		fmt.Fprintf(os.Stderr, "Error: configuration file is empty\n")
		os.Exit(1)
	}

	// Generate a formatted prompt for the recommendations function
	prompt := fmt.Sprintf(`Error detected in configuration file '%s':

Configuration File Contents:

'%v'

---

Resolution Steps:
1. Identify the attribute(s) or syntax causing the error.
2. Refer to the Terraform or OpenTofu documentation for valid syntax and attribute usage the resources.
3. Correct the invalid attribute(s), fix the syntax, or remove the invalid attributes if they are unnecessary.
4. Reformat the corrected resource block if needed.

Example Corrected Configuration:
resource "type" "name" {
  # Explanation of attribute1's purpose
  attribute1 = "value1"
  
  # Optional comment for attribute2
  attribute2 = "value2"
}

Please review and update the configuration file as outlined above to resolve the issue.`, filePath, config)

	// Pass the prompt and file content to GetRecommendations
	recommendations, err := internal.GetRecommendations(string(config), nil, model, tool, prompt, addr)
	if err != nil {
		u.LogErrorAndExit(err)
	}

	internal.PrettyPrint(recommendations)
}
