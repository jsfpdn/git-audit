// Package server manages the git-audit gRPC server.
package server

import (
	"context"

	"github.com/jsfpdn/git-audit/pkg/changelog"
	pb "github.com/jsfpdn/git-audit/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GitAuditServer is a gRPC server.
type GitAuditServer struct {
	client changelog.Client

	pb.UnimplementedChangelogServiceServer
}

// NewGitAuditServer returns a gRPC server that wraps the GitAudit client.
// This server than has to be registered by the `RegisterChangelogServiceServer`.
func NewGitAuditServer(client changelog.Client) *GitAuditServer {
	return &GitAuditServer{client: client}
}

// GetChangelog returns the changelog of repository specified by pb.ChangelogRequest sent.
func (s *GitAuditServer) GetChangelog(ctx context.Context, req *pb.ChangelogRequest) (*pb.ChangelogResponse, error) {
	clog, err := s.client.GetChangelog(ctx, req.Owner, req.Repo, req.SHA)
	if err != nil {
		// TODO: refactor changelog.client so that it returns proper status codes.
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	pbCLog := make([]*pb.Commit, len(clog))
	for i, c := range clog {
		pbCLog[i] = &pb.Commit{SHA: c.SHA, Message: c.Message}
	}

	return &pb.ChangelogResponse{
		Request: req,
		Commits: pbCLog,
	}, nil
}
