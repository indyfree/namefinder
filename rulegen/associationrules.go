package rulegen

import "fmt"

// Association Rule in the form {A => B}, where A and B can be itemsets
type AssociationRule struct {
	A          []string `json:"a"`
	B          []string `json:"b"`
	Support    float64  `json:"support"`
	Confidence float64  `json:"confidence"`
	Lift       float64  `json:"lift"`
}

// TODO explain why to declare
type AssociationRules []AssociationRule

func (a AssociationRule) String() string {
	return fmt.Sprintf("%s => %s, sup: %f, conf: %f, lift: %f", a.A, a.B, a.Support, a.Confidence, a.Lift)
}
