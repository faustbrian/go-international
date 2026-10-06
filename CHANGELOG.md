# Changelog

All notable changes and dataset updates are recorded here.

## [Unreleased]

### Changed

- Advance International to `/v4` and adopt public `go-wire/v3` v3.0.0.
  Canonical and retained wire adapters now use Wire v3's nominal `Format` and
  categorized codec errors; migrate International and Wire imports together.
  Preserve ordinary supported-format round-trip fixtures, unsupported-format
  sentinel behavior, and the facade lifecycle. Intentionally inherit Wire v3's
  MessagePack aggregate-value and key-comparison-work admission limits and
  corrected YAML block/folded scalar handling; older accepted inputs or bytes
  are not universally preserved. MessagePack work-limit rejection matches
  `wire.ErrSizeLimit` and leaves the decode target unchanged. Wire codec
  messages are privacy-safe while their categories remain inspectable through
  `errors.Is` and `errors.As`. Preserve released v1, v2, and v3 API snapshots.

## [3.0.2] - 2026-10-06

### Changed

- Adopt public `go-config/v2` v2.0.0 for configuration integration tests,
  retaining strict scalar decoding and atomic validation assertions. No
  International public API or production import changes.

## [3.0.1] - 2026-10-04

### Changed

- Adopt `golang.org/x/text` v0.42.0 and its selected `x/sync` v0.23.0
  dependency. Explicit postal NFC normalization now applies corrected
  composition after Hangul, across intervening starters and for supplementary
  characters. Parsing and stored values remain unchanged unless callers
  request normalization. Keep the unchanged language registry data while
  identifying the active parser version in language and locale provenance.
- Adopt `go-config` v1.1.0, `go-wire` v1.0.1, and PGX v5.11.0 while
  retaining scalar decoding and legacy and successor adapter contracts.
  Applications implementing PGX `Rows` directly must provide its new
  `TypeMap` method; International uses `pgtype` scalar maps instead.

## [3.0.0] - 2026-10-02

### Changed

- Advance International to `/v3` for package-owned parse and fixture diagnostic
  classifications. Unrecognized caller kinds and reasons are replaced rather
  than echoed; callers requiring free-form diagnostics must own their redacted
  application errors. Preserve the published v1 and v2 API snapshots.
- Redact dataset, transport, response-body, and output-write failure details
  from generator diagnostics while preserving caller cancellation identity.
  Generator requests and output writes inherit command-interrupt cancellation.

## [2.0.0] - 2026-10-02

### Changed

- Refresh the checksum-pinned SIX current currency list to 2026-09-17 while
  retaining the historic list dated 2026-01-01. Generation accepts independently
  published lists; currency codes, names, minor units and status remain unchanged.
  Currency provenance records both publication dates and the refreshed checksum.
- Require Go 1.27.0 for the module and repository verification.
- Prepare the `github.com/faustbrian/go-international/v2` module on main and
  adopt published `go-validation/v2` v2.0.0. Canonical and retained Validation
  factories now return that major's nominal `Validator[string]` type.
- Preserve the v1 API baseline and release history. Applications must migrate
  International and Validation imports together; existing v1 consumers remain
  on their independently selected v1 dependencies.

## [1.1.0] - 2026-09-09

### Added

- Add target-oriented `adapters/postgres`, `adapters/validation`, and
  `adapters/wire` packages with behavior-preserving legacy compatibility paths.

### Changed

- Adopt public `go-validation` v1.1.0 while keeping International rules pure
  and synchronous.
- Document the adapter migration, support interval, rollback order, ownership,
  package selection, and error identities.

- Advance shared tooling to the checksum-verified `go-library-tools` v1.4.0
  release and align local configuration, inventory, cohesion, repository,
  online specification, workflow, and implementation gates.
- Resolve `go-config` and `go-wire` through their canonical public v1.0.0
  module archives and `go-validation` through public v1.1.0.
- Publish complete schema-v2 cohesion metadata and versioned Golib ecosystem
  navigation for the international module, its target adapters, and its test
  companion.

- Adopt the checksum-verified `go-library-tools` v1.3.0 CLI, add the local
  `make cohesion` validation entry point, and pin reusable-workflow cohesion
  enforcement to its final immutable revision.

- Govern standards-backed behavior through the
  [specification decision register](docs/specification-decisions.md), pinned
  source and update authorities, maintained-peer conformance bindings, and a
  canonical-tooling CI check. Initial decision records:
  `INTERNATIONAL-DEC-001 sha256:29ed9c6e8a623b4771cf13e9894d35f2ea1c81834ee2a49e49de8e5fbd6ecb25`,
  `INTERNATIONAL-DEC-002 sha256:83725f09367f86e94d18084b30feb57eb83178b6224c407ad99f0c06da41e68c`,
  `INTERNATIONAL-DEC-003 sha256:390292b5428e8c579d9e6010e690b664a24a859e075332b7e0023a72bf641e21`,
  `INTERNATIONAL-DEC-004 sha256:0b2904fb1cdd8a666bfc0f054f196b83fcd2cfa11c300c6df621c479b48605ce`,
  and `INTERNATIONAL-DEC-005 sha256:51c6c7168eadc56cc241187229200beb85bc392fc47f6e41a8b9d9e47884481c`.

- Adopt the checksum-verified `go-library-tools` v1.2.0 CLI and immutable
  shared workflow for specification governance, replacing the temporary
  source-built specification job while preserving module policy,
  package-owned checks, and content-addressed mutation evidence.

### Documentation

- Replace noisy CLDR GitHub-feed monitoring with Unicode's versioned stable
  distribution authority while retaining the four pinned CLDR 48.2 source
  payloads and explicit dataset review policy.
- Point ecosystem and package-family navigation at the immutable v1.4.0
  documentation set.
- Record the behavior-neutral review of the CLDR 49 alpha 2 release feed while
  retaining the pinned CLDR 48.2 dataset and its decision bindings.
- Record the behavior-neutral 2026-09-03 review of unchanged SIX ISO 4217
  List One and List Three data, and monitor their deterministic XML payloads
  instead of nondeterministic landing-page chrome.

- Replace the archived monorepo link with package-owned documentation.

## [1.0.0] - 2026-08-26

### Changed

- Refresh the reviewed zero-mutant identity for the extracted wire package
  without weakening the exact mutation contract.

- Exclude intentional nested modules from root local-proxy archives so local,
  bootstrap, CI, and public module checksums describe the same source
  boundary.

- Track the pinned documentation-tool lockfile so clean CI checkouts install
  the exact validated cspell dependency.

- Reconcile standalone dependency checksums against deterministic current
  module archives so CI, local verification, and release consumers resolve
  identical content.

- Harden standalone documentation validation with deterministic spelling and
  link checks, package-specific documentation gates, and repository-local
  contributor guidance.

### Documentation

- Link the package README to package-owned documentation.
- Keep the initial `v1.0.0` scope under Unreleased until a tag is published.

### Changed

- Publish the module from its standalone `github.com/faustbrian/go-international` identity while preserving its documented API and behavior.
- Replace obsolete owned-module pseudo-version pins with the monorepo's local
  `v0.0.0` source-proxy coordinates; release tooling continues to emit exact
  `v1.0.0` dependency versions.
- Repair malformed UTF-8 parse-error reasons before returning diagnostics and
  classify oversized dataset review inputs as resource-limit failures.
- Add opt-in country-aware postal-code syntax validation compatible with the
  pinned `brick/postcode` 0.5.0 rules used by legacy Postal callers.
- Restore source-resolvable owned dependency versions after local zero-version
  metadata made the international module unusable outside the repository.
- Add bounded, locale-neutral full Unicode lowercasing for protocols that
  require multi-rune and context-sensitive lowercase mappings.
- Replace nonexistent owned-module v0.1.0 tags with available main revisions
  so clean consumers can resolve the international module.
- Run API compatibility against the repository's isolated owned-module graph
  instead of stale published checksums.
- Refresh owned-module checksums against the final consolidated archives.
- Use deterministic execution counts for default fuzz smoke campaigns while
  allowing explicit duration overrides for extended fuzzing.
- Normalize standalone module metadata against the canonical owned dependency
  graph, including complete checksums for clean consumer resolution.
- Bound and sanitize public parse-error kinds so hostile diagnostics cannot
  panic or echo overlong caller-controlled labels.
- Isolate Make-based release gates from ignored local Go workspaces so local
  verification uses the dependency versions pinned in `go.mod`, matching CI.
- Replace the unpublished validation revision with its available successor
  and record its module checksums for reproducible clean-checkout builds.
- Require 100% mutation coverage and efficacy across all country, subdivision,
  language, locale, currency, phone, and postal acceptance and canonicalization
  implementations, in addition to selected semantic mutants.
- Keep provenance and documentation gates portable on clean CI runners without
  requiring ripgrep outside the declared Go toolchain.

### v1.0.0 scope

The following initial scope is included in `v1.0.0`.

- Establish typed country, subdivision, language, locale, currency, phone, and
  postal primitives.
- Pin CLDR 48.2, IANA 2026-06-14, ISO 4217 2026-01-01, and libphonenumber
  v9.0.32-compatible metadata.
- Add strict text, JSON, SQL, pgx, config, wire, and validation
  integration.
- Add deterministic generation, provenance, exact coverage, race, fuzz,
  benchmark, and privacy hardening gates.
- Preserve authoritative mappings for reused country and currency numeric
  identifiers and reject ambiguous historical policy expansions.
- Add explicit historic text, JSON, and SQL decoding without weakening strict
  default codecs.
- Add source-labelled independent fixture vectors and differential checks for
  country, locale, currency, phone, and package-policy postal behavior.
- Add explicit Unicode and resource-budget evidence, concurrent metadata tests,
  and a clean advisory NilAway gate.

[1.1.0]: https://github.com/faustbrian/go-international/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/faustbrian/go-international/releases/tag/v1.0.0

[3.0.1]: https://github.com/faustbrian/go-international/compare/v3.0.0...v3.0.1
[3.0.0]: https://github.com/faustbrian/go-international/compare/v2.0.0...v3.0.0
[2.0.0]: https://github.com/faustbrian/go-international/compare/v1.1.0...v2.0.0
