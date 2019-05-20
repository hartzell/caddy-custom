package main

//go:generate go run gen.go


import (
	"github.com/mholt/caddy/caddy/caddymain"
)

func main() {
	// optional: disable telemetry
	caddymain.EnableTelemetry = enableTelemetry()
	caddymain.Run()
}
