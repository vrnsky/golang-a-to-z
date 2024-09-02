package pocketlog_test

import "golang-a-to-z/chapter3/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello %s", "world")
	//Output: Hello, world
}