package main

import (
	"github.com/kelseyhightower/envconfig"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"os"
	"testing"
)

func testCase(t *testing.T, test string, expected string) {
	var conf Configuration
	err := envconfig.Process("sshnotify", &conf)
	if err != nil {
		t.Errorf(err.Error())
	}
	input, readfixturerr := ioutil.ReadFile("../test/fixtures/" + test)
	if readfixturerr != nil {
		t.Errorf("Failed to read fixture %s: %s", test, readfixturerr)
	}
	f, tempfileerr := ioutil.TempFile(os.TempDir(), test)
	if tempfileerr != nil {
		t.Errorf("Write temporary fixture file: %s", tempfileerr)
	}
	defer os.Remove(f.Name())
	_, err = f.Write(input)
	if err != nil {
		t.Errorf("Failed to write temporary fixture %s: %s", f.Name(), err)
	}
	conf.PamConfig = f.Name()
	callerr := installPamConfig(conf)
	if callerr != nil {
		t.Errorf("Failed to read result %s: %s", test, callerr)
	}
	expect, expreaderr := ioutil.ReadFile("../test/fixtures/" + expected)
	if expreaderr != nil {
		t.Errorf("Failed to read result %s: %s", test, expreaderr)
	}
	result, resreaderr := ioutil.ReadFile(f.Name())
	if resreaderr != nil {
		t.Errorf("Failed to read result %s: %s", test, resreaderr)
	}
	So(string(result), ShouldResemble, string(expect))
}

func errCase(t *testing.T, testfile string) {
	var conf Configuration
	err := envconfig.Process("sshnotify", &conf)
	if err != nil {
		t.Errorf(err.Error())
	}
	conf.PamConfig = testfile
	conf.Directory = "/var/lib/toolbox/ssh-notify/"
	So(installPamConfig(conf), ShouldNotBeNil)
}

func nilCase(t *testing.T, test string) {
	testCase(t, test, test)
}

func TestPamConfig(t *testing.T) {
	Convey("Leave be if the only thing in the file", t, func() { nilCase(t, "only") })
	Convey("Leave the file in place if it has the line already", t, func() { nilCase(t, "already") })
	Convey("Leave the file if it starts with a tab", t, func() { nilCase(t, "tabstart") })
	Convey("Install if not present in default setup", t, func() { testCase(t, "default", "already") })
	Convey("Error if file not found", t, func() { errCase(t, "notexist") })
}
