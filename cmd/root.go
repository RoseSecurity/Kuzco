// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

// Package cmd provides the CLI commands for Kuzco.
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
	dryrun   bool
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
		_ = cmd.Help() // Print help to explain the required flags
		return
	}

	// Validate that the specified model exists
	if internal.IsClaudeModel(model) {
		if err := internal.ValidateClaudeModel(model); err != nil {
			u.LogErrorAndExit(err)
		}
	} else {
		if err := internal.ValidateOllamaModel(model, addr); err != nil {
			u.LogErrorAndExit(err)
		}
	}
	_ = cmd.Help()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		u.LogErrorAndExit(err)
	}
}
