package associationrules

import (
	"sort"
)

// Construct 1-Rule
func ConstructRules(set Itemset) []AssociationRule {
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
