# Specification conformance matrix

The [specification decision register](../docs/specification-decisions.md) owns
the interpretations behind these bindings. Source versions, integrity pins,
and update authorities are in `sources.tsv` and `monitoring.json`.

| Decision | Sources | Behavioral evidence | Maintained-peer evidence |
|---|---|---|---|
| INTERNATIONAL-DEC-001 | CLDR 48.2 region/subdivision data; SIX List One 2026-09-17 and List Three 2026-01-01 | Current versus historic status and numeric-reuse tests; governed vectors; parser fuzzing | Pinned x/text country and currency comparisons |
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

## Upstream review history

### 2026-09-09

- The Unicode CLDR GitHub feed added the unrelated
  `production/2026-09-08-1807z` snapshot. The feed mixes stable releases,
  prereleases, and production snapshots, so exact-byte monitoring produced a
  repository-wide failure without identifying a change to the selected CLDR
  48.2 inputs.
- Replace that noisy feed with Unicode's versioned CLDR 48 stable-distribution
  record. The four exact CLDR 48.2 source payloads remain pinned and monitored,
  and adopting a later stable CLDR release remains an explicit dataset review
  rather than a gate on unrelated pull requests.

### 2026-09-03

- The Unicode CLDR release feed added the `release-49-alpha2` prerelease,
  published on 2026-09-03. Three consecutive fetches were byte-identical at
  17,063 bytes with SHA-256
  `eceecf01dca1d5000ec42abad2c45e614033e3b5b1f4d0513a355bb19384095a`.
- The four selected CLDR 48.2 region, mapping, subdivision, and English-name
  payloads remain byte-for-byte identical to their pinned checksums. The
  prerelease therefore does not change the accepted dataset, decision, or
  conformance bindings.

- SIX ISO 4217 List One is reviewed at 47,491 bytes with SHA-256
  `33139b438657d1cee116ba737807ea71d19d6de4b90f799a09c56f0cc6a1b0ff`,
  and List Three remains 30,638 bytes with SHA-256
  `98fde2423cdb916dd59dcf5fe96222edad8fa198d865c1c83dbc464b9cc52387`.
  List One declares 2026-09-17; List Three independently declares 2026-01-01.
  The refreshed current list retains all 178 selected current records without
  changing codes, numeric mappings, names or minor units. Generation requires
  both publication dates but does not require synchronized publication.
  The currency acceptance decision and semantic conformance vectors are unchanged.
- The retired `iso4217-releases` monitor targeted a redirected SIX landing
  page whose request-specific chrome is not a deterministic change signal.
  Current monitoring binds separate change-authority entries to the exact
  List One and List Three XML payloads, so either authoritative data change
  requires review without pinning arbitrary landing-page bytes.
