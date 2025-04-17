package sbpl

import "github.com/syumai/sbx/internal/sliceutil"

func NewNetworkAllOperation(allowed bool, filters []*NetworkFilter) *Operation {
	return &Operation{
		Type:    OperationTypeNetworkAll,
		Allowed: allowed,
		Filters: sliceutil.Map(filters, func(f *NetworkFilter) Filter { return f }),
	}
}

func NewNetworkInboundOperation(allowed bool, filters []*NetworkFilter) *Operation {
	return &Operation{
		Type:    OperationTypeNetworkInbound,
		Allowed: allowed,
		Filters: sliceutil.Map(filters, func(f *NetworkFilter) Filter { return f }),
	}
}

func NewNetworkOutboundOperation(allowed bool, filters []*NetworkFilter) *Operation {
	return &Operation{
		Type:    OperationTypeNetworkOutbound,
		Allowed: allowed,
		Filters: sliceutil.Map(filters, func(f *NetworkFilter) Filter { return f }),
	}
}
