# Contract schemas

> Generated types for [AXONE contracts schemas](https://github.com/axone-protocol/contracts).

[![version](https://img.shields.io/github/v/release/axone-protocol/axone-contract-schema?style=for-the-badge&logo=github)](https://github.com/axone-protocol/axone-contract-schema/releases)
[![build](https://img.shields.io/github/actions/workflow/status/axone-protocol/axone-contract-schema/build.yml?branch=main&label=build&style=for-the-badge&logo=github)](https://github.com/axone-protocol/axone-contract-schema/actions/workflows/build.yml)
[![lint](https://img.shields.io/github/actions/workflow/status/axone-protocol/axone-contract-schema/lint.yml?branch=main&label=lint&style=for-the-badge&logo=github)](https://github.com/axone-protocol/axone-contract-schema/actions/workflows/lint.yml)
[![test](https://img.shields.io/github/actions/workflow/status/axone-protocol/axone-contract-schema/test.yml?branch=main&label=test&style=for-the-badge&logo=github)](https://github.com/axone-protocol/axone-contract-schema/actions/workflows/test.yml)
[![conventional commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge&logo=conventionalcommits)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg?style=for-the-badge)](https://github.com/semantic-release/semantic-release)
[![contributor covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg?style=for-the-badge)](https://github.com/axone-protocol/.github/blob/main/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg?style=for-the-badge)](https://opensource.org/licenses/BSD-3-Clause)

## Purpose

This repository contains [AXONE contract schemas](https://github.com/axone-protocol/contracts) and enables the generation of JSON Schema types in multiple programming languages, including:

- TypeScript ([ts](ts/))

## Usage

### Requirements

- [mage](https://magefile.org/) 1.15+
- [Go](https://golang.org/) 1.22+

### Available commands

```bash
Targets:
  build:go           build go schema for the given contract schema.
  build:ts           build typescript schema for the given contract schema.
  bumpVersion:go     bumps the version of the go packages with the given version.
  bumpVersion:ts     bumps the version of the typescript packages with the given version.
  publish:ts         publishes the typescript packages for the given schema.
  schema:clean       remove temporary files.
  schema:download    download contracts schemas at a given ref.
  schema:generate    build and generate contracts json schemas at the given ref.
  schema:readme      generate contracts readme on all target.
```

Download and generate schema for a specific version:

```bash
mage schema:generate v5.0.0
```

Build targeted language schema:

```bash
mage build:ts axone-objectarium
```

## You want to get involved? 😍

Please check out AXONE health files :

- [Contributing](https://github.com/axone-protocol/.github/blob/main/CONTRIBUTING.md)
- [Code of conduct](https://github.com/axone-protocol/.github/blob/main/CODE_OF_CONDUCT.md)
