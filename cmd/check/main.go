package check

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/kasey/go-legnahc/changelog"
)

type githubConf struct {
	FragmentListingEnv string // name of env var containing a list of file fragments
}

func parseArgs(args []string) (*changelog.Config, *githubConf, error) {
	flags := flag.NewFlagSet("check", flag.ContinueOnError)
	c := &changelog.Config{RepoConfig: changelog.RepoConfig{Owner: "prysmaticlabs", Repo: "prysm"}}
	flags.StringVar(&c.RepoPath, "repo", "", "Path to the git repository")
	flags.StringVar(&c.ChangesDir, "changelog-dir", "changelog", "Path to the directory containing changelog fragments for each commit")
	flags.StringVar(&c.RepoConfig.MainRev, "main-rev", "origin/develop", "Main branch tip revision")
	flags.StringVar(&c.Branch, "branch", "HEAD", "branch tip revision")
	envCfg := &githubConf{}
	flags.StringVar(&envCfg.FragmentListingEnv, "fragment-env", "", "Name of the environment variable containing a list of changelog fragments")
	flags.Parse(args)
	if envCfg.FragmentListingEnv != "" {
		return nil, envCfg, nil
	}

	if c.RepoPath == "" {
		return c, nil, fmt.Errorf("repo is required")
	}
	return c, nil, nil
}

func Run(ctx context.Context, args []string) error {
	cfg, envCfg, err := parseArgs(args)
	if err != nil {
		return err
	}
	if envCfg != nil {
		return checkFragments(envCfg)
	}
	parent, commits, err := changelog.BranchCommits(cfg, cfg.RepoConfig.MainRev, cfg.Branch)
	if err != nil {
		return err
	}
	fmt.Printf("upstream branch parent commit: %s\n", parent.Id())
	tail := commits[len(commits)-1]
	log.Printf("looking for changelog fragment between upstream commit %s and branch %s %s", parent.Id(), cfg.Branch, tail.Id())
	frag, err := changelog.FindFragment(cfg.ChangesDir, *parent, *tail)
	if err != nil {
		return fmt.Errorf("could not find changelog fragment in branch: %w", err)
	}
	fmt.Printf("found fragment path: %s\n", frag.Path)
	return nil
}

func checkFragments(envCfg *githubConf) error {
	listBlob := os.Getenv(envCfg.FragmentListingEnv)
	if listBlob == "" {
		return fmt.Errorf("no fragments found in env var %s", envCfg.FragmentListingEnv)
	}
	filePaths := strings.Split(listBlob, "\n")
	if len(filePaths) == 0 {
		return fmt.Errorf("no fragments found in env var %s", envCfg.FragmentListingEnv)
	}
	for _, p := range filePaths {
		b, err := os.ReadFile(p)
		if err != nil {
			return fmt.Errorf("could not read fragment file at %s: %w", p, err)
		}
		lines := strings.Split(string(b), "\n")
		parsed := changelog.ParseFragment(lines, "")
		for k, v := range parsed {
			if len(v) == 0 {
				delete(parsed, k)
			}
		}
		if err := changelog.ValidSections(parsed); err != nil {
			return fmt.Errorf("fragment %s is invalid: %w", p, err)
		}
	}
	return nil
}
