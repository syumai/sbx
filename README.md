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

or download binary from [releases](https://github.com/syumai/sbx/releases) page.

## Usage

```
sbx [flags] <command> [command-args...]
sbx [flags] -- <command> [command-flags] [command-args...]
```

### Flags

- By default, `sbx` denies all operations.
  - To allow all operations for investigation purposes, use the `--allow-all` flag.
- You can allow operations by specifying the corresponding flags.
- You can deny operations by specifying the corresponding flags with `deny-` prefix.
- `-all` flags are boolean flags that allow / deny all operations for the corresponding operation type.
  - Another flags requires a path arguments as comma-separated values.
- `-network` flags supports only settings below.
  - only `ip` protocol.
  - only `localhost` or `*` for host.

#### Special Flag

- `--allow-all`               Allow all operations (without this flag, deny all operations by default)

#### File Operations
- `--allow-file`              Allow file operations
- `--deny-file`               Deny file operations
- `--allow-file-all`          Allow all file operations
- `--deny-file-all`           Deny all file operations

##### File Read Operations
- `--allow-file-read`         Allow file read operations
- `--deny-file-read`          Deny file read operations
- `--allow-file-read-all`     Allow all file read operations
- `--deny-file-read-all`      Deny all file read operations

##### File Write Operations
- `--allow-file-write`        Allow file write operations
- `--deny-file-write`         Deny file write operations
- `--allow-file-write-all`    Allow all file write operations
- `--deny-file-write-all`     Deny all file write operations

#### Network Operations
- `--allow-network-all`       Allow all network operations
- `--deny-network-all`        Deny all network operations
- `--allow-network-inbound`   Allow inbound network operations
- `--deny-network-inbound`    Deny inbound network operations
- `--allow-network-outbound`  Allow outbound network operations
- `--deny-network-outbound`   Deny outbound network operations

#### Process Operations
- `--allow-process-exec`      Allow process execution
- `--deny-process-exec`       Deny process execution
- `--allow-process-exec-all`  Allow all process execution
- `--deny-process-exec-all`   Deny all process execution

#### Sysctl Operations
- `--allow-sysctl-read`       Allow sysctl read
- `--deny-sysctl-read`        Deny sysctl read
- `--allow-sysctl-read-all`   Allow all sysctl read
- `--deny-sysctl-read-all`    Deny all sysctl read

### Example

* Allow read operation for current directory.

```console
sbx --allow-file-read . ls .
# same as above
sbx --allow-file-read='.' ls .

# with command flags
sbx --allow-file-read='.' -- ls -l .
```

* Allow network operation for `localhost:8080`.

```console
sbx --allow-network='localhost:8080' curl http://localhost:8080
```

* Allow network operation for remote host.
  - Allow read access to the `/opt/local` directory to retrieve CA certificates. (This example uses homebrew-installed `curl`.)

```console
sbx --allow-network='*:443' --allow-file-read='/opt/local' curl https://syum.ai/ascii
```

## License

MIT
