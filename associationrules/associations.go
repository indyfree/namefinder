package associationrules

import (
	"fmt"
	"log"
	"sort"
	"time"
)

// Mine associationrules from transactions
func Mine(t []Itemset, minsup float64, minconf float64) []AssociationRule {
	// 1. Find frequent itemsets
	alphabet := determineAlphabet(t)
	fsets := Apriori(t, alphabet, minsup)

	// 2. Determine strong rules from frequent itemsets
	srules := mineStrongRules(fsets, minconf)
	return srules
}

// Generate rules which have enough confidence (strong rules)
func mineStrongRules(fsets []FrequentItemset, minconf float64) []AssociationRule {
	defer timeTrack(time.Now(), "mineStrongRules")
	smap := make(map[string]float64) // Map for fast support lookup
	srules := make([]AssociationRule, 0)

	for _, f := range fsets {
		smap[fmt.Sprintf("%s", f.items)] = f.support
		if len(f.items) > 1 {
			srules = append(srules, strongRulesOfSet(f, smap, minconf)...)
		}
	}
	return srules
}

func strongRulesOfSet(fset FrequentItemset, smap map[string]float64, minconf float64) []AssociationRule {
	srules := make([]AssociationRule, 0)
	for i, item := range fset.items {
		conf := fset.support / smap[fmt.Sprintf("%s", item)]
		if conf > minconf {
			srules = append(srules, generateRule(i, fset.items, fset.support, conf))
		}
	}
	return srules
}

func generateRule(index int, set Itemset, support, conf float64) AssociationRule {
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
	defer timeTrack(time.Now(), "determineAlphabet")
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

// Profiling Purposes
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
