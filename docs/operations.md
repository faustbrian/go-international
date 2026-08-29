# Performance, FAQ, and troubleshooting

Lookups use immutable generated maps and require no network. Benchmark with the
shared tooling benchmark gate; protect regressions by comparing `ns/op`,
allocations, and bulk conversion throughput on the same Go version and
hardware. Phone parsing is intentionally heavier than fixed-code lookup.

## Local verification

`make ci` delegates to the pinned shared tool, which reproduces every blocking
release gate: formatting, vet, Staticcheck, strict golangci-lint, tests, exact
coverage, race detection, generated-data drift, provenance and license checks,
reviewed mutation checks, documentation, API compatibility, vulnerability
scanning, workflow linting, fuzz smoke, and benchmarks. NilAway runs separately
as an explicitly advisory signal. The generated-data, provenance, and
documentation checks remain package-owned operations in
`verification/package.mk`; all other gates are owned by the shared tool.

The generated-data operation acquires checksum-pinned authoritative inputs.
Core identifier parsing, validation, lookup, and formatting remain offline and
do not invoke the generator or perform network requests.

**Why was a lowercase code rejected?** Strict `Parse` preserves the boundary.
Use an explicit canonicalization API only when your contract permits it.

**Why is a historic code rejected?** Enable the corresponding parse option and
retain its status. Do not silently alias it.

**Why does national phone parsing fail?** Supply `RegionHint`; the package does
not infer locale or country.

**Why is a possible number invalid?** Its shape is plausible but current
metadata does not recognize it as valid. Neither result proves ownership.

**Why can’t postal parsing confirm a city?** Postal values are bounded opaque
values. Address validation, search, and provider rules belong in Postal.

**When should `postal.ValidSyntax` be used?** Only when a contract explicitly
requires the pinned country-format compatibility result. Bound untrusted input
first and do not interpret `true` as proof that the code exists or is
deliverable.

For generated drift, verify network access, upstream checksum changes, and
license terms before updating a pin. For config or SQL failures, inspect the
typed safe error and confirm the input is a string/byte value, not a number.
