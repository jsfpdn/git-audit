// Package cmd providing necessary commands.
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/jsfpdn/git-audit/pkg/changelog"
	"github.com/spf13/cobra"
)

var (
	owner   *string
	repo    *string
	sha     *string
	verbose *bool
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Generate changelog of the provided public GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		client := changelog.NewClient(nil)
		changelog, err := client.GetChangelog(context.Background(), *owner, *repo, *sha)
		if err != nil {
			fmt.Printf("could not get changelog: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("changelog for github.com/%s/%s from %s\n", *owner, *repo, *sha)
		for _, c := range changelog {
			if *verbose {
				fmt.Println(c)
			} else {
				fmt.Printf("%s %s\n", c.SHA, c.FirstLine())
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	owner = clientCmd.Flags().StringP("owner", "o", "", "owner of the repository")
	repo = clientCmd.Flags().StringP("repository", "r", "", "repository name")
	sha = clientCmd.Flags().StringP("sha", "s", "HEAD", "SHA of the commit")
	verbose = clientCmd.Flags().BoolP("verbose", "v", false, "verbose shows full commit message")
	_ = clientCmd.MarkFlagRequired("owner")
	_ = clientCmd.MarkFlagRequired("repository")
}
