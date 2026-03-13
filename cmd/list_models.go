// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/spf13/cobra"

	u "github.com/RoseSecurity/kuzco/pkg/utils"
)

// ListOllamaModelsCmd lists Ollama models
var listOllamaModelsCmd = &cobra.Command{
	Use:                "models",
	Short:              "Lists available Ollama models",
	Long:               `This command lists Ollama models`,
	Example:            "list models",
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
	Run: func(cmd *cobra.Command, args []string) {
		models, err := u.ListOllamaModels(addr)
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
				fmt.Fprintf(&modelList, "- %s\n", model)
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

// listClaudeModelsCmd lists available Claude models via the Claude Code CLI
var listClaudeModelsCmd = &cobra.Command{
	Use:     "claude-models",
	Short:   "Lists available Claude models",
	Long:    `This command lists Claude models available through the Claude Code CLI`,
	Example: "list claude-models",
	Run: func(cmd *cobra.Command, args []string) {
		claudeModels := []string{
			"claude-sonnet-4-20250514",
			"claude-haiku-4-5-20251001",
			"claude-opus-4-20250514",
		}

		var modelList strings.Builder
		modelList.WriteString("### Available Claude Models\n\n")

		for _, model := range claudeModels {
			fmt.Fprintf(&modelList, "- %s\n", model)
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
	listOllamaModelsCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
	listCmd.AddCommand(listOllamaModelsCmd)
	listCmd.AddCommand(listClaudeModelsCmd)
	listOllamaModelsCmd.DisableFlagParsing = false
}
