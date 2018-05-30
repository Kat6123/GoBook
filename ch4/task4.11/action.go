package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// Read do 'GET' request, parse json and return it as issueInfo.
func (i issue) Read() (issueInfo, error) {
	resp, err := http.Get(i.Url())
	if err != nil {
		return issueInfo{}, fmt.Errorf("error when read issue: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return issueInfo{}, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Is it not obvious?
	var result issueInfo
	err = result.Decode(resp)
	return result, err
}

// Close do 'PATCH' request to change issue state and return issueInfo as the result.
func (i issue) Close() (issueInfo, error) {
	// Create request.
	jsonBody := `{"state": "closed"}`

	req, err := http.NewRequest(
		http.MethodPatch, i.Url(),
		bytes.NewBufferString(jsonBody))
	if err != nil {
		return issueInfo{}, fmt.Errorf("can't create request: %v", err)
	}

	// Set headers.
	req.SetBasicAuth(i.User.Name, i.User.Password)
	// Is it necessary to set content-type?
	req.Header.Set("Content-Type", "application/json")

	// Make request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issueInfo{}, fmt.Errorf("error while doing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return issueInfo{}, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	var result issueInfo
	err = result.Decode(resp)
	return result, err
}

// Create do 'POST' request to create issue and return issueInfo as the result.
func (i issue) Create() (issueInfo, error) {
	// Create request.
	req, err := http.NewRequest(
		http.MethodPost, i.createUrl(),
		bytes.NewBuffer(i.Body))
	if err != nil {
		return issueInfo{}, fmt.Errorf("can't create request: %v", err)
	}

	// Set headers.
	req.SetBasicAuth(i.User.Name, i.User.Password)
	// Is it necessary to set content-type?
	req.Header.Set("Content-Type", "application/json")

	// Make request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issueInfo{}, fmt.Errorf("error while doing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return issueInfo{}, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	var result issueInfo
	err = result.Decode(resp)
	return result, err
}

// Update do 'PATCH' request to update issue with issue body and return issueInfo as the result.
func (i issue) Update() (issueInfo, error) {
	// Create request.
	req, err := http.NewRequest(
		http.MethodPatch, i.Url(),
		bytes.NewBuffer(i.Body))
	if err != nil {
		return issueInfo{}, fmt.Errorf("can't create request: %v", err)
	}

	// Set headers.
	req.SetBasicAuth(i.User.Name, i.User.Password)
	// Is it necessary to set content-type?
	req.Header.Set("Content-Type", "application/json")

	// Make request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return issueInfo{}, fmt.Errorf("error while doing request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return issueInfo{}, fmt.Errorf("get issue failed with status: %s", resp.Status)
	}

	// Parse JSON.
	var result issueInfo
	err = result.Decode(resp)
	return result, err
}
