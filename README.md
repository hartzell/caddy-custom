# My Custom Caddy

I need a few plugins for my Caddy.  It's easy enough to build it
myself, following the [Caddy build][caddy-build] instructions.

This is the result.

Build it like so:

```
go build
```

## TODO

> Shall I whip up an go gen thingy that adds that line? I think you
> should then be able to build with `go build -tags=coredns` and let
> it add that line. (I think)
>
> From a [CoreDNS thread][coredns-adding-plugins].

[caddy-build]: https://github.com/mholt/caddy#build
[coredns-adding-plugins]: https://caddy.community/t/building-coredns-automatically/481/10
