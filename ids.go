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
		id = strings.TrimSpace(id)
		if id == "" || seen[id] {
			continue
		}
		seen[id] = true
		out = append(out, id)
	}
	return out
}
