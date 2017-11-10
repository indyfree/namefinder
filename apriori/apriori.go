package apriori

func Run(t []Transaction, itemsets []Itemset, minsup float64) []Itemset {
	fsets := FrequentItemsets(t, itemsets, minsup)
	result := fsets

	for len(fsets) > 0 {
		candidates := GenerateCandidates(fsets)
		fsets = FrequentItemsets(t, candidates, minsup)
		result = append(result, fsets...)
	}
	return result
}

func FrequentItemsets(t []Transaction, itemsets []Itemset, minsup float64) []Itemset {
	frequent := make([]Itemset, 0)

	for _, set := range itemsets {
		count := 0
		for _, t := range t {
			if t.ContainsSet(set) {
				count++
			}
		}
		sup := float64(count) / float64(len(t))
		if sup >= minsup {
			frequent = append(frequent, set)
		}
	}
	return frequent
}

// TODO Use Channels!
func GenerateCandidates(itemsets []Itemset) []Itemset {
	candidates := make([]Itemset, 0)

	for i := 0; i < len(itemsets); i++ {
		for j := i + 1; j < len(itemsets); j++ {
			c := CombineItemset(itemsets[i], itemsets[j])
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
