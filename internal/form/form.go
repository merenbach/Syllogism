package form

// A Form of syllogism.
type Form int

// TODO: use iota
// TODO: should be able to swap ordering of any of these and retain same functionality once refactor is complete
//       since we won't be relying on integer values anymore
const (
	Undefined   Form = (-1)
	SomeAIsB         = 0
	SomeAIsNotB      = 1
	AllAIsB          = 2
	NoAIsB           = 3
	AIsT             = 4
	AIsNotT          = 5
)

// IsNegative determines if this form is negative.
func (t Form) IsNegative() bool {
	return t%2 == 1
}

// func (t Type) String() string {
// 	switch t {
// 	case TypeGeneralTerm:
// 		return "general term"
// 	case TypeDesignator:
// 		return "designator"
// 	default:
// 		return "undetermined type"
// 	}
// }
