package associationrules

import (
	"time"
)

func Apriori(transactions []Itemset, alphabet []Itemset, minsup float64) []FrequentItemset {
	defer timeTrack(time.Now(), "Apriori")

	//Find frequent 1-Itemsets first
	candidates := setToChannel(alphabet)
	fsets := FrequentItemsets(transactions, minsup, candidates)
	results := fsets

	// Generate candidates from subsequent itemsets and find frequent ones
	for len(fsets) > 0 {
		candidates = GenerateCandidates(fsets)
		fsets = FrequentItemsets(transactions, minsup, candidates)
		results = append(results, fsets...)
	}
	return results
}

// TODO Channels! Searching for frequent can happen concurrent
func FrequentItemsets(transactions []Itemset, minsup float64, candidates <-chan Itemset) []FrequentItemset {
	fsets := make([]FrequentItemset, 0)
	for c := range candidates {
		count := 0
		for _, t := range transactions {
			if t.ContainsSet(c) {
				count++
			}
		}
		sup := float64(count) / float64(len(transactions))
		if sup >= minsup {
			item := c // Copy
			fset := FrequentItemset{&item, sup}
			fsets = append(fsets, fset)
		}
	}
	return fsets
}

// Channel Generator Pattern
func GenerateCandidates(fsets []FrequentItemset) <-chan Itemset {
	ch := make(chan Itemset, len(fsets)*len(fsets))
	go func() {
		for i := 0; i < len(fsets); i++ {
			for j := i + 1; j < len(fsets); j++ {
				cset := CombineItemset(*fsets[i].items, *fsets[j].items)
				if cset != nil {
					ch <- cset
				}
			}
		}
		close(ch)
	}()
	return ch
}

// Only combine itemsets that are different at the last index
// Apriori premise
// TODO Make Method, changing own type?
func CombineItemset(a Itemset, b Itemset) Itemset {
	if len(a) != len(b) || a[len(a)-1] == b[len(b)-1] {
		return nil
	}

	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return nil
		}
	}
	return append(a, b[len(b)-1])
}

// Helper Functions
func setToChannel(items []Itemset) <-chan Itemset {
	c := make(chan Itemset, len(items))
	for _, v := range items {
		c <- v
	}
	close(c)
	return c
}
