package main

import (
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
	"strings"
)

// RunInPath runs a command in a specific path.
func RunInPath(path string, cmd string, args ...string) error {
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
func EnsureGit() {
	if err := sh.Run("command", "-v", "git"); err != nil {
		panic("git is not installed")
	}
}

// EnsureCargo ensures that cargo is installed, if not it panics.
func EnsureCargo() {
	if err := sh.Run("command", "-v", "cargo"); err != nil {
		panic("cargo is not installed")
	}
}

// EnsureCargoMake ensures that cargo-make is installed, if not, it tries to install it.
func EnsureCargoMake() {
	EnsureCargo()

	if err := sh.Run("cargo", "make", "--help"); err == nil {
		return
	}

	sh.Run("cargo", "install", "--force", "cargo-make")
}
