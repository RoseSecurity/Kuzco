// Copyright (c) RoseSecurity
// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

// PrettyPrint formats and displays markdown content with styled headers and recommendations.
// It applies custom styling using lipgloss and renders markdown using glamour.
// The output includes a header, the rendered markdown content, and a footer.
func PrettyPrint(markdownContent string) {
	// Create a style for headers
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF79C6")).
		Background(lipgloss.Color("#282A36")).
		Padding(1, 2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#FF79C6")).
		Align(lipgloss.Center).
		Width(75)

	// Create a style for recommendations
	commentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#8BE9FD")).
		Italic(true).
		Margin(1, 2)

	header := headerStyle.Render("Kuzco Recommendations")

	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		log.Fatal(err)
	}

	renderedContent, err := renderer.Render(markdownContent)
	if err != nil {
		log.Fatal(err)
	}

	footer := commentStyle.Render("Make sure to adjust these attributes according to your specific use case.")

	output := strings.Join([]string{header, renderedContent, footer}, "\n\n")

	fmt.Println(output)
}
