// Package subdivision provides ISO 3166-2-style identifiers backed by pinned
// CLDR validity and English display metadata.
// Parse defaults to current records; deleted records require an explicit option.
// Identifiers, suffixes and English display metadata are distinct values. Country
// resolution uses the separately pinned country table and can return absence.
package subdivision

//go:generate go run ../cmd/international-generate -subdivision-output data_generated.go
