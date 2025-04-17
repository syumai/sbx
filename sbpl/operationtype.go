package sbpl

import "fmt"

type OperationType int

const (
	OperationTypeUnknown OperationType = iota
	OperationTypeFileAll
	OperationTypeFileRead
	OperationTypeFileWrite
	OperationTypeNetworkAll
	OperationTypeNetworkInbound
	OperationTypeNetworkOutbound
	OperationTypeProcessExec
	OperationTypeProcessExecNoSandbox
	OperationTypeSysctlRead
)

func (t OperationType) String() string {
	switch t {
	case OperationTypeFileAll:
		return "file*"
	case OperationTypeFileRead:
		return "file-read*"
	case OperationTypeFileWrite:
		return "file-write*"
	case OperationTypeNetworkAll:
		return "network*"
	case OperationTypeNetworkInbound:
		return "network-inbound"
	case OperationTypeNetworkOutbound:
		return "network-outbound"
	case OperationTypeProcessExec, OperationTypeProcessExecNoSandbox:
		return "process-exec"
	case OperationTypeSysctlRead:
		return "sysctl-read"
	}
	panic(fmt.Sprintf("unexpected operation type: %d", t))
}
