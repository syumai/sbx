package sbpl

import (
	"fmt"
	"path/filepath"
	"strings"
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

func NewLiteralPathFilter(path string) Filter {
	return &PathFilter{
		Type: PathFilterTypeLiteral,
		Path: path,
	}
}

func NewSubpathPathFilter(path string) (Filter, error) {
	if strings.HasPrefix(path, "/") {
		return &PathFilter{
			Type: PathFilterTypeSubpath,
			Path: path,
		}, nil
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	return &PathFilter{
		Type: PathFilterTypeSubpath,
		Path: abs,
	}, nil
}
