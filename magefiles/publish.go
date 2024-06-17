package main

import (
	"fmt"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"os"
	"path/filepath"
)

type BumpVersion mg.Namespace

func (BumpVersion) Ts(version string) error {
	fmt.Println("🔖 Bump typescript packages version")

	ensureYarn()

	err := isVersionTagValid(version)
	if err != nil {
		return err
	}

	packages, err := os.ReadDir(TS_DIR)
	if err != nil {
		return fmt.Errorf("failed to read typescript directory '%s': %w", TS_DIR, err)
	}

	for _, pkg := range packages {
		if !pkg.IsDir() {
			continue
		}
		fmt.Printf("    ➡️ Bumping %s to version %s\n", pkg.Name(), version)
		err := sh.Run("yarn",
			"--cwd", filepath.Join(TS_DIR, pkg.Name()),
			"version",
			"--new-version", version,
			"--allow-same-version",
			"--no-git-tag-version")
		if err != nil {
			return err
		}
	}
	return nil
}
