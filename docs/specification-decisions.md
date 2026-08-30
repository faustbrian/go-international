# Specification decisions

This register records observable choices where a standard, registry, or
maintained compatibility profile does not by itself select one application
behavior. The machine bindings, authority pins, update monitoring, and
append-only history live under [`specification/`](../specification/README.md).

## INTERNATIONAL-DEC-001: Current and historic registry acceptance

Status `resolved`; owner `international maintainers`; classification `optional behavior`;
decision scope `application-policy`; specification `Unicode CLDR region validity data`;
version `CLDR 48.2`; source authority `cldr-region-48.2`; authority URL
https://raw.githubusercontent.com/unicode-org/cldr/release-48-2/common/validity/region.xml;
section `common/validity/region.xml and common/supplemental/supplementalData.xml`;
requirement strength `informative`.

Additional authoritative source: `{"id":"cldr-subdivision-48.2","version":"CLDR 48.2","url":"https://raw.githubusercontent.com/unicode-org/cldr/release-48-2/common/validity/subdivision.xml","specifications":["Unicode CLDR subdivision validity data"]}`

Additional authoritative source: `{"id":"cldr-region-mappings-48.2","version":"CLDR 48.2","url":"https://raw.githubusercontent.com/unicode-org/cldr/release-48-2/common/supplemental/supplementalData.xml","specifications":["Unicode CLDR region validity data"]}`

Additional authoritative source: `{"id":"cldr-subdivision-names-48.2","version":"CLDR 48.2","url":"https://raw.githubusercontent.com/unicode-org/cldr/release-48-2/common/subdivisions/en.xml","specifications":["Unicode CLDR subdivision validity data"]}`

Additional authoritative source: `{"id":"iso4217-list-one-2026-01-01","version":"2026-01-01","url":"https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml","specifications":["ISO 4217 currency lists"]}`

Additional authoritative source: `{"id":"iso4217-list-three-2026-01-01","version":"2026-01-01","url":"https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-three.xml","specifications":["ISO 4217 currency lists"]}`

| Field | Decision |
|---|---|
| Issue | Registry datasets contain current, reserved, deleted, transitional, user-assigned, and historic identifiers without defining one application acceptance default. |
| Credible interpretations | Accept every registry record by default. Or: Accept current identifiers by default and require explicit options for non-current identifiers. |
| Known peer behavior | Pinned x/text comparisons agree on current country and currency mappings but do not define this package's historic persistence policy. |
| Selected behavior | Strict parsers accept current identifiers by default; options-bearing parse and persistence APIs opt into named non-current statuses, and ambiguous reused numeric identifiers are rejected. |
| Rationale | Explicit status policy prevents a dataset update from silently reinterpreting persisted identifiers while retaining deliberate migration access to historic records. |
| Security consequences | Unknown identifiers and ambiguous numeric reuse fail closed instead of being assigned an unintended authority record. |
| Resource consequences | Lookups use immutable bounded generated maps and perform no network access. |
| Compatibility consequences | Additions are normally compatible; removals, status transitions, and numeric reassignments require dataset-diff and release review. |
| Wire consequences | Text, JSON, SQL, configuration, validation, pgx, and wire decoders share the same default-current and explicit-historic policy. |
| Executable evidence | `TestHistoricalAndUserAssignedCodesRequireOptIn`; `TestDeletedSubdivisionRequiresOptIn`; `TestHistoricCurrenciesRequireOptInAndPreserveWithdrawalText`; `TestReusedNumericCountryCodesPreserveTheirMappedIdentity` |
| Fixture and fuzz evidence | `internationaltest/vectors.go`; `FuzzTextParsers` |
| Maintained-peer evidence | `country/differential_test.go`; `currency/differential_test.go` |
| Public APIs | `country.ParseWithOptions`; `subdivision.ParseWithOptions`; `currency.ParseWithOptions`; `Status` |
| Documentation | `docs/specification-decisions.md`; `docs/reference.md`; `docs/provenance.md` |
| Upstream status | CLDR and SIX source snapshots plus their release surfaces are monitored independently; no upstream issue is open for the package-owned acceptance default. |
| Reconsider when | A source registry changes status semantics, numeric reuse rules, or supplies a normative application acceptance profile. |

## INTERNATIONAL-DEC-002: BCP 47 identity canonicalization and fallback

Status `resolved`; owner `international maintainers`; classification
`interoperability policy`; decision scope `application-policy`; specification
`BCP 47 language tags`; version `RFC 5646`; source authority `bcp47-rfc5646`;
authority URL https://www.rfc-editor.org/rfc/rfc5646.txt; section
`Sections 2.1.1, 2.2.9, and 4.5`; requirement strength `SHOULD`.

Additional authoritative source: `{"id":"iana-language-2026-06-14","version":"Registry 2026-06-14 via x/text v0.40.0","url":"https://raw.githubusercontent.com/golang/text/v0.40.0/internal/language/tables.go","specifications":["IANA Language Subtag Registry snapshot"]}`

| Field | Decision |
|---|---|
| Issue | A well-formed BCP 47 tag can preserve its caller spelling or be canonicalized, while application fallback is not the same operation as language-range matching. |
| Credible interpretations | Canonicalize every accepted tag during parsing and infer a fallback chain. Or: Preserve valid caller spelling and expose canonicalization and fallback as separate explicit operations. |
| Known peer behavior | Pinned golang.org/x/text v0.40.0 agrees on parsed structure and canonical results for the governed vectors; the package adds explicit source preservation and fallback policy. |
| Selected behavior | Parse validates and preserves the accepted source string; Canonical returns the pinned registry-aware canonical tag, and Fallback requires an explicit none, parent, or language policy. |
| Rationale | Separating validation, canonicalization, and fallback keeps persisted identity stable and prevents a lossy application choice from occurring implicitly. |
| Security consequences | Malformed, invalid UTF-8, overlong, and over-segmented tags fail before retention; no locale or resource is resolved from a tag. |
| Resource consequences | Tags are limited to 255 bytes and 32 segments and parsing performs no network access. |
| Compatibility consequences | Canonical results can change only through a reviewed registry or x/text update; preserved source strings remain stable. |
| Wire consequences | Text, JSON, SQL, and integration encoders retain the accepted spelling unless the caller explicitly serializes Canonical output. |
| Executable evidence | `TestParsePreservesValidSpellingAndCanonicalizationIsExplicit`; `TestLocalePartsRemainAvailableWithoutLossyFallback`; `TestFallbackRequiresAnExplicitPolicy`; `TestCanonicalizationDifferentialAgainstPinnedXText` |
| Fixture and fuzz evidence | `internationaltest/vectors.go`; `FuzzTextParsers` |
| Maintained-peer evidence | `locale/differential_test.go` |
| Public APIs | `locale.Parse`; `locale.Tag.Canonical`; `locale.Tag.Fallback` |
| Documentation | `docs/specification-decisions.md`; `docs/reference.md`; `docs/provenance.md` |
| Upstream status | RFC 5646 errata and the current IANA registry are monitored separately from the pinned x/text-derived snapshot. |
| Reconsider when | BCP 47, its errata, or the IANA registry changes canonicalization requirements, or the package adopts a distinct matching profile. |

## INTERNATIONAL-DEC-003: Phone identity and metadata classification

Status `resolved`; owner `international maintainers`; classification
`implementation-defined behavior`; decision scope `application-policy`;
specification `libphonenumber metadata profile`; version
`v9.0.32 via nyaruka/phonenumbers v1.8.1`; source authority
`libphonenumber-metadata`; authority URL
https://raw.githubusercontent.com/nyaruka/phonenumbers/v1.8.1/data/metadata.xml.gz;
section `data/metadata.xml.gz numbering-plan metadata`; requirement strength
`informative`.

| Field | Decision |
|---|---|
| Issue | Numbering metadata distinguishes parseability, possibility, validity, type, display formatting, and canonical E.164 identity, but applications may incorrectly collapse them into one result. |
| Credible interpretations | Treat every parseable or possible number as valid and use formatted text as identity. Or: Keep canonical E.164 identity separate from extension, possibility, validity, type, and display formatting. |
| Known peer behavior | The wrapper is differentially checked against pinned nyaruka/phonenumbers behavior derived from libphonenumber v9.0.32 for public example ranges. |
| Selected behavior | International parsing produces canonical E.164 identity with a separate extension; national parsing requires an explicit region, and Possible, Valid, Type, and display formats remain distinct metadata results. |
| Rationale | Separate results prevent routing or authorization from treating a formatting choice or a merely possible number as a validated identity. |
| Security consequences | Inputs and extensions are bounded, diagnostics and default formatting are redacted, and metadata parsing performs no network access. |
| Resource consequences | Input is limited to 128 bytes, extensions to 20 bytes, and all metadata is immutable in-process data. |
| Compatibility consequences | A metadata update can change possibility, validity, type, or formatting and therefore requires differential review and release notes. |
| Wire consequences | Persistence encodes canonical E.164 identity and extension separately; display formats never replace identity. |
| Executable evidence | `TestParseInternationalNumberSeparatesCanonicalAndDisplayForms`; `TestNationalParsingRequiresExplicitRegionHint`; `TestPossibleAndValidAreDistinctMetadataDecisions`; `TestDifferentialAgainstPinnedLibphonenumber` |
| Fixture and fuzz evidence | `internationaltest/vectors.go`; `FuzzPhoneAndPostalBoundedParsing` |
| Maintained-peer evidence | `phone/differential_test.go` |
| Public APIs | `phone.Parse`; `phone.ParseE164`; `phone.Number.Possible`; `phone.Number.Valid`; `phone.Number.Type`; `phone.Number.Format` |
| Documentation | `docs/specification-decisions.md`; `docs/reference.md`; `docs/provenance.md` |
| Upstream status | The pinned metadata payload and current libphonenumber release feed are monitored; no local divergence is known for the governed vectors. |
| Reconsider when | The pinned dependency changes its metadata contract or a replacement numbering-plan authority and maintained implementation are adopted. |

## INTERNATIONAL-DEC-004: Opaque postal identity and optional syntax profile

Status `resolved`; owner `international maintainers`; classification `optional behavior`;
decision scope `application-policy`; specification `Brick postcode compatibility profile`;
version `0.5.0 at ead386982c31d825843e80ab86a1919eca1a1ad5`; source authority
`brick-postcode-0.5.0`; authority URL
`https://raw.githubusercontent.com/brick/postcode/ead386982c31d825843e80ab86a1919eca1a1ad5/composer.json`;
section `Formatter implementations and formatter test corpus`; requirement
strength `informative`.

| Field | Decision |
|---|---|
| Issue | Country-specific syntax rules can provide compatibility screening but cannot establish postal-code existence, deliverability, locality, or address correctness. |
| Credible interpretations | Apply country syntax during every postal parse and imply address validity. Or: Keep postal values opaque and expose pinned syntax matching as an explicit optional operation. |
| Known peer behavior | ValidSyntax is checked against the pinned Brick postcode 0.5.0 formatter corpus while ordinary Parse deliberately accepts bounded printable values outside that syntax profile. |
| Selected behavior | Parse stores a bounded printable UTF-8 value with explicit country context without a syntax claim; ValidSyntax separately applies the pinned Brick-compatible country rule after ASCII space, dash, and case handling. |
| Rationale | The split avoids turning a maintained peer's formatter rules into an unsupported deliverability or address-validation claim. |
| Security consequences | Postal values are redacted from default formatting, invalid UTF-8 and control characters fail, and callers must bound transport input before syntax validation. |
| Resource consequences | Stored values are limited to 32 bytes and syntax checks use fixed precompiled country rules without network access. |
| Compatibility consequences | Ordinary parsing remains stable when the optional syntax corpus changes; a syntax-boundary change requires corpus, compatibility, and changelog review. |
| Wire consequences | Encoded identity preserves the country and raw value; syntax validity is not serialized as an authoritative fact. |
| Executable evidence | `TestParsePreservesCallerValueAndCountryWithoutSyntaxClaims`; `TestValidSyntaxMatchesLegacyPostcodeContract`; `TestValidSyntaxMatchesPinnedBrickPostcodeCorpus` |
| Fixture and fuzz evidence | `postal/syntax_compatibility_test.go`; `FuzzPhoneAndPostalBoundedParsing` |
| Maintained-peer evidence | `postal/syntax_compatibility_test.go` |
| Public APIs | `postal.Parse`; `postal.Code.Normalize`; `postal.ValidSyntax`; `postal.SyntaxDataset` |
| Documentation | `docs/specification-decisions.md`; `docs/reference.md`; `docs/provenance.md` |
| Upstream status | The exact Brick commit and current tags feed are monitored; the package makes no claim that the peer profile is an official postal authority. |
| Reconsider when | The compatibility profile changes materially or an authoritative, licensed, maintainable syntax source is adopted. |

## INTERNATIONAL-DEC-005: Canonical language identifier representation

Status `resolved`; owner `international maintainers`; classification
`interoperability policy`; decision scope `application-policy`; specification
`IANA Language Subtag Registry snapshot`; version
`Registry 2026-06-14 via x/text v0.40.0`; source authority
`iana-language-2026-06-14`; authority URL
https://raw.githubusercontent.com/golang/text/v0.40.0/internal/language/tables.go;
section `Language subtag records and Preferred-Value mappings`; requirement
strength `informative`.

| Field | Decision |
|---|---|
| Issue | The registry contains two-letter and three-letter language identifiers, deprecated aliases, and preferred values, so one language can have multiple historical representations. |
| Credible interpretations | Accept every registered alias and preserve the supplied representation. Or: Expose one canonical lowercase ISO 639 identity and provide an explicit ISO 639-3 conversion boundary. |
| Known peer behavior | The pinned x/text-derived registry tables supply the maintained two-letter, three-letter, and preferred-value mappings used by the generated dataset. |
| Selected behavior | Parse accepts the canonical lowercase ISO 639-1 code when one exists or the canonical three-letter-only identifier; ParseISO3 converts a registered ISO 639-3 identifier to that identity, and deprecated aliases are rejected. |
| Rationale | One canonical identity prevents aliases from comparing or persisting as different languages while retaining explicit three-letter interoperability. |
| Security consequences | Only bounded lowercase ASCII identifiers with an exact registry mapping are accepted; Unicode lookalikes and unknown aliases fail. |
| Resource consequences | Parsing uses immutable generated maps and performs no network access. |
| Compatibility consequences | A preferred-value or ISO mapping change requires generated-data, persistence, and release review. |
| Wire consequences | Text, JSON, SQL, configuration, validation, pgx, and wire encoders emit the canonical package identity rather than a deprecated alias. |
| Executable evidence | `TestParseCanonicalISO639LanguageAndConvertISO3`; `TestParseISO3UsesAuthoritativeMapping`; `TestCanonicalThreeLetterOnlyLanguageRoundTrips`; `TestLanguageParsingIsStrictAndBounded` |
| Fixture and fuzz evidence | `internationaltest/vectors.go`; `FuzzTextParsers` |
| Public APIs | `language.Parse`; `language.ParseISO3`; `language.Code.ISO3` |
| Documentation | `docs/specification-decisions.md`; `docs/reference.md`; `docs/provenance.md` |
| Upstream status | The exact x/text-derived snapshot and the current IANA registry are monitored; no local alias extension is defined. |
| Reconsider when | The IANA registry or pinned x/text tables change preferred-value or ISO 639 mapping semantics. |
