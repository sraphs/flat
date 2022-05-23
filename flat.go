// Package flat provides flatten and unflatten method for map
package flat

import (
	"strings"

	"github.com/sraphs/go/x/strcase"
)

var DefaultOption = &Option{
	Separator: ".",
}

type Option struct {
	Case      Case   // "", "lower" or "upper", defaults to "" to no case.
	Separator string // "-" | "_" | "." etc..., defaults to "."
}

func (o *Option) GetSeparator() string {
	if o.Separator == "" {
		return "."
	}

	return o.Separator
}

type Case int32

const (
	CaseNone Case = iota
	CaseLower
	CaseUpper
	CaseCamel
	CaseSnake
	CasePascal
)

func (c Case) to(s string) string {
	switch c {
	case CaseLower:
		return strings.ToLower(s)
	case CaseUpper:
		return strings.ToUpper(s)
	case CaseCamel:
		return strcase.ToCamel(s)
	case CaseSnake:
		return strcase.ToSnake(s)
	case CasePascal:
		return strcase.ToPascal(s)

	}

	return s
}
