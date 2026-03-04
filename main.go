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

	fmt.Println(mapping)

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
