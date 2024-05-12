package cmd

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/sanurb/mkrepo/internal/config"
	"github.com/sanurb/mkrepo/internal/github"
	"github.com/spf13/cobra"
	"os"
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
		spinner, _ := pterm.DefaultSpinner.Start("Processing your repository...")
		if err := github.ProcessRepository(cfg); err != nil {
			spinner.Fail("Error: " + err.Error()) // Use PTerm to show error
			os.Exit(1)
		}
		spinner.Success("Repository created and setup successfully!")
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfg.TemplateName, "template", "t", "bare-minimum", "Template repository name")
	rootCmd.PersistentFlags().StringVarP(&cfg.Description, "description", "d", "Short Sweet Headline 🎇🎉", "Description of the repository")
	rootCmd.PersistentFlags().BoolVarP(&cfg.IsPublic, "public", "p", false, "Make the repository public (default is private)")
	rootCmd.PersistentFlags().StringVarP(&cfg.OrgName, "org", "o", "", "Specify the organization name under which to create the repository")
}

func Execute() error {
	_ = pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("mkrepo", pterm.NewStyle(pterm.FgCyan)),
	).Render()
	return rootCmd.Execute()
}
