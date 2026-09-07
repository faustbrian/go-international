# Five-minute quickstarts

## Adapters

```go
import (
    internationalpostgres "github.com/faustbrian/go-international/adapters/postgres"
    internationalvalidation "github.com/faustbrian/go-international/adapters/validation"
    internationalwire "github.com/faustbrian/go-international/adapters/wire"
    "github.com/faustbrian/go-international/country"
    validation "github.com/faustbrian/go-validation"
    "github.com/faustbrian/go-wire"
    "github.com/jackc/pgx/v5/pgtype"
)

fi, err := country.Parse("FI")
if err != nil {
    return err
}

typeMap := pgtype.NewMap()
internationalpostgres.Register(typeMap) // before concurrent map use

ctx, err := validation.NewContext(validation.DefaultLimits())
if err != nil {
    return err
}
if err := internationalvalidation.Country().Validate(ctx, "FI").Err(); err != nil {
    return err
}

payload, err := internationalwire.Encode(wire.FormatJSON, fi)
```

The [executable adapter example](../adapter_successor_test.go) is compiled and
run by the root test suite. These operations are synchronous and retain no
caller-owned values or contexts.

## Country

```go
fi, err := country.Parse("FI")               // strict alpha-2
alpha3, err := country.ParseAlpha3("FIN")    // distinct identifier type
fi, _ = alpha3.Alpha2()                      // authoritative conversion
name := country.Name(fi, textlanguage.English) // display metadata only
```

Use `Canonicalize("fi")` only when accepting case variants intentionally.
Historic, reserved, and user-assigned records require `ParseWithOptions`.

## Phone

```go
fi, _ := country.Parse("FI")
number, err := phone.Parse("040 123 4567", phone.ParseOptions{RegionHint: fi})
if err != nil { return err }
canonical := number.E164()
possible, valid := number.Possible(), number.Valid()
display, _ := number.Format(phone.FormatInternational)
```

`valid` is numbering-plan metadata, not proof of ownership or reachability.
Default string formatting is redacted. Persist using JSON, SQL, or
`MarshalText`; extensions are retained separately.

## Locale

```go
tag, err := locale.Parse("zh-hant-tw-u-ca-gregory")
canonical, err := tag.Canonical()
parent, ok := canonical.Fallback(locale.FallbackParent)
```

Parsing preserves accepted caller spelling. Canonicalization and fallback are
separate operations; no automatic truncation or environment detection occurs.

## Currency

```go
eur, err := currency.Parse("EUR")
numeric, _ := eur.Numeric()
minorUnits, specified := eur.MinorUnits()
```

Currency identity is metadata, not money arithmetic. Historic codes require
`ParseWithOptions`; metadata never changes persisted monetary values.

## Postal

```go
fi, _ := country.Parse("FI")
code, err := postal.Parse(" 00100 ", fi)
normalized, err := code.Normalize(postal.NormalizeOptions{
    Spaces: postal.SpacesCollapseASCII,
    Case: postal.CaseUpperASCII,
    Unicode: postal.UnicodeNFC,
})
```

The country context is mandatory and preserved. Parsing and normalization make
no claim about syntax, locality, deliverability, or carrier acceptance.
