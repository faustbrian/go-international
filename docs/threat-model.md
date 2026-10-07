# International threat model

## Version and scope

Model version: 1. Reviewed source boundary: International `v3.0.0`, commit
`36c54fea1e040d757af0e7d927f2e4aeb02883f5`, module
`github.com/faustbrian/go-international/v3`, Go 1.27.0. The
[module manifest](../go.mod) and [package inventory](../modules.json) define
the dependency and package boundary. This document describes that source;
it does not change parsing, dependency pins, or release behavior.

The [current v4 supplement](threat-model-v4.md) records subsequent changes and
their disclosure disposition. This v3 source attribution remains historical.

International maintainers own the library contracts, generated datasets,
dependency review, and this model. Application owners own ingress limits,
authorization, persistence policy, presentation, and handling of personal data.
Generation operators own their network client, filesystem permissions, and
chosen output paths. Private reporting remains governed by the
[security policy](../SECURITY.md).

## Assets, actors, and trust boundaries

The protected assets are identifier and mapping integrity, predictable parsing
work, stable persistence semantics, generated-source integrity, and privacy of
phone and postal values. The relevant untrusted actor can supply application
input or malformed encoded values; upstream changes or compromised dependency
and dataset distribution can affect a future maintainer update. These are
threat assumptions, not reports of a known compromise.

| Boundary | Ownership and existing controls |
|---|---|
| Application input to core values | Country, currency, subdivision, and language parsers enforce fixed representations and registry policies. Locale and phone parsing apply explicit input budgets before dependency parsing. Postal parsing stores bounded printable UTF-8 with explicit country context. See [country](../country/country.go), [currency](../currency/currency.go), [subdivision](../subdivision/subdivision.go), [language](../language/language.go), [locale](../locale/locale.go), [phone](../phone/phone.go), and [postal](../postal/postal.go). |
| Typed values to persistence and integrations | [Scalar codecs](../internal/codec/codec.go) restrict JSON and SQL forms; [phone](../phone/encoding.go) and [postal](../postal/encoding.go) preserve their explicit persistence representations. [Postgres registration](../adapters/postgres/postgres.go) mutates only a caller-owned type map and must precede concurrent use. [Validation](../adapters/validation/validation.go) is synchronous parser-backed policy; [wire dispatch](../internal/wireadapter/dispatch.go) delegates supported formats to pinned dependencies and rejects unsupported formats. |
| Compiled metadata to consumers | Immutable values and independent returned slices avoid handing consumers mutable dataset ownership. Display names, phone classification, and postal syntax are metadata, not authorization or identity verification. Existing decisions are in the [specification register](specification-decisions.md). |
| Maintainer acquisition to generated source | The [generator](../internal/generate/main.go) requests pinned sources, bounds reads, verifies checksums before parsing, validates dataset structure, sorts output, quotes generated string literals, and formats Go source. The operator chooses output paths; the command is a maintainer tool, not a service for untrusted requests. |
| Dependency and dataset updates to a release | Maintainers review pinned versions, licenses, semantic differences, and compatibility under the [provenance procedure](provenance.md), [dataset report](dataset-report.md), and [engineering policy](../AGENTS.md). Consumers select and deploy their own dependency updates. |

Core parsing, lookup, normalization, and metadata formatting are in-process.
They do not invoke the generator, open connections, update datasets, log, or
emit telemetry. Optional adapters do not own database connections or perform
connection I/O. The retained `internationalpgx`, `internationalvalidation`,
and `internationalwire` paths have the compatibility boundaries documented in
[integrations](integrations.md); they are not separate online services.

This model does not cover a compromised application process, unsafe/reflection
mutation, caller callbacks or custom serialization behavior, database/server
security, transport authentication, or release credentials. Those remain
under their respective application and operational owners; a typed value does
not create a sandbox around caller code.

## Input, ambiguity, and privacy controls

The authoritative [resource budgets and Unicode evidence](verification.md#resource-budgets-and-unicode-threat-model)
cover parser bytes and segments, extension storage, codecs, diagnostics,
dataset review, and generated inputs. This model does not duplicate those
limits. [Unicode lowercasing](../case.go) additionally requires a caller-owned
input and expanded-output budget, with a per-call transformer.

Strict fixed-code parsers reject non-ASCII lookalikes and unexpected casing.
Locale parsing preserves accepted spelling; canonicalization and fallback
are explicit. Postal normalization is separately selected and ordinary postal
values remain opaque. These policies reduce ambiguity but do not provide
universal Unicode spoofing protection. Applications must select comparison
and display policy for their own context rather than treating normalized text
or localized names as authoritative identifiers.

Default-current acceptance and explicit historic options prevent silently
loading non-current country, subdivision, and currency records. Numeric
parsers reject ambiguity under the selected status policy and retain their
authoritative mapping. Existing [decisions 001, 002, and 005](specification-decisions.md)
describe the compatibility rationale; metadata updates do not authorize
automatic reinterpretation of stored application records.

[Parse errors](../errors.go) retain only package-defined classifications rather
than arbitrary caller reasons. Default phone and postal `String` and `GoString`
are redacted. Explicit value access, formatting, text/JSON encoding, and SQL
values intentionally expose data needed for application use: redacted default
formatting is not encryption, storage protection, or a guarantee against
application logging. Applications must redact those outputs, limit transport
input before decoding, and avoid personal values in reports and fixtures.
General wire envelopes and dependency errors are not a promise that every
caller-controlled object or field is private.

`phone.Valid` and `Possible` are distinct pinned metadata results, not ownership,
reachability, or identity proof. `postal.ValidSyntax` is an optional pinned
compatibility operation, not proof of existence, locality, address correctness,
or deliverability. Unlike bounded `postal.Parse`, its public contract requires
the caller to bound untrusted input before syntax normalization. It compiles
only fixed package-owned patterns. See [phone](../phone/phone.go),
[postal syntax](../postal/syntax.go), and [decisions 003 and 004](specification-decisions.md).

## Generation and update operations

Generation is a distinct trust boundary with network and filesystem I/O.
The production [generator](../internal/generate/main.go) uses fixed HTTPS
source URLs and a timed HTTP client; checksums bind the accepted bytes rather
than attesting that an upstream source is semantically correct. Ordinary
library input does not select a dataset URL. Pins, source authorities,
licenses, and update responsibilities are recorded in [provenance](provenance.md).

The [generation command](../cmd/international-generate/main.go) passes interrupt
cancellation to requests. `RunContext` checks cancellation before acquisition
and before each output write. Successful response bodies are closed through
the download path; bounded reads and checksum verification precede parsing.
Transport, body-read, and output-write failure details are redacted while
cancellation/deadline identity remains available.

Internal injected HTTP doers, fetch functions, and writers are trusted
collaborators, not untrusted extension interfaces. Their owners must preserve
valid HTTP response/body contracts and implement appropriate bounds and
cancellation. Request context does not forcibly interrupt a custom doer or
body reader that ignores it. The writer has no context argument; the check
before a write cannot interrupt a write already in progress. Generation also
does not promise rollback or an atomic multi-file transaction: earlier output
writes may remain after later failure. Operators must use trusted output paths,
inspect the complete generated diff, and avoid publishing partial output.

Maintainers own dependency updates, source-pin changes, license review, semantic
dataset review, and release compatibility. Phone metadata and postal syntax
have their own documented upstream update procedures. Pinned data deliberately
does not refresh at runtime; application owners decide when to adopt a release
and whether stored data requires revalidation or migration.

## Accepted contract limitations and residual risks

The following records describe limits already selected by maintained contracts
and specification decisions, not newly approved vulnerability exceptions. The
owner identifies who must manage the remaining risk; this model does not
invent an individual acceptance or an additional support commitment.

| Residual risk and existing basis | Owner | Rationale and mitigation | Review condition |
|---|---|---|---|
| Stale or semantically incorrect upstream metadata; [provenance](provenance.md) and decisions 001–005 | International maintainers for pins; application owners for adoption | Offline, reproducible behavior intentionally uses pinned snapshots. Review semantic diffs, licenses, differential vectors, advisories, and compatibility before updates; revalidation of persisted application data remains explicit. | A relevant upstream advisory, dataset change, preferred mapping change, or dependency update. |
| Classification or syntax mistaken for identity/authorization; [decisions 003 and 004](specification-decisions.md) | Application owners | A library cannot prove phone ownership or postal deliverability. Keep metadata decisions separate from authentication, contact verification, and provider/address policy. | A consumer uses these results for authorization, routing, or a new correctness claim. |
| Caller-visible personal data and Unicode presentation ambiguity; [phone](../phone/phone.go), [postal](../postal/postal.go), and [security policy](../SECURITY.md) | Application owners | Explicit access and serialization are necessary for use; opaque postal values retain caller text. Redact logs/telemetry, apply transport bounds, protect storage, and choose explicit normalization/comparison policies. | A new output/logging path, comparison policy, or privacy-sensitive integration. |
| Work outside parser limits, including unbounded caller input to `ValidSyntax` or generic envelopes; [postal syntax](../postal/syntax.go) and [wire dispatch](../internal/wireadapter/dispatch.go) | Application owners | Library-local limits do not bound request acquisition or every surrounding operation. Bound ingress and syntax inputs, and review pinned decoder limits for the selected envelope. | A new input surface, envelope format, or decoder/dependency behavior change. |
| Cooperative cancellation and partial generator writes; [generator](../internal/generate/main.go) | Generation operators and International maintainers | Trusted I/O seams and ordinary file writes avoid claiming universal interruption or transaction guarantees. Use bounded cooperative clients, trusted destinations, complete-diff review, and drift checks before publication. | A new custom I/O collaborator, output workflow, or atomicity/cancellation requirement. |
| Trusted dependency/toolchain or source-pin compromise; [module manifest](../go.mod) and [engineering policy](../AGENTS.md) | International maintainers | Pins and checksums constrain selected bytes, not all upstream intent. Review dependency/source changes and apply relevant release security checks; consumers assess their deployment boundary. | A relevant advisory, pin/toolchain update, or changed release trust boundary. |

## Evidence and review maintenance

[Verification](verification.md) maps existing behavior to tests, governed
fixtures, differential comparisons, bounded-input cases, and release gates.
Its dated benchmark and mutation results retain their original versions and
environments; they are not a fresh Go 1.27 security assessment. The source
identity above identifies the modeled release, not evidence that every
security threat has been tested or eliminated.

Existing source, release, and ordinary clean-consumer evidence may establish
the behaviors and delivery boundaries they actually exercise. They do not
constitute a broader security certification. This documentation update does
not freshly run vulnerability/secret scanners, hostile-input campaigns,
external-service checks, or runtime verification, and does not claim their
results. It is not a module release or a new tag.

Maintainers should update this versioned model when public acceptance,
privacy, dependency, dataset, adapter, or generation trust boundaries change.
Any new vulnerability exception or materially broader acceptance decision
requires its actual owner and rationale; no such decision is inferred from
this document or from a green ordinary consumer test.
