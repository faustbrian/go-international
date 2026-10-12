# API and standards reference

| Package | Identity/standard | Default acceptance |
|---|---|---|
| `country` | ISO 3166-1 alpha-2/alpha-3/numeric | distinct current types |
| `subdivision` | ISO 3166-2-derived CLDR identifiers | current records |
| `language` | ISO 639 plus IANA BCP 47 data | canonical ISO 639-1 or ISO 639-3-only identifiers |
| `locale` | BCP 47 | bounded registry-aware tags |
| `currency` | ISO 4217 alphabetic/numeric | distinct active types |
| `phone` | E.164 and libphonenumber metadata | bounded parseable numbers |
| `postal` | country-contextual opaque value; optional pinned syntax validation | bounded UTF-8 without controls |
| `adapters/postgres` | pgx integration | eleven International scalar types registered as PostgreSQL text |
| `adapters/validation` | Golib Validation integration | twelve pure synchronous parser-backed string rules |
| `adapters/wire` | Golib Wire integration | bounded JSON, XML, YAML, TOML, and MessagePack dispatch |

Country, subdivision, currency, language and locale `Parse` preserve accepted
identifier spelling. Phone parsing instead snapshots provider-normalized identity
and display; `ParseCallingCode` also stores a numeric value, so its String result
omits any accepted leading zeroes.
Language `ParseISO3` converts the provider-recognized three-letter representation,
but currently can return deprecated spellings such as `iw` for `heb`. Strict Parse
and default persistence decoding reject those spellings: this is an existing
conversion/round-trip limitation, not endorsed canonical behavior. Locale
`Tag.Language` delegates its explicit base through the same conversion and can
expose the same deprecated spelling. Language text, JSON and SQL encoders preserve
that stored spelling rather than repairing it; successful encoding therefore does
not imply acceptance by the strict corresponding decoder.
APIs named `Canonicalize`, `Canonical`, or `Normalize` perform explicit
transformations.
Display names and phone formatting are presentation metadata, never identity.

`postal.ValidSyntax` is an explicit compatibility operation backed by
`postal.SyntaxDataset`. It removes ASCII spaces and dashes and canonicalizes
ASCII case before applying the pinned country rule. It does not change
`postal.Parse`: ordinary postal values remain opaque, and syntax validity does
not establish existence, deliverability, locality, or address correctness.
Callers must apply their transport-specific input bound before syntax
validation.

`international.LowercaseUnicode` performs explicit locale-neutral full
Unicode lowercasing for protocol normalization. The caller supplies the
maximum input and expanded-output byte size. Invalid UTF-8 and either
resource-limit breach fail without returning partially transformed text. Output
size is checked after transformation, not as a peak-allocation limit.
Each invocation owns its case transformer and is safe to call concurrently.

Every scalar implements strict text, JSON, `database/sql.Scanner`, and
`driver.Valuer` contracts. Zero means absent: text marshal fails, JSON emits
`null`, SQL emits `NULL`, and decoding null resets the destination. Failed
decodes leave it unchanged. Nullable database columns may therefore use either
the zero value or an application-owned presence wrapper when absent and
explicit zero must be distinguished.

`Status` distinguishes unknown, official, reserved, transitional, deleted,
user-assigned, and historic records. Non-current acceptance is always opt-in.
Country alpha-2, alpha-3, and numeric parsers each provide an options-bearing
form so an accepted historic representation retains its authoritative status.
Options-bearing text, JSON, and SQL decode methods preserve that opt-in policy
when loading historic country, subdivision, or currency values. Postal Parse can
retain such a country context, but its default persistence decoder uses current-only
country parsing and has no historic options method. Reassigned
numeric identifiers retain their authoritative alphabetic mapping in memory;
an options-bearing numeric parse rejects the value when the selected current
and historic statuses would make that mapping ambiguous.
`DatasetDiff` classifies additions, removals, aliases, status changes, and
metadata changes for update review; a retained key can have several classifications.
Its generic comparison rejects duplicate keys but does not validate individual
record contents. Provenance validation checks descriptive field shape, not
content authentication.

Language-only fallback may return the same tag with true; callers must not use
that boolean alone as a loop-progress condition.

Core deliberately excludes message catalogs, content negotiation, money arithmetic, address validation,
geocoding, tax, sanctions, delivery, and locale detection.

The legacy `internationalpgx`, `internationalvalidation`, and
`internationalwire` packages remain importable compatibility paths. Their
target-oriented successors preserve signatures and behavior; new code should
select the successor imports above.

`internationaltest` exports independently mutable copies of governed country,
subdivision, language, locale, currency, phone, and postal vectors. Each vector family has a
source constant, and the package tests evaluate frozen expected results rather
than deriving expectations from the generated tables. Country mappings are
also compared across the complete official set with `x/text`; stable currency
and locale behavior and all public phone vectors have differential checks.
