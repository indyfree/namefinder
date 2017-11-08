package main

// Association Rule in the form {A => B}, where A and B can be itemsets
// Why public attributes? (Capitals)
type AssociationRule struct {
	A          []string
	B          []string
	Support    int
	Confidence float64
	Lift       float64
}
