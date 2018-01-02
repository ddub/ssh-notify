package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func installPamConfig(conf Configuration) error {
	// read the existing file into lines array
	input, err := ioutil.ReadFile(conf.PamConfig)
	if err != nil {
		return err
	}
	lines := strings.Split(strings.Trim(string(input), "\n"), "\n")

	// setup the pam configuration, we make an array to join together for
	// pattern recognition or with spaces for appending
	session := []string{
		conf.AuthInterface,
		conf.AuthControl,
		conf.AuthExec,
		conf.Directory + conf.Binary,
		conf.Directory + conf.Config}

	// check to see if the configuration is already present
	pattern := strings.Join(session, "\\s+")
	found := false
	for _, line := range lines {
		match, _ := regexp.MatchString(pattern, line)
		if match {
			found = true
			continue
		}
	}

	// if not present then add the configuration with the new line appended
	if found == false {
		f, err := os.OpenFile(conf.PamConfig, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err = f.WriteString(fmt.Sprintf("%s", strings.Join(session, " "))); err != nil {
			return err
		}
	}
	return nil
}
