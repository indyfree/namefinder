package associationrules

import (
	"fmt"
	"log"
	"sort"
	"time"
)

func GetRules(t []Itemset, minsup float64, minconf float64) []AssociationRule {
	defer timeTrack(time.Now(), "GetRules")
	alphabet := DetermineAlphabet(t)
	fsets := Apriori(t, alphabet, minsup)

	// Lookup map for support values
	smap := make(map[string]float64)

	// Find strong rules
	defer timeTrack(time.Now(), "FindStrongRules")
	srules := make([]AssociationRule, 0)
	for _, f := range fsets {
		smap[fmt.Sprintf("%s", *f.items)] = f.support
		rules := ConstructRules(*f.items)
		for _, r := range rules {
			conf := f.support / smap[fmt.Sprintf("%s", r.A)]
			if conf >= minconf {
				r.Confidence = conf
				r.Support = f.support
				srules = append(srules, r)
			}
		}
	}
	return srules
}

// Construct 1-Rule
func ConstructRules(set Itemset) []AssociationRule {
	if len(set) <= 1 { // No rules with only 1 element
		return nil
	}

	rules := make([]AssociationRule, len(set))
	for i, _ := range set {
		rule := AssociationRule{A: Itemset{set[i]}}
		b := make([]string, len(set)-1)
		k := 0
		for j := 0; j < len(set); j++ {
			if j != i {
				b[k] = set[j]
				k++
			}
		}
		rule.B = b
		rules[i] = rule
	}
	return rules
}

// TODO: break up nested for loop
// Find items which the transactions consist of, return sorted itemset
func DetermineAlphabet(transactions []Itemset) Itemset {
	defer timeTrack(time.Now(), "DetermineAlphabet")
	alphabet := make(Itemset, 0)
	for _, t := range transactions {
		for _, titem := range t {
			if !alphabet.Contains(titem) {
				alphabet = append(alphabet, titem)
			}
		}
	}
	sort.Strings(alphabet)
	return alphabet
}

// Profiling Purposes
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
