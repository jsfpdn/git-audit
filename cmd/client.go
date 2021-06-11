// Package cmd providing necessary commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Generate changelog of the provided public GitHub repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client called")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	// TODO: utilize the given flags
	_ = clientCmd.Flags().StringP("owner", "o", "", "owner of the repository")
	_ = clientCmd.Flags().StringP("repository", "r", "", "repository name")
	_ = clientCmd.Flags().StringP("commit", "c", "HEAD", "commit to")
	_ = clientCmd.MarkFlagRequired("owner")
	_ = clientCmd.MarkFlagRequired("repository")
}
