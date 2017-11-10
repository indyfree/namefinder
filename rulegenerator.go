package main

import (
	"math/rand"
	"time"

	"github.com/indyfree/namefinder/apriori"
)

// TODO use pointers?
func GenerateTransactions(n int, items []string) []apriori.Transaction {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	transactions := make([]apriori.Transaction, n)

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
