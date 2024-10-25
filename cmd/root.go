package cmd

import (
	"fmt"
	"os"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

var (
	filePath string
	tool     string
	model    string
	prompt   string
	addr     string
)

var rootCmd = &cobra.Command{
	Use:   "kuzco",
	Short: "Intelligently analyze your Terraform and OpenTofu configurations",
	Long:  `Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations and fixes for boosting efficiency, security, and performance.`,
	Run:   Banner,
}

func init() {
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(recommendCmd)
	rootCmd.AddCommand(fixCmd)
}

func Banner(cmd *cobra.Command, args []string) {
	// Check if the required flag is set
	ascii := figlet4go.NewAsciiRender()
	options := figlet4go.NewRenderOptions()
	color, err := figlet4go.NewTrueColorFromHexString("FF00FF")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating color: %v\n", err)
		os.Exit(1)
	}
	options.FontColor = []figlet4go.Color{
		color, // Magenta
	}
	banner, err := ascii.Render("Kuzco")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating ASCII banner: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(banner)
	cmd.Help() // Print help to explain the required flags
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
