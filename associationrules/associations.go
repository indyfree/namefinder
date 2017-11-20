package associationrules

import (
	"fmt"
	"sort"
)

// Mine performs assocation rule mining using the Apriori algorithm.
// This function is the main entry point to retrieve rules from a transaction databse.
func Mine(t []Itemset, minsup float64, minconf float64) []AssociationRule {
	// 1. Find frequent itemsets
	alphabet := determineAlphabet(t)
	fsets := Apriori(t, alphabet, minsup)

	// 2. Determine strong rules from frequent itemsets
	srules := mineStrongRules(fsets, minconf)
	return srules
}

// Generate rules that have a confidence that exceeds the minconf threshold.
func mineStrongRules(fsets []FrequentItemset, minconf float64) []AssociationRule {
	smap := make(map[string]float64) // Map for fast support lookup
	srules := make([]AssociationRule, 0)

	for _, f := range fsets {
		smap[fmt.Sprintf("%s", f.items)] = f.support
		// It is only possible to generate a rules of itemsets with size > 1
		if len(f.items) > 1 {
			srules = append(srules, strongSetRules(f, smap, minconf)...)
		}
	}
	return srules
}

// strongSetRules generates all strong rules of one frequent itemset
func strongSetRules(fset FrequentItemset, smap map[string]float64, minconf float64) []AssociationRule {
	srules := make([]AssociationRule, 0)
	for i, item := range fset.items {
		conf := fset.support / smap[fmt.Sprintf("%s", item)]
		if conf > minconf {
			srules = append(srules, constructRule(i, fset.items, fset.support, conf))
		}
	}
	return srules
}

func constructRule(index int, set Itemset, support, conf float64) AssociationRule {
	rule := AssociationRule{A: Itemset{set[index]}}
	b := make(Itemset, 0)
	for j := 0; j < len(set); j++ {
		if j != index {
			b = append(b, set[j])
		}
	}
	rule.B = b
	rule.Support = support
	rule.Confidence = conf
	return rule
}

// TODO: break up nested for loop
// Find items which the transactions consist of, return sorted itemset
func determineAlphabet(transactions []Itemset) Itemset {
	alphabet := make(Itemset, 0)
	for _, t := range transactions {
		for _, titem := range t {
			if !alphabet.Contains(titem) {
				alphabet = append(alphabet, titem)
			}
		}
	}
	sort.Strings(alphabet)
	return alphabet
}
