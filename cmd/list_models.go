package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"

	u "github.com/RoseSecurity/kuzco/pkg/utils"
)

// ListModelsCmd lists Ollama models
var listModelsCmd = &cobra.Command{
	Use:                "models",
	Short:              "Lists available Ollama models",
	Long:               `This command lists Ollama models`,
	Example:            "list models",
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
	Run: func(cmd *cobra.Command, args []string) {
		models, err := u.ListModels(addr)
		if err != nil {
			u.LogErrorAndExit(err)
		}

		// Prepare a markdown list of models
		var modelList strings.Builder
		modelList.WriteString("### Available Ollama Models\n\n")

		if len(models) == 0 {
			modelList.WriteString("No models available.\n")
		} else {
			for _, model := range models {
				modelList.WriteString(fmt.Sprintf("- %s\n", model))
			}
		}

		renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
		if err != nil {
			u.LogErrorAndExit(fmt.Errorf("error creating markdown renderer: %v", err))
		}

		rendered, err := renderer.Render(modelList.String())
		if err != nil {
			u.LogErrorAndExit(fmt.Errorf("error rendering markdown: %v", err))
		}

		fmt.Print(rendered)
	},
}

func init() {
	listModelsCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
	listCmd.AddCommand(listModelsCmd)
	listModelsCmd.DisableFlagParsing = false
}
