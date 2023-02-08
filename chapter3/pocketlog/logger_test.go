package pocketlog_test

import "vrnsky.io/m/v2/pocketlog"

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello %s", "world")
	// Output: Hello, world
}
