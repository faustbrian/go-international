# International v4 security supplement

## Version and scope

Model version: 2. Reviewed on 2026-10-07 against public International v4.0.0,
commit `7511086bcf28a8f8d745da1c0a570c5fa688de5a`, module
`github.com/faustbrian/go-international/v4`, Go 1.27. The
[original v3 model](threat-model.md) retains its historical source attribution;
this supplement does not relabel that review as v4 evidence.

The original assets, trust boundaries, input policies and owned residual risks
remain applicable except for the changes below. Maintainers own package
classifications, dataset and dependency pins, and adapter defaults. Application
owners own ingress limits, explicit personal-value output, comparison policy,
custom callbacks and log sinks. Generation operators own trusted output paths
and visibility of command diagnostics. The [security policy](../SECURITY.md)
defines private reporting and coordinated disclosure.

## Changed boundaries and current evidence

| Public change | Current boundary and verification |
|---|---|
| v3.0.1 normalization and dependencies | Explicit postal NFC normalization adopts corrected `x/text` behavior. Parsing and stored values do not automatically normalize. Applications must select comparison policy rather than treating normalized text as identity proof. See [postal](../postal/postal.go) and [verification](verification.md). |
| v3.0.2 Config integration | Config v2 adoption affects test composition, not an implicit runtime configuration or environment reader. Core values remain synchronous and perform no configuration acquisition. |
| v4.0.0 Wire integration | Both canonical and retained adapters use public Wire v3 types and safe producer defaults. MessagePack aggregate-value and key-comparison work limits refuse excess work without changing the decode target. Errors retain inspectable categories without default payload text; YAML scalar corrections are intentional compatibility changes. See [dispatch](../internal/wireadapter/dispatch.go) and [migration](migration.md). |

Wire adapters intentionally serialize application-selected values. Their safe
defaults do not sanitize arbitrary caller envelopes, custom serialization code
or a later log sink. Explicit phone/postal serialization remains a deliberate
personal-data release boundary, not encryption or permission to log values.

Current source has successful [main CI](https://github.com/faustbrian/go-international/actions/runs/37473523902)
and [release rehearsal](https://github.com/faustbrian/go-international/actions/runs/37473584849).
Retained source reviews and an actual clean public v4 consumer cover fifteen
packages, five wire formats, canonical/retained adapter parity, PostgreSQL and
Validation composition, redacted error classification, and MessagePack refusal
with unchanged targets. These are current release-boundary evidence; historical
benchmark or mutation prose is not presented as a fresh v4 execution.

## Diagnostic and generator disclosure disposition

This source assessment establishes bounded hardening and dependency adoption,
not an International-specific vulnerability requiring a new advisory. Earlier
constructor documentation described redaction more broadly than its actual
behavior; the following version distinctions are therefore important.

| Boundary | Reviewed released behavior | First changed International release and disposition |
|---|---|---|
| Public parse-error constructor | v1.0.0, v1.1.0 and v2.0.0 retain a bounded UTF-8 prefix of caller-supplied reason text and a permitted caller-selected kind. `Error` formats those fields; `Unwrap` exposes only the invalid-value sentinel. Package-owned parsers inspected in these releases pass literal classifications rather than rejected values or dependency messages. | v3.0.0 restricts both fields to package-defined classifications. Earlier exposure requires application code to supply sensitive or untrusted diagnostic text and then expose the error. No automatic attacker-to-protected-data path was established. |
| Fixture diagnostics | In those three earlier releases, fixture formatting renders the selected kind, but never the underlying error text. Owned fixture helpers select fixed labels. | v3.0.0 substitutes unknown kinds with a safe classification. This hardens trusted tooling input; downstream tools still own arbitrary labels and diagnostic visibility. |
| Dataset generation | Those earlier releases render some malformed dataset fields and wrap transport, body-read or output-write errors. Fixed HTTPS sources, a 30-second HTTP timeout, 8 MiB admission and checksum-before-parsing already bound production acquisition. Output paths are operator-selected. | v3.0.0 adds classified diagnostics and request/command cancellation. This is explicit maintainer-operation hardening: ordinary library input does not invoke generation or choose a dataset URL. Custom I/O cooperation and partial output remain operational limitations, not a remote application ingress route. |
| Explicit NFC normalization | v3.0.1 adopts corrected normalization after Hangul, across intervening starters and supplementary characters. Normalization is explicitly selected, not implicit parsing. | Correctness/dependency update. Security impact depends on an application's use of normalized values as security identities; International does not promise that property. |
| Wire adoption | International v3.0.0 pins Wire v1.0.0; v3.0.1 and v3.0.2 pin Wire v1.0.1. v4.0.0 first adopts Wire v3.0.0. | v4.0.0 is the first International adoption of these producer controls, not necessarily the first producer fix. Wire vulnerability/version dispositions remain producer-owned; a separate International advisory or arbitrary custom-serializer privacy guarantee is not inferred. |

The reviewed versions above identify actual behavior, not a fabricated
affected-version vulnerability range. No severity, exploitation, incident,
absence of historical impact, or universal safety of earlier versions is
asserted. Repository advisory listings alone cannot establish those facts.
Reopen private triage if a supported integration demonstrates protected-data
exposure, exploitable parser work or a security-sensitive normalization mismatch.

The released [v2 constructor](https://github.com/faustbrian/go-international/blob/v2.0.0/errors.go),
[v2 ingress codecs](https://github.com/faustbrian/go-international/blob/v2.0.0/internal/codec/codec.go),
[v3 changes](https://github.com/faustbrian/go-international/compare/v2.0.0...v3.0.0),
and [v3.0.1 changes](https://github.com/faustbrian/go-international/compare/v3.0.0...v3.0.1)
provide the historical source boundary for this assessment.

## Residual ownership and review conditions

The [v3 residual-risk records](threat-model.md#accepted-contract-limitations-and-residual-risks)
remain in force, including explicit output, metadata staleness, syntax policy,
caller-owned normalization, cooperative generation and non-atomic writes.

| Owner | Rationale and mitigation | Review condition |
|---|---|---|
| International maintainers | Fixed classifications, pinned defaults and reviewed dependency updates protect owned boundaries without sandboxing application code. Preserve diagnostic, adapter and refusal regression tests. | Changed diagnostics, dependencies, formats or input surfaces. |
| Application owners | Explicit serialization and comparison are necessary application operations. Apply ingress limits, personal-data redaction, access/retention policy and explicit identity comparison rules. | New output, log, authorization or normalization-based identity integration. |
| Generation operators and maintainers | Generation is opt-in offline tooling with trusted paths. Use bounded cooperative I/O, restrict command logs, inspect the complete generated diff and do not publish partial output. | New collaborators, output workflows, remote inputs or atomicity requirements. |
| Wire producer and application owners | International composes producer defaults rather than replacing its parser or surrounding application policy. Preserve safe defaults and inspect relevant producer advisories. | Producer findings, decoder changes, custom encoding or newly supported formats. |
