package golang

import (
	_ "embed"
	"sort"
	"strconv"
	"text/template"

	"github.com/open-feature/cli/internal/generate"
	"github.com/open-feature/cli/internal/generate/types"

	"github.com/iancoleman/strcase"
)

// TmplData contains the Golang-specific data and the base data for the codegen.
type TmplData struct {
	*types.BaseTmplData
	GoPackage string
}

type genImpl struct {
	goPackage string
}

// BaseTmplDataInfo provides the base template data for the codegen.
func (td *TmplData) BaseTmplDataInfo() *types.BaseTmplData {
	return td.BaseTmplData
}

// supportedFlagTypes is the flag types supported by the Go template.
var supportedFlagTypes = map[types.FlagType]bool{
	types.FloatType:  true,
	types.StringType: true,
	types.IntType:    true,
	types.BoolType:   true,
	types.ObjectType: false,
}

func (*genImpl) SupportedFlagTypes() map[types.FlagType]bool {
	return supportedFlagTypes
}

//go:embed golang.tmpl
var golangTmpl string

// Go Funcs BEGIN

func flagVarName(flagName string) string {
	return strcase.ToCamel(flagName)
}

func flagInitParam(flagName string) string {
	return strconv.Quote(flagName)
}

func openFeatureType(t types.FlagType) string {
	switch t {
	case types.IntType:
		return "Int"
	case types.FloatType:
		return "Float"
	case types.BoolType:
		return "Boolean"
	case types.StringType:
		return "String"
	default:
		return ""
	}
}

func supportImports(flags []*types.FlagTmplData) []string {
	var res []string
	if len(flags) > 0 {
		res = append(res, "\"context\"")
		res = append(res, "\"github.com/open-feature/go-sdk/openfeature\"")
	}
	sort.Strings(res)
	return res
}

func defaultValueLiteral(flag *types.FlagTmplData) string {
	switch flag.Type {
	case types.StringType:
		return strconv.Quote(flag.DefaultValue)
	default:
		return flag.DefaultValue
	}
}

func typeString(flagType types.FlagType) string {
	switch flagType {
	case types.StringType:
		return "string"
	case types.IntType:
		return "int64"
	case types.BoolType:
		return "bool"
	case types.FloatType:
		return "float64"
	default:
		return ""
	}
}

// Go Funcs END

// Generate generates the Go flag accessors for OpenFeature.
func (g *genImpl) Generate(input types.Input) error {
	funcs := template.FuncMap{
		"FlagVarName":         flagVarName,
		"FlagInitParam":       flagInitParam,
		"OpenFeatureType":     openFeatureType,
		"SupportImports":      supportImports,
		"DefaultValueLiteral": defaultValueLiteral,
		"TypeString":          typeString,
	}
	td := TmplData{
		BaseTmplData: input.BaseData,
		GoPackage:    g.goPackage,
	}
	return generate.GenerateFile(funcs, golangTmpl, &td)
}

// Params are parameters for creating a Generator
type Params struct {
	GoPackage string
}

// NewGenerator creates a generator for Go.
func NewGenerator(params Params) types.Generator {
	return &genImpl{
		goPackage: params.GoPackage,
	}
}
