package model

import validation "github.com/go-ozzo/ozzo-validation"

// if condition is true validation is required, otherwise it's not
func RequiredIf(cond bool) validation.RuleFunc {
	return func (value interface{}) error {
		if cond {
			return validation.Validate(
				value, validation.Required,
			)
		}

		return nil
	}
} 