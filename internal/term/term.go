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

// Other term type, used in an error message.
// TODO: note that we only ever expect non-default cases here.
// TODO: there is probably a more Golang-idiomatic way to accomplish our goal
func (t Type) Other() Type {
	switch t {
	case TypeGeneralTerm:
		return TypeDesignator
	case TypeDesignator:
		return TypeGeneralTerm
	default:
		return TypeUndetermined
	}
}

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
