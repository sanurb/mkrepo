package cmd

import (
	"fmt"
	"os"

	"github.com/sanurb/mkrepo/internal/config"
	"github.com/sanurb/mkrepo/internal/github"
	"github.com/spf13/cobra"
)

var cfg config.RepoConfig

var rootCmd = &cobra.Command{
	Use:   "mkrepo <repoName> [templateName] [description]",
	Short: "mkrepo creates a new repository from the command line.",
	Long: `mkrepo: Ready to go repos from the CLI 🚀💡
Creates a new repository, optionally based on a template and with a specified description.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.RepoName = args[0]
		if err := github.ProcessRepository(cfg); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.TemplateName, "template", "t", "bare-minimum", "Template repository name")
	rootCmd.PersistentFlags().StringVarP(&cfg.Description, "description", "d", "Short Sweet Headline 🎇🎉", "Description of the repository")
	rootCmd.PersistentFlags().BoolVarP(&cfg.IsPublic, "public", "p", false, "Make the repository public (default is private)")
}

func Execute() error {
	return rootCmd.Execute()
}
