# Security policy

Report vulnerabilities through the
[private GitHub security-advisory workflow](https://github.com/faustbrian/go-international/security/advisories/new).
Do not include real phone numbers, postal codes, credentials, or customer data.

Severity classification, acknowledgement and remediation targets, embargo
handling, advisories, and coordinated releases follow the shared
[ecosystem vulnerability-management policy](https://github.com/faustbrian/go-library-tools/blob/main/docs/ecosystem/security/vulnerability-management.md).

The [versioned threat model](docs/threat-model.md) documents repository-specific
trust boundaries, control ownership, and maintained contract limitations.

Parsers bound bytes, locale segments, Unicode normalization, extensions,
metadata, and diagnostics. Contracts reject invalid UTF-8 rather than repairing
it. Parse diagnostics retain only package-defined reason classifications;
unrecognized caller-provided reasons and fixture classifications are replaced
rather than echoed. Generated inputs are HTTPS-fetched, size-limited,
checksum-pinned, deterministically transformed, license-reviewed, and
drift-checked. Malformed dataset values and transport or response-body failure
details are omitted from generator diagnostics, as are output-write failure
details. The generation command passes interrupt cancellation through every
dataset request in addition to its fixed HTTP timeout and checks cancellation
before each generated output write.

Threats considered include Unicode confusables and normalization mismatch,
numeric ambiguity, regex denial of service, metadata poisoning, dependency
compromise, and sensitive-value leakage. Identifiers are not identity proof;
phone validity is not ownership, and postal syntax is not deliverability.

Exact byte, segment, diagnostic, dataset, and generator budgets and their
hostile-input evidence are maintained in `docs/verification.md`. Unsupported
display locales return no metadata rather than panicking or inferring a locale.

Applications should redact phone and postal values from logs and telemetry,
limit request bodies before decoding, use explicit locale policies, review data
diffs, run vulnerability scanning, and promptly update pinned metadata after an
upstream security advisory.
