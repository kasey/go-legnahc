package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/kasey/go-legnahc/changelog"
)

func popFlags() (*changelog.Config, error) {
	c := &changelog.Config{RepoConfig: changelog.RepoConfig{Owner: "prysmaticlabs", Repo: "prysm"}}
	flag.StringVar(&c.RepoPath, "repo", "", "Path to the git repository")
	flag.StringVar(&c.ChangesDir, "changelog-dir", "changelog", "Path to the directory containing changelog fragments for each commit")
	flag.StringVar(&c.Tag, "tag", "", "Tag anchor changelog")
	flag.StringVar(&c.PreviousPath, "prev", "CHANGELOG.md", "Path to current changelog in the repo. This will be pulled from HEAD")
	flag.BoolVar(&c.Cleanup, "cleanup", false, "Remove the changelog fragment files after generating the changelog")
	flag.Parse()
	if c.RepoPath == "" {
		return c, fmt.Errorf("repo is required")
	}
	if c.Tag == "" {
		return c, fmt.Errorf("tag is required")
	}
	if c.PreviousPath == "" {
		return c, fmt.Errorf("prev is required")
	}
	return c, nil
}

func main() {
	ctx := context.Background()
	cfg, err := popFlags()
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := generateChangelog(ctx, cfg); err != nil {
		fmt.Println(err)
		return
	}
}

func generateChangelog(ctx context.Context, cfg *changelog.Config) error {
	out, err := changelog.Release(ctx, cfg)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
