package associationrules

import "fmt"

// AssociationRule are the form {A => B}, where A and B are itemsets that can consist of one or more items.
type AssociationRule struct {
	A          Itemset `json:"a"`
	B          Itemset `json:"b"`
	Support    float64 `json:"support"`
	Confidence float64 `json:"confidence"`
}

func (a AssociationRule) String() string {
	return fmt.Sprintf("%s => %s, sup: %f, conf: %f", a.A, a.B, a.Support, a.Confidence)
}

// Itemset is a collection of Items that are the basis for transactions and association rules.
type Itemset []string

func (s Itemset) ContainsSet(b Itemset) bool {
	for _, item := range b {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

func (s Itemset) Contains(item string) bool {
	for _, v := range s {
		if item == v {
			return true
		}
	}
	return false
}

// FrequentItemset is a wrapper type of Itemset to save its support
type FrequentItemset struct {
	items   Itemset
	support float64
}

func (f FrequentItemset) String() string {
	return fmt.Sprintf("%s:%f", f.items, f.support)
}
