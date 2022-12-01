//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
	Second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func parseTime(input string) (Time, error) {

	parts := strings.Split(input, ":")

	if len(parts) != 3 {
		return Time{}, &TimeParseError{"invalid time format", input}
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid hour input: %v", err), input}
	}

	if hour < 0 || hour > 24 {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid hour input: %v", err), input}
	}

	minutes, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid hour input: %v", err), input}
	}

	if minutes < 0 || minutes > 59 {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid minutes input: %v", err), input}
	}

	seconds, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid seconds input: %v", err), input}
	}

	if seconds < 0 || seconds > 59 {
		return Time{}, &TimeParseError{fmt.Sprintf("invalid seconds input: %v", err), input}
	}

	return Time{hour, minutes, seconds}, nil

}
