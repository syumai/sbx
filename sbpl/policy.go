package sbpl

import "strings"

func NewPolicy(operations []*Operation) *Policy {
	return &Policy{
		Operations: operations,
	}
}

type Policy struct {
	Allowed    bool
	Operations []*Operation
}

func (p *Policy) String() string {
	body := []string{"(version 1)"}
	if p.Allowed {
		body = append(body, "(allow default)")
	} else {
		body = append(body, "(deny default)")
	}
	for _, operation := range p.Operations {
		body = append(body, operation.String())
	}
	return strings.Join(body, "\n")
}
