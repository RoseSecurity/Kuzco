// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// OllamaModel represents the structure of an Ollama model
type OllamaModel struct {
	Name string `json:"name"`
}

// ListOllamaModels lists available models from Ollama
func ListOllamaModels(addr string) ([]string, error) {
	if _, err := url.Parse(addr); err != nil {
		return nil, fmt.Errorf("invalid address provided: %v", err)
	}

	resp, err := http.Get(fmt.Sprintf("%s/api/tags", addr))
	if err != nil {
		return nil, fmt.Errorf("error fetching models from Ollama: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to retrieve models: status code %d", resp.StatusCode)
	}

	var result struct {
		Models []OllamaModel `json:"models"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	var modelNames []string
	for _, model := range result.Models {
		modelNames = append(modelNames, model.Name)
	}

	return modelNames, nil
}
