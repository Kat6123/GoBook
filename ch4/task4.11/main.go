package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
)

var (
	// User vars.
	name, pwd string

	// Mode vars.
	readMode, createMode, updateMode, closeMode bool

	// Issue vars.
	number      int
	owner, repo string
	body        *bytes.Buffer
	iisue       issue

	// Editor name flag.
	editorN string
)

func main() {
	setFlags()
	if err := parseArgs(); err != nil {
		log.Fatalf("error when parse flags: %v", err)
	}

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
	flag.BoolVar(&readMode, "read", true, "Read issue")
	flag.BoolVar(&createMode, "create", false, "Create issue")
	flag.BoolVar(&updateMode, "update", false, "Update issue")
	flag.BoolVar(&closeMode, "close", false, "Close issue")

	// Set issue flags.
	flag.IntVar(&number, "n", 0, "Issue number")
	flag.StringVar(&owner, "o", "", "Repo owner")
	flag.StringVar(&repo, "r", "", "Repo")

	// Set editor name flag.
	flag.StringVar(&editorN, "editor", "vim", "Preferred editor")
}

func parseArgs() error {
	flag.Parse()

	if flag.NArg() == 0 && (createMode || updateMode) {
		buf, err := editor(editorN)
		if err != nil {
			return fmt.Errorf("error while using editor: %v", err)
		}
		body = buf
	} else {
		body = bytes.NewBufferString(flag.Arg(0))
	}

	return nil
}

func manage() (*issueInfo, error) {
	switch {
	case createMode:
		info, err := iisue.Create()
		if err != nil {
			return nil, fmt.Errorf("create issue failed: %v", err)
		}
		return info, nil
	case updateMode:
		info, err := iisue.Update()
		if err != nil {
			return nil, fmt.Errorf("update issue failed: %v", err)
		}
		return info, nil
	case closeMode:
		info, err := iisue.Close()
		if err != nil {
			return nil, fmt.Errorf("close issue failed: %v", err)
		}
		return info, nil
	case readMode:
		info, err := iisue.Read()
		if err != nil {
			return nil, fmt.Errorf("read issue failed: %v", err)
		}
		return info, nil
	}

	return nil, nil
}
