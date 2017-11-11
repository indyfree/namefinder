package associationrules

import (
	"testing"
)

func TestApriori(t *testing.T) {
	cases := []struct {
		t      []Itemset
		items  []Itemset
		minsup float64
		want   []Itemset
	}{
		{t: []Itemset{{"A", "B"}, {"A", "D"}, {"B", "C"}, {"B", "D"}, {"A", "C", "D"},
			{"A", "C", "E"}, {"A", "B", "E"}, {"A", "C", "D", "E"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}},
			minsup: 0.25,
			want: []Itemset{{"A"}, {"B"}, {"C"}, {"D"}, {"E"}, {"A", "B"}, {"A", "C"},
				{"A", "D"}, {"A", "E"}, {"C", "D"}, {"C", "E"}, {"A", "C", "D"}, {"A", "C", "E"}}},
		{t: []Itemset{{"A", "B"}, {"B", "C"}, {"A", "D"}, {"B", "D"}},
			items:  []Itemset{{"A"}, {"B"}, {"C"}, {"D"}},
			minsup: 0.5,
			want:   []Itemset{{"A"}, {"B"}, {"D"}}},
		{t: []Itemset{{"Peter", "Gracie"}, {"Jack", "Barley"}, {"Max", "Tom"}, {"Tom", "Barley"}, {"Barley", "Jack"}},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 0.6,
			want:   []Itemset{{"Barley"}}},
		{t: []Itemset{},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 0.0,
			want:   []Itemset{}},
		{t: []Itemset{{"Peter"}, {"Gracie"}},
			items:  []Itemset{{"Peter"}, {"Gracie"}, {"Jack"}, {"Barley"}, {"Tom"}},
			minsup: 1.0,
			want:   []Itemset{}},
	}

	for _, c := range cases {
		fsets := Apriori(c.t, c.items, c.minsup)
		got := GetItemset(fsets)
		if !IsEqual(c.want, got) {
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

	for _, c := range testcases {
		fsets := FrequentItemsets(c.t, c.items, c.minsup)
		got := GetItemset(fsets)
		if !IsEqual(c.want, got) {
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

// Helper function
func IsEqual(a []Itemset, b []Itemset) bool {
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
