package associationrules

import (
	"fmt"
	"sort"
)

func GetRules(t []Itemset, minsup float64, minconf float64) []AssociationRule {
	alphabet := FindAlphabet(t)
	fsets := Apriori(t, alphabet, minsup)
	srules := FindStrongRules(fsets, minconf)
	return srules
}

func FindStrongRules(fsets []FrequentItemset, minconf float64) []AssociationRule {
	result := make([]AssociationRule, 0)

	// Lookup map for support value
	smap := make(map[string]float64)
	for _, f := range fsets {
		smap[fmt.Sprintf("%s", *f.items)] = f.support
	}

	for _, f := range fsets {
		rules := ConstructRules(*f.items)
		for _, r := range rules {
			conf := f.support / smap[fmt.Sprintf("%s", r.A)]
			if conf >= minconf {
				r.Confidence = conf
				r.Support = f.support
				result = append(result, r)
			}
		}
	}
	return result
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
