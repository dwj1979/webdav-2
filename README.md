# Simple WebDAV Server

This is a simple WebDAV server in Go.

It only serves on localhost and does'nt handle TLS or authentication.
It is meant to be deployed with a reverse proxy
like [Nginx][nginx] or [Caddy][caddy] to handle those aspects.

[nginx]: https://nginx.org/
[caddy]: https://caddyserver.com/

## Installing the App

1. First you have to [install Go.][go]

[go]: https://golang.org/doc/install

3. Run `go get github.com/edbedbe/webdav` to download
the latest source.

2. Run `go install github.com/edbedbe/webdav`
to install the `webdav` command.

## How to Use It

There are two optional arguments.

1. `-port [8080]`: the port to serve on at localhost
2. `-dir [.]`: the directory to serve over WebDAV

For example:

```
$ webdav -port 9000 -dir /tmp/webdav
```

