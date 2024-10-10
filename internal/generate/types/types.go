// Package types contains all the common types and interfaces for generating flag accessors.
package types

// FlagType are the primitive types of flags.
type FlagType int

// Collection of the different kinds of flag types
const (
	UnknownFlagType FlagType = iota
	IntType
	FloatType
	BoolType
	StringType
	ObjectType
)

// FlagTmplData is the per-flag specific data.
// It represents a common interface between Mendel source and codegen file output.
type FlagTmplData struct {
	Name         string
	Type         FlagType
	DefaultValue string
	Docs         string
}

// BaseTmplData is the base for all OpenFeature code generation.
type BaseTmplData struct {
	OutputPath string
	Flags      []*FlagTmplData
}

type TmplDataInterface interface {
	// BaseTmplDataInfo returns a pointer to a BaseTmplData struct containing
	// all the relevant information needed for metadata construction.
	BaseTmplDataInfo() *BaseTmplData
}

type Input struct {
	BaseData *BaseTmplData
}

// Generator provides interface to generate language specific, strongly-typed flag accessors.
type Generator interface {
	Generate(input Input) error
	SupportedFlagTypes() map[FlagType]bool
}
