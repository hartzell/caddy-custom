package main

import (
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = enableTelemetry()
	caddymain.Run()
}
