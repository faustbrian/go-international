# Migration guide

## Version 3 diagnostic redaction

Use `github.com/faustbrian/go-international/v3` consistently across an
application. Validation remains on published `go-validation/v2` v2.0.0.
International v1, v2, and v3 values have distinct nominal Go identities.

`NewParseError` retains only package-defined kind and reason classifications.
Callers that supplied free-form diagnostic text must use their own explicitly
redacted application errors; package parse failures still match
`international.ErrInvalid` through `errors.Is`. `FormatFixtureError` replaces
unrecognized fixture kinds with `value`.

Released v1 and v2 API bytes remain in `api/baseline.txt` and `api/v2.txt`.
The current v3 API is maintained separately in `api/v3.txt`; no
version-specific source directory or branch is required.

## Version 2 module adoption

The published v2 module uses `github.com/faustbrian/go-international/v2` and
its package paths together
with `github.com/faustbrian/go-validation/v2` v2.0.0 on Go 1.27. The canonical
`adapters/validation` and retained `internationalvalidation` factories both
return Validation v2's `Validator[string]`; their synchronous successful-value
behavior is unchanged.

Update all International imports in one application composition. International
v1 and v2 values are distinct nominal Go types, not interchangeable aliases.
Keep consumers using the v1 Money, Localized, and Knapsack contracts on their
existing International v1 dependency until those consumers deliberately adopt
matching major identities. Tools' retained v1 fixture is likewise not evidence
of v2 adoption. Select public versions from their stable tags and releases.

The released v1 API bytes remain in `api/baseline.txt`; the released v2 API is
retained separately in `api/v2.txt`. No version-specific source directory or
branch is required.

## Target-oriented adapters within v1

No call signature changes are required. Change the import path as follows:

| Existing import | Successor import |
| --- | --- |
| `github.com/faustbrian/go-international/internationalpgx` | `github.com/faustbrian/go-international/adapters/postgres` |
| `github.com/faustbrian/go-international/internationalvalidation` | `github.com/faustbrian/go-international/adapters/validation` |
| `github.com/faustbrian/go-international/internationalwire` | `github.com/faustbrian/go-international/adapters/wire` |

The PostgreSQL successor's package identifier is `internationalpostgres`.
Rename `internationalpgx.Register` selectors accordingly, or explicitly alias
the successor import as `internationalpgx` while migrating. The Validation and
wire successors retain their legacy package identifiers.

The existing paths remain supported for the longer of 180 days after public
successor availability and two subsequently published stable root-module minor
releases. First adopt public `go-validation v1.1.0`, then adopt International
v1.1.0 and migrate imports independently. Before publication, rollback the
coherent owner change normally. After publication, keep both import families
resolvable and use a forward patch; never delete or move the published tag.

## From `cline/intl`

Replace universal string codes with the matching package type. Move implicit
case repair to explicit `Canonicalize`/`Canonical` calls. Replace locale
truncation with a selected fallback policy. Audit zero and null behavior before
changing database columns.

## From country-list packages

Persist `country.Code`, not a display name. Replace hand-maintained maps with
`Alpha3`, `Numeric`, `Name`, and `DatasetProvenance`. Decide explicitly whether
historic, reserved, or user-assigned values are accepted. When importing such
values from JSON or SQL, use the options-bearing decode methods; the default
interfaces intentionally continue to accept current identifiers only.

## From Brick PhoneNumber or wrappers

Parse using `phone.Parse` with an explicit region hint for national input.
Persist E.164 plus the separate extension through the built-in codecs. Replace
single “valid” booleans with `Possible` and `Valid`, and remove any ownership,
SMS capability, or reachability inference.

## Rollout

Inventory stored values, run a dry parse that records only aggregate counts,
classify rejected obsolete aliases, choose opt-in policies, backfill canonical
representations transactionally, and dual-read before switching writes. Never
log rejected phone or postal values. Pin the package and dataset versions in
the migration record.
