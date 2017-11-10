package rulegen

import (
	"math/rand"
	"time"
)

type Transactions [][]string
type Itemset []string
type Itemsets []Itemset

// TODO refactor?
func (a Itemset) isEqual(b Itemset) bool {
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

func (a Itemsets) isEqual(b Itemsets) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].isEqual(b[i]) {
			return false
		}
	}
	return true
}

// TODO use pointers?
func GenerateTransactions(number int, items []string) Transactions {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	transactions := make([][]string, number)
	for i := 0; i < number; i++ {
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

func FrequentItemsets(t Transactions, items []string, minsup float64) Itemsets {
	frequent := make(Itemsets, 0)

	for _, item := range items {
		count := 0
		for _, t := range t {
			if ItemInTransaction(item, t) {
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

func ItemInTransaction(item string, transaction []string) bool {
	for _, v := range transaction {
		if item == v {
			return true
		}
	}
	return false
}
