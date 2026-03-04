package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/abdul-local/mockdata/data"
)

func main() {

	var inputFile, outputPath string

	var help bool
	flag.StringVar(&inputFile, "i", "", "Input json file")
	flag.StringVar(&inputFile, "input", "", "Input json file")
	flag.StringVar(&outputPath, "o", "", "Output json file")
	flag.StringVar(&outputPath, "output", "", "Output json file")
	flag.BoolVar(&help, "h", false, "Show help!")
	flag.BoolVar(&help, "help", false, "Show help!")

	flag.Parse()

	if help || inputFile == "" || outputPath == "" {
		printUsage()
		os.Exit(1)
	}
	// validate input
	if err := validationInput(inputFile); err != nil {
		fmt.Printf("invalid input: %s\n", err.Error())
		os.Exit(1)

	}

	// validate output
	if err := validationOutput(outputPath); err != nil {
		fmt.Printf("invalid output: %s\n", err.Error())
		os.Exit(1)

	}

	// prosess input

	var mapping map[string]string

	if err := readInput(inputFile, &mapping); err != nil {
		fmt.Printf("failed reading input: %s\n", err.Error())
		os.Exit(1)
	}

	if err := validateType(mapping); err != nil {
		fmt.Printf("invalid type: %s\n", err.Error())
		os.Exit(1)

	}
	// membuat data palsu
	result, err := generateOutput(mapping)
	if err != nil {
		fmt.Printf("failed generating output: %s\n", err.Error())
		os.Exit(0)
	}

	// menulis hasil ke file
	if err := writeOutput(outputPath, result); err != nil {
		fmt.Printf("failed writing output: %s\n", err.Error())
		os.Exit(0)
	}

	fmt.Println("success run ")

}

func printUsage() {
	fmt.Println("Usage: fakegen [-i | --input] <input file> [-o | --output] <output path>")
	fmt.Println("-i, --input: Input of JSON file as a template")
	fmt.Println("-o, --output: Output JSON file for the generated data")
}

func validationInput(input string) error {
	if _, err := os.Stat(input); os.IsNotExist(err) {
		return err

	}
	return nil
}

func validationOutput(input string) error {
	if _, err := os.Stat(input); os.IsNotExist(err) {
		return nil

	}
	fmt.Println("Output file already exist")
	confirmOverwrite()
	return nil
}

func confirmOverwrite() {

	fmt.Print("Are you sure you want to overwrite the file? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))
	if response != "y" && response != "yes" {
		fmt.Println("Aborting...")
		os.Exit(1)
	}

}

func readInput(path string, mapping *map[string]string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	if mapping == nil {
		return errors.New("mapping is null")
	}
	// read file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fileByte, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileByte) == 0 {
		return errors.New("file is empty")
	}
	if err = json.Unmarshal(fileByte, &mapping); err != nil {
		return err
	}
	return nil
}

func validateType(mapping map[string]string) error {
	for _, value := range mapping {
		if !data.Supported[value] {
			return fmt.Errorf("%s type is not supported", value)
		}

	}
	return nil
}

func generateOutput(mapping map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, value := range mapping {
		result[key] = data.Generate(value)
	}
	return result, nil
}

func writeOutput(path string, result map[string]any) error {
	if path == "" {
		return errors.New("path is empty")
	}
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC // 25121024
	file, err := os.OpenFile(path, flags, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// marshal result dengan indentasi
	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// tulis ke file
	_, err = file.Write(resultBytes)
	if err != nil {
		return err
	}

	return nil
}
