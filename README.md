# international

[![CI](https://github.com/faustbrian/go-international/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/faustbrian/go-international/actions/workflows/ci.yml)
[![CodeQL](https://img.shields.io/badge/CodeQL-required-blue)](https://github.com/faustbrian/go-international/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Mutation](https://img.shields.io/badge/mutation-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Documentation](https://img.shields.io/badge/docs-checked_in_CI-blue)](docs/)
[![Go Reference](https://pkg.go.dev/badge/github.com/faustbrian/go-international.svg)](https://pkg.go.dev/github.com/faustbrian/go-international)
[![Release](https://img.shields.io/github/v/release/faustbrian/go-international?sort=semver)](https://github.com/faustbrian/go-international/releases)
[![Go](https://img.shields.io/badge/go-1.26.6-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Typed, immutable international identifiers and metadata for Go services.
Countries, subdivisions, languages, locales, currencies, phone numbers, and
postal values remain distinct types with strict parsing, explicit
canonicalization, offline behavior, and versioned dataset provenance.

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
in the [verification report](docs/verification.md).

Requires Go 1.26.6 or newer. Licensed under MIT; dataset licenses remain with
their upstream publishers.

## Documentation

Start with the [documentation index](docs/README.md) for datasets, provenance,
integration, migration, and operations guidance.
