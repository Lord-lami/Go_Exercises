package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function

// DivideFood takes the a fodderCalculator and the number of cows as an integer
// It returns the amount of fodder needed for each cow and an error from the fodderCalculator or nil.
func DivideFood(fodderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	totalFodderAmount, err1 := fodderCalculator.FodderAmount(numberOfCows)
	if err1 != nil {return 0, err1}

	fatteningFactor, err2 := fodderCalculator.FatteningFactor()
	if err2 != nil {return 0, err2}

	return (totalFodderAmount * fatteningFactor) / float64(numberOfCows), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function

// ValidateInputAndDivideFood takes the fodderCalculator and the number of cows as an integer
// It returns the amount of food needed and an error if the number of cows is below 1 or nil if otherwise.
func ValidateInputAndDivideFood(fooderCalculator FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows > 0 {return DivideFood(fooderCalculator, numberOfCows)}
	return 0, errors.New("invalid number of cows")
}

// InvalidCowsError describes the error of invalid number of cows better.
type InvalidCowsError struct {
	nbOfCows int
	message string
}

// Error method of InvalidCowsError prints the invalid cow number and error message
func (ice *InvalidCowsError) Error() string {
	return fmt.Sprintf("%v cows are invalid: %s", ice.nbOfCows, ice.message)
}

// TODO: define the 'ValidateNumberOfCows' function

func ValidateNumberOfCows(numberOfCows int) error {
	switch {
	case numberOfCows < 0:
		return &InvalidCowsError{numberOfCows, "there are no negative cows"}
	case numberOfCows == 0:
		return &InvalidCowsError{0, "no cows don't need food"}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
