package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	// User vars.
	name, pwd string

	// Mode vars.
	read, create, update, close bool

	// Issue vars.
	number      int
	owner, repo string
	body        []byte
	iisue       issue
)

func main() {
	setFlags()
	parseArgs()

	iisue = issue{
		owner, repo,
		number, body,
		&user{name, pwd}}

	info, err := manage()
	if err != nil {
		log.Fatalf("error during issue manage: %v", err)
	}

	fmt.Println(info)
}

func setFlags() {
	// Set user flags.
	flag.StringVar(&name, "u", "", "User name")
	flag.StringVar(&pwd, "p", "", "User password")

	// Set mode flags.
	flag.BoolVar(&read, "read", true, "Read issue")
	flag.BoolVar(&create, "create", false, "Create issue")
	flag.BoolVar(&update, "update", false, "Update issue")
	flag.BoolVar(&close, "close", false, "Close issue")

	// Set issue flags.
	flag.IntVar(&number, "n", 0, "Issue number")
	flag.StringVar(&owner, "o", "", "Repo owner")
	flag.StringVar(&repo, "r", "", "Repo")
}

func parseArgs() {
	flag.Parse()

	if flag.NArg() == 0 {
		// open editor
	} else {
		body = []byte(flag.Arg(0))
	}
}

func manage() (issueInfo, error) {
	switch {
	case create:
		info, err := iisue.Create()
		if err != nil {
			return info, fmt.Errorf("create issue failed: %v", err)
		}
		return info, nil
	case update:
		info, err := iisue.Update()
		if err != nil {
			return info, fmt.Errorf("update issue failed: %v", err)
		}
		return info, nil
	case close:
		info, err := iisue.Close()
		if err != nil {
			return info, fmt.Errorf("close issue failed: %v", err)
		}
		return info, nil
	case read:
		info, err := iisue.Read()
		if err != nil {
			return info, fmt.Errorf("read issue failed: %v", err)
		}
		return info, nil
	}

	return issueInfo{}, nil
}
