package web

import (
	"bytes"
	"github.com/inhies/go-tld"
	"unicode"
	"unicode/utf8"
)

var (
	// A-Z, a-z, 0-9, and hyphen
	domainTable = &unicode.RangeTable{
		R16: []unicode.Range16{
			{'-', '-', 1},
			{'0', '9', 1},
			{'A', 'Z', 1},
			{'a', 'z', 1},
		},
		LatinOffset: 4,
	}
)

// Checks for a valid domain name
func IsDomain(p []byte) bool {
	// Domain rules:
	// - 253 character total length max
	// - 63 character label max
	// - 127 sub-domains
	// - Characters a-z, A-Z, 0-9, and -
	// - Labels may not start or end with -
	// - TLD may not be all numeric

	// Check for max length.
	// NOTE: Invalid unicode will count as a 1 byte rune, but we'll catch that
	// later.
	if utf8.RuneCount(p) > 252 {
		return false
	}

	// First we split by label
	domain := bytes.Split(p, []byte("."))

	// 127 sub-domains max (not including TLD)
	if len(domain) > 127 {
		return false
	}

	// Check each domain for valid characters
	for _, subDomain := range domain {
		length := len(subDomain)
		// Check for a domain with two periods next to eachother.
		if length < 1 {
			return false
		}

		// Check 63 character max.
		if length > 62 {
			return false
		}

		// Check that label doesn't start or end with hyphen.
		r, size := utf8.DecodeRune(subDomain)
		if r == utf8.RuneError && size == 1 {
			// Invalid rune
			return false
		}

		if r == '-' {
			return false
		}

		r, size = utf8.DecodeLastRune(subDomain)
		if r == utf8.RuneError && size == 1 {
			// Invalid rune
			return false
		}

		if r == '-' {
			return false
		}

		// Now we check each rune individually to make sure its valid unicode
		// and an acceptable character.
		for i := 0; i < length; {
			if subDomain[i] < utf8.RuneSelf {
				// Check if it's a valid domain character
				if !unicode.Is(domainTable, rune(subDomain[i])) {
					return false
				}
				i++
			} else {
				r, size := utf8.DecodeRune(subDomain[i:])
				if size == 1 {
					// All valid runes of size 1 (those
					// below RuneSelf) were handled above.
					// This must be a RuneError.
					return false
				}
				// Check if it's a valid domain character
				if !unicode.Is(domainTable, r) {
					return false
				}
				i += size
			}
		}
	}

	// We have all valid unicode characters, now make sure the TLD is real.
	domainTLD := domain[len(domain)-1]
	if tld.Valid(domainTLD) {
		return true
	}

	// Not sure how we got here, but lets return false just in case.
	return false
}
