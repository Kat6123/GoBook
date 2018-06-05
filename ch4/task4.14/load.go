package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

const (
	itemsPERPAGE = 30

	githubAPI  = "https://api.github.com/repos/golang/go/"
	contribURL = githubAPI + "contributors"
)

// load contributors and return new slice.
// When reload contributors if number more than  capacity of base array then will make new array,
// otherwise use old slice.
func (contr *contributors) load() error {
	pageNum, err := pageNumber(contribURL)
	if err != nil {
		return fmt.Errorf("extract page number of contributors: %v", err)
	}

	log.Printf("total number of contributors pages %d", pageNum)
	contr.fitPageSize(pageNum)

	for i := 1; i <= pageNum; i++ {
		if err := contr.loadPage(i); err != nil {
			return fmt.Errorf("load %d Page of contributors: %v", i, err)
		}
		log.Printf("load %d Page: total %d contributors", i, len(*contr))
	}

	return nil
}

func (c *contributors) loadPage(num int) error {
	pageURL := contribURL + "?page=" + strconv.Itoa(num)
	resp, err := request(pageURL)
	if err != nil {
		return fmt.Errorf("request contributors: %v", err)
	}
	defer resp.Body.Close()

	temp := make(contributors, 0, itemsPERPAGE)
	if err := json.NewDecoder(resp.Body).Decode(&temp); err != nil {
		return fmt.Errorf("json decoding has failed: %v", err)
	}
	*c = append(*c, temp...)
	return nil
}

func (contrPtr *contributors) fitPageSize(pages int) {
	contribs := *contrPtr

	if cap(contribs) < pages*itemsPERPAGE {
		c = make(contributors, 0, pages*itemsPERPAGE)
	}

	*contrPtr = contribs[:0]
}
