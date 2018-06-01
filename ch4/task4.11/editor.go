package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const tmppath = "/imanager"

// editor create temp file, invoke specified editor to edit it.
// Return pointer to bytes.Buffer where temp file content was written.
func editor(name string) (*bytes.Buffer, error) {
	tmp, err := createTmp()
	if err != nil {
		return nil, err
	}

	if err := invokeEditor(name, tmp); err != nil {
		return nil, err
	}

	buf, err := ioutil.ReadFile(tmp)
	if err != nil {
		return nil, fmt.Errorf("error when read from temp file: %v", err)
	}
	return bytes.NewBuffer(buf), nil
}

// createTmp creates temp file in os.Tmp() dir with 'tmppath' name.
// If exists tan truncate its content.
func createTmp() (string, error) {
	fpath := os.TempDir() + tmppath

	f, err := os.Create(fpath)
	if err != nil {
		return "", fmt.Errorf("temp file creation failed: %v", err)
	}
	defer f.Close()

	return fpath, nil
}

// invokeEditor exec command to start editor.
func invokeEditor(name, tmpfile string) error {
	cmd := exec.Command(name, tmpfile)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("start editor failed: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("error while editing: %v", err)
	}
	return nil
}
