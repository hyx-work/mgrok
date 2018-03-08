package httpProxy

import (
	"flag"
)

type Options struct {
	config string
}

func parseArgs() (opts *Options) {
	config := flag.String(
		"config",
		"httpProxy.yaml",
		"Path to mgrok configuration file. (default: mgrok.yaml)",
	)

	opts = &Options{
		config: *config,
	}

	return opts
}
