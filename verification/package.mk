.PHONY: generated-check dataset-snapshot dataset-diff provenance docs conformance

conformance:
	go test ./country ./currency ./language ./locale ./phone ./postal ./subdivision

generated-check:
	./scripts/check-generated.sh

dataset-snapshot:
	go run ./cmd/international-dataset-review -snapshot data/dataset-snapshot.json

dataset-diff:
	@test -n "$(BEFORE)" -a -n "$(AFTER)" || \
		{ echo 'usage: make dataset-diff BEFORE=old.json AFTER=new.json' >&2; exit 1; }
	go run ./cmd/international-dataset-review -before "$(BEFORE)" -after "$(AFTER)"

provenance:
	./scripts/check-provenance.sh

docs:
	./scripts/check-docs.sh
