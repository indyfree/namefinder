package rulegen

import (
	"math/rand"
	"time"
)

func GenerateTransactions(number int, items []string) [][]string {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	transactions := make([][]string, number)
	for i := 0; i < number; i++ {
		tLength := r.Intn(len(items)-1) + 2

		t := make([]string, tLength)
		perm := r.Perm(len(items))
		for j := range t {
			t[j] = items[perm[j]]
		}
		transactions[i] = t
	}
	return transactions
}
