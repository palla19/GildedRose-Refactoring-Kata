package gildedrose

import (
	"strings"
)

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.Name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		item.SellIn = item.SellIn - 1
		item.Quality = item.Quality + getDelta(item)
		if item.Quality > 50 {
			item.Quality = 50
		} else if item.Quality < 0 {
			item.Quality = 0
		}
	}
}

func getDelta(item *Item) int {
	isConjured, strippedName := isConjured(item)

	absDelta := 1
	if item.SellIn < 0 {
		absDelta = absDelta * 2
	}
	if isConjured {
		absDelta = absDelta * 2
	}

	isDegrading := true
	switch strippedName {
	case "Aged Brie":
		isDegrading = false
	case "Backstage passes to a TAFKAL80ETC concert":
		isDegrading = false
		if item.SellIn < 0 {
			isDegrading = true
			absDelta = item.Quality
		} else if item.SellIn < 5 {
			absDelta = absDelta * 3
		} else if item.SellIn < 10 {
			absDelta = absDelta * 2
		}
	}

	if isDegrading {
		return -1 * absDelta
	}
	return absDelta
}

func isConjured(item *Item) (bool, string){
	return strings.HasPrefix(item.Name, "Conjured"), strings.TrimPrefix(item.Name, "Conjured ")
}
