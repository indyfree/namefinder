package main

import (
	"math/rand"
	"time"

	ar "github.com/indyfree/namefinder/associationrules"
)

func GenerateTransactions(n int, items []string) []ar.Itemset {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	transactions := make([]ar.Itemset, n)

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
