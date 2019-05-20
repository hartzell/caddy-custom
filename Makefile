caddy: main.go custom.go
	go build

custom.go: gen_custom.go
	env CADDY_ENABLE_TELEMETRY="${CADDY_ENABLE_TELEMETRY}" \
		CADDY_PLUGINS="${CADDY_PLUGINS}" \
		go generate

clean:
	-rm caddy custom.go
