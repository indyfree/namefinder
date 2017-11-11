package associationrules

import "fmt"

// Association Rule in the form {A => B}, where A and B can be itemsets
type AssociationRule struct {
	A          []string `json:"a"`
	B          []string `json:"b"`
	Support    float64  `json:"support"`
	Confidence float64  `json:"confidence"`
	Lift       float64  `json:"lift"`
}

// TODO explain why to declare
type AssociationRules []AssociationRule

func (a AssociationRule) String() string {
	return fmt.Sprintf("%s => %s, sup: %f, conf: %f, lift: %f", a.A, a.B, a.Support, a.Confidence, a.Lift)
}

type Itemset []string

type FrequentItemset struct {
	items   *Itemset
	support float64
}

func (f FrequentItemset) String() string {
	return fmt.Sprintf("%s:%f", *f.items, f.support)
}

// TODO refactor!
func (s Itemset) Equals(b Itemset) bool {
	if len(s) != len(b) {
		return false
	}
	for i := range s {
		if s[i] != b[i] {
			return false
		}
	}
	return true
}

func (s Itemset) ContainsSet(itemset Itemset) bool {
	for _, item := range itemset {
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
