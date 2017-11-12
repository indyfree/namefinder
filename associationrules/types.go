package associationrules

import "fmt"

//
// Association Rules
//
// In the form {A => B}, where A and B can be itemsets
type AssociationRule struct {
	A          Itemset `json:"a"`
	B          Itemset `json:"b"`
	Support    float64 `json:"support"`
	Confidence float64 `json:"confidence"`
	Lift       float64 `json:"lift"`
}

func (a AssociationRule) String() string {
	return fmt.Sprintf("%s => %s, sup: %f, conf: %f, lift: %f", a.A, a.B, a.Support, a.Confidence, a.Lift)
}

//
// Itemset
//
// Association rules base on items in transactions
//
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

//
// Frequent Itemset
//
// Wrapper type to capsule support
type FrequentItemset struct {
	items   *Itemset
	support float64
}

func (f FrequentItemset) String() string {
	return fmt.Sprintf("%s:%f", *f.items, f.support)
}
