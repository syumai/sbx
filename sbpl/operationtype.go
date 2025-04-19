package sbpl

import "fmt"

type OperationType int

const (
	OperationTypeUnknown OperationType = iota
	OperationTypeFile
	OperationTypeFileAll
	OperationTypeFileRead
	OperationTypeFileWrite
	OperationTypeNetwork
	OperationTypeNetworkAll
	OperationTypeNetworkInbound
	OperationTypeNetworkOutbound
	OperationTypeProcessExec
	OperationTypeProcessExecAll
	OperationTypeProcessExecNoSandbox
	OperationTypeSysctlRead
)

func (t OperationType) String() string {
	switch t {
	case OperationTypeFile, OperationTypeFileAll:
		return "file*"
	case OperationTypeFileRead:
		return "file-read*"
	case OperationTypeFileWrite:
		return "file-write*"
	case OperationTypeNetwork, OperationTypeNetworkAll:
		return "network*"
	case OperationTypeNetworkInbound:
		return "network-inbound"
	case OperationTypeNetworkOutbound:
		return "network-outbound"
	case OperationTypeProcessExec, OperationTypeProcessExecNoSandbox, OperationTypeProcessExecAll:
		return "process-exec"
	case OperationTypeSysctlRead:
		return "sysctl-read"
	}
	panic(fmt.Sprintf("unexpected operation type: %d", t))
}
