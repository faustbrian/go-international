# Source documentation audit

## Pre-edit scope and acceptance

Base `592e3779ef639424ad5204e6ca4b82239f25ea74`, clean detached main-source worktree.
This worker owns Tier A source comments/docs and the confirmed Tier C repair
of three declared generation routes; original main and other work remain untouched. Outcome: useful actual Unicode/identity/status/metadata/phone/
postal/persistence/adapter/resource contracts and adjacent rationale, without
runtime/dependency/authority/fixture or publication changes. Direct consumers
are the core identifier families, wire/validation/Postgres adapters and retained
facades plus Config/Validation/Wire public contracts named by the catalog.

Named risks: confusing parser acceptance with normalization/metadata authority,
implicit locale fallback and inferred regions, Unicode byte/rune/output bounds,
absence/null persistence versus SQL NULL, phone possible/valid versus ownership,
postal syntax versus deliverability, mutable inputs/returned copies/provider
callbacks and cache lifecycle, generated producer mismatch and frozen corpora.
Smallest proof: all-source AST/token/literal identity and exact directives except
the three approved target corrections, formatting,
native21 package docs and selected existing Examples/docs/API; offline source/
spec/local links/markers/protected bytes/pinned staged secrets and ordinary
complete TierA final-diff review. No race/fuzz/conformance/services/performance/
online dataset-refresh campaign for unchanged executable contracts. All Go only
after explicit root grant with pinned1.27.2 and three owned disposable caches.

## Reproducible complete inventory

`git ls-files "*.go"`, `modules.json`, `packages.json` and native source-docs
inventory/check yield105 tracked Go files, one v4 module and21 packages. Initial
input is70 exported/interface/field omissions plus one external-generator review,
not a score for prose quality. No prior source-audit file was found.

- `.`: `adapter_legacy_coexistence_test.go`, `adapter_successor_test.go`, `benchmark_test.go`, `case.go`, `case_test.go`, `concurrency_test.go`, `config_integration_test.go`, `dataset.go`, `encoding_roundtrip_test.go`, `errors.go`, `fuzz_test.go`, `international_test.go`, `provenance.go`, `status.go`.
- `adapters/postgres`: `postgres.go`, `postgres_test.go`.
- `adapters/validation`: `validation.go`, `validation_test.go`.
- `adapters/wire`: `fuzz_test.go`, `wire.go`, `wire_test.go`, `wire_v3_test.go`.
- `cmd/international-dataset-review`: `main.go`.
- `cmd/international-generate`: `main.go`.
- `country`: `country.go`, `country_test.go`, `data_generated.go`, `dataset.go`, `differential_test.go`, `doc.go`, `encoding.go`, `example_test.go`, `internal_test.go`, `provenance.go`.
- `currency`: `currency.go`, `currency_test.go`, `data_generated.go`, `dataset.go`, `differential_test.go`, `doc.go`, `encoding.go`, `example_test.go`, `internal_test.go`, `provenance.go`.
- `internal/codec`: `codec.go`, `codec_test.go`.
- `internal/datasetreview`: `review.go`, `review_test.go`.
- `internal/generate`: `adversarial_test.go`, `country_test.go`, `currency_test.go`, `fuzz_test.go`, `main.go`, `run_test.go`, `subdivision_test.go`.
- `internal/wireadapter`: `dispatch.go`, `dispatch_test.go`.
- `internationalpgx`: `register.go`, `register_test.go`.
- `internationaltest`: `fixtures.go`, `fixtures_test.go`, `vectors.go`, `vectors_test.go`.
- `internationalvalidation`: `validation.go`, `validation_test.go`.
- `internationalwire`: `fuzz_test.go`, `wire.go`, `wire_test.go`.
- `language`: `encoding.go`, `internal_test.go`, `language.go`, `language_test.go`, `provenance.go`.
- `locale`: `differential_test.go`, `encoding.go`, `example_test.go`, `internal_test.go`, `locale.go`, `locale_test.go`, `provenance.go`.
- `phone`: `differential_test.go`, `encoding.go`, `example_test.go`, `internal_test.go`, `mapping_test.go`, `phone.go`, `phone_test.go`, `provenance.go`.
- `postal`: `encoding.go`, `example_test.go`, `internal_test.go`, `normalization_regression_test.go`, `postal.go`, `postal_test.go`, `syntax.go`, `syntax_compatibility_test.go`, `syntax_test.go`.
- `subdivision`: `data_generated.go`, `dataset.go`, `doc.go`, `encoding.go`, `internal_test.go`, `provenance.go`, `subdivision.go`, `subdivision_test.go`.

## Missing comments versus source questions

Review all public declarations and important private parser/canonicalization/
dataset/generator/codec/dispatch constraints with test/helper/fuzz/benchmark/
example coverage. Missing field comments are structural; actual default/nil/
zero/absence/alias/status policy, UTF8/normalization, returned-value ownership,
callback/IO behavior, generated authority and resource guarantees require source
inspection rather than generic filler. Preserve all historical registers/pins
and normative expected behavior; record real limitations without endorsing them.

## Generation, variants and native selection

Three local generated data files remain exact, with pinned CLDR48.2/SIX inputs
and genuine producer under internal/generate. Declared go:generate directives in
country/currency/subdivision doc.go named a non-main internal/generate package
at the base; the native regeneration script uses cmd/international-generate.
The coordinator confirmed the exact three target corrections as trusted generation
configuration (Tier C), retaining flags and output names. They were corrected
only after observing all three pre-fix non-main failures in the granted session.
Producer executable bodies and generated bytes remain frozen; one independent
complete-final-diff review is required. The externally generated Brick postcode0.5.0
corpus retains exact source pin/license/header and its manual availability review;
docs/provenance.md identifies source extraction inputs but no owned corpus
regeneration executable was found. No generated file is manually rewritten.

Retained internationalwire/internationalvalidation/internationalpgx facades and
intentional provider/type differences are audited independently. No platform/
build-tag split is present. Configured native API is apidiff with api/v4.txt;
do not regenerate immutable signatures as GoDoc text. Native docs invokes
verification/package.mk docs and scripts/check-docs.sh; it renders all packages
and rejects unfinished prose, but does not run Examples. Select existing Examples
separately. Generated-data/provenance recipes perform pinned upstream acquisition
and comparison; inspect their producer/frozen bytes without executing a dataset
campaign for untouched behavior. The changed declared routes instead require
focused real package-directory invocation and native authenticated generation/drift
proof, with exact frozen output and dataset-snapshot comparison. Coordinator owns
shared enforcement, independent generation review and delivery.

## Semantic coverage and remaining questions

- Root shared status/errors/casing/provenance/diff contracts: enum membership is
  not authority; diagnostic classifications omit caller values; transformations
  own their Caser but output admission follows allocation; mutable provenance
  copies share transformation slices, and generic record comparison validates
  uniqueness/count rather than digest/status/content shape. Datasetreview's
  stricter schema, ordered IDs and digest validation remain intentionally separate.
- Country/subdivision/currency maps, options, numeric ambiguity, projections and
  fresh All/History/Provenance slices were traced to genuine generators and tests.
  History stores deduplicated name/date projections, not entity or epoch identity;
  date lists are not independently deduplicated. No source authority was refreshed.
- Language/locale use pinned x/text Raw/BCP47/DeprecatedBase profiles, not an
  invented complete registry-conformance claim. Explicit components avoid likely
  inference; parent removes extensions first, language fallback can return itself.
  Observed ParseISO3("heb"/"yid"/"mol") returns "iw"/"ji"/"mo", all rejected by
  strict Parse; this conversion/round-trip limitation does not satisfy the declared
  canonical identity policy. Locale Tag.Language is a direct consumer through
  Raw base.ISO3; language MarshalText/MarshalJSON/Value preserve the resulting
  deprecated spelling, while strict decoders reject it. That propagation is a
  source trace, not an extra runtime observation. Its runtime repair is outside
  this source pass.
  Canonical locale output is not re-admitted through original input limits. Policy rationale for these differences
  remains unanswered, not replaced by a purported standards requirement.
- Phone snapshots provider objects/results and separates region hint, identity,
  extension, display, type, possibility and validity. Postal stores non-control
  UTF-8 (not every rune must be printable), applies NFC/spaces/case in order and
  rechecks output bounds. Optional unbounded syntax preserves pinned country-specific
  terminal-line-feed cases; callers own transport admission. Postal persistence
  cannot round-trip an historic country context through its current-only decoder.
- JSON/SQL absence and failed-decode receiver preservation, strict/current versus
  explicit historic variants and borrowed/copy boundaries were traced in every
  encoding file. SQL bytes are copied before adapter admission. Decode receivers
  require exclusive caller ownership; immutable snapshots do not make shared
  mutation methods safe. Wire delegates target failure semantics to the codec;
  Validation borrows context but does not poll cancellation, retaining option values.
- Postgres mutates only caller-owned maps and must register before concurrency;
  retained facades forward behavior while independently replaceable Wire sentinels
  preserve package-local ownership. Actual published Config/Validation/Wire major
  imports are frozen, not replaced by sibling working directories.
- All test sources were inventoried for actual assertions, helpers, hostile seeds,
  conditional paths and examples: root persistence/status/privacy/provenance/diff/
  casing/Config/concurrency tests; family boundary/history/zero/copy/differential
  tests; adapter coexistence/codec/error/target tests; source-labelled fixture
  helpers/vectors; internal generator/adversarial/cancellation/snapshot tests.
  Seven fuzz functions now distinguish panic resilience from the successor parity
  oracle, explicit payload skips and missing semantic assertions. Benchmark notes
  distinguish combined parse/canonical work and excluded All setup/corpus-size
  effects. All seeds, assertions, timed bodies and five benchmark functions remain
  unchanged; no new campaign was executed. Six existing Examples are the selected live docs proof.
- Generated output ownership: country joins CLDR region/mappings, currency joins
  SIX independently dated current/history lists, subdivision joins CLDR validity/
  English names. The CLI delegates interrupt context; six bounded authenticated
  payloads precede parsing, three relative output flags choose writes in fixed
  order. Writes are not atomic and CPU transformation is not cooperatively polled.
  Native generated-check separately compares compiled semantic snapshot. External
  Brick corpus extraction remains a real manual regeneration-availability question;
  its exact producer attribution, source pin, license and bytes are preserved.

The selected format/render/API/Examples/generation checks and preserved inputs
are recorded below; secret checks and review are retained in the local handoff. Historical reports and pins
are not rewritten or claimed as freshly executed runtime qualification.

## Bounded native outcomes

Pinned Go1.27.2 session observed all three old package-directory generation routes
fail because internal/generate is not main, before editing their targets. The
corrected actual go generate package commands ran in an owned contained source
copy and regenerated country/currency/subdivision byte-identically. Native
scripts/check-generated.sh authenticated the six original checksum-pinned inputs
and passed all generated-output and semantic snapshot comparisons. Existing
internal/generate TestRun checks passed for invocation, cancellation and failure
boundaries; no new producer, data version or source acquisition policy was added.

Disposable parser/scanner proof preserved all105 comment-free ASTs/signatures,
literals and raw significant tokens exactly; only the three approved generator
routes differed. Parsed comment groups rendered, all21 native package
entrypoints and package-owned docs passed; six existing Examples passed. Native
apidiff passed with all signature baselines untouched. Offline specification
validation retained five decisions in one module. Source checker reports zero
errors and the one real external corpus regeneration-availability review.

Small non-network observations independently confirmed language-only fallback
returning itself with true, postal format-rune acceptance despite non-printability,
and historic postal context failing default text decode. These document existing
limitations separately from unchanged runtime/conformance proof. ParseISO3 alias
results are noted above; normative policy and append-only history were preserved.
Final editorial and canonical package clarifications passed affected11-file
AST/signature/literal/raw-token/directive identity, comment rendering and five
refreshed native entrypoints (root, datasetreview, internationaltest, language and
phone). The other16 native renders, Examples, authenticated generation/Run tests
and API proof retain their original session identities, not repeated execution.

Protected API/data/specification inputs, generated corpus/tables, dependencies,
catalog/tool/CI configuration, native scripts/licenses and tracked historical
verification evidence remain exact against the base (78 files). Existing runtime/race/fuzz/
conformance/performance/service/consumer campaigns were not rerun for comments.
All three exact caches, copied source, probe, helpers, logs and rendered outputs
were removed and absence checked. A read-only module-cache permission prevented
the first removal; making only the exact owned directory tree writable resolved
it. A final disposable render-capture filename error also removed its session
immediately; corrected capture and all five affected renders passed, with exact
cleanup verified. Coordinator owns independent generation review and delivery.


## Review-driven seam clarification

The first independent complete-final-diff review requested changes against diff
`f5649db81f7cc1e88aeec5da73cd693e8c983eed022f71372e4e16b26c0f42ac`:
locale Language inherited the provider-conversion limitation, and language text/
SQL encoder comments still promised canonical output while emitting stored text.
The correction documents exact language/locale entrypoint and text/JSON/SQL
encoding-to-strict-decoding seams, without altering code, canonical identity
policy, frozen history/pins or genuine generation. The original review remains
requested-changes evidence, not rewritten as approval. Affected source identity/
comment and native language/locale rendering are selected; unchanged Examples,
API, generation and other package renders remain their original evidence.


The review-driven refresh passed all105 AST/signature/literal/raw-token checks
without semicolon normalization and preserved every directive except the three
approved targets. All491 comment groups parsed/rendered; final native language
and locale docs exposed the corrected Code, encoder and extraction caveats.
The owned refresh and all three caches/helpers/rendered outputs were removed
with absence verified. Fresh whole-diff review is required before commit.

## Objective Enforcement

The checksum-pinned public Tools v2.4.1 verifier now supplies
`make source-docs-check` as a prerequisite of the existing local and CI Make
routes. Ordinary shared CI selects strict source-documentation checking and
requires its result. Native prerequisites, operation budgets, security and
release controls are retained. Hosted exact-head evidence is recorded by the
coordination ledger; configuration alone does not establish green delivery.
