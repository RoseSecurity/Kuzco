package internal

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

func prettyPrint(markdownContent string) {
	// Create a style for headers
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF79C6")).
		Background(lipgloss.Color("#282A36")).
		Padding(1, 2).
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#FF79C6"))

	// Create a style for recommendations
	commentStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#8BE9FD")).
		Italic(true).
		Margin(1, 2)

	header := headerStyle.Render("# Kuzco Recommendations")

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
