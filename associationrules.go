package main

import "fmt"

// Association Rule in the form {A => B}, where A and B can be itemsets
type AssociationRule struct {
	A          []string
	B          []string
	Support    int
	Confidence float64
	Lift       float64
}

func (a AssociationRule) String() string {
	return fmt.Sprintf("%s => %s, sup: %d, conf: %f, lift: %f", a.A, a.B, a.Support, a.Confidence, a.Lift)
}
