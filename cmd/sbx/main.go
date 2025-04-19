package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
	"github.com/syumai/sbx/sbpl"
	"github.com/urfave/cli/v3"
)

const (
	flagAllowFile    = "allow-file"
	flagDenyFile     = "deny-file"
	flagAllowFileAll = "allow-file-all"
	flagDenyFileAll  = "deny-file-all"

	flagAllowFileRead = "allow-file-read"
	flagDenyFileRead  = "deny-file-read"

	flagAllowFileWrite = "allow-file-write"
	flagDenyFileWrite  = "deny-file-write"

	flagAllowNetwork    = "allow-network"
	flagDenyNetwork     = "deny-network"
	flagAllowNetworkAll = "allow-network-all"
	flagDenyNetworkAll  = "deny-network-all"

	flagAllowNetworkInbound = "allow-network-inbound"
	flagDenyNetworkInbound  = "deny-network-inbound"

	flagAllowNetworkOutbound = "allow-network-outbound"
	flagDenyNetworkOutbound  = "deny-network-outbound"

	flagAllowProcessExec    = "allow-process-exec"
	flagDenyProcessExec     = "deny-process-exec"
	flagAllowProcessExecAll = "allow-process-exec-all"
	flagDenyProcessExecAll  = "deny-process-exec-all"

	flagAllowSysctlRead = "allow-sysctl-read"
	flagDenySysctlRead  = "deny-sysctl-read"
)

var operationTypeByFlagMap = map[string]sbpl.OperationType{
	flagAllowFile:    sbpl.OperationTypeFile,
	flagDenyFile:     sbpl.OperationTypeFile,
	flagAllowFileAll: sbpl.OperationTypeFileAll,
	flagDenyFileAll:  sbpl.OperationTypeFileAll,

	flagAllowFileRead: sbpl.OperationTypeFileRead,
	flagDenyFileRead:  sbpl.OperationTypeFileRead,

	flagAllowFileWrite: sbpl.OperationTypeFileWrite,
	flagDenyFileWrite:  sbpl.OperationTypeFileWrite,

	flagAllowNetwork:    sbpl.OperationTypeNetwork,
	flagDenyNetwork:     sbpl.OperationTypeNetwork,
	flagAllowNetworkAll: sbpl.OperationTypeNetworkAll,
	flagDenyNetworkAll:  sbpl.OperationTypeNetworkAll,

	flagAllowNetworkInbound: sbpl.OperationTypeNetworkInbound,
	flagDenyNetworkInbound:  sbpl.OperationTypeNetworkInbound,

	flagAllowNetworkOutbound: sbpl.OperationTypeNetworkOutbound,
	flagDenyNetworkOutbound:  sbpl.OperationTypeNetworkOutbound,

	flagAllowProcessExec:    sbpl.OperationTypeProcessExec,
	flagDenyProcessExec:     sbpl.OperationTypeProcessExec,
	flagAllowProcessExecAll: sbpl.OperationTypeProcessExecAll,
	flagDenyProcessExecAll:  sbpl.OperationTypeProcessExecAll,

	flagAllowSysctlRead: sbpl.OperationTypeSysctlRead,
	flagDenySysctlRead:  sbpl.OperationTypeSysctlRead,
}

func main() {
	flags := []cli.Flag{
		&cli.StringFlag{Name: flagAllowFile, Usage: "allow file* operation"},
		&cli.StringFlag{Name: flagDenyFile, Usage: "deny file* operation"},
		&cli.BoolFlag{Name: flagAllowFileAll, Usage: "allow all file* operation"},
		&cli.BoolFlag{Name: flagDenyFileAll, Usage: "deny all file* operation"},

		&cli.StringFlag{Name: flagAllowFileRead, Usage: "allow file-read operation"},
		&cli.StringFlag{Name: flagDenyFileRead, Usage: "deny file-read operation"},

		&cli.StringFlag{Name: flagAllowFileWrite, Usage: "allow file-write operation"},
		&cli.StringFlag{Name: flagDenyFileWrite, Usage: "deny file-write operation"},

		&cli.StringFlag{Name: flagAllowNetwork, Usage: "allow network* operation"},
		&cli.StringFlag{Name: flagDenyNetwork, Usage: "deny network* operation"},
		&cli.BoolFlag{Name: flagAllowNetworkAll, Usage: "allow all network* operation"},
		&cli.BoolFlag{Name: flagDenyNetworkAll, Usage: "deny all network* operation"},

		&cli.StringFlag{Name: flagAllowNetworkInbound, Usage: "allow network-inbound operation"},
		&cli.StringFlag{Name: flagDenyNetworkInbound, Usage: "deny network-inbound operation"},

		&cli.StringFlag{Name: flagAllowNetworkOutbound, Usage: "allow network-outbound operation"},
		&cli.StringFlag{Name: flagDenyNetworkOutbound, Usage: "deny network-outbound operation"},

		&cli.StringFlag{Name: flagAllowProcessExec, Usage: "allow process-exec operation"},
		&cli.StringFlag{Name: flagDenyProcessExec, Usage: "deny process-exec operation"},
		&cli.BoolFlag{Name: flagAllowProcessExecAll, Usage: "allow all process-exec operation"},
		&cli.BoolFlag{Name: flagDenyProcessExecAll, Usage: "deny all process-exec operation"},

		&cli.StringFlag{Name: flagAllowSysctlRead, Usage: "allow sysctl-read operation"},
		&cli.StringFlag{Name: flagDenySysctlRead, Usage: "deny sysctl-read operation"},
	}
	cmd := &cli.Command{
		Name:  "sbx",
		Flags: sliceutil.Map(flags, func(flag cli.Flag) cli.Flag { return flag }),
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var (
				operations       []*sbpl.Operation
				isNetworkAllowed = false
			)
			for _, flag := range flags {
				if !flag.IsSet() {
					continue
				}
				flagName := flag.Names()[0]
				operationType := operationTypeByFlagMap[flagName]
				allowed := strings.HasPrefix(flagName, "allow-")
				value := cmd.String(flagName)
				values := strings.Split(value, ",")
				addressFilters, err := func() ([]*sbpl.NetworkFilterAddress, error) {
					switch operationType {
					case sbpl.OperationTypeNetwork, sbpl.OperationTypeNetworkInbound, sbpl.OperationTypeNetworkOutbound:
						return sliceutil.MapWithError(values, func(v string) (*sbpl.NetworkFilterAddress, error) {
							host, port, found := strings.Cut(v, ":")
							if !found {
								return nil, fmt.Errorf("address must be in the format of host:port: %s", v)
							}
							return sbpl.NewNetworkFilterAddress(host, port)
						})
					default:
						return nil, nil
					}
				}()
				if err != nil {
					return err
				}
				filters, err := sliceutil.MapWithError(values, func(v string) (sbpl.Filter, error) {
					switch operationType {
					case sbpl.OperationTypeFileAll, sbpl.OperationTypeNetworkAll, sbpl.OperationTypeProcessExecAll:
						return nil, nil
					case sbpl.OperationTypeFile, sbpl.OperationTypeFileRead, sbpl.OperationTypeFileWrite:
						return sbpl.NewSubpathPathFilter(v)
					case sbpl.OperationTypeNetwork, sbpl.OperationTypeNetworkInbound, sbpl.OperationTypeNetworkOutbound:
						isNetworkAllowed = true
						return sbpl.NewNetworkFilter(
							false,                        // support only remote
							sbpl.NetworkFilterProtocolIP, // support only ip
							addressFilters,
						), nil
					case sbpl.OperationTypeProcessExec, sbpl.OperationTypeProcessExecNoSandbox:
						return sbpl.NewLiteralPathFilter(v), nil
					default:
						panic(fmt.Sprintf("unexpected operation type: %s", operationType))
					}
				})
				if err != nil {
					return err
				}
				filters = sliceutil.Filter(filters, func(f sbpl.Filter) bool {
					return f != nil
				})
				operations = append(operations, &sbpl.Operation{
					Type:    operationType,
					Allowed: allowed,
					Filters: filters,
				})
			}
			command := cmd.Args().First()
			commandPath, err := exec.LookPath(command)
			if err != nil {
				return fmt.Errorf("command not found: %s", command)
			}
			operations = append(operations, &sbpl.Operation{
				Type:    sbpl.OperationTypeProcessExec,
				Allowed: true,
				Filters: []sbpl.Filter{
					sbpl.NewLiteralPathFilter(commandPath),
				},
			})
			policy := sbpl.NewPolicy(isNetworkAllowed, operations).String()
			return sandboxExec(ctx, policy, commandPath, cmd.Args().Tail()...)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
