# Parsing, validation, and integrations

Parse at untrusted boundaries, store the typed value, validate application
policy separately, and format only at presentation boundaries. A parse error
means malformed or unaccepted identity. `Possible` and `Valid` are separate
phone metadata decisions. Postal parsing is only bounded storage.

JSON and SQL use the scalar contracts directly. Call
`internationalpostgres.Register(conn.TypeMap())` from `adapters/postgres`
before concurrent connection use. `config` discovers `UnmarshalText`
automatically. `adapters/validation` provides string rules for every primitive;
`Phone` checks parseability while `ValidPhone` also checks current metadata.

Default `UnmarshalText`, `UnmarshalJSON`, and `Scan` methods accept current
identifiers only. Applications that deliberately persist historic country,
subdivision, or currency values must call the corresponding `WithOptions`
method at the storage boundary. This keeps obsolete-code acceptance explicit
instead of turning every configuration or database decode into an alias policy.

`adapters/wire` supports strict bounded JSON, XML, YAML, TOML, and MessagePack
dispatch. SOAP, CBOR, and BSON return `ErrUnsupportedFormat`
because their current generic adapters cannot guarantee lossless immutable
scalar round trips.

The released `internationalpgx`, `internationalvalidation`, and
`internationalwire` imports remain behavior-preserving compatibility paths for
the documented deprecation interval. Their successor calls have identical
signatures. The wire sentinels share one initial error identity, while each
package continues to honor replacement of its own exported variable only.

`money` may consume `currency.Code`; this module never imports money.
Coordinates and spatial algorithms remain in `geo`. Optional UI formatting
may use `x/text` directly.

Database schemas should use bounded text columns matching the documented
maximum and a nullable column for absence. Do not index localized country names
as identifiers. Phone and postal values are personal data: avoid logs,
telemetry, traces, and error payloads.
