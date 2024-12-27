// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

// Mock the HTTP client
func TestLatestRelease(t *testing.T) {
	mockResponse := Release{TagName: "v1.1.0"}
	mockData, _ := json.Marshal(mockResponse)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockData)
	}))
	defer server.Close()

	originalURL := "https://api.github.com/repos/RoseSecurity/Kuzco/releases/latest"
	defer func() { originalURL = originalURL }() // Reset after test

	resp, err := http.Get(server.URL)
	assert.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var release Release
	err = json.Unmarshal(body, &release)
	assert.NoError(t, err)
	assert.Equal(t, "v1.1.0", release.TagName)
}

func TestUpdateKuzco(t *testing.T) {
	// Redirect output to capture the print statements
	var buf bytes.Buffer
	color.Output = &buf

	latestVersion := "1.1.0"
	updateKuzco(latestVersion)

	output := buf.String()
	expectedMessage := "\nYour version of Kuzco is out of date. The latest version is 1.1.0\n\n"
	assert.Contains(t, output, expectedMessage)
}
