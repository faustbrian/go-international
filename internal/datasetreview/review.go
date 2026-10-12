// Package datasetreview creates bounded semantic snapshots of generated data.
package datasetreview

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	international "github.com/faustbrian/go-international/v4"
	"github.com/faustbrian/go-international/v4/country"
	"github.com/faustbrian/go-international/v4/currency"
	"github.com/faustbrian/go-international/v4/subdivision"
)

const (
	// SchemaVersion identifies the deterministic snapshot representation.
	SchemaVersion = 1
	// MaxSnapshotBytes bounds review artifact decoding and diagnostics.
	MaxSnapshotBytes = 2 << 20
)

// Snapshot is the compatibility-relevant projection of every generated table.
type Snapshot struct {
	// Schema selects the supported snapshot representation.
	Schema int `json:"schema"`
	// Country contains sorted country record projections.
	Country []international.Record `json:"country"`
	// Subdivision contains sorted subdivision record projections.
	Subdivision []international.Record `json:"subdivision"`
	// Currency contains sorted currency record projections.
	Currency []international.Record `json:"currency"`
}

// Report classifies semantic changes independently for each generated dataset.
type Report struct {
	// Country classifies changes between the country projections.
	Country international.DatasetDiff `json:"country"`
	// Subdivision classifies changes between the subdivision projections.
	Subdivision international.DatasetDiff `json:"subdivision"`
	// Currency classifies changes between the currency projections.
	Currency international.DatasetDiff `json:"currency"`
}

// Current returns a fresh snapshot of the compiled generated tables.
func Current() Snapshot {
	return Snapshot{
		Schema:      SchemaVersion,
		Country:     country.DatasetRecords(),
		Subdivision: subdivision.DatasetRecords(),
		Currency:    currency.DatasetRecords(),
	}
}

// Encode validates before writing deterministic indented JSON to a borrowed
// writer. It does not close the writer or roll back bytes after a write error.
func Encode(writer io.Writer, snapshot Snapshot) error {
	if writer == nil {
		return fmt.Errorf("%w: snapshot writer is required", international.ErrInvalidDataset)
	}
	if err := validate(snapshot); err != nil {
		return err
	}
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(snapshot); err != nil {
		return fmt.Errorf("encode dataset snapshot: %w", err)
	}
	return nil
}

// Decode reads one bounded document from a borrowed reader without closing it.
// Unknown fields, trailing data, unsupported schemas and invalid record ordering
// or fingerprints fail; successful slices belong to the decoded snapshot.
func Decode(reader io.Reader) (Snapshot, error) {
	if reader == nil {
		return Snapshot{}, fmt.Errorf("%w: snapshot reader is required", international.ErrInvalidDataset)
	}
	payload, err := io.ReadAll(io.LimitReader(reader, MaxSnapshotBytes+1))
	if err != nil {
		return Snapshot{}, fmt.Errorf("read dataset snapshot: %w", err)
	}
	if len(payload) > MaxSnapshotBytes {
		return Snapshot{}, international.ErrResourceLimit
	}
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var snapshot Snapshot
	if err := decoder.Decode(&snapshot); err != nil {
		return Snapshot{}, fmt.Errorf("%w: decode snapshot", international.ErrInvalidDataset)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return Snapshot{}, fmt.Errorf("%w: trailing snapshot data", international.ErrInvalidDataset)
	}
	if err := validate(snapshot); err != nil {
		return Snapshot{}, err
	}
	return snapshot, nil
}

// Diff validates and classifies changes between two semantic snapshots.
func Diff(before, after Snapshot) (Report, error) {
	if err := validate(before); err != nil {
		return Report{}, err
	}
	if err := validate(after); err != nil {
		return Report{}, err
	}
	countryDiff, err := international.DiffRecords(before.Country, after.Country)
	subdivisionDiff, subdivisionErr := international.DiffRecords(before.Subdivision, after.Subdivision)
	currencyDiff, currencyErr := international.DiffRecords(before.Currency, after.Currency)
	return Report{Country: countryDiff, Subdivision: subdivisionDiff, Currency: currencyDiff},
		errors.Join(err, subdivisionErr, currencyErr)
}

func validate(snapshot Snapshot) error {
	if snapshot.Schema != SchemaVersion {
		return fmt.Errorf("%w: unsupported snapshot schema", international.ErrInvalidDataset)
	}
	for _, dataset := range []struct {
		name    string
		records []international.Record
	}{
		{name: "country", records: snapshot.Country},
		{name: "subdivision", records: snapshot.Subdivision},
		{name: "currency", records: snapshot.Currency},
	} {
		if err := validateRecords(dataset.records); err != nil {
			return fmt.Errorf("%w: invalid %s records", err, dataset.name)
		}
	}
	return nil
}

func validateRecords(records []international.Record) error {
	if len(records) == 0 {
		return international.ErrInvalidDataset
	}
	if len(records) > international.MaxDatasetRecords {
		return international.ErrResourceLimit
	}
	previous := ""
	for _, record := range records {
		if record.ID == "" {
			return international.ErrInvalidDataset
		}
		if record.ID <= previous {
			return international.ErrInvalidDataset
		}
		if record.Status > international.StatusHistoric {
			return international.ErrInvalidDataset
		}
		if len(record.Fingerprint) != hex.EncodedLen(sha256Bytes) {
			return international.ErrInvalidDataset
		}
		if _, err := hex.DecodeString(record.Fingerprint); err != nil {
			return international.ErrInvalidDataset
		}
		previous = record.ID
	}
	return nil
}

const sha256Bytes = 32
