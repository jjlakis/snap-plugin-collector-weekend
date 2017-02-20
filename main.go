package main

import (
	"os"

	// Import the snap plugin library
	"github.com/intelsdi-x/snap/control/plugin"
	// Import our collector plugin implementation
	"github.com/jjlakis/snap-plugin-collector-weekend/weekend"
)

func main() {
	// Provided:
	//   the definition of the plugin metadata
	//   the implementation satisfying plugin.CollectorPlugin

	// Define metadata about Plugin
	meta := weekend.Meta()

	// Start a collector
	plugin.Start(meta, new(weekend.Weekend), os.Args[1])
}
