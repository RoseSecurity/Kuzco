// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd commands lists models
var listCmd = &cobra.Command{
	Use:                "list",
	Short:              "Lists available LLM models",
	Long:               `Queries and lists available LLM models for Kuzco use`,
	FParseErrWhitelist: struct{ UnknownFlags bool }{UnknownFlags: false},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
