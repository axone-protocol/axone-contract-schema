# OKP4 law-stone schema

> Generated typescript types for [okp4-law-stone contract](https://github.com/okp4/contracts/tree/v2.1.0/contracts/okp4-law-stone).

[![version](https://img.shields.io/github/v/release/okp4/okp4-contract-schema?style=for-the-badge&logo=github)](https://github.com/okp4/okp4-contract-schema/releases)
[![build](https://img.shields.io/github/actions/workflow/status/okp4/okp4-contract-schema/build.yml?branch=main&label=build&style=for-the-badge&logo=github)](https://github.com/okp4/okp4-contract-schema/actions/workflows/build.yml)
[![lint](https://img.shields.io/github/actions/workflow/status/okp4/okp4-contract-schema/lint.yml?branch=main&label=lint&style=for-the-badge&logo=github)](https://github.com/okp4/okp4-contract-schema/actions/workflows/lint.yml)
[![test](https://img.shields.io/github/actions/workflow/status/okp4/okp4-contract-schema/test.yml?branch=main&label=test&style=for-the-badge&logo=github)](https://github.com/okp4/okp4-contract-schema/actions/workflows/test.yml)
[![conventional commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg?style=for-the-badge&logo=conventionalcommits)](https://conventionalcommits.org)
[![semantic-release](https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg?style=for-the-badge)](https://github.com/semantic-release/semantic-release)
[![contributor covenant](https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg?style=for-the-badge)](https://github.com/okp4/.github/blob/main/CODE_OF_CONDUCT.md)
[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg?style=for-the-badge)](https://opensource.org/licenses/BSD-3-Clause)

## Usage 

First add or install the module to your existing project using either `yarn` or `npm`. 

```bash
yarn add @okp4/law-stone-schema
```

or 
```bash
npm install --save @okp4/law-stone-schema
```

Then import wanted type : 

```typescript
import type { InstantiateMsg } from "@okp4/law-stone-schema";
```

---

# okp4-law-stone Schema

```txt
undefined
```

# Law Stone

## Overview

The `okp4-law-stone` smart contract aims to provide GaaS (i.e. Governance as a Service) in any [Cosmos blockchains](https://cosmos.network/) using the [CosmWasm](https://cosmwasm.com/) framework and the [Logic](https://docs.okp4.network/modules/next/logic) OKP4 module.

This contract is built around a Prolog program describing the law by rules and facts. The law stone is immutable, this means it can only been questioned, there is no update mechanisms.

The `okp4-law-stone` responsibility is to guarantee the availability of its rules in order to question them, but not to ensure the rules application.

To ensure reliability over time, the associated Prolog program is stored and pinned in a `okp4-objectarium` contract. Moreover, all the eventual loaded files must be stored in a `okp4-objectarium` contract as well, allowing the contract to pin them.

To be able to free the underlying resources (i.e. objects in `okp4-objectarium`) if not used anymore, the contract admin can break the stone.

➡️ Checkout the [examples](https://github.com/okp4/contracts/tree/main/contracts/okp4-law-stone/examples/) for usage information.

| Abstract            | Extensible | Status         | Identifiable            | Custom Properties | Additional Properties | Access Restrictions | Defined In                                                               |
| :------------------ | :--------- | :------------- | :---------------------- | :---------------- | :-------------------- | :------------------ | :----------------------------------------------------------------------- |
| Can be instantiated | No         | Unknown status | Unknown identifiability | Forbidden         | Allowed               | none                | [okp4-law-stone.json](schema/okp4-law-stone.json "open original schema") |

## okp4-law-stone Type

unknown ([okp4-law-stone](okp4-law-stone.md))


