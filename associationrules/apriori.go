package associationrules

import (
	"sync"
	"time"
)

func Apriori(transactions []Itemset, alphabet Itemset, minsup float64) []FrequentItemset {
	defer timeTrack(time.Now(), "Apriori")

	//Find frequent 1-Itemsets first
	candidates := candidates(alphabet)
	fsets := FrequentItemsets(transactions, minsup, candidates)
	results := fsets

	// Generate candidates from subsequent itemsets and find frequent ones
	for len(fsets) > 0 {
		candidates := GenerateCandidates(fsets)
		fsets = FrequentItemsets(transactions, minsup, candidates)
		results = append(results, fsets...)
	}
	return results
}

func FrequentItemsets(transactions []Itemset, minsup float64, candidates <-chan Itemset) []FrequentItemset {
	results := make(chan FrequentItemset, cap(candidates))

	// Worker Pool to concurrently scan transactions
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// Only Reading from transaction DB, no mutex needed
		go frequentItemWorker(&wg, transactions, minsup, candidates, results)
	}
	wg.Wait()
	close(results)

	// Collect results
	fsets := make([]FrequentItemset, 0)
	for fset := range results {
		fsets = append(fsets, fset)
	}
	return fsets
}

// Go Worker to determine which of the candidates are frequent
func frequentItemWorker(wg *sync.WaitGroup, transactions []Itemset, minsup float64,
	candidates <-chan Itemset, fsets chan<- FrequentItemset) {
	for c := range candidates {
		sup := calculateSupport(transactions, c)
		if sup >= minsup {
			fsets <- FrequentItemset{c, sup}
		}
	}
	wg.Done()
}

func calculateSupport(transactions []Itemset, set Itemset) float64 {
	count := 0
	for _, t := range transactions {
		if t.ContainsSet(set) {
			count++
		}
	}
	return float64(count) / float64(len(transactions))
}

// Channel Generator Pattern
func GenerateCandidates(fsets []FrequentItemset) <-chan Itemset {
	ch := make(chan Itemset, len(fsets)*len(fsets))
	go func() {
		for i := 0; i < len(fsets); i++ {
			for j := i + 1; j < len(fsets); j++ {
				cset := CombineItemset(fsets[i].items, fsets[j].items)
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
func candidates(items []string) <-chan Itemset {
	c := make(chan Itemset, len(items)*len(items))
	for _, v := range items {
		c <- Itemset{v}
	}
	close(c)
	return c
}
