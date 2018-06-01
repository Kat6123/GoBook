package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// Read does 'GET' request, parse json and return it as issueInfo.
func (i issue) Read() (*issueInfo, error) {
	resp, err := http.Get(i.url())
	if err != nil {
		return nil, fmt.Errorf("error when read issue: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	return decode(resp)
}

// Close does 'PATCH' request to change issue state and return issueInfo as the result.
func (i issue) Close() (*issueInfo, error) {
	// Create request.
	jsonBody := `{"state": "closed"}`

	req, err := http.NewRequest(
		http.MethodPatch, i.url(),
		bytes.NewBufferString(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("can't create request: %v", err)
	}

	resp, err := makeRequest(req, i.User)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	return decode(resp)
}

// Create does 'POST' request to create issue and return issueInfo as the result.
func (i issue) Create() (*issueInfo, error) {
	// Create request.
	req, err := http.NewRequest(
		http.MethodPost, i.createUrl(),
		i.Body)
	if err != nil {
		return nil, fmt.Errorf("can't create request: %v", err)
	}

	resp, err := makeRequest(req, i.User)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	return decode(resp)
}

// Update does 'PATCH' request to update issue with issue body and return issueInfo as the result.
func (i issue) Update() (*issueInfo, error) {
	// Create request.
	req, err := http.NewRequest(
		http.MethodPatch, i.url(),
		i.Body)
	if err != nil {
		return nil, fmt.Errorf("can't create request: %v", err)
	}

	resp, err := makeRequest(req, i.User)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	return decode(resp)
}

func makeRequest(req *http.Request, u *user) (*http.Response, error) {
	// Set headers.
	req.SetBasicAuth(u.Name, u.Password)
	req.Header.Set("Content-Type", "application/json")

	// Make request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while doing request: %v", err)
	}

	return resp, nil
}
