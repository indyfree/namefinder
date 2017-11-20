package associationrules

import (
	"log"
	"sync"
	"time"
)

// Apriori runs the Apriori Algorithm from Agrawal and Sikrant (1994) on a transaction database.
// Go concurrency features are used to improve efficiency the proposed Algorithm.
func Apriori(transactions []Itemset, alphabet Itemset, minsup float64) []FrequentItemset {
	defer timeTrack(time.Now(), "Apriori")

	//Find frequent 1-Itemsets first
	candidates := candidates(alphabet)
	fsets := FrequentItemsets(transactions, minsup, candidates)
	results := fsets

	// Perform the two Apriori steps until no more frequent itemsets are discovered
	for len(fsets) > 0 {
		candidates := GenerateCandidates(fsets)
		fsets = FrequentItemsets(transactions, minsup, candidates)
		results = append(results, fsets...)
	}
	return results
}

// FrequentItemsets return the candidates that exceed the support threshold.
func FrequentItemsets(transactions []Itemset, minsup float64, candidates <-chan Itemset) []FrequentItemset {
	results := make(chan FrequentItemset, cap(candidates))

	// Worker Pool to concurrently scan transactions
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// Worker only reads from transaction db, thus no mutex needed
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

// frequentItemWorker is a Go worker that determines the candidates that are frequent
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

// GenerateCandidates constructs the possible candidates given the
// frequent itemsets of the previous step. It uses the channel generator
// pattern to facilitate concurrent Candidate Genration and Support Lookup.
func GenerateCandidates(fsets []FrequentItemset) <-chan Itemset {
	ch := make(chan Itemset, len(fsets)*len(fsets))
	go func() {
		for i := 0; i < len(fsets); i++ {
			for j := i + 1; j < len(fsets); j++ {
				cset := combineItemset(fsets[i].items, fsets[j].items)
				if cset != nil {
					ch <- cset
				}
			}
		}
		close(ch)
	}()
	return ch
}

// According to the Apriori premise, candidate itemsets can only
// be a combination of frequequent itemsets that differ at the last index.
func combineItemset(a Itemset, b Itemset) Itemset {
	if len(a) != len(b) || a[len(a)-1] == b[len(b)-1] {
		return nil
	}

	for i := 0; i < len(a)-1; i++ {
		if a[i] != b[i] {
			return nil
		}
	}

	citems := make([]string, len(b)+1)
	for j := range a {
		citems[j] = a[j]
	}
	citems[len(b)] = b[len(b)-1]
	return citems
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

// Debug: Profiling Purposes
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
