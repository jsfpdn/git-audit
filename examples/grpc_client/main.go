// Package main shows an example of calling the gRPC server from client.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jsfpdn/git-audit/pkg/changelog"
	pb "github.com/jsfpdn/git-audit/proto"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 5 {
		log.Fatalf("usage: %s <gRPC_addr> <owner> <repo> <sha>\n", os.Args[0])
	}

	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not create connection: %v", err)
	}

	client := pb.NewChangelogServiceClient(conn)

	req := &pb.ChangelogRequest{
		Owner: os.Args[2],
		Repo:  os.Args[3],
		SHA:   os.Args[4],
	}
	res, err := client.GetChangelog(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get changelog: %v", err)
	}
	fmt.Println(res.Request)

	for _, c := range res.Commits {
		fmt.Println(changelog.Commit{SHA: c.SHA, Message: c.Message})
	}
}
