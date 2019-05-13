package main

import (
	"github.com/mholt/caddy/caddy/caddymain"
	
	// plug in plugins here, for example:
	_ "github.com/jung-kurt/caddy-cgi"
	_ "github.com/hacdias/caddy-webdav"
	// _ "import/path/here"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}

