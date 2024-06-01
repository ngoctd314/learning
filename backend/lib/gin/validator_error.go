package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidatorError struct {
	Key       string `json:"key"`
	Value     any    `json:"value"`
	Condition string `json:"condition"`
	Message   string `json:"message"`
}

func NewValidatorError(e validator.FieldError) ValidatorError {
	condition := func() string {
		if param := e.Param(); param != "" {
			return fmt.Sprintf("%s [%s]", e.ActualTag(), param)
		}
		return fmt.Sprintf("%s", e.ActualTag())
	}
	message := func() string {
		m := map[string]string{
			"lte":   "value is too large",
			"lt":    "value is too large",
			"gte":   "value is too small",
			"gt":    "value is too small",
			"oneof": "not in the whitelist",
			"eq":    "value must be equal",
			"ne":    "value must be not equal",
		}
		if v, ok := m[e.Tag()]; ok {
			return v
		}

		return "unknown"
	}

	return ValidatorError{
		Key:       e.Namespace(),
		Value:     e.Value(),
		Condition: condition(),
		Message:   message(),
	}
}
