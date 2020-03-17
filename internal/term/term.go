package term

// A Type is either a general term or a designator.
type Type int

// TODO: should be able to swap ordering of any of these and retain same functionality once refactor is complete
//       since we won't be relying on integer values anymore
const (
	// UndeterminedType represents an undetermined type.
	UndeterminedType Type = iota

	// GeneralTermType represents a general term.
	GeneralTermType

	// DesignatorType represents a designator.
	DesignatorType
)

func (t Type) String() string {
	switch t {
	case GeneralTermType:
		return "general term"
	case DesignatorType:
		return "designator"
	default:
		return "undetermined type"
	}
}
