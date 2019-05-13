# My Custom Caddy

I need a few plugins for my Caddy, it's easy enough to build it
myself, following the [caddy-build] instructions.

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
> From: [coredns-adding-plugins]

[caddy-build]: https://github.com/mholt/caddy#build
[coredns-adding-pluggins]: https://caddy.community/t/building-coredns-automatically/481/10