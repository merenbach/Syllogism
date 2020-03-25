package article

// A Type of article.
type Type int

// TODO: should be able to swap ordering of any of these and retain same functionality once refactor is complete
//       since we won't be relying on integer values anymore
const (
	// TypeNone represents no specific type.
	// TODO: is there a better way to name TypeNone?
	TypeNone Type = iota

	// TypeA represents a type preceded by the article "a."
	TypeA

	// TypeAn represents a type preceded by the article "an."
	TypeAn

	// TypeSm represents some quantity.
	TypeSm
)

const (
	// WordA represents the word "a."
	WordA = "a"

	// WordAn represents the word "an."
	WordAn = "an"

	// WordSm represents the word "sm."
	WordSm = "sm"

	// WordThe represents the word "the."
	WordThe = "the"
)

// TypeFromString returns an article type based on the provided string.
func TypeFromString(s string) Type {
	switch s {
	case WordA:
		return TypeA
	case WordAn:
		return TypeAn
	case WordSm:
		return TypeSm
	default:
		return TypeNone
	}
}

func (t Type) String() string {
	switch t {
	case TypeA:
		return WordA + " "
	case TypeAn:
		return WordAn + " "
	case TypeSm:
		return WordSm + " "
	default:
		return ""
	}
}
