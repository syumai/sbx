# sbx

- `sbx` is an easy-to-use command-line tool for running commands with macOS sandbox-exec policies using flag-based interface.
- This command is heavily inspired by [littledivy](https://github.com/littledivy)'s [sh-deno](https://github.com/littledivy/sh-deno).

## Features

- Easy allow/deny configuration for common operations (file, network, process, sysctl, etc.)
- Supports both of relative and absolute path filtering.

## Important Notes

- **This command is using deprecated feature (sandbox-exec).**
- **This command is experimental and unstable.**

## Notes

- This command allows access to dylibs because it's required for process exec.
- When you specify `-network` flag, this command allows access to unix-socket.

## Installation

```
go install github.com/syumai/sbx/cmd/sbx@latest
```

## Usage

```
sbx [flags] -- <command> [args...]
```

### Flags

- By default, `sbx` denies all operations.
- You can allow operations by specifying the corresponding flags.
- You can deny operations by specifying the corresponding flags with `deny-` prefix.
- `-all` flags are boolean flags that allow / deny all operations for the corresponding operation type.
  - Another flags requires a path arguments as comma-separated values.
- `-network` flags supports only settings below.
  - only `ip` protocol.
  - only `localhost` or `*` for host.

#### Allow Operations
- `--allow-file-all`          Allow all file operations
- `--allow-file-read`         Allow file read operations
- `--allow-file-write`        Allow file write operations
- `--allow-network-all`       Allow all network operations
- `--allow-network-inbound`   Allow inbound network operations
- `--allow-network-outbound`  Allow outbound network operations
- `--allow-process-exec`      Allow process execution
- `--allow-sysctl-read`       Allow sysctl read

#### Deny Operations
- `--deny-file-all`           Deny all file operations
- `--deny-file-read`          Deny file read operations
- `--deny-file-write`         Deny file write operations
- `--deny-network-all`        Deny all network operations
- `--deny-network-inbound`    Deny inbound network operations
- `--deny-network-outbound`   Deny outbound network operations
- `--deny-process-exec`       Deny process execution
- `--deny-sysctl-read`        Deny sysctl read

### Example

* Allow read operation for current directory.

```
$ sbx --allow-file-read="." ls .
```

* Allow network operation for `localhost:8080`.

```
$ sbx --allow-network="localhost:8080" curl http://localhost:8080
```

* Allow network operation for remote host.
  - Allow read access to the `/opt/local` directory to retrieve CA certificates. (This example uses homebrew-installed `curl`.)

```
$ sbx --allow-network="*:443" --allow-file-read="/opt/local" curl https://syum.ai/ascii
```

## License

MIT
