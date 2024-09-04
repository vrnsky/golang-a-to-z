package pocketlog_test

import (
	"golang-a-to-z/chapter3/pocketlog"
	"os"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello %s", "world")
	//Output: Hello world
}