package main

import "codegen/cmd"

var (
	// Overridden by Go Releaser at build time
	version    = "dev"
	commit     = "HEAD"
	date       = "unknown"
)

func main() {
	cmd.Execute(version, commit, date)
}
