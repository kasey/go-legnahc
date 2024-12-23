package main

import (
	"context"
	"testing"

	"github.com/kasey/go-legnahc/changelog"
)

func TestWalk(t *testing.T) {
	cfg := &changelog.Config{
		RepoConfig: changelog.RepoConfig{
			Owner:   "prysmaticlabs",
			Repo:    "prysm",
			MainRev: "origin/develop",
		},
		RepoPath:   "/home/kasey/src/prysmaticlabs/prysm",
		ChangesDir: "changelog",
	}
	err := ensureChangelog(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
}
