package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const gitAPI = "https://api.github.com/"

type user struct {
	Name     string
	Password string
}

// Is it wrong that issue contains user?
type issue struct {
	Owner  string
	Repo   string
	Number int
	Body   []byte
	User   *user
}

type issueInfo struct {
	Url       string `json:"html_url"`
	Title     string
	Labels    []string
	State     string
	Body      string
	Assignees []string
}

func (i issue) createUrl() string {
	return fmt.Sprintf(gitAPI+"repos/%s/%s/issues", i.Owner, i.Repo)
}

func (i issue) Url() string {
	return fmt.Sprintf(gitAPI+"repos/%s/%s/issues/%d", i.Owner, i.Repo, i.Number)
}

func (i *issueInfo) Decode(resp *http.Response) error {
	if err := json.NewDecoder(resp.Body).Decode(i); err != nil {
		return fmt.Errorf("json decoding has failed: %v", err)
	}
	return nil
}
