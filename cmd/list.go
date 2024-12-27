// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd commands lists models
var listCmd = &cobra.Command{
	Use:                "list",
	Short:              "Lists available Ollama models",
	Long:               `Queries and lists available Ollama models for Kuzco use`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
