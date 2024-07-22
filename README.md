# cURL for Go



## Installation

```sh
go install github.com/jcbhmr/go-curl/v8/cmd/curl@latest
```

```sh
go get github.com/jcbhmr/go-curl/v8
```

## Usage

```sh
curl --help
```

```go
// TODO: Add libcurl
```

## How it works

The AIO `curl` binary in [`internal/curl/`](./internal/curl) is provided by [ahgamut/superconfigure](https://github.com/ahgamut/superconfigure) in the [`web/curl/`](https://github.com/ahgamut/superconfigure/tree/main/web/curl) subfolder. This binary runs on Windows, macOS, Linux, and more on both x86-64 and AArch64 architectures. This large binary is `//go:embed`-ed into the [`cmd/curl/`](./cmd/curl) command and extracted to `$XDG_DATA_HOME` upon its first invocation. It then replaces itself (the `go install`-ed binary) with a symlink to the big `curl` binary in `$XDG_DATA_HOME`.

## Development

- [ ] Add raw Go bindings to `libcurl`
- [ ] Add GitHub Actions to run `go test`

ℹ There aren't versioned URLs on [cosmo.zip/pub/cosmos/bin](https://cosmo.zip/pub/cosmos/bin). 🤷‍♀️
