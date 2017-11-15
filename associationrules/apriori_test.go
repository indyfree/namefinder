package associationrules

import (
	"testing"
)

func TestApriori(t *testing.T) {
	cases := []struct {
		t      []Itemset
		items  Itemset
		minsup float64
		want   []Itemset
	}{
		{t: []Itemset{{"A", "B"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"A", "C", "D"},
			{"A", "C", "E"}, {"A", "B", "E"}, {"A", "C", "D", "E"}},
			items:  Itemset{"A", "B", "C", "D", "E"},
			minsup: 0.25,
			want: []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}, {"A", "B"}, {"A", "C"},
				{"A", "D"}, {"A", "E"}, {"C", "D"}, {"C", "E"}, {"A", "C", "D"}, {"A", "C", "E"}}},
		{t: []Itemset{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
			items:  Itemset{"A", "B", "C", "D"},
			minsup: 0.5,
			want:   []Itemset{{"A"}, {"B"}, {"D"}}},
		{t: []Itemset{{"Peter", "Gracie"}, {"Jack", "Barley"}, {"Max", "Tom"}, {"Tom", "Barley"}, {"Jack", "Barley"}},
			items:  Itemset{"Peter", "Gracie", "Jack", "Barley", "Tom"},
			minsup: 0.4,
			want:   []Itemset{{"Barley"}, {"Jack"}, {"Tom"}, {"Barley Jack"}}},
		{t: []Itemset{},
			items:  Itemset{"Peter", "Gracie", "Jack", "Barley", "Tom"},
			minsup: 0.0,
			want:   []Itemset{}},
		{t: []Itemset{{"Peter"}, {"Gracie"}},
			items:  Itemset{"Peter", "Gracie", "Jack", "Barley", "Tom"},
			minsup: 1.0,
			want:   []Itemset{}},
	}

	for _, c := range cases {
		fsets := Apriori(c.t, c.items, c.minsup)
		got := GetItemset(fsets)
		if !hasSameSets(c.want, got) {
			t.Errorf("Apriori(%q, %f): \n got: %q\n, want: %q", c.t, c.minsup, got, c.want)
		}
	}

}

func GetItemset(fsets []FrequentItemset) []Itemset {
	items := make([]Itemset, len(fsets))
	for i, v := range fsets {
		items[i] = *v.items
	}
	return items
}

func TestFrequentItemSets(t *testing.T) {
	testcases := []struct {
		t      []Itemset
		items  []Itemset
		minsup float64
		want   []Itemset
	}{
		{t: []Itemset{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"A", "B"}},
			minsup: 0.5,
			want:   []Itemset{{"A"}, {"B"}, {"D"}}},
		{t: []Itemset{},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
			minsup: 0.0,
			want:   []Itemset{}},
	}

	c := testcases[0]
	candidates := make(chan Itemset, len(c.items))
	for _, v := range c.items {
		candidates <- v
	}
	close(candidates)
	fsets := FrequentItemsets(c.t, c.minsup, candidates)
	got := GetItemset(fsets)
	if !hasSameSets(c.want, got) {
		t.Errorf("FrequentItemSets(%q, %f) == %q, want %q", c.t, c.minsup, got, c.want)
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

//
// Test helper functions
//

// Set order does not matter
func hasSameSets(a []Itemset, b []Itemset) bool {
	if len(a) != len(b) {
		return false
	}
	for _, v := range b {
		if !contains(a, v) {
			return false
		}
	}
	return true
}

func contains(a []Itemset, b Itemset) bool {
	for i, v := range a {
		if a[i].Equals(v) {
			return true
		}
	}
	return false
}

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
