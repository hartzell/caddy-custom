# My Custom Caddy

I need a few plugins for my Caddy.  It's easy enough to build it
myself, following the [Caddy build][caddy-build] instructions.

But then I wanted to make it more flexible and repeatable.

This is the result.

Build it like so:

```
go build -tags "cgi webdav"
```

## Build tags

### Plugins

- cgi -- enables https://github.com/jung-kurt/caddy-cgi
- docker-proxy -- enables https://github.com/lucaslorentz/caddy-docker-proxy
- filter -- enables https://github.com/echocat/caddy-filter
- git -- enables https://github.com/abiosoft/caddy-git
- s3browser -- enables https://github.com/techknowlogick/caddy-s3browser
- webdav -- enables https://github.com/hacdias/caddy-webdav

### Telemetry

- telemetry -- enables Caddy's telemetry feature

## FreeBSD port

See my [freebsd-ports][hartzell-freebsd-ports] tree, useful with [portshaker].

*Et voila!*

[caddy-build]: https://github.com/mholt/caddy#build
[coredns-adding-plugins]: https://caddy.community/t/building-coredns-automatically/481/10
[hartzell-freebsd-ports]: https://github.com/hartzell/freebsd-ports
[portshaker]: https://github.com/smortex/portshaker
