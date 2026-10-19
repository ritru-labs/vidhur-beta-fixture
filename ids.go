// Package ids normalizes identifiers supplied by batch-import callers.
package ids

import (
	"strings"
)

// Normalize returns nonempty distinct identifiers in first-seen order.
// Leading and trailing whitespace must not change identifier identity.
func Normalize(input []string) []string {
	out := make([]string, 0, len(input))
	seen := make(map[string]bool)
	for _, id := range input {
		trimmedID := strings.TrimSpace(id)
		if trimmedID == "" || seen[trimmedID] {
			continue
		}
		seen[trimmedID] = true
		out = append(out, trimmedID)
	}
	return out
}
