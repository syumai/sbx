package sbpl

import (
	"fmt"
)

type PathFilterType int

const (
	PathFilterTypeUnknown PathFilterType = iota
	PathFilterTypeLiteral
	PathFilterTypeSubpath
)

type PathFilter struct {
	Type PathFilterType
	Path string
}

func (f *PathFilter) String() string {
	switch f.Type {
	case PathFilterTypeLiteral:
		return fmt.Sprintf("(literal %q)", f.Path)
	case PathFilterTypeSubpath:
		return fmt.Sprintf("(subpath %q)", f.Path)
	}
	panic(fmt.Sprintf("unexpected path filter type: %d", f.Type))
}

func NewPathFilter(path string) Filter {
	return &PathFilter{
		Type: PathFilterTypeSubpath,
		Path: path,
	}
}
