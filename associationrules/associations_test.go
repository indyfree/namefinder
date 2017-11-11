package associationrules

import (
	"testing"
)

func TestFindAlphabet(t *testing.T) {
	cases := []struct {
		in   []Itemset
		want []Itemset
	}{
		{in: []Itemset{{"A", "B"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"A", "C", "D"},
			{"A", "C", "E"}, {"A", "B", "E"}, {"A", "C", "D", "E"}},
			want: []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
		},
		{in: []Itemset{{"Gracie", "Barley", "Bonny"}, {"Max", "Jackey", "Bonny"}},
			want: []Itemset{{"Barley"}, {"Bonny"}, {"Gracie"}, {"Jackey"}, {"Max"}},
		},
		{in: []Itemset{},
			want: []Itemset{},
		},
	}

	for _, c := range cases {
		got := FindAlphabet(c.in)
		if !IsEqual(c.want, got) {
			t.Errorf("FindAlphabet() got: %q want: %q", got, c.want)
		}
	}
}

func TestConstructRules(t *testing.T) {
	cases := []struct {
		in   Itemset
		want []AssociationRule
	}{
		{in: Itemset{"A", "B", "E"},
			want: []AssociationRule{AssociationRule{A: Itemset{"A"}, B: Itemset{"B", "E"}},
				AssociationRule{A: Itemset{"B"}, B: Itemset{"A", "E"}}, AssociationRule{A: Itemset{"E"}, B: Itemset{"A", "B"}}},
		},
		{in: Itemset{"A", "B"},
			want: []AssociationRule{AssociationRule{A: Itemset{"A"}, B: Itemset{"B"}},
				AssociationRule{A: Itemset{"B"}, B: Itemset{"A"}}},
		},
		{in: Itemset{"A"},
			want: []AssociationRule{},
		},
	}

	for _, c := range cases {
		got := ConstructRules(c.in)
		if !equalRules(got, c.want) {
			t.Errorf("ConstructRules(%q) == \n%q want: \n%q", c.in, got, c.want)
		}
	}
}

func TestGetRules(t *testing.T) {
	cases := []struct {
		t       []Itemset
		minsup  float64
		minconf float64
		want    []AssociationRule
	}{
		{t: []Itemset{{"A", "B"}, {"B", "C"}, {"B", "D"}, {"A", "B", "D"}},
			minsup:  0.50,
			minconf: 0.50,
			want: []AssociationRule{
				AssociationRule{A: Itemset{"A"}, B: Itemset{"B"}, Support: 0.5, Confidence: 1},
				AssociationRule{A: Itemset{"B"}, B: Itemset{"A"}, Support: 0.5, Confidence: 0.5},
				AssociationRule{A: Itemset{"B"}, B: Itemset{"D"}, Support: 0.5, Confidence: 0.5},
				AssociationRule{A: Itemset{"D"}, B: Itemset{"B"}, Support: 0.5, Confidence: 1},
			},
		},
	}

	for _, c := range cases {
		got := GetRules(c.t, c.minsup, c.minconf)
		if !equalRules(got, c.want) {
			t.Errorf("GetRules(): \n got: %q\n, want: %q", got, c.want)
		}
	}
}

func equalRules(a []AssociationRule, b []AssociationRule) bool {
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
