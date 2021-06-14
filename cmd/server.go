// Package cmd providing necessary commands.
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/jsfpdn/git-audit/pkg/changelog"
	server "github.com/jsfpdn/git-audit/pkg/server"
	pb "github.com/jsfpdn/git-audit/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port *int

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		client := changelog.NewClient(nil)
		pb.RegisterChangelogServiceServer(grpcServer, server.NewGitAuditServer(client))

		log.Printf("git-audit gRPC server listening on :%d", *port)
		err = grpcServer.Serve(listener)
		if err != nil {
			fmt.Printf("could not serve: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	port = serverCmd.Flags().IntP("port", "p", 8080, "port for the gRPC server to listen on")
}
