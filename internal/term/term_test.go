package term

import "fmt"

func ExampleUndeterminedType() {
	fmt.Println(UndeterminedType)
	// Output: undetermined type
}

func ExampleGeneralTermType() {
	fmt.Println(GeneralTermType)
	// Output: general term
}

func ExampleDesignatorType() {
	fmt.Println(DesignatorType)
	// Output: designator
}
