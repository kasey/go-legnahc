package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"regexp"
	"time"
)

var versionRE = regexp.MustCompile(`^#+ \[(v\d+\.\d+\.\d+)\]`)

const preamble = `# Changelog

All notable changes to this project will be documented in this file.

The format is based on Keep a Changelog, and this project adheres to Semantic Versioning.`

const preambleLines = 5

var changelogSections = []string{"Added", "Changed", "Deprecated", "Removed", "Fixed", "Security", "Uncategorized"}

var changelogNames = map[string]string{
	"changelog/added":      "Added",
	"changelog/changed":    "Changed",
	"changelog/deprecated": "Deprecated",
	"changelog/removed":    "Removed",
	"changelog/fixed":      "Fixed",
	"changelog/security":   "Security",
}

type config struct {
	Repo     string
	Tag      string
	Previous string
}

type changelog struct {
	prevBody    string
	prevVersion string
	cfg         config
	sections    map[string][]commit
	repo        repoConfig
}

type previous struct {
	path string
	fd   *os.File
}

func versionFromLine(line string) string {
	if !versionRE.MatchString(line) {
		return ""
	}
	return versionRE.FindStringSubmatch(line)[1]
}

func loadBase(cfg config) (*changelog, error) {
	c := &changelog{
		cfg:      cfg,
		sections: make(map[string][]commit),
		repo:     prysmRepo,
	}
	fd, err := os.Open(cfg.Previous)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scn := bufio.NewScanner(fd)
	scn.Text()
	for scn.Scan() {
		line := scn.Text()
		if c.prevVersion != "" {
			c.prevBody += "\n" + line
			continue
		}
		v := versionFromLine(line)
		if v != "" {
			c.prevVersion = v
			c.prevBody += line
		}
	}
	if c.prevVersion == "" {
		return nil, fmt.Errorf("no version found")
	}
	return c, nil
}

func (c *changelog) generate(ctx context.Context) (string, error) {
	gh, err := getClient(ctx)
	if err != nil {
		return "", err
	}
	commits, err := commitsAfter(c.cfg.Repo, c.prevVersion)
	if err != nil {
		return "", err
	}
	for _, ct := range commits {
		l, err := ct.getLabel(ctx, gh)
		if err != nil {
			return "", err
		}
		if l == "" {
			l = "Uncategorized"
		}
		c.sections[l] = append(c.sections[l], ct)
	}
	body := preamble + "\n\n" + c.header()
	for _, s := range changelogSections {
		cts, ok := c.sections[s]
		if !ok || len(cts) == 0 {
			continue
		}
		body += c.formatSection(s, cts)
	}
	return body + "\n" + c.prevBody + "\n", nil
}

func (c *changelog) formatSection(name string, commits []commit) string {
	section := "\n\n### " + name + "\n"
	for _, c := range commits {
		section += "\n" + c.bulletPoint()
	}
	return section
}

func (c *changelog) header() string {
	// ## [v5.1.1](https://github.com/prysmaticlabs/prysm/compare/v5.1.0...v5.1.1) - 2024-10-15
	return fmt.Sprintf("## [%s](https://github.com/%s/%s/compare/%s...%s) - %s",
		c.cfg.Tag,
		c.repo.Owner, c.repo.Repo,
		c.prevVersion, c.cfg.Tag,
		time.Now().Format("2006-01-02"),
	)
}
