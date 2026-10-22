# Contributing

## Before Editing

1. Read [`AGENTS.md`](AGENTS.md) and the affected package's documentation.
2. Run ordinary Go tests for the affected package before editing.
3. Identify owned dependencies and reverse dependants in `modules.json`.
4. Preserve unrelated work and generated/corpus provenance.

## Changes

Keep commits focused and conventional. Update every affected changelog with
the behavior and migration impact. Public API changes require compatibility
evidence and documentation. Specification behavior requires a decision record,
fixture coverage, and interoperability evidence.

New direct dependencies and dependency updates must follow the
[dependency governance policy](AGENTS.md#dependencies-and-supply-chain). Package-local
update bots are forbidden; the root policy owns every module and action update.

Specification-backed changes must follow the
[specification governance contract](AGENTS.md#design), update
the affected entries in the [specification decision register](docs/specification-decisions.md)
and its machine bindings, and complete the Specification Decisions section of
the pull request template. An unresolved interpretation or stale source pin is
release-blocking; peer behavior cannot silently select policy.

Selected mutation gates must finish with zero surviving viable mutants.

Do not fork shared verification workflows, add permanent replacements, machine-specific
paths, bypass flags, broad mutation exclusions, or aggregate quality metrics
that hide a failing package.

## Verification

Use Go 1.27.2 for development and CI; the public minimum remains Go 1.27.0.
An ordinary checkout is sufficient for package tests and a focused active fuzz
smoke; no sibling repository or goal ledger is required:

```bash
go test ./...
go test . -run '^$' -fuzz '^FuzzTextParsers$' -fuzztime=2s
```

Select the native target and package affected by your change. Ordinary Go tests
already exercise deterministic fuzz seeds. For shared checks, install the
public `golib` v2.4.1 release declared in `.golib.yaml` and verify its checksums;
CI uses that same release, not a sibling source build:

```bash
make check
make ci
```

Both commands use bounded local checks. Optional focused mutation analysis can
select a single affected package, for example:

```bash
golib mutation --module . --package ./locale
```

Use `make check-full` for explicit full assurance, including the online
specification check, full configured fuzz budgets, API compatibility, conformance,
authoritative generated-data drift and provenance. Dataset changes must also
follow [the update procedure](docs/dataset-report.md) and
[source provenance](docs/provenance.md); core identifier operations stay offline.
Run affected adapter integration and direct consumer tests for public-contract
changes. Explicit full and release qualification retain their selected gates;
ordinary scheduled shared CI remains bounded.
Report unavailable or failing checks without claiming release readiness.
CI/configuration-only changes do not require a library release.

## Selected Monthly Assurance

The owner-level [selected-assurance caller](.github/workflows/assurance.yml) is
configured for main only, on day 4 of each month at 05:41 UTC. It selects one
of the 6 current declared root-module fuzz targets using the shared UTC-week
rotation, with 30 seconds of active exploration. The selected risks are:
identifier, phone/postal and persistence parsing, generated XML/ranges,
and wire dispatch panic resilience.
These targets use in-process inputs; they do not update upstream datasets or
qualify generated-data drift/provenance, full conformance or direct consumers.

Base-free ordinary checks execute deterministic seeds, not active exploration;
committed-range CI selects affected active smoke. Manual assurance requires an
exact anchored selector from the root fuzz operation in `.golib.yaml`, for
example `^FuzzTextParsers$`.
Empty, unknown and malformed selections refuse. The shared job bounds setup
and execution to ten minutes, active workers and CPU to two, owned memory to
2 GiB, native output to 1 MiB and private-cache lifetime. Failure, cancellation,
exhausted budgets and cleanup refusal remain attributable failures. This lane
does not run mutation or replace Required, original full fuzz budgets,
explicit `make check-full`, or relevant release qualification.

## Adding A Module

Follow [repository structure policy](AGENTS.md#repository-structure). New modules
require an explicit purpose, ownership boundary, dependency review, package
catalog entry, full quality gates, documentation, changelog, license, security
policy, compatibility plan, and release dry-run.

## Source Documentation

Install the checksum-pinned public verifier declared in `.golib.yaml`.
`make source-docs-check` validates objective Go comment, marker, link and
generated-producer structure; `make check`, `make check-full` and `make ci`
require it before their existing checks. Shared CI runs the same strict
contract and includes its result in Required. This is structural enforcement,
not a substitute for source-truth review or behavioral verification.
