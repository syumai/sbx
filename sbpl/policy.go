package sbpl

import (
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
)

func NewPolicy(allowAllOperations bool, isNetworkAllowed bool, operations []*Operation) *Policy {
	return &Policy{
		AllowAllOperations: allowAllOperations,
		IsNetworkAllowed:   isNetworkAllowed,
		Operations:         operations,
	}
}

type Policy struct {
	AllowAllOperations bool
	IsNetworkAllowed   bool
	Operations         []*Operation
}

func (p *Policy) String() string {
	body := []string{
		"(version 1)",
		`(import "bsd.sb")`,
	}
	if p.AllowAllOperations {
		body = append(body, "(allow default)")
	} else {
		body = append(body, "(deny default)")
	}
	// allow access to dylibs because it's required for process exec
	body = append(body,
		`(allow file-read*
	(subpath "/opt/local/lib")
	(subpath "/usr/lib")
	(subpath "/usr/local/lib")
	)`)
	if p.IsNetworkAllowed {
		body = append(body,
			// allow access to unix-socket when network is allowed
			`(allow network* (local unix-socket))`,
		)
	}
	body = append(body, sliceutil.MapStringer(p.Operations)...)
	return strings.Join(body, "\n")
}
