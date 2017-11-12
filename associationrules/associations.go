package associationrules

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func GetRules(t []Itemset, minsup float64, minconf float64) []AssociationRule {
	defer timeTrack(time.Now(), "GetRules")
	alphabet := FindAlphabet(t)
	fsets := Apriori(t, alphabet, minsup)

	// Lookup map for support values
	smap := make(map[string]float64)

	// Find strong rules
	defer timeTrack(time.Now(), "FindStrongRules")
	srules := make([]AssociationRule, 0)
	for _, f := range fsets {
		smap[fmt.Sprintf("%s", *f.items)] = f.support
		rules := ConstructRules(*f.items)
		for _, r := range rules {
			conf := f.support / smap[fmt.Sprintf("%s", r.A)]
			if conf >= minconf {
				r.Confidence = conf
				r.Support = f.support
				srules = append(srules, r)
			}
		}
	}
	return srules
}

// Construct 1-Rule
func ConstructRules(set Itemset) []AssociationRule {
	if len(set) <= 1 { // No rules with only 1 element
		return nil
	}

	rules := make([]AssociationRule, len(set))
	for i, _ := range set {
		rule := AssociationRule{A: Itemset{set[i]}}
		b := make([]string, len(set)-1)
		k := 0
		for j := 0; j < len(set); j++ {
			if j != i {
				b[k] = set[j]
				k++
			}
		}
		rule.B = b
		rules[i] = rule
	}
	return rules
}

// TODO: break up nested for loop
// Find items which the transactions consist of, return sorted itemset
func FindAlphabet(transactions []Itemset) []Itemset {
	defer timeTrack(time.Now(), "FindAlphabet")
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

// Profiling Purposes
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
