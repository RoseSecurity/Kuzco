package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
	"github.com/mattn/go-colorable"
)

// ANSI color codes
const (
	ColorReset     = "\033[0m"
	ColorGreen     = "\033[32m"
	ColorYellow    = "\033[33m"
	ColorRed       = "\033[31m"
	ColorBold      = "\033[1m"
	ColorUnderline = "\033[4m"
)

type LlamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type LlamaResponse struct {
	Recommendations string `json:"response"`
}

func GetRecommendations(resourceType string, unusedAttrs []string, model string, addr string) (string, error) {
	prompt := fmt.Sprintf(
		"For the Terraform resource type '%s', the following attributes are unused: %v. Suggest which attributes should be enabled, formatted as Terraform code with the recommended parameters, and include a concise comment for each explaining why it should be used. For each concise comment, format the output as a Terraform resource with each recommendation as a comment above the corresponding parameter.",
		resourceType, unusedAttrs)

	requestBody := LlamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Initialize the spinner with a custom message
	spinnerText := "Yzma’s got nothing on this! Fetching those Terraform recommendations now!"
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Color("magenta")
	s.Writer = colorable.NewColorableStdout() // Ensure colors are supported on Windows
	s.Suffix = " " + spinnerText

	// Start the spinner and print the message
	fmt.Printf("%s%s%s ", ColorBold+ColorGreen, spinnerText, ColorReset)
	s.Start()

	// Make the HTTP request
	resp, err := http.Post(fmt.Sprintf("%s/api/generate", addr), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		s.Stop()
		fmt.Println()
		return "", fmt.Errorf("error making request to API: %v", err)
	}
	defer resp.Body.Close()

	// Stop the spinner after the request is done
	s.Stop()
	fmt.Println()

	var rawResponse bytes.Buffer
	_, err = rawResponse.ReadFrom(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	var llamaResp LlamaResponse
	if err := json.NewDecoder(&rawResponse).Decode(&llamaResp); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	// Print recommendations with color formatting
	fmt.Printf("%sRecommendations:%s\n", ColorBold+ColorYellow, ColorReset)
	fmt.Println(llamaResp.Recommendations)

	return llamaResp.Recommendations, nil
}
