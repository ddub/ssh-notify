package main

import (
	"fmt"
	"github.com/ashwanthkumar/slack-go-webhook"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Configuration struct {
	SlackWebhook string
	SlackChannel string
	SlackIcon    string
}

func main() {
	config := Configuration{
		SlackChannel: "#general",
		SlackIcon:    "https://i.imgur.com/lhy3U3N.jpg",
	}
	yamlFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Printf("yamlFile.Get err #%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	hostname, hosterr := os.Hostname()
	if hosterr != nil {
		log.Printf("Error finding hostname: %v", hosterr)
	}

    event := os.Getenv("PAM_TYPE")
	message := "login"
	if event == "close_session" {
		message = "logout"
	}

    username := os.Getenv("PAM_USER")
	if username == "" {
		username = "Unknown user"
	}
    remote := os.Getenv("PAM_RHOST")
	if remote == "" {
		remote = "an unknown location"
	}

	payload := slack.Payload{
		Text:    fmt.Sprintf("SSH %s: %s from %s on %s", message, username, remote, hostname),
		Channel: config.SlackChannel,
		IconUrl: config.SlackIcon,
	}
	slackerr := slack.Send(config.SlackWebhook, "", payload)
	if len(slackerr) > 0 {
		log.Fatalf("error: %s\n", slackerr)
	}
}
