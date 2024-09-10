package generator

import (
	"bytes"
	"fmt"
	"os"
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

type Generator interface {
	Generate(input Input) error
	SupportedFlagTypes() map[FlagType]bool
}

func GenerateFile(funcs template.FuncMap, outputPath string, contents string, data TmplDataInterface) error {
	contentsTmpl, err := template.New("contents").Funcs(funcs).Parse(contents)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := contentsTmpl.Execute(&buf, data); err != nil {
		return err
	}

	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer f.Close()
	writtenBytes, err := f.Write(buf.Bytes())
	if err != nil {
		return err
	}
	if writtenBytes != buf.Len() {
		return fmt.Errorf("error writing entire file %v: writtenBytes != expectedWrittenBytes", outputPath)
	}

	return nil
}
