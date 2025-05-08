package snakecase

type some_type struct{} // want "snake_case name 'some_type', use MixedCaps or mixedCaps rather than underscore when writing multi-word names."

type SomeType struct {
	some_field string // want "snake_case name 'some_field', use MixedCaps or mixedCaps rather than underscore when writing multi-word names."

	goodField int
}

func some_function() {} // want "snake_case name 'some_function', use MixedCaps or mixedCaps rather than underscore when writing multi-word names."

func some_Function() {} // want "snake_case name 'some_Function', use MixedCaps or mixedCaps rather than underscore when writing multi-word names."

func goodFunction() {}

func normalFunction() {
	some_var := 1 // want "snake_case name 'some_var', use MixedCaps or mixedCaps rather than underscore when writing multi-word names."

	goodVar := 2

	// Test single letter names (these should be allowed)
	a := 3

	// Use variables to avoid unused variable errors
	_, _, _ = some_var, goodVar, a
}

// Special cases that should be allowed
func _privatePrefix() {}

func t() {} // Single letter function

func goodPrintFunction(format string, args ...interface{}) {}
