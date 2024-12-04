package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v67/github"
)

const tokenEnv = "GITHUB_TOKEN"

var errNoToken = errors.New("no token")

type repoConfig struct {
	Owner string `json:"owner"`
	Repo  string `json:"repo"`
}

var prysmRepo = repoConfig{Owner: "prysmaticlabs", Repo: "prysm"}

type ghwrap struct {
	client *github.Client
}

func (w *ghwrap) prLabel(ctx context.Context, prNum int) (string, error) {
	log.Printf("Retrieving github label for PR #%d", prNum)
	pr, _, err := w.client.PullRequests.Get(ctx, prysmRepo.Owner, prysmRepo.Repo, prNum)
	if err != nil {
		return "", err
	}
	return changelogLabel(pr.Labels), nil
}

func changelogLabel(labels []*github.Label) string {
	for _, l := range labels {
		name := l.GetName()
		cl, ok := changelogNames[name]
		if ok {
			return cl
		}
	}
	return ""
}

func getClient(ctx context.Context) (*ghwrap, error) {
	token, err := getToken()
	if err != nil {
		return nil, err
	}
	client := github.NewClient(nil).WithAuthToken(strings.TrimRight(string(token), " \n"))

	return &ghwrap{
		client: client,
	}, nil
}

// func getToken() (string, error) {
func getToken() ([]byte, error) {
	path := os.Getenv(tokenEnv)
	if path == "" {
		//return "", errors.New("GITHUB_TOKEN not set")
		return nil, errors.New("GITHUB_TOKEN not set")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		//return "", errNoToken
		return nil, errNoToken
	}
	//return string(b), nil
	return b, nil
}
