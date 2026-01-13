package main

import "github.com/arfadmuzali/restui/cmd"

var (
	version = "dev"
	// commit  = "none"
	// date    = "unknown"
)

func main() {
	cmd.Execute(version)
}
