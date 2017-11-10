package rulegen

import (
	"testing"
)

// TODO: Why here small attributes?
type testpair struct {
	transactions [][]string
	min_sup      float64
	min_conf     float64
	rules        AssociationRules
}

// var cases = []testpair{
// 	{[][]string{{"A", "B"}, {"B", "C"}, {"A", "B", "C"}, {"A", "B"}}, 1.0, 0.5, 0.5,
// 		ruleserver.AssociationRules{ruleserver.AssociationRule{[]string{"A"}, []string{"B"}, 1, 0.75, 0.75}}},
// }

func TestGenerateTransactions(t *testing.T) {
	itemset := []string{"A", "B", "C", "D"}
	tnumber := 100
	transactions := GenerateTransactions(tnumber, itemset)

	if len(transactions) != tnumber {
		t.Errorf("Number of generated transaction does not match, want: %d, got: %d",
			tnumber, len(transactions))
	}
}

func TestFrequentItemSets(t *testing.T) {
	cases := []struct {
		t      []Transaction
		items  []string
		minsup float64
		want   []Itemset
	}{
		{[]Transaction{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"D", "B"}},
			[]string{"A", "B", "C", "D"}, 0.5, []Itemset{{"A"}, {"B"}, {"D"}}},
	}
	for _, c := range cases {
		got := FrequentItemsets(c.t, c.items, c.minsup)
		if !equalSets(c.want, got) {
			t.Errorf("FrequentItemSets() == %q, want %q", got, c.want)
		}
	}
}

func equalSets(a []Itemset, b []Itemset) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}
