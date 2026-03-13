// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-colorable"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// validClaudeModelPrefixes contains the known valid Claude model prefixes.
var validClaudeModelPrefixes = []string{
	"claude-sonnet-",
	"claude-haiku-",
	"claude-opus-",
}

// IsClaudeModel returns true if the model name indicates Claude Code should be used.
func IsClaudeModel(model string) bool {
	return strings.HasPrefix(model, "claude")
}

// ValidateClaudeModel checks if the specified model matches a known Claude model prefix.
func ValidateClaudeModel(model string) error {
	for _, prefix := range validClaudeModelPrefixes {
		if strings.HasPrefix(model, prefix) {
			return nil
		}
	}
	return fmt.Errorf("unknown Claude model '%s'. Valid model prefixes: %v", model, validClaudeModelPrefixes)
}

// GetClaudeRecommendations generates recommendations using Claude Code's CLI.
func GetClaudeRecommendations(resourceType string, unusedAttrs []string, tool string, prompt string) (string, error) {
	if err := validateClaudeInstalled(); err != nil {
		return "", err
	}

	formattedPrompt := buildClaudePrompt(resourceType, unusedAttrs, tool, prompt)
	return execClaude(formattedPrompt)
}

// GetClaudeFix sends a diagnostic prompt to Claude Code for fixing configuration errors.
func GetClaudeFix(prompt string) (string, error) {
	if err := validateClaudeInstalled(); err != nil {
		return "", err
	}

	return execClaude(prompt)
}

func validateClaudeInstalled() error {
	_, err := exec.LookPath("claude")
	if err != nil {
		return fmt.Errorf("claude CLI not found in PATH. Install Claude Code: https://docs.anthropic.com/en/docs/claude-code")
	}
	return nil
}

func buildClaudePrompt(resourceType string, unusedAttrs []string, tool string, userPrompt string) string {
	tool = cases.Title(language.English, cases.NoLower).String(tool)

	if userPrompt != "" {
		return fmt.Sprintf(`Unused attributes for '%s' resource '%s': %v

'%s'

Example output:
resource "type" "name" {
  # Enables feature X for improved security
  attribute1 = value1

  # Optimizes performance by setting Y
  attribute2 = value2
}`, tool, resourceType, unusedAttrs, userPrompt)
	}

	return fmt.Sprintf(`Unused attributes for '%s' resource '%s': %v

For each attribute that should be enabled:
1. Recommend it as Terraform code
2. Add a brief comment explaining its purpose
3. Format as a resource block with comments above uncommented parameters

Example output:
resource "type" "name" {
  # Enables feature X for improved security
  attribute1 = value1

  # Optimizes performance by setting Y
  attribute2 = value2
}`, tool, resourceType, unusedAttrs)
}

func execClaude(prompt string) (string, error) {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	_ = s.Color("magenta")
	s.Writer = colorable.NewColorableStdout()
	s.Suffix = " Pull the lever, Kronk!"
	fmt.Printf("%s%s%s ", ColorBold+ColorGreen, s.Suffix, ColorReset)
	s.Start()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, "claude", "--output-format", "text")
	cmd.Stdin = strings.NewReader(prompt)
	output, err := cmd.CombinedOutput()

	s.Stop()
	fmt.Println()

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("claude command timed out after 2 minutes")
		}
		return "", fmt.Errorf("error executing claude: %v\n%s", err, string(output))
	}

	return strings.TrimSpace(string(output)), nil
}
