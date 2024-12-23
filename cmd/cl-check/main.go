package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kasey/go-legnahc/changelog"
)

func popFlags() (*changelog.Config, error) {
	c := &changelog.Config{RepoConfig: changelog.RepoConfig{Owner: "prysmaticlabs", Repo: "prysm", MainRev: "origin/develop"}}
	flag.StringVar(&c.RepoPath, "repo", "", "Path to the git repository")
	flag.StringVar(&c.ChangesDir, "changelog-dir", "changelog", "Path to the directory containing changelog fragments for each commit")
	flag.Parse()
	if c.RepoPath == "" {
		return c, fmt.Errorf("repo is required")
	}
	return c, nil
}

func main() {
	ctx := context.Background()
	cfg, err := popFlags()
	if err != nil {
		errExit(err)
	}
	if err := ensureChangelog(ctx, cfg); err != nil {
		errExit(err)
	}
}

func errExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func ensureChangelog(ctx context.Context, cfg *changelog.Config) error {
	parent, commits, err := changelog.BranchCommits(cfg, cfg.RepoConfig.MainRev, "HEAD")
	if err != nil {
		return err
	}
	fmt.Printf("upstream branch parent commit: %s\n", parent.Id())
	tail := commits[len(commits)-1]
	log.Printf("looking for changelog fragment between upstream commit %s and HEAD %s", parent.Id(), tail.Id())
	frag, err := changelog.FindFragment(cfg.ChangesDir, *parent, *tail)
	if err != nil {
		return fmt.Errorf("could not find changelog fragment in branch: %w", err)
	}
	fmt.Printf("found fragment path: %s\n", frag.Path)
	return nil
}
