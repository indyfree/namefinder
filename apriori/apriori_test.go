package apriori

import (
	"testing"
)

func TestApriori(t *testing.T) {
	cases := []struct {
		t      []Transaction
		items  []Itemset
		minsup float64
		want   []Itemset
	}{
		{t: []Transaction{{"A", "B"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"A", "C", "D"},
			{"A", "C", "E"}, {"A", "B", "E"}, {"A", "C", "D", "E"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
			minsup: 0.25,
			want: []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}, {"A", "B"}, {"A", "C"},
				{"A", "D"}, {"A", "E"}, {"C", "D"}, {"C", "E"}, {"A", "C", "D"}, {"A", "C", "E"}}},
		{t: []Transaction{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}},
			minsup: 0.5,
			want:   []Itemset{{"A"}, {"B"}, {"D"}}},
		{t: []Transaction{{"Peter", "Gracie"}, {"Jack", "Barley"}, {"Max", "Tom"}, {"Tom", "Barley"}, {"Barley", "Jack"}},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 0.6,
			want:   []Itemset{{"Barley"}}},
		{t: []Transaction{},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 0.0,
			want:   []Itemset{}},
		{t: []Transaction{{"Peter"}, {"Gracie"}},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 1.0,
			want:   []Itemset{}},
	}

	for _, c := range cases {
		got := Run(c.t, c.items, c.minsup)
		if !equalSets(c.want, got) {
			t.Errorf("Apriori(%q, %f): \n got: %q\n, want: %q", c.t, c.minsup, got, c.want)
		}
	}
}

func TestFrequentItemSets(t *testing.T) {
	testcases := []struct {
		t      []Transaction
		items  []Itemset
		minsup float64
		want   []Itemset
	}{
		{t: []Transaction{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"A", "B"}},
			minsup: 0.5,
			want:   []Itemset{{"A"}, {"B"}, {"D"}}},
		{t: []Transaction{},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
			minsup: 0.0,
			want:   []Itemset{}},
	}

	for _, c := range testcases {
		got := FrequentItemsets(c.t, c.items, c.minsup)
		if !equalSets(c.want, got) {
			t.Errorf("FrequentItemSets(%q, %f) == %q, want %q", c.t, c.minsup, got, c.want)
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
		if !c.want.Equals(got) {
			t.Errorf("CombineItemset(%q, %q) == %q, want %q", c.in[0], c.in[1], got, c.want)
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
