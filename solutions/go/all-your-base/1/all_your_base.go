// Package allyourbase provides a tools for number base conversion.
package allyourbase

import (
	"errors"
)

// pow takes an int nbr and multiplies it by itself exp times.
func pow(nbr, exp int) int {
	result := 1
	for ; exp > 0; exp /= 2 {
		if exp%2 == 1 {
			result *= nbr
		}
		nbr *= nbr
	}
	return result
}

// convertToDecimal takes an inputDigits int slice and an inputBase integer
// it returns an integer of the number represented by inputDigits as an integer.
func convertToDecimal(inputDigits []int, inputBase int) int {
	inputNbr := 0
	for i, digit := range inputDigits {
		exp := len(inputDigits) - (i + 1)
		inputNbr += digit * pow(inputBase, exp)
	}
	return inputNbr
}

// convertFromDecimal takes a decimal(base 10 number) and an outputBase integer
// it converts the number to a left-right slice of digits in the prescribed
// outputBase.
func convertFromDecimal(decimal int, outputBase int) (outputXgits []int) {
	if decimal == 0 {
		outputXgits = []int{0}
	}
	for ; decimal > 0; decimal /= outputBase {
		outputXgits = append([]int{decimal % outputBase}, outputXgits...)
	}
	return
}

// ConvertToBase takes an inputBase, inputDigits and outputBase
// It converts the number represented by inputDigits from its inputBase to the outputBase
// It returns a non-nil error if the inputBase or outputBase is less than 2 or
// if the inputDigits are not a valid numbers in the inputBase.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}
	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	decimal := convertToDecimal(inputDigits, inputBase)
	return convertFromDecimal(decimal, outputBase), nil
}
