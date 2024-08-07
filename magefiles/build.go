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
	if schema == "axone-cognitarium" {
		fmt.Println("🚪 Skipping axone-cognitarium schema since codegen generation is failing")
		return nil
	}

	fmt.Printf("⚙️ Generate go types for %s\n", schema)

	ensureGoCodegen()

	_, dest := schemaDestination(schema, GO_DIR)
	if err := os.MkdirAll(dest, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	err := sh.Run("go-codegen", "generate",
		"messages",
		filepath.Join(SCHEMA_DIR, fmt.Sprintf("%s.json", schema)),
		"-o", filepath.Join(dest, "schema.go"),
		"--package-name", "schema")
	if err != nil {
		return fmt.Errorf("failed to generate go types: %w", err)
	}

	fmt.Println("👨‍💻 Generate go client")
	err = sh.Run("go-codegen", "generate",
		"query-client",
		filepath.Join(SCHEMA_DIR, fmt.Sprintf("%s.json", schema)),
		"-o", filepath.Join(dest, "client.go"),
		"--package-name", "schema")
	if err != nil {
		return fmt.Errorf("failed to generate go client: %w", err)
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
