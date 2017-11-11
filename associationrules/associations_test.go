package associationrules

import "testing"

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
		if !equalSets(c.want, got) {
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
		if !equalRules(c.want, got) {
			t.Errorf("ConstructRules(%q) == \n%q want: \n%q", c.in, got, c.want)
		}
	}
}
