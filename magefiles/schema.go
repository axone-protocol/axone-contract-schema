//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const (
	CONTRACTS_REPOSITORY = "https://github.com/axone-protocol/contracts.git"
	CONTRACTS_TMP_DIR    = "tmp/contracts"
)

type Schema mg.Namespace

// Download download contracts schemas
func (Schema) Download(ref string) error {
	mg.Deps(Schema.Clean)

	fmt.Println("📥 Downloading contracts schemas")

	ensureGit()

	if err := sh.Run("git",
		"clone",
		"--depth", "1",
		"--branch", ref,
		CONTRACTS_REPOSITORY,
		CONTRACTS_TMP_DIR); err != nil {
		return err
	}

	return nil
}

// Generate build and generate contracts json schema
func (s Schema) Generate(ref string) error {
	mg.Deps(mg.F(Schema.Download, ref))

	defer func() {
		s.Clean()
	}()

	fmt.Println("🔨 Generating contracts json schema")

	ensureCargo()

	return nil
}

func (Schema) Clean() error {
	fmt.Println("🧹 Cleaning schema temporary files")
	return sh.Rm(CONTRACTS_TMP_DIR)
}

func ensureGit() {
	if err := sh.Run("command", "-v", "git"); err != nil {
		panic("git is not installed")
	}
}

func ensureCargo() {
	if err := sh.Run("command", "-v", "cargo"); err != nil {
		panic("cargo is not installed")
	}
}
