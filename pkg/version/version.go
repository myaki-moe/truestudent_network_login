package version

var (
	commit    = "unknown"
	version   = "unknown"
	buildDate = "unknown"
)

func GetCommit() string {
	return commit
}

func GetVersion() string {
	return version
}

func GetBuildDate() string {
	return buildDate
}
