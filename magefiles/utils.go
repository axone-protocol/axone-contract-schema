//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
)

// RunInPath runs a command in a specific path.
func runInPath(path string, cmd string, args ...string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	defer func() {
		dirs := filepath.SplitList(path)
		os.Chdir(strings.Repeat("../", len(dirs)))
	}()

	return sh.Run(cmd, args...)
}

// EnsureGit ensures that git is installed, if not it panics.
func ensureGit() {
	if err := sh.Run("git", "--help"); err != nil {
		panic("git is not installed")
	}
}

// EnsureCargo ensures that cargo is installed, if not it panics.
func ensureCargo() {
	if err := sh.Run("cargo", "--help"); err != nil {
		panic("cargo is not installed")
	}
}

// EnsureCargoMake ensures that cargo-make is installed, if not, it tries to install it.
func ensureCargoMake() {
	ensureCargo()

	if err := sh.Run("cargo", "make", "--help"); err == nil {
		return
	}

	sh.Run("cargo", "install", "--force", "cargo-make")
}

// EnsureQuicktype ensures that quicktype is installed, if not it panics.
func ensureQuicktype() {
	if err := sh.Run("quicktype", "--help"); err != nil {
		panic("quicktype is not installed")
	}
}
