package slogo

import (
	"os"

	"github.com/OpenSLO/go-sdk/pkg/openslo"
	"github.com/OpenSLO/go-sdk/pkg/openslosdk"
)

// ReadYAMLFile reads multiple OpenSLO objects from a YAML file.
// The YAML file can contain multiple objects separated by '---'.
func ReadYAMLFile(filename string) ([]openslo.Object, error) {
	return readFile(filename, openslosdk.FormatYAML)
}

// ReadJSONFile reads multiple OpenSLO objects from a JSON file.
// The JSON file should contain an array of objects.
func ReadJSONFile(filename string) ([]openslo.Object, error) {
	return readFile(filename, openslosdk.FormatJSON)
}

// WriteYAMLFile writes multiple OpenSLO objects to a YAML file.
// Objects will be separated by '---' in the output file.
func WriteYAMLFile(filename string, objects ...openslo.Object) error {
	return writeFile(filename, openslosdk.FormatYAML, objects...)
}

// WriteJSONFile writes multiple OpenSLO objects to a JSON file.
// Objects will be written as a JSON array.
func WriteJSONFile(filename string, objects ...openslo.Object) error {
	return writeFile(filename, openslosdk.FormatJSON, objects...)
}

// readFile reads multiple OpenSLO objects from a file with the specified format.
func readFile(filename string, format openslosdk.ObjectFormat) ([]openslo.Object, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return openslosdk.Decode(f, format)
}

// writeFile writes multiple OpenSLO objects to a file with the specified format.
func writeFile(filename string, format openslosdk.ObjectFormat, objects ...openslo.Object) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return openslosdk.Encode(f, format, objects...)
}
