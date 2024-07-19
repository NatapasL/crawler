package namecleaner

import (
	"strings"
)

type NameCleaner struct {
	regexpPattern *RegexpPattern
}

func NewNameCleaner(regexpPattern *RegexpPattern) *NameCleaner {
	return &NameCleaner{regexpPattern}
}

func (nc NameCleaner) Clean(name string) string {
	for _, rx := range nc.regexpPattern.GetRegexpList() {
		name = rx.ReplaceAllString(name, "")
	}

	return nc.standardizeSpaces(name)
}

func (NameCleaner) standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
