package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
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

func (BumpVersion) Go(version string) error {
	fmt.Println("🔖 Bump go packages version")

	err := isVersionTagValid(version)
	if err != nil {
		return err
	}

	majorVersion := strings.Split(version, ".")[0]

	packages, err := os.ReadDir(GO_DIR)
	if err != nil {
		return fmt.Errorf("failed to read go directory '%s': %w", GO_DIR, err)
	}

	for _, pkg := range packages {
		if !pkg.IsDir() {
			continue
		}

		moduleName, err := outputInPath(filepath.Join(GO_DIR, pkg.Name()), "bash", "-c",
			"go mod edit -json | jq -r '.Module.Path'")
		if err != nil {
			return fmt.Errorf("failed to get module name: %w", err)
		}

		fmt.Printf("    ➡️ Bumping %s to version %s\n", moduleName, version)
		moduleNameUnversioned := regexp.
			MustCompile(`/v[0-9]+$`).
			ReplaceAllString(moduleName, "")
		moduleNameVersioned := fmt.Sprintf("%s/%s", moduleNameUnversioned, majorVersion)

		fmt.Printf("    🔬 Updating module path to %s\n", moduleNameVersioned)
		err = runInPath(filepath.Join(GO_DIR, pkg.Name()), "go",
			"mod",
			"edit",
			"-module",
			moduleNameVersioned)
		if err != nil {
			return err
		}
	}
	return nil
}

type Publish mg.Namespace

func (Publish) Ts(schema string, ref string) error {
	mg.Deps(mg.F(Build.Ts, schema))

	fmt.Println("🚀 Publishing typescript packages")

	name := strings.TrimPrefix(schema, "axone-")
	dest := filepath.Join(TS_DIR, fmt.Sprintf("%s-schema", name))

	if strings.HasPrefix(ref, "refs/tags/v") {
		return sh.Run("yarn",
			"--cwd", dest,
			"publish",
			"--access=public",
			"--no-git-tag-version",
			"--non-interactive",
			"--tag", "latest")
	} else {
		date := time.Now().Format("20060102150405")
		return sh.Run("yarn",
			"--cwd", dest,
			"publish",
			"--access=public",
			"--no-git-tag-version",
			"--non-interactive",
			"--prerelease",
			"--preid", fmt.Sprintf("next.%s", date),
			"--tag", "next")
	}
}
