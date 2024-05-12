package cmd

import (
	"github.com/sanurb/mkrepo/internal/github"
	"github.com/spf13/cobra"
)

var orgListLimit int

var listOrgsCmd = &cobra.Command{
	Use:   "list-orgs",
	Short: "List GitHub organizations",
	Long:  `Lists the GitHub organizations associated with the authenticated user.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return github.ListOrganizations(orgListLimit)
	},
}

func init() {
	rootCmd.AddCommand(listOrgsCmd)
	listOrgsCmd.Flags().IntVarP(&orgListLimit, "limit", "L", 30, "Maximum number of organizations to list")
}
