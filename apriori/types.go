package apriori

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

type Transaction []string
type Itemset []string

type FrequentItemset struct {
	items   *Itemset
	support float64
}

func (f FrequentItemset) String() string {
	return fmt.Sprintf("%s:%f", *f.items, f.support)
}

// TODO refactor!
func (a Itemset) Equals(b Itemset) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (t Transaction) ContainsSet(itemset Itemset) bool {
	for _, item := range itemset {
		if !t.Contains(item) {
			return false
		}
	}
	return true
}

func (t Transaction) Contains(item string) bool {
	for _, v := range t {
		if item == v {
			return true
		}
	}
	return false
}
