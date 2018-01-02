package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func writeConfig(conf Configuration) error {
	content, err := yaml.Marshal(&conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(conf.Directory + conf.Config, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
