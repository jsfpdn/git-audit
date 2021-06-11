// Package cmd providing necessary commands.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// TODO: utilize the given flag
	_ = serverCmd.Flags().IntP("port", "p", 8080, "port for the gRPC server to listen on")
}
