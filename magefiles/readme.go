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

// generate readmes for the given contract in all target languages.
func generateReadmes(contract string) error {
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

	for _, languagePath := range []string{GO_DIR, TS_DIR} {
		if err := generate(languagePath, params); err != nil {
			return err
		}
	}
	return nil
}

func generate(languagePath string, params ReadmeParams) error {
	fmt.Printf("       ➡️ %s readme\n", languagePath)
	tmplPath := filepath.Join(languagePath, "README.md.template")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template '%s': %w", tmplPath, err)
	}

	readmeFile, err := os.Create(filepath.Join(languagePath, params.SchemaName, "README.md"))
	if err != nil {
		return fmt.Errorf("could not create readme file: %w", err)
	}
	defer readmeFile.Close()

	return tmpl.Execute(readmeFile, params)
}
