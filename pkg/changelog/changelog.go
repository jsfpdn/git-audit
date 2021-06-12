// Package changelog implements basic auditing methods on public repositories.
package changelog

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

// Client wraps a github client.
type Client struct {
	client *github.Client
}

// NewClient returns new github client.
func NewClient(client *http.Client) Client {
	return Client{client: github.NewClient(client)}
}

// GetChangelog of a repository github.com/<owner>/<repo> from <sha> to HEAD on the default branch.
// If sha is HEAD, then the commit at HEAD is returned.
func (c *Client) GetChangelog(ctx context.Context, owner, repo, sha string) ([]Commit, error) {
	repository, _, err := c.client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, fmt.Errorf("could not get repository %s/%s: %v", owner, repo, err)
	}

	defaultBranch := repository.GetDefaultBranch()
	if defaultBranch == "" {
		return nil, fmt.Errorf("could not find commits: %s/%s has no default branch", owner, repo)
	}

	cmp, res, err := c.client.Repositories.CompareCommits(ctx, owner, repo, sha, defaultBranch)
	if err != nil {
		if res.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("no commit with SHA '%s' in %s/%s", sha, owner, repo)
		}
		return nil, fmt.Errorf("could not compare commits: %v", err)
	}
	defer res.Body.Close()

	status := cmp.GetStatus()
	if status != "ahead" && status != "identical" {
		return nil, fmt.Errorf("commit %s has status '%s' (expected 'ahead' or 'identical')", sha, status)
	}

	changelog := []Commit{}
	if status == "ahead" {
		// Get all the commits between the commit with the provided SHA and the HEAD.
		for _, c := range cmp.Commits {
			changelog = append(changelog, Commit{
				SHA:     *c.SHA,
				Message: c.Commit.GetMessage(),
			})
		}
	}

	changelog = append(changelog, Commit{
		SHA:     *cmp.BaseCommit.SHA,
		Message: cmp.BaseCommit.Commit.GetMessage(),
	})

	return changelog, nil
}
