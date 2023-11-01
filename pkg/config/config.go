package config

import (
	"flag"
	"time"
)

// Parse parse cmdline args to map
func Parse() map[string]any {
	verbosePtr := flag.Bool("v", false, "verbose output")
	runOncePtr := flag.Bool("r", false, "run once")
	forceLoginPtr := flag.Bool("f", false, "force login")
	checkDurationPtr := flag.Duration("d", time.Minute*5, "connectivity check duration")
	retryCountPtr := flag.Int("c", 3, "retry count")

	flag.Parse()

	ret := map[string]any{}
	ret["verbose"] = *verbosePtr
	ret["runOnce"] = *runOncePtr
	ret["forceLogin"] = *forceLoginPtr
	ret["checkDuration"] = *checkDurationPtr
	ret["retryCount"] = *retryCountPtr

	return ret
}
