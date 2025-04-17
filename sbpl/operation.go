package sbpl

import (
	"fmt"
	"strings"

	"github.com/syumai/sbx/internal/sliceutil"
)

type Operation struct {
	Type    OperationType
	Allowed bool
	Filters []Filter
}

func (o *Operation) String() string {
	filters := strings.Join(sliceutil.MapStringer(o.Filters), " ")
	allowed := "allow"
	if !o.Allowed {
		allowed = "deny"
	}
	body := []string{allowed, o.Type.String()}
	if filters != "" {
		body = append(body, filters)
	}
	if o.Type == OperationTypeProcessExecNoSandbox {
		body = append(body, "(with no-sandbox)")
	}
	return fmt.Sprintf("(%s)", strings.Join(body, " "))
}
