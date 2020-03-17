package term

// A Type of term.
type Type int

// TODO: should be able to swap ordering of any of these and retain same functionality once refactor is complete
//       since we won't be relying on integer values anymore
const (
	// TypeUndetermined represents an undetermined type.
	TypeUndetermined Type = iota

	// TypeGeneralTerm represents a general term.
	TypeGeneralTerm

	// TypeDesignator represents a designator.
	TypeDesignator
)

func (t Type) String() string {
	switch t {
	case TypeGeneralTerm:
		return "general term"
	case TypeDesignator:
		return "designator"
	default:
		return "undetermined type"
	}
}
