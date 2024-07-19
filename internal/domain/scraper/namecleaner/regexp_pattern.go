package namecleaner

import (
	"regexp"
	"sort"
	"strings"
)

type RegexpPattern struct {
	patterns []string
}

func NewRegexpPattern(patterns []string) *RegexpPattern {
	regexpPattern := RegexpPattern{patterns}
	regexpPattern.sort()

	return &regexpPattern
}

func (rp *RegexpPattern) sort() {
	sortedPatterns := make([]string, len(rp.patterns))
	copy(sortedPatterns, rp.patterns)

	sort.Slice(sortedPatterns, func(i, j int) bool {
		a := rp.patternPoint(sortedPatterns[i])
		b := rp.patternPoint(sortedPatterns[j])
		return b > a
	})

	rp.patterns = sortedPatterns
}

func (rp RegexpPattern) patternPoint(pattern string) int {
	if strings.Contains(pattern, `^ `) || strings.Contains(pattern, ` $`) {
		return 5
	}

	if strings.Contains(pattern, `\[`) {
		return 4
	}

	if strings.Contains(pattern, `\(`) {
		return 3
	}

	return 0
}

func (rp RegexpPattern) GetRegexpList() []*regexp.Regexp {
	var rxs []*regexp.Regexp
	for _, pattern := range rp.patterns {
		rx := regexp.MustCompile(pattern)
		rxs = append(rxs, rx)
	}

	return rxs
}
