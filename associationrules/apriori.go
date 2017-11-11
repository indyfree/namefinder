package associationrules

func Apriori(transactions []Itemset, alphabet []Itemset, minsup float64) []FrequentItemset {
	// Find frequent 1-Itemsets first
	fsets := FrequentItemsets(transactions, alphabet, minsup)
	result := fsets

	// Generate candidates from frequent k-1 Itemsets and find frequent ones
	for len(fsets) > 0 {
		candidates := GenerateCandidates(fsets)
		fsets = FrequentItemsets(transactions, candidates, minsup)
		result = append(result, fsets...)
	}
	return result
}

func FrequentItemsets(transactions []Itemset, itemsets []Itemset, minsup float64) []FrequentItemset {
	frequent := make([]FrequentItemset, 0)

	for i, set := range itemsets {
		count := 0
		for _, t := range transactions {
			if t.ContainsSet(set) {
				count++
			}
		}
		sup := float64(count) / float64(len(transactions))
		if sup >= minsup {
			frequent = append(frequent, FrequentItemset{&itemsets[i], sup})
		}
	}
	return frequent
}

// TODO Use Channels!
func GenerateCandidates(fsets []FrequentItemset) []Itemset {
	candidates := make([]Itemset, 0)

	for i := 0; i < len(fsets); i++ {
		for j := i + 1; j < len(fsets); j++ {
			c := CombineItemset(*fsets[i].items, *fsets[j].items)
			if c != nil {
				candidates = append(candidates, c)
			}
		}
	}
	return candidates
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
