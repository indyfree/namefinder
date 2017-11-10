package rulegen

import (
	"math/rand"
	"time"
)

type Transaction []string
type Itemset []string

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

// TODO use pointers?
func GenerateTransactions(n int, items []string) []Transaction {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	transactions := make([]Transaction, n)
	for i := 0; i < n; i++ {
		tLength := r.Intn(len(items)-1) + 2
		perm := r.Perm(len(items))

		t := make([]string, tLength)
		for j := range t {
			t[j] = items[perm[j]]
		}
		transactions[i] = t
	}
	return transactions
}

func FrequentItemsets(t []Transaction, itemsets []Itemset, minsup float64) []Itemset {
	frequent := make([]Itemset, 0)

	for _, set := range itemsets {
		count := 0
		for _, t := range t {
			if t.ContainsSet(set) {
				count++
			}
		}
		sup := float64(count) / float64(len(t))
		if sup >= minsup {
			frequent = append(frequent, set)
		}
	}
	return frequent
}

func Apriori(t []Transaction, items []string, minsup float64) []Itemset {
	frequent1 := FrequentItemsets(t, items, minsup)
	candidates1 := GenerateCandidates(frequent1)
	//frequent2 := FrequentItemsets(t, candidates1, minsup)
	//candidates2 := GenerateCandidates(frequent2)
	return candidates1
}

// TODO Use Channels!
func GenerateCandidates(itemsets []Itemset) []Itemset {
	candidates := make([]Itemset, 0)

	for i := 0; i < len(itemsets); i++ {
		for j := i + 1; j < len(itemsets); j++ {
			c := append(itemsets[i], itemsets[j]...)
			candidates = append(candidates, c)
		}
	}
	return candidates
}
