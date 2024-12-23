package changelog

import (
	"testing"

	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TestLogParse(t *testing.T) {
	fixture := `Fix Deadline Again During Rollback (#14686)

	* fix it again
	
	* CHANGELOG
	`
	var h plumbing.Hash
	o := &object.Commit{
		Hash:    h,
		Message: fixture,
	}
	_, err := parseCommit(o)
	if err != nil {
		t.Fatal(err)
	}
}
