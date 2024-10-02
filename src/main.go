package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"

	flagmanifest "codegen/docs/schema/v0"
	generator "codegen/src/generators"
	"codegen/src/generators/golang"

	jsonschema "github.com/santhosh-tekuri/jsonschema/v5"
)

var flagManifestPath = flag.String("flag_manifest_path", "", "Path to the flag manifest.")
var moduleName = flag.String("module_name", "", "Name of the module to be generated.")
var outputPath = flag.String("output_path", "", "Output path for the codegen")

var stringToFlagType = map[string]generator.FlagType{
	"string":  generator.StringType,
	"boolean": generator.BoolType,
	"float":   generator.FloatType,
	"integer": generator.IntType,
	"object":  generator.ObjectType,
}

func getDefaultValue(defaultValue interface{}, flagType generator.FlagType) string {
	switch flagType {
	case generator.BoolType:
		return strconv.FormatBool(defaultValue.(bool))
	case generator.IntType:
		//the conversion to float64 instead of integer typically occurs
		//due to how JSON is parsed in Go. In Go's encoding/json package,
		//all JSON numbers are unmarshaled into float64 by default when decoding into an interface{}.
		return strconv.FormatFloat(defaultValue.(float64), 'g', -1, 64)
	case generator.FloatType:
		return strconv.FormatFloat(defaultValue.(float64), 'g', -1, 64)
	case generator.StringType:
		return defaultValue.(string)
	default:
		return ""
	}
}

func unmarshalFlagManifest(data []byte, supportedFlagTypes map[generator.FlagType]bool) (*generator.BaseTmplData, error) {
	dynamic := make(map[string]interface{})
	err := json.Unmarshal(data, &dynamic)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	sch, err := jsonschema.CompileString(flagmanifest.SchemaPath, flagmanifest.Schema)
	if err != nil {
		return nil, fmt.Errorf("error compiling JSON schema: %v", err)
	}
	if err = sch.Validate(dynamic); err != nil {
		return nil, fmt.Errorf("error validating JSON schema: %v", err)
	}
	// All casts can be done directly since the JSON is already validated by the schema.
	iFlags := dynamic["flags"]
	flags := iFlags.(map[string]interface{})
	btData := generator.BaseTmplData{
		OutputDir: path.Dir(*outputPath),
	}
	for flagKey, iFlagData := range flags {
		flagData := iFlagData.(map[string]interface{})
		flagTypeString := flagData["flagType"].(string)
		flagType := stringToFlagType[flagTypeString]
		if !supportedFlagTypes[flagType] {
			log.Printf("Skipping generation of flag %q as type %v is not supported for this language", flagKey, flagTypeString)
			continue
		}
		docs := flagData["description"].(string)
		defaultValue := getDefaultValue(flagData["defaultValue"], flagType)
		btData.Flags = append(btData.Flags, &generator.FlagTmplData{
			Name:         flagKey,
			Type:         flagType,
			DefaultValue: defaultValue,
			Docs:         docs,
		})
	}
	return &btData, nil
}

func loadData(manifestPath string, supportedFlagTypes map[generator.FlagType]bool) (*generator.BaseTmplData, error) {
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return nil, fmt.Errorf("error reading contents from file %q", manifestPath)
	}
	return unmarshalFlagManifest(data, supportedFlagTypes)
}

func generate(gen generator.Generator) error {
	btData, err := loadData(*flagManifestPath, gen.SupportedFlagTypes())
	if err != nil {
		return err
	}
	input := generator.Input{
		BaseData: btData,
	}
	return gen.Generate(input)
}

// command line params working example
// -flag_manifest_path "sample/golang/golang_sample.json" -output_path "sample/golang/golang_sample.go"
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
