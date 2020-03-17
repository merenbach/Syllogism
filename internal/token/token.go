package token

// A Type of token.
type Type int

const (
	// TypeReserved denotes a reserved term.
	TypeReserved Type = iota

	// TypeLineNumber denotes a line number.
	TypeLineNumber

	// TypeSlash denotes a slash.
	TypeSlash

	// TypeQuantifier denotes a quantifier.
	TypeQuantifier

	// TypeNegation denotes a negation (i.e., no/not).
	TypeNegation

	// TypeCopula denotes a copula (i.e., is/are).
	TypeCopula

	// TypeTerm denotes a term.
	TypeTerm
)
