//go:build (darwin && cgo) || linux || mage
// +build darwin,cgo linux mage

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Build mg.Namespace

// Ts build typescript schema for the given contract schema.
func (Build) Ts(schema string) error {
	mg.Deps(mg.F(Clean.Ts, schema))

	fmt.Printf("⚙️ Generate typescript types for %s\n", schema)

	ensureTsCodegen()

	name, dest := schemaDestination(schema, TS_DIR)

	err := sh.Run("ts-codegen", "generate",
		"--schema", filepath.Join(SCHEMA_DIR, schema),
		"--out", filepath.Join(dest, "gen-ts"),
		"--typesOnly",
		"--no-bundle",
		"--name", name,
	)
	if err != nil {
		return fmt.Errorf("failed to generate typescript types: %w", err)
	}

	fmt.Println("🔨 Building typescript")

	err = sh.Run("yarn", "--cwd", dest)
	if err != nil {
		return err
	}

	return sh.Run("yarn", "--cwd", dest, "build")
}

// Go build go schema for the given contract schema.
func (Build) Go(schema string) error {
	fmt.Printf("⚙️ Generate go types for %s\n", schema)

	ensureQuicktype()

	_, dest := schemaDestination(schema, GO_DIR)

	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	err := sh.Run("bash", "-c",
		fmt.Sprintf("quicktype -s schema %s -o %s --lang go --package schema",
			filepath.Join(SCHEMA_DIR, schema, "*.json"),
			filepath.Join(dest, "schema.go")))
	if err != nil {
		return fmt.Errorf("failed to generate go types: %w", err)
	}

	fmt.Println("🔨 Building go")
	return runInPath(dest, "go", "build")
}

type Clean mg.Namespace

func (Clean) Ts(schema string) error {
	fmt.Printf("🧹 Cleaning generated typescript files for %s\n", schema)

	_, dest := schemaDestination(schema, TS_DIR)

	return sh.Run("yarn", "--cwd", dest, "clean")
}

func schemaDestination(schema, root string) (name string, destination string) {
	name = strings.TrimPrefix(schema, "axone-")
	destination = filepath.Join(root, fmt.Sprintf("%s-schema", name))
	return
}
