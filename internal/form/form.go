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
	// = 6???
)

// IsNegative determines if this form is negative.
func (t Form) IsNegative() bool {
	// TODO: don't rely on numeric values here?
	// switch t {
	// case SomeAIsNotB:
	// 	fallthrough
	// case NoAIsB:
	// 	fallthrough
	// case AIsNotT:
	// 	return true
	// }
	return t%2 == 1
}

// Copula associated with this form.
// TODO: add some tests!
func (t Form) Copula() string {
	switch t {
	case SomeAIsB:
		return "  is"
	case SomeAIsNotB:
		return "  is not"
	case AllAIsB:
		return "*  is"
	case NoAIsB:
		return "*  is"
	case AIsT:
		return "+  is"
	case AIsNotT:
		return "+  is not"
	case 6: // TODO: identity
		return "+  = "
	case 7: // TODO: not equal identity (meant to be slash equals)
		return "+   = / = "
	default:
		return ""
	}
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
