package sbpl

import "strings"

func NewDocument(operations []*Operation) *Document {
	return &Document{
		Operations: operations,
	}
}

type Document struct {
	Allowed    bool
	Operations []*Operation
}

func (d *Document) String() string {
	body := []string{"(version 1)"}
	if d.Allowed {
		body = append(body, "(allow default)")
	} else {
		body = append(body, "(deny default)")
	}
	for _, operation := range d.Operations {
		body = append(body, operation.String())
	}
	return strings.Join(body, "\n")
}
