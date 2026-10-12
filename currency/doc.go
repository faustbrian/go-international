// Package currency provides ISO 4217 identity and versioned metadata. It does
// not implement money arithmetic or localized amount formatting.
// Parse defaults to active records; history requires an explicit option. Numeric
// identifiers retain their alphabetic mapping and reject policy-induced ambiguity.
// An unavailable minor-unit exponent differs from a specified zero exponent.
// Historic date text is a projection, not an epoch-selection API.
package currency

//go:generate go run ../cmd/international-generate -currency-output data_generated.go
