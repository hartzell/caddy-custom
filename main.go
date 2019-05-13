package main

import (
	"github.com/mholt/caddy/caddy/caddymain"

	// plug in plugins here, for example:
	// _ "import/path/here"
	_ "github.com/abiosoft/caddy-git"
	_ "github.com/hacdias/caddy-webdav"
	_ "github.com/jung-kurt/caddy-cgi"
	_ "github.com/techknowlogick/caddy-s3browser"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
