package main

import (
	"bufio"
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

var prCommitRE = regexp.MustCompile(`^(.*) \(#(\d+)\)$`)

type commit struct {
	pr    int
	title string
	sha   string
}

func (c commit) bulletPoint() string {
	return "- " + strings.TrimRight(c.title, " .") + "." +
		" [[PR]](" + c.prURL() + ")"
}

func (c commit) prURL() string {
	return "https://github.com/prysmaticlabs/prysm/pull/" + strconv.Itoa(c.pr)
}

func (c commit) getLabel(ctx context.Context, gh *ghwrap) (string, error) {
	return gh.prLabel(ctx, c.pr)
}

func commitsAfter(repoPath string, version string) ([]commit, error) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, err
	}
	/*
		w, err := r.Worktree()
		if err != nil {
			return nil, err
		}
		if err = w.Pull(&git.PullOptions{RemoteName: "origin"}); err != nil {
			return nil, err
		}
	*/
	endTag, err := r.Tag(version)
	if err != nil {
		return nil, err
	}
	endObj, err := r.CommitObject(endTag.Hash())
	if err != nil {
		return nil, err
	}
	endTime := endObj.Author.When
	iter, err := r.Log(&git.LogOptions{Since: &endTime})
	if err != nil {
		return nil, err
	}
	commits := make([]commit, 0)
	err = iter.ForEach(func(c *object.Commit) error {
		cm, err := parseCommit(c)
		if err != nil {
			return err
		}
		commits = append(commits, cm)
		return nil
	})
	return commits, err
}

func parseCommit(c *object.Commit) (commit, error) {
	_, firstLine, err := bufio.ScanLines([]byte(c.Message), true)
	if err != nil {
		return commit{}, err
	}
	first := string(firstLine)
	if !prCommitRE.MatchString(first) {
		return commit{}, fmt.Errorf("could not parse format of commit message: %s", c.Message)
	}
	m := prCommitRE.FindStringSubmatch(first)
	pri, err := strconv.Atoi(m[2])
	return commit{
		sha:   c.Hash.String(),
		title: m[1],
		pr:    pri,
	}, err
}
