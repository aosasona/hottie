# Hottie

A web server that serves static websites and provides hot reloading functionality through Server-Sent Events (SSE).

_(Pardon the name, I just had to.)_

# Installation

## Homebrew

```bash
brew tap aosasona/hottie && brew install hottie
```

## Go

```bash
go install github.com/aosasona/hottie
```

## Binary/executable

You can also download any of the binaries from the [releases page](https://github.com/aosasona/hottie/releases).

## Building from source

```bash
git clone https://github.com/aosasona/hottie.git
cd hottie
go build
```

> You need the Go compiler/toolchain to build this package from source code.

# Usage

```bash
path/to/hottie -dir target_directory -port 9000
```

Available flags:

- `addr`: address to serve the files on, default is `127.0.0.1`
- `dir`: the directory ton serve, default is `.`
- `port`: port to bind the server to, default is `3000`
- `reload`: enable or disable hot-reloading, `true` (enabled) by default
- `open`: open the URL in the default browser after starting server, `false` (disabled) by default
