// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

// ANSI color codes shared across LLM backends
const (
	ColorReset = "\033[0m"
	ColorGreen = "\033[32m"
	ColorBold  = "\033[1m"
)

// GetRecommendations routes to the appropriate LLM backend based on the model name.
func GetRecommendations(resourceType string, unusedAttrs []string, model string, tool string, prompt string, addr string) (string, error) {
	if IsClaudeModel(model) {
		return GetClaudeRecommendations(resourceType, unusedAttrs, tool, prompt)
	}
	return GetOllamaRecommendations(resourceType, unusedAttrs, model, tool, prompt, addr)
}

// GetFix routes fix prompts to the appropriate LLM backend based on the model name.
func GetFix(prompt, model, addr string) (string, error) {
	if IsClaudeModel(model) {
		return GetClaudeFix(prompt)
	}
	return GetOllamaFix(prompt, model, addr)
}
