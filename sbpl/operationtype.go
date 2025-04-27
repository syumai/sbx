package sbpl

import "fmt"

type OperationType int

const (
	OperationTypeUnknown OperationType = iota
	OperationTypeFile
	OperationTypeFileRead
	OperationTypeFileWrite
	OperationTypeNetwork
	OperationTypeNetworkInbound
	OperationTypeNetworkOutbound
	OperationTypeProcessExec
	OperationTypeProcessExecNoSandbox
)

func (t OperationType) String() string {
	switch t {
	case OperationTypeFile:
		return "file*"
	case OperationTypeFileRead:
		return "file-read*"
	case OperationTypeFileWrite:
		return "file-write*"
	case OperationTypeNetwork:
		return "network*"
	case OperationTypeNetworkInbound:
		return "network-inbound"
	case OperationTypeNetworkOutbound:
		return "network-outbound"
	case OperationTypeProcessExec, OperationTypeProcessExecNoSandbox:
		return "process-exec"
	}
	panic(fmt.Sprintf("unexpected operation type: %d", t))
}
