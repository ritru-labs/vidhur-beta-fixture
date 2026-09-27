# Vidhur supervised beta fixture

Private, isolated acceptance repository for G-001. This seed intentionally
contains a whitespace-normalization bug and a pre-existing failing regression.
It has no third-party Go dependencies or credentials.

Bounded repair: fix `Normalize` in `ids.go` to trim whitespace, drop empty IDs,
deduplicate after trimming, preserve first-seen order and leave input unchanged.
Keep the existing tests, public API and module file unchanged.

Regression: `go test -json -count=1 -run ^TestNormalizeWhitespaceAndStableOrder$ .`
Full checks: `go test -count=1 ./...` and `go vet ./...`.

A Vidhur-generated repair must arrive as a draft PR for human review. Seed setup
and host tests are operator preparation, not live Vidhur execution evidence.
The shared work tracker lives in ritru-labs/vidhur at
`docs/implementation/G-001-TASK-TRACKER.md` (R08/R12/R13).
