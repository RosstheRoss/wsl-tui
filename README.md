# wsl-tui

A simple terminal interface to interact with having several [WSL](https://learn.microsoft.com/en-gb/windows/wsl/) distributions installed.

## Config file

See [example config](./config_example.toml).

## Requirements

- `go` 1.23
- `wsl` if on Windows
- `SSH` if on any other platform or on Windows built with the `nohost` tag.

### Building

```bash
go build -ldflags "-s -w" -o wsl-tui
```

To use SSH on Windows instead of WSL, instead use the following:

```bash
go build -ldflags "-s -w" -tags nohost -o wsl-tui
```

## License

[MIT](./LICENSE)
