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

func FrequentItemsets(t []Transaction, items []string, minsup float64) []Itemset {
	frequent := make([]Itemset, 0)

	for _, item := range items {
		count := 0
		for _, t := range t {
			if t.Contains(item) {
				count++
			}
		}
		sup := float64(count) / float64(len(t))
		if sup >= minsup {
			frequent = append(frequent, []string{item})
		}
	}
	return frequent
}
