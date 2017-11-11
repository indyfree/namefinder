package associationrules

import (
	"sort"
)

// TODO: break up nested for loop
// Find items which the transactions consist of, return sorted itemset
func FindAlphabet(transactions []Itemset) []Itemset {
	items := make(Itemset, 0)
	for _, t := range transactions {
		for _, titem := range t {
			if !items.Contains(titem) {
				items = append(items, titem)
			}
		}
	}

	// TODO: Return []Itemsets for Apriori to process, change?
	sort.Strings(items)
	itemset := make([]Itemset, len(items))
	for i, item := range items {
		is := Itemset{item}
		itemset[i] = is
	}
	return itemset
}
