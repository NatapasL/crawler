package siamintershop

import "manga-crawler/internal/domain/scraper/namecleaner"

type siamintershopNameCleanerModule struct {
	NameCleaner *namecleaner.NameCleaner
}

type siamintershopNameCleanerDependencies struct {
	Patterns []string
}

func initializeSiamintershopNameCleanerModule(
	deps siamintershopNameCleanerDependencies,
) siamintershopNameCleanerModule {
	regexpPattern := namecleaner.NewRegexpPattern(
		namecleaner.RegexpPatternDependencies{Patterns: deps.Patterns},
	)
	nameCleaner := namecleaner.NewNameCleaner(namecleaner.NameCleanerDependencies{RegexpPattern: regexpPattern})

	return siamintershopNameCleanerModule{NameCleaner: nameCleaner}
}
