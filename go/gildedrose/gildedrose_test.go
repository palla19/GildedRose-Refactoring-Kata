package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_AgedBrie(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Aged Brie", 0, 0},
		{"Aged Brie", -1, 0},
		{"Aged Brie", 1, 0},
		{"Aged Brie", -1, 49},
		{"Aged Brie", 1, 50},
	}

	var expected = []*gildedrose.Item{
		{"Aged Brie", -1, 2},
		{"Aged Brie", -2, 2},
		{"Aged Brie", 0, 1},
		{"Aged Brie", -2, 50},
		{"Aged Brie", 0, 50},
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("Expected %v but got %v (test index %d)", expected[i], item, i)
		}
	}
}

func Test_BackstagePasses(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 48},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
	}

	var expected = []*gildedrose.Item{
		{"Backstage passes to a TAFKAL80ETC concert", 9, 12},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 13},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 11},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("Expected %v but got %v (test index %d)", expected[i], item, i)
		}
	}
}

func Test_Sulfuras(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Sulfuras, Hand of Ragnaros", 1, 80},
	}

	var expected = []*gildedrose.Item{
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Sulfuras, Hand of Ragnaros", 1, 80},
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("Expected %v but got %v (test index %d)", expected[i], item, i)
		}
	}
}

func Test_RandomItem(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Item 1", 0, 50},
		{"Item 2", -1, 50},
		{"Item 3", -1, 1},
		{"Item 4", 1, 0},
		{"Item 5", 1, 50},
	}

	var expected = []*gildedrose.Item{
		{"Item 1", -1, 48},
		{"Item 2", -2, 48},
		{"Item 3", -2, 0},
		{"Item 4", 0, 0},
		{"Item 5", 0, 49},
	}

	gildedrose.UpdateQuality(items)

	for i, item := range items {
		if *item != *expected[i] {
			t.Errorf("Expected %v but got %v (test index %d)", expected[i], item, i)
		}
	}
}
