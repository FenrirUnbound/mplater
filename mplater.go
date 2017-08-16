package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/urfave/cli"
)

func parseKeyValues(list []string) (map[string]string, error) {
	result := make(map[string]string)

	for _, item := range list {
		parts := strings.Split(item, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Invalid input value: %v", item)
		}

		key, value := parts[0], parts[1]
		result[key] = value
	}

	return result, nil
}

func primeTemplateEngine(file string) (*template.Template, error) {
	absolutePath, err := filepath.Abs(file)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFiles(absolutePath)

	return tmpl, nil
}

func process(c *cli.Context) error {
	templateValues, err := parseKeyValues(c.StringSlice("x"))
	if err != nil {
		return err
	}

	outputFile, _ := filepath.Abs(c.String("o"))
	fileWriter, err := os.Create(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered while creating the output file: %v", err)
		return err
	}

	tmpl, err := primeTemplateEngine(c.String("i"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered while priming the template engine: %v", err)
		return err
	}

	err = tmpl.Execute(fileWriter, templateValues)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encountered while completing the template: %v", err)
		return err
	}

	fileWriter.Close()
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "mplater"
	app.Usage = "A template filler-upper"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Usage: "Input template file",
		},
		cli.StringFlag{
			Name:  "output, o",
			Usage: "Output file",
			Value: "output.txt",
		},
		cli.StringSliceFlag{
			Name:  "keyvalue, x",
			Usage: "Key-Value pair",
		},
	}

	app.Action = process

	app.Run(os.Args)
}
