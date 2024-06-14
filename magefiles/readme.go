package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type ReadmeParams struct {
	Docs         string
	Name         string
	ContractName string
	SchemaName   string
	Ref          string
}

// generate readmes for the given contract in all target languages
func generateReadme(contract string) error {
	fmt.Printf("    📄 Generating %s readme\n", contract)

	readme, err := os.ReadFile(filepath.Join(CONTRACTS_TMP_DIR, "docs", fmt.Sprintf("%s.md", contract)))
	if err != nil {
		return fmt.Errorf("could not read contract readme: %w", err)
	}

	name := strings.TrimPrefix(contract, "axone-")
	schemaName := fmt.Sprintf("%s-schema", name)

	params := ReadmeParams{
		Docs:         string(readme),
		Name:         name,
		ContractName: contract,
		SchemaName:   schemaName,
		Ref:          tag(),
	}

	return ts(params)
}

// generate typescript readmes
func ts(params ReadmeParams) error {
	fmt.Println("       ➡️ Typescript readme")
	tmplPath := filepath.Join("ts", "README.md.template")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template '%s': %w", tmplPath, err)
	}

	readmeFile, err := os.Create(filepath.Join("ts", params.SchemaName, "README.md"))
	if err != nil {
		return fmt.Errorf("could not create readme file: %w", err)
	}
	defer readmeFile.Close()

	return tmpl.Execute(readmeFile, params)
}
