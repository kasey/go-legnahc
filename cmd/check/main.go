package check

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/kasey/go-legnahc/changelog"
)

func parseArgs(args []string) (*changelog.Config, error) {
	flags := flag.NewFlagSet("check", flag.ContinueOnError)
	c := &changelog.Config{RepoConfig: changelog.RepoConfig{Owner: "prysmaticlabs", Repo: "prysm", MainRev: "origin/develop"}}
	flags.StringVar(&c.RepoPath, "repo", "", "Path to the git repository")
	flags.StringVar(&c.ChangesDir, "changelog-dir", "changelog", "Path to the directory containing changelog fragments for each commit")
	flags.StringVar(&c.RepoConfig.MainRev, "main-rev", "origin/develop", "Main branch tip revision")
	flags.Parse(args)
	if c.RepoPath == "" {
		return c, fmt.Errorf("repo is required")
	}
	return c, nil
}

func Run(ctx context.Context, args []string) error {
	cfg, err := parseArgs(args)
	if err != nil {
		return err
	}
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
