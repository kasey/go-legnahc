package main

import (
	"context"
	"flag"
	"fmt"
)

func popFlags() (config, error) {
	c := config{}
	flag.StringVar(&c.Repo, "repo", "", "Path to the git repository")
	flag.StringVar(&c.Tag, "tag", "", "Tag anchor changelog")
	flag.StringVar(&c.Previous, "prev", "", "Current changelog file path, new version will be prepended to the contents")
	flag.Parse()
	if c.Repo == "" {
		return c, fmt.Errorf("repo is required")
	}
	if c.Tag == "" {
		return c, fmt.Errorf("tag is required")
	}
	if c.Previous == "" {
		return c, fmt.Errorf("prev is required")
	}
	return c, nil
}

func main() {
	ctx := context.Background()
	if err := generateChangelog(ctx); err != nil {
		fmt.Println(err)
		return
	}
}
func generateChangelog(ctx context.Context) error {
	cfg, err := popFlags()
	if err != nil {
		return err
	}
	chl, err := loadBase(cfg)
	if err != nil {
		return err
	}
	out, err := chl.generate(ctx)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
