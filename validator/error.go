package validator

import (
	"fmt"

	"github.com/vektah/gqlparser/errors"
)

type ErrorOption func(err *errors.Validation)

func Message(msg string, args ...interface{}) ErrorOption {
	return func(err *errors.Validation) {
		err.Message += fmt.Sprintf(msg, args...)
	}
}

func SuggestListQuoted(prefix string, typed string, suggestions []string) ErrorOption {
	suggested := SuggestionList(typed, suggestions)
	return func(err *errors.Validation) {
		if len(suggested) > 0 {
			err.Message += " " + prefix + " " + QuotedOrList(suggested...) + "?"
		}
	}
}

func SuggestListUnquoted(prefix string, typed string, suggestions []string) ErrorOption {
	suggested := SuggestionList(typed, suggestions)
	return func(err *errors.Validation) {
		if len(suggested) > 0 {
			err.Message += " " + prefix + " " + OrList(suggested...) + "?"
		}
	}
}

func Suggestf(suggestion string, args ...interface{}) ErrorOption {
	return func(err *errors.Validation) {
		err.Message += " Did you mean " + fmt.Sprintf(suggestion, args...) + "?"
	}
}