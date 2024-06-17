package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/magefile/mage/sh"
)

// RunInPath runs a command in a specific path.
func runInPath(path string, cmd string, args ...string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	defer func() {
		dirs := filepath.SplitList(path)
		_ = os.Chdir(strings.Repeat("../", len(dirs)))
	}()

	return sh.Run(cmd, args...)
}

func isVersionTagValid(tag string) error {
	pattern := `v(?P<Major>0|(?:[1-9]\d*))(?:\.(?P<Minor>0|(?:[1-9]\d*))(?:\.(?P<Patch>0|(?:[1-9]\d*)))?(?:\-(?P<PreRelease>[0-9A-Z\.-]+))?(?:\+(?P<Meta>[0-9A-Z\.-]+))?)?`
	re, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("failed to compile regex: %w", err)
	}

	if !re.MatchString(tag) {
		return fmt.Errorf("tag version '%s' is not valid, should be 'vX.Y.Z'", tag)
	}

	return nil
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

	if err := sh.Run("cargo", "make", "--help"); err != nil {
		if err := sh.Run("cargo", "install", "cargo-make"); err != nil {
			panic(fmt.Sprintf("failed to install cargo-make: %v", err))
		}
	}
}

// EnsureQuicktype ensures that quicktype is installed, if not it panics.
func ensureQuicktype() {
	if err := sh.Run("quicktype", "--help"); err != nil {
		panic("quicktype is not installed")
	}
}

// EnsureYarn ensures that yarn is installed, if not it panics.
func ensureYarn() {
	if err := sh.Run("yarn", "--help"); err != nil {
		panic("yarn is not installed")
	}
}
