package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
	"github.com/syumai/sbx/sbpl"
	"github.com/urfave/cli/v3"
)

const (
	flagAllowFile            = "allow-file"
	flagDenyFile             = "deny-file"
	flagAllowFileRead        = "allow-file-read"
	flagDenyFileRead         = "deny-file-read"
	flagAllowFileWrite       = "allow-file-write"
	flagDenyFileWrite        = "deny-file-write"
	flagAllowNetwork         = "allow-network"
	flagDenyNetwork          = "deny-network"
	flagAllowNetworkInbound  = "allow-network-inbound"
	flagDenyNetworkInbound   = "deny-network-inbound"
	flagAllowNetworkOutbound = "allow-network-outbound"
	flagDenyNetworkOutbound  = "deny-network-outbound"
	flagAllowProcessExec     = "allow-process-exec"
	flagDenyProcessExec      = "deny-process-exec"
	flagAllowSysctlRead      = "allow-sysctl-read"
	flagDenySysctlRead       = "deny-sysctl-read"
)

var operationTypeByFlagMap = map[string]sbpl.OperationType{
	flagAllowFile:            sbpl.OperationTypeFileAll,
	flagDenyFile:             sbpl.OperationTypeFileAll,
	flagAllowFileRead:        sbpl.OperationTypeFileRead,
	flagDenyFileRead:         sbpl.OperationTypeFileRead,
	flagAllowFileWrite:       sbpl.OperationTypeFileWrite,
	flagDenyFileWrite:        sbpl.OperationTypeFileWrite,
	flagAllowNetwork:         sbpl.OperationTypeNetworkAll,
	flagDenyNetwork:          sbpl.OperationTypeNetworkAll,
	flagAllowNetworkInbound:  sbpl.OperationTypeNetworkInbound,
	flagDenyNetworkInbound:   sbpl.OperationTypeNetworkInbound,
	flagAllowNetworkOutbound: sbpl.OperationTypeNetworkOutbound,
	flagDenyNetworkOutbound:  sbpl.OperationTypeNetworkOutbound,
	flagAllowProcessExec:     sbpl.OperationTypeProcessExec,
	flagDenyProcessExec:      sbpl.OperationTypeProcessExec,
	flagAllowSysctlRead:      sbpl.OperationTypeSysctlRead,
	flagDenySysctlRead:       sbpl.OperationTypeSysctlRead,
}

func main() {
	flags := []*cli.StringFlag{
		{Name: flagAllowFile, Usage: "allow file* operation"},
		{Name: flagDenyFile, Usage: "deny file* operation"},
		{Name: flagAllowFileRead, Usage: "allow file-read operation"},
		{Name: flagDenyFileRead, Usage: "deny file-read operation"},
		{Name: flagAllowFileWrite, Usage: "allow file-write operation"},
		{Name: flagDenyFileWrite, Usage: "deny file-write operation"},
		{Name: flagAllowNetwork, Usage: "allow network* operation"},
		{Name: flagDenyNetwork, Usage: "deny network* operation"},
		{Name: flagAllowNetworkInbound, Usage: "allow network-inbound operation"},
		{Name: flagDenyNetworkInbound, Usage: "deny network-inbound operation"},
		{Name: flagAllowNetworkOutbound, Usage: "allow network-outbound operation"},
		{Name: flagDenyNetworkOutbound, Usage: "deny network-outbound operation"},
		{Name: flagAllowProcessExec, Usage: "allow process-exec operation"},
		{Name: flagDenyProcessExec, Usage: "deny process-exec operation"},
		{Name: flagAllowSysctlRead, Usage: "allow sysctl-read operation"},
		{Name: flagDenySysctlRead, Usage: "deny sysctl-read operation"},
	}
	cmd := &cli.Command{
		Name:  "sbx",
		Flags: sliceutil.Map(flags, func(flag *cli.StringFlag) cli.Flag { return flag }),
		Action: func(ctx context.Context, cmd *cli.Command) error {
			var operations []*sbpl.Operation
			for _, flag := range flags {
				if !flag.IsSet() {
					continue
				}
				operationType := operationTypeByFlagMap[flag.Name]
				allowed := strings.HasPrefix(flag.Name, "allow-")
				value := cmd.String(flag.Name)
				values := strings.Split(value, ",")
				addressFilters, err := func() ([]*sbpl.NetworkFilterAddress, error) {
					switch operationType {
					case sbpl.OperationTypeNetworkAll, sbpl.OperationTypeNetworkInbound, sbpl.OperationTypeNetworkOutbound:
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
				filters := sliceutil.Map(values, func(v string) sbpl.Filter {
					switch operationType {
					case sbpl.OperationTypeFileAll, sbpl.OperationTypeFileRead, sbpl.OperationTypeFileWrite:
						return sbpl.NewPathFilter(v)
					case sbpl.OperationTypeNetworkAll, sbpl.OperationTypeNetworkInbound, sbpl.OperationTypeNetworkOutbound:
						return sbpl.NewNetworkFilter(
							false,                        // support only remote
							sbpl.NetworkFilterProtocolIP, // support only ip
							addressFilters,
						)
					case sbpl.OperationTypeProcessExec, sbpl.OperationTypeProcessExecNoSandbox:
						return sbpl.NewPathFilter(v)
					default:
						panic(fmt.Sprintf("unexpected operation type: %s", operationType))
					}
				})
				operations = append(operations, &sbpl.Operation{
					Type:    operationType,
					Allowed: allowed,
					Filters: filters,
				})
			}
			policy := sbpl.NewPolicy(operations).String()
			fmt.Println(policy)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
