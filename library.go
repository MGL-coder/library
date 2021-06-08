package library

import (
	"errors"
	"math"
	"strings"
	"unicode"
)

func ChangeCase(str string) string {
	letters := []rune(str)
	if unicode.IsLower(letters[0]) {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}
}

func ToSolveQuadraticFunc(a, b, c float64) (float64, float64, error) {
	d := b*b - 4*a*c
	if d < 0 {
		return 0, 0, errors.New("no solution exists for the given quadratic equation")
	}
	return (-b - math.Sqrt(d)) / (2 * a), (-b + math.Sqrt(d)) / (2 * a), nil
}