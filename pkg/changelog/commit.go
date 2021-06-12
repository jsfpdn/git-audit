// Package changelog implements basic auditing methods on public repositories.
package changelog

import (
	"fmt"
	"strings"
)

// Commit represents a single commit in a GitHub repository.
type Commit struct {
	SHA, Message string
}

func (c Commit) String() string {
	return fmt.Sprintf("%s %s%s", c.SHA, c.FirstLine(), c.CommitBody())
}

// FirstLine returns the first line of the commit message.
func (c Commit) FirstLine() string {
	lines := strings.Split(c.Message, "\n")
	if len(lines) > 0 {
		return strings.TrimRight(lines[0], "\n")
	}
	return ""
}

// CommitBody returns body of a commit message without the first line returned from FirstLine().
// Body of the message is already shifted to the right with a single '\t'.
func (c Commit) CommitBody() string {
	lines := strings.Split(c.Message, "\n")
	if len(lines) > 2 {
		// Commit message has an empty line between the first line and the following body,
		// which we skip.
		return strings.Join(lines[1:], "\n\t")
	}
	return ""
}
