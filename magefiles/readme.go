//go:build mage
// +build mage

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

func generateReadme(contract string) error {
	fmt.Printf("    📄 Generating %s readme\n", contract)

	readme, err := os.ReadFile(filepath.Join(CONTRACTS_TMP_DIR, "docs", fmt.Sprintf("%s.md", contract)))
	name := strings.TrimPrefix(contract, "axone-")
	schemaName := fmt.Sprintf("%s-schema", name)
	if err != nil {
		return err
	}

	params := ReadmeParams{
		Docs:         string(readme),
		Name:         name,
		ContractName: contract,
		SchemaName:   schemaName,
		Ref:          tag(),
	}

	return ts(params)
}

func ts(params ReadmeParams) error {
	fmt.Println("       ➡️ Typescript readme")
	tmpl, err := template.ParseFiles(filepath.Join("ts", "README.md.template"))
	if err != nil {
		return err
	}

	readmeFile, err := os.Create(filepath.Join("ts", params.SchemaName, "README.md"))
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	return tmpl.Execute(readmeFile, params)
}
