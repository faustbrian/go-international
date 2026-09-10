# international

[![CI](https://github.com/faustbrian/go-international/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/faustbrian/go-international/actions/workflows/ci.yml)
[![CodeQL](https://img.shields.io/badge/CodeQL-required-blue)](https://github.com/faustbrian/go-international/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Mutation](https://img.shields.io/badge/mutation-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Documentation](https://img.shields.io/badge/docs-checked_in_CI-blue)](docs/)
[![Go Reference](https://pkg.go.dev/badge/github.com/faustbrian/go-international.svg)](https://pkg.go.dev/github.com/faustbrian/go-international)
[![Release](https://img.shields.io/github/v/release/faustbrian/go-international?sort=semver)](https://github.com/faustbrian/go-international/releases)
[![Go](https://img.shields.io/badge/go-1.27.0-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Typed, immutable international identifiers and metadata for Go services.
Countries, subdivisions, languages, locales, currencies, phone numbers, and
postal values remain distinct types with strict parsing, explicit
canonicalization, offline behavior, and versioned dataset provenance.

```sh
go get github.com/faustbrian/go-international@v1.1.0
```

```go
finland, err := country.Parse("FI")
if err != nil { return err }

tag, err := locale.Parse("fi-FI")
if err != nil { return err }

number, err := phone.Parse("040 123 4567", phone.ParseOptions{
    RegionHint: finland,
})
```

The zero value of every scalar means absent. Text encoding rejects absent
values; JSON and SQL encode them as `null`/`NULL`. Parsing never performs
country inference, locale detection, delivery validation, identity claims, or
runtime network access.

Start with the [five-minute quickstarts](docs/quickstart.md), then read the
[API and standards reference](docs/reference.md), [integration guide](docs/integrations.md),
and [security model](SECURITY.md). Dataset versions and licenses are documented
in [provenance](docs/provenance.md); the checked semantic baseline and update
classification procedure are in the [dataset report](docs/dataset-report.md).
The requirement-to-test mapping, resource budgets, and local gate evidence are
in the [verification report](docs/verification.md). Observable interpretations
are recorded in the [specification decision register](docs/specification-decisions.md).
Shared construction, ownership, lifecycle, and composition expectations are in
the versioned [Golib ecosystem index](https://github.com/faustbrian/go-library-tools/blob/v1.4.0/docs/ecosystem/README.md)
and its [Foundations family](https://github.com/faustbrian/go-library-tools/blob/v1.4.0/docs/ecosystem/design-language.md#package-families-and-selection).

The module is stable and active and requires Go 1.27.0 or newer. Its APIs are
stateless: callers retain all configuration, buffers, pgx maps, and validation
contexts, and the module starts no background work. Parse and validation
failures preserve `international.ErrInvalid` or
`international.ErrResourceLimit`; unsupported wire formats preserve the
calling adapter's `ErrUnsupportedFormat` sentinel.

## Package map

| Package | Use |
| --- | --- |
| root | Shared status, parse-error, resource-limit, normalization, and dataset-diff contracts |
| `country`, `subdivision` | ISO 3166 and governed subdivision identities |
| `language`, `locale` | ISO 639 and bounded BCP 47 identities |
| `currency` | ISO 4217 alphabetic and numeric identities |
| `phone`, `postal` | Bounded, privacy-safe phone and country-bound postal values |
| `adapters/postgres` | pgx text registration on a caller-owned type map |
| `adapters/validation` | Pure International rules for Golib Validation |
| `adapters/wire` | Strict bounded JSON, XML, YAML, TOML, and MessagePack dispatch |
| `internationaltest` | Test-only governed fixtures and assertions |

The former `internationalpgx`, `internationalvalidation`, and
`internationalwire` paths remain supported compatibility facades; new code
should use the target-oriented adapters. See the
[compiler-checked adapter example](adapter_successor_test.go) and the
[migration guide](docs/migration.md).

Licensed under MIT; dataset licenses remain with their upstream publishers.
For help, see [Support](SUPPORT.md). Report vulnerabilities through the
private process in [Security](SECURITY.md).

## Documentation

Start with the [documentation index](docs/README.md) for datasets, provenance,
integration, migration, and operations guidance.
