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

func (a AssociationRule) Equals(b AssociationRule) bool {
	if !a.A.Equals(b.A) || !a.B.Equals(b.B) {
		return false
	} else if a.Support != b.Support || a.Confidence != a.Confidence || a.Lift != b.Lift {
		return false
	}
	return true
}

//
// Itemset
//
// Association rules base on items in transactions
//
type Itemset []string

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
// Wrapper type to capsule according support
type FrequentItemset struct {
	items   *Itemset
	support float64
}

func (f FrequentItemset) String() string {
	return fmt.Sprintf("%s:%f", *f.items, f.support)
}
