package main

import (
	"github.com/caddyserver/caddy/caddy/caddymain"
)

func main() {
	caddymain.EnableTelemetry = enableTelemetry()
	caddymain.Run()
}
