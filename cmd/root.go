package cmd

import (
	"fmt"
	"os"

	"github.com/RoseSecurity/kuzco/internal"
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
	Long:  `Intelligently analyze your Terraform and OpenTofu configurations to receive personalized recommendations for boosting efficiency, security, and performance.`,
	Run:   runAnalyzer,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.Flags().StringVarP(&filePath, "file", "f", "", "Path to the Terraform and OpenTofu file (required)")
	rootCmd.Flags().StringVarP(&tool, "tool", "t", "terraform", "Specifies the configuration tooling for configurations. Valid values include: `terraform` and `opentofu`")
	rootCmd.Flags().StringVarP(&model, "model", "m", "llama3.2", "LLM model to use for generating recommendations")
	rootCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "User prompt for guiding the response format of the LLM model")
	rootCmd.Flags().StringVarP(&addr, "address", "a", "http://localhost:11434", "IP Address and port to use for the LLM model (ex: http://localhost:11434)")
}

func runAnalyzer(cmd *cobra.Command, args []string) {
	// Check if the required flag is set
	if filePath == "" {
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
		return
	}

	// Validate that the specified model exists in Ollama
	if err := internal.ValidateModel(model, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Model validation error: %v\n", err)
		os.Exit(1)
	}

	// Proceed with the main logic if all required flags are set
	if err := internal.Run(filePath, tool, model, prompt, addr); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
