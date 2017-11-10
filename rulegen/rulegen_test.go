package rulegen

import (
	"fmt"
	"testing"
)

// TODO: Why here small attributes?
type testpair struct {
	transactions [][]string
	min_sup      float64
	min_conf     float64
	rules        AssociationRules
}

var cases = []struct {
	t      []Transaction
	items  []Itemset
	minsup float64
	want   []Itemset
}{
	{[]Transaction{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
		[]Itemset{{"A"}, {"B"}, {"C"}, {"D"}}, 0.0, []Itemset{{"A"}, {"B"}, {"C"}, {"D"}}},
	{[]Transaction{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
		[]Itemset{{"A"}, {"B"}, {"C"}, {"D"}}, 0.5, []Itemset{{"A"}, {"B"}, {"D"}}},
	{[]Transaction{{"Hund", "Katze"}, {"Maus", "Kind"}, {"Vater", "Mutter"}, {"Mutter", "Kind"}, {"Kind", "Maus"}},
		[]Itemset{{"Hund"}, {"Katze"}, {"Maus"}, {"Kind"}, {"Mutter"}}, 0.6, []Itemset{{"Kind"}}},
	{[]Transaction{},
		[]Itemset{{"Hund"}, {"Katze"}, {"Maus"}, {"Kind"}, {"Mutter"}}, 0.0, []Itemset{}},
	{[]Transaction{{"Hund"}, {"Katze"}},
		[]Itemset{{"Hund"}, {"Katze"}, {"Maus"}, {"Kind"}, {"Mutter"}}, 1.0, []Itemset{}},
}

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
	for _, c := range cases {
		got := FrequentItemsets(c.t, c.items, c.minsup)
		if !equalSets(c.want, got) {
			t.Errorf("FrequentItemSets() == %q, want %q", got, c.want)
		}
	}
}

func TestCombineItemset(t *testing.T) {
	testcases := []struct {
		in   []Itemset
		want Itemset
	}{
		{[]Itemset{{"A"}, {"A"}}, nil},
		{[]Itemset{{"A", "B", "C"}, {"A", "B", "D"}}, Itemset{"A", "B", "C", "D"}},
		{[]Itemset{{"A", "B", "C"}, {"A", "C", "D"}}, nil},
	}
	for _, c := range testcases {
		got := CombineItemset(c.in[0], c.in[1])
		fmt.Println(got, c.want)
		if !c.want.Equals(got) {
			t.Errorf("CombineItemset() == %q, want %q", got, c.want)
		}
	}

}

func TestApriori(t *testing.T) {
	c := cases[0]
	fmt.Println(Apriori(c.t, c.items, c.minsup))
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
