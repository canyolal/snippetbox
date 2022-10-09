package validator

import (
	"strings"
	"unicode/utf8"
)

// Defina a new Validator type which ocntains a map of validation errors for our
// form fields.

type Validator struct {
	FieldErrors map[string]string
}

// Valid() returns true if the FieldErrors map doesn't contain any errors.
func (v *Validator) Valid() bool {
	return len(v.FieldErrors) == 0
}

// AddFieldErrors() adds an error message to the FieldErrors map (so long as no
// entry already exists for the given key.)
func (v *Validator) AddFieldErrors(key, message string) {
	if v.FieldErrors == nil {
		v.FieldErrors = make(map[string]string)
	}
	if _, exists := v.FieldErrors[key]; !exists {
		v.FieldErrors[key] = message
	}
}

// CheckField() adds an error message to the FieldErrors map only if a
// validation check is not "ok".
func (v *Validator) CheckField(ok bool, key, message string) {
	if !ok {
		v.AddFieldErrors(key, message)
	}
}

// NotBlank() returns true if a value is not an empty string.
func (v *Validator) NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// MaxChars() returns true if a value contains no more than n chararacters.
func (v *Validator) MaxChars(value string, n int) bool {
	return utf8.RuneCountInString(value) <= n
}

// PermittedInt() returns true if a value is in a list of permitted integers.
func (v *Validator) PermittedInt(value int, permittedValues ...int) bool {
	for i := range permittedValues {
		if permittedValues[i] == value {
			return true
		}
	}
	return false
}
