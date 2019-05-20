# My Custom Caddy

I need a few plugins for my Caddy.  It's easy enough to build it
myself, following the [Caddy build][caddy-build] instructions.

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

Create `/usr/ports/www/caddy-custom`, copy the contents of the port
directory into it, `cd` into it, and `make install`.

*Et voila!*

## TODO

> Shall I whip up an go gen thingy that adds that line? I think you
> should then be able to build with `go build -tags=coredns` and let
> it add that line. (I think)
>
> From a [CoreDNS thread][coredns-adding-plugins].

[caddy-build]: https://github.com/mholt/caddy#build
[coredns-adding-plugins]: https://caddy.community/t/building-coredns-automatically/481/10
