package generator

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
)

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
	OutputDir string
	Flags     []*FlagTmplData
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

// GenerateFile receives data for the Go template engine and outputs the contents to the file.
func GenerateFile(funcs template.FuncMap, filename string, contents string, data TmplDataInterface) error {
	contentsTmpl, err := template.New("contents").Funcs(funcs).Parse(contents)
	if err != nil {
		return fmt.Errorf("error initializing template: %v", err)
	}

	var buf bytes.Buffer
	if err := contentsTmpl.Execute(&buf, data); err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	f, err := os.Create(path.Join(data.BaseTmplDataInfo().OutputDir, filename))
	if err != nil {
		return fmt.Errorf("error creating file %q: %v", filename, err)
	}
	defer f.Close()
	writtenBytes, err := f.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("error writing contents to file %q: %v", filename, err)
	}
	if writtenBytes != buf.Len() {
		return fmt.Errorf("error writing entire file %v: writtenBytes != expectedWrittenBytes", filename)
	}

	return nil
}
