// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	u "github.com/RoseSecurity/kuzco/pkg/utils"
	"github.com/briandowns/spinner"
	"github.com/mattn/go-colorable"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// OllamaRequest represents a request to the Ollama API
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaResponse represents a response from the Ollama API
type OllamaResponse struct {
	Recommendations string `json:"response"`
}

// sendOllamaRequest sends a prompt to the Ollama API and returns the response text.
func sendOllamaRequest(prompt, model, addr string) (string, error) {
	requestBody := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Initialize and start the spinner
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	_ = s.Color("magenta")
	s.Writer = colorable.NewColorableStdout() // Ensure colors are supported on Windows
	s.Suffix = " Pull the lever, Kronk!"
	fmt.Printf("%s%s%s ", ColorBold+ColorGreen, s.Suffix, ColorReset)
	s.Start()

	// Make the HTTP request
	resp, err := http.Post(fmt.Sprintf("%s/api/generate", addr), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		s.Stop()
		fmt.Println()
		return "", fmt.Errorf("error making request to Ollama API: %v", err)
	}
	defer resp.Body.Close()

	// Stop the spinner after the request is done
	s.Stop()
	fmt.Println()

	// Check HTTP status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("Ollama API returned status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	// Read and decode the response
	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	if strings.TrimSpace(ollamaResp.Recommendations) == "" {
		return "", fmt.Errorf("Ollama returned an empty response")
	}

	return ollamaResp.Recommendations, nil
}

// GetOllamaRecommendations generates recommendations using Ollama's API.
func GetOllamaRecommendations(resourceType string, unusedAttrs []string, model string, tool string, prompt string, addr string) (string, error) {
	tool = cases.Title(language.English, cases.NoLower).String(tool)
	var formattedPrompt string
	if prompt == "" {
		formattedPrompt = fmt.Sprintf(`Unused attributes for '%s' resource '%s': %v

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
	} else {
		formattedPrompt = fmt.Sprintf(`Unused attributes for '%s' resource '%s': %v

'%s'

Example output:
resource "type" "name" {
  # Enables feature X for improved security
  attribute1 = value1

  # Optimizes performance by setting Y
  attribute2 = value2
}`, tool, resourceType, unusedAttrs, prompt)
	}

	return sendOllamaRequest(formattedPrompt, model, addr)
}

// GetOllamaFix sends a pre-formatted diagnostic prompt to Ollama for fixing configuration errors.
func GetOllamaFix(prompt, model, addr string) (string, error) {
	return sendOllamaRequest(prompt, model, addr)
}

// ValidateOllamaModel checks if the specified model exists in Ollama
func ValidateOllamaModel(model, addr string) error {
	// Get a list of available models from Ollama
	resp, err := http.Get(fmt.Sprintf("%s/api/tags", addr))
	if err != nil {
		return fmt.Errorf("error fetching models from Ollama: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to retrieve models: status code %d", resp.StatusCode)
	}

	// Parse the response body
	var modelsResp struct {
		Models []u.OllamaModel `json:"models"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&modelsResp); err != nil {
		return fmt.Errorf("error decoding models response: %v", err)
	}

	// Check if the requested model is in the list of models
	for _, availableModel := range modelsResp.Models {
		if availableModel.Name == model {
			return nil
		}
	}

	// If the model is not found, return an error and list available models
	var availableModelNames []string
	for _, availableModel := range modelsResp.Models {
		availableModelNames = append(availableModelNames, availableModel.Name)
	}
	return fmt.Errorf("model '%s' not found. Available models: %v", model, availableModelNames)
}
