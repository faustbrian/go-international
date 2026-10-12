// Package country provides strict ISO 3166-1 identifiers and CLDR-backed
// display metadata. Names are display values and never stable identifiers.
// Parse defaults to official records; options admit other stored statuses without
// silently mapping to successors. Canonicalize only repairs ASCII letter casing.
// All and DatasetRecords return independent slices; names use pinned x/text display
// fallback rather than defining application presentation or locale policy.
package country

//go:generate go run ../cmd/international-generate -country-output data_generated.go
