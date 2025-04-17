package sbpl

import (
	"github.com/syumai/sbx/internal/sliceutil"
)

func NewFileAllOperation(allowed bool, paths []string) *Operation {
	return &Operation{
		Type:    OperationTypeFileAll,
		Allowed: true,
		Filters: sliceutil.Map(paths, NewSubpathPathFilter),
	}
}

func NewFileReadOperation(allowed bool, paths []string) *Operation {
	return &Operation{
		Type:    OperationTypeFileRead,
		Allowed: allowed,
		Filters: sliceutil.Map(paths, NewSubpathPathFilter),
	}
}

func NewFileWriteOperation(allowed bool, paths []string) *Operation {
	return &Operation{
		Type:    OperationTypeFileWrite,
		Allowed: allowed,
		Filters: sliceutil.Map(paths, NewSubpathPathFilter),
	}
}
