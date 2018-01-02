package main

import (
	"time"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Configuration struct {
	// How often to re-apply, zero means exit after first install
	Loop      int `default:0`
	// File locations
	PamConfig string `envconfig:"pam_config" default:"/etc/pam.d/systemd-user"`
	Source    string
	// Directory prefixes all of the files following
	Directory string `envconfig:"dir" default:"/opt/ssh-notify/"`
	Binary    string `default:"notify"`
	Config    string `default:"config.yaml"`
	// PAM configuration
	AuthControl   string `envconfig:"auth_control" default:"optional"`
	AuthInterface string `envconfig:"auth_interface" default:"session"`
	AuthExec      string `envconfig:"auth_exec" default:"pam_exec.so"`
	// Slack configuration
	SlackWebHook string `envconfig:"slack_webhook"`
	SlackChannel string `envconfig:"slack_channel"`
}

func main() {
	// parse environment variables into Configuration struct
	var conf Configuration
	err := envconfig.Process("sshnotify", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	// check directory validity,
	// otherwise assume Binary and Config values
	// are a fully qualified path
	if len(conf.Directory) < 1 ||
		string([]rune(conf.Directory)[0]) != "/" ||
		string([]rune(conf.Directory)[len(conf.Directory)-1:]) != "/" {
		log.Printf("Removing common path prefix")
		conf.Directory = ""
	}

    for {
		// Writing the config file
		err = writeConfig(conf)
		if err != nil {
			log.Fatalf("failed to write config: " + err.Error())
		}

		// Copy the notification binary
		err = copyFile(conf.Source, conf.Directory+conf.Binary)
		if err != nil {
			log.Fatalf("failed to install binary file: " + err.Error())
		}

		// Configure the system pluggable authentication module
		err = installPamConfig(conf)
		if err != nil {
			log.Fatalf("failed to configure PAM: " + err.Error())
		}
		if conf.Loop == 0 {
			break
		}
		time.Sleep(time.Duration(conf.Loop) * time.Second)
	}
}
