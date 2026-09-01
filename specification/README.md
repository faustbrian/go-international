# Specification conformance matrix

The [specification decision register](../docs/specification-decisions.md) owns
the interpretations behind these bindings. Source versions, integrity pins,
and update authorities are in `sources.tsv` and `monitoring.json`.

| Decision | Sources | Behavioral evidence | Maintained-peer evidence |
|---|---|---|---|
| INTERNATIONAL-DEC-001 | CLDR 48.2 region/subdivision data and SIX ISO 4217 lists dated 2026-01-01 | Current versus historic status and numeric-reuse tests; governed vectors; parser fuzzing | Pinned x/text country and currency comparisons |
| INTERNATIONAL-DEC-002 | RFC 5646 and the IANA snapshot represented by x/text v0.40.0 | Source preservation, canonicalization, fallback, governed vectors, and parser fuzzing | Pinned x/text language-tag comparison |
| INTERNATIONAL-DEC-003 | libphonenumber v9.0.32 metadata through nyaruka/phonenumbers v1.8.1 | Identity, region, possibility, validity, formatting, governed vectors, and bounded-input fuzzing | Pinned libphonenumber wrapper comparison |
| INTERNATIONAL-DEC-004 | Brick postcode 0.5.0 at `ead386982c31d825843e80ab86a1919eca1a1ad5` | Opaque parsing, explicit normalization, optional syntax rules, and bounded-input fuzzing | Complete pinned Brick formatter corpus |
| INTERNATIONAL-DEC-005 | IANA registry snapshot represented by x/text v0.40.0 | Canonical two-letter and three-letter-only identity, ISO 639-3 conversion, strict parsing, governed vectors, and parser fuzzing | Maintained x/text-derived registry tables; differential behavior not separately assessed |

ISO 3166 labels describe the public identifier family, but the executable
acceptance dataset is the pinned CLDR projection. E.164 describes the phone
wire identity, while numbering-plan classification is governed by the pinned
libphonenumber metadata profile. Postal syntax is a maintained-peer
compatibility profile, not an official postal authority or deliverability
claim.
