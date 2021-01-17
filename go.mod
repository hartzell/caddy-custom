module github.com/hartzell/caddy-custom

go 1.15

// https://github.com/lucas-clemente/quic-go/issues/2614
replace github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.18.0

require (
	github.com/abiosoft/caddy-git v0.0.0-20190703061829-f8cc2f20c9e7
	github.com/caddyserver/caddy v1.0.5
	github.com/echocat/caddy-filter v0.14.0
	github.com/hacdias/caddy-webdav v1.1.0
	github.com/jung-kurt/caddy-cgi v1.11.4
	github.com/techknowlogick/caddy-s3browser v1.0.0
)
