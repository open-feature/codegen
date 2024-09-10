package main

import (
	generator "codegen/generators"
	"codegen/generators/golang"
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"strconv"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

const flagManifestSchema = "../docs/schema/v0/flag_manifest.json"

var flagManifestPath = flag.String("flag_manifest_path", "", "Path to the flag manifest.")
var moduleName = flag.String("module_name", "", "Name of the module to be generated.")
var outputPath = flag.String("output_path", "", "Output path for the codegen")

var stringToFlagType = map[string]generator.FlagType{
	"string":  generator.StringType,
	"boolean": generator.BoolType,
	"float":   generator.FloatType,
	"integer": generator.IntType,
}

func getDefaultValue(defaultValue interface{}, flagType generator.FlagType) string {
	switch flagType {
	case generator.BoolType:
		return strconv.FormatBool(defaultValue.(bool))
	case generator.IntType:
		return strconv.FormatInt(defaultValue.(int64), 10)
	case generator.FloatType:
		return strconv.FormatFloat(defaultValue.(float64), 'g', -1, 64)
	case generator.StringType:
		return defaultValue.(string)
	}
	return ""
}

func unmarshalFlagManifest(data []byte) (*generator.BaseTmplData, error) {
	dynamic := make(map[string]interface{})
	err := json.Unmarshal(data, &dynamic)
	if err != nil {
		return nil, err
	}

	sch, err := jsonschema.Compile(flagManifestSchema)
	if err != nil {
		return nil, err
	}
	if err = sch.Validate(dynamic); err != nil {
		return nil, err
	}
	iFlags := dynamic["flags"]
	flags := iFlags.(map[string]interface{})
	btData := generator.BaseTmplData{
		OutputDir: path.Dir(*outputPath),
	}
	for flagKey, iFlagData := range flags {
		flagData := iFlagData.(map[string]interface{})
		flagType := stringToFlagType[flagData["flag_type"].(string)]
		docs := flagData["description"].(string)
		defaultValue := getDefaultValue(flagData["default_value"], flagType)
		btData.Flags = append(btData.Flags, &generator.FlagTmplData{
			Name:         flagKey,
			Type:         flagType,
			DefaultValue: defaultValue,
			Docs:         docs,
		})
	}
	return &btData, nil
}

func loadData(manifestPath string) (*generator.BaseTmplData, error) {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error reading contents from file %q", manifestPath)
	}
	return unmarshalFlagManifest(data)
}

func generate(gen generator.Generator) error {
	btData, err := loadData(*flagManifestPath)
	if err != nil {
		return err
	}
	input := generator.Input{
		BaseData: btData,
	}
	return gen.Generate(input)
}

func main() {
	flag.Parse()
	_, filename := path.Split(*outputPath)
	params := golang.Params{
		File: filename,
		// Probably some conversion applied here, toLower and remove special characters.
		GoPackage: *moduleName,
	}
	gen := golang.NewGenerator(params)
	err := generate(gen)
	if err != nil {
		fmt.Printf("error generating flag accesssors: %v\n", err)
	}
}
