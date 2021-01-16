package cgi

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
)

// handlerType is a middleware type that can handle CGI requests
type handlerType struct {
	next  httpserver.Handler
	rules []ruleType
	root  string // same as root, but absolute path
}

// ruleType represents a CGI handling rule; it is parsed from the cgi directive
// in the Caddyfile
type ruleType struct {
	// Glob patterns to match in order to apply rule
	matches []string // glob patterns, [1..n]
	// Match exceptions
	exceptions []string
	// Name of executable script or binary
	exe string // [1]
	// Working directory (default, current Caddy working directory)
	dir string // [0..1]
	// Arguments to submit to executable
	args []string // [0..n]
	// Environment key value pairs ([0]: key, [1]: value) for this particular app
	envs [][2]string // [0..n]
	// Environment keys to pass through for all apps
	passEnvs []string // [0..n]
	// Environment keys to send with empty values
	emptyEnvs []string // [0..n]
	// True to return inspection page rather than call CGI executable
	inspect bool
	// True to pass all environment variables to CGI executable
	passAll bool
}
