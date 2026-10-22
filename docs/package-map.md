# Package map

[Documentation index](README.md) · [Repository entry](../README.md)

This map follows the current source modules, not a synchronized release.
Use the repository README for published installation versions; an unpublished
source major is not installable merely because its package appears here.
Internal packages and non-releasable fixtures are intentionally excluded.

## github.com/faustbrian/go-international/v5

Typed, immutable international identifiers and metadata for Go services. Countries, subdivisions, languages, locales, currencies, phone numbers, and postal values remain distinct types with strict parsing, explicit canonicalization, offline behavior, and versioned dataset provenance.

| Import path | Responsibility |
| --- | --- |
| [`github.com/faustbrian/go-international/v5`](../case.go) | Use for shared international status, parse-error, resource-limit, Unicode normalization, and dataset-diff contracts. |
| [`github.com/faustbrian/go-international/v5/adapters/postgres`](../adapters/postgres/postgres.go) | Use to register International scalar values as PostgreSQL text on caller-owned pgx type maps. |
| [`github.com/faustbrian/go-international/v5/adapters/validation`](../adapters/validation/validation.go) | Use to validate International string values through the Golib Validation contract. |
| [`github.com/faustbrian/go-international/v5/adapters/wire`](../adapters/wire/wire.go) | Use to encode and decode International values through strict bounded Golib wire formats. |
| [`github.com/faustbrian/go-international/v5/country`](../country/doc.go) | Use for immutable ISO 3166 country identifiers, explicit historic-status policy, canonicalization, and display metadata. |
| [`github.com/faustbrian/go-international/v5/currency`](../currency/doc.go) | Use for immutable ISO 4217 currency identifiers, numeric mappings, minor-unit metadata, and explicit historic-status policy. |
| [`github.com/faustbrian/go-international/v5/internationaltest`](../internationaltest/fixtures.go) | Use only in tests for governed international fixtures and reusable assertions. |
| [`github.com/faustbrian/go-international/v5/language`](../language/encoding.go) | Use for immutable ISO 639 language identifiers and IANA registry-backed canonicalization. |
| [`github.com/faustbrian/go-international/v5/locale`](../locale/encoding.go) | Use for bounded BCP 47 locale tags, canonicalization, matching, and explicit fallback. |
| [`github.com/faustbrian/go-international/v5/phone`](../phone/encoding.go) | Use for bounded phone-number parsing, E.164 identity, metadata validation, and privacy-safe formatting. |
| [`github.com/faustbrian/go-international/v5/postal`](../postal/encoding.go) | Use for country-bound opaque postal values, explicit normalization, and opt-in syntax validation. |
| [`github.com/faustbrian/go-international/v5/subdivision`](../subdivision/doc.go) | Use for immutable ISO 3166-2-derived subdivision identifiers and explicit historic-status policy. |

Choose the next guide in the [documentation index](README.md)
or return to the [repository entry](../README.md) for installation and ownership.
