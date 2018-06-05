package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/peterhellberg/link"
)

// Don't forget to close the response.
func request(url string) (*http.Response, error) {
	// Create request.
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("create GET request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("make request to %s: %v", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("reuest to %s has failed with status: %s", url, resp.Status)
	}

	return resp, nil
}

func pageNumber(uri string) (int, error) {
	resp, err := request(uri)
	if err != nil {
		return 0, fmt.Errorf("request page number: %v", err)
	}
	defer resp.Body.Close()

	linkHeader := link.ParseResponse(resp)

	// Get url from link Header with 'last' relation.
	if lastRelation, ok := linkHeader["last"]; ok {
		lastPageURI, err := url.Parse(lastRelation.URI)
		if err != nil {
			return 0, fmt.Errorf("error while parse last page uri: %v", err)
		}

		// ? If empty string then error raise when strconv will parse to int
		page := lastPageURI.Query().Get("page")

		res, err := strconv.Atoi(page)
		if err != nil {
			return 0, fmt.Errorf("error while parse page param in last page URI: %v", err)
		}
		return res, nil
	}

	return 0, nil
}
