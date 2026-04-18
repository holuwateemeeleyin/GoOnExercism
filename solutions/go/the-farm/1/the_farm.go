package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood handles the core calculation logic.
func DivideFood(fc FodderCalculator, cows int) (float64, error) {
	amount, err := fc.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}

	// Logic: (Total Fodder * Factor) / Number of Cows
	return (amount * factor) / float64(cows), nil
}

// ValidateInputAndDivideFood adds a simple check before calling DivideFood.
func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, cows)
}

// InvalidCowsError represents a custom error for cow validation logic.
type InvalidCowsError struct {
	cows    int
	message string
}

// Error formats the InvalidCowsError to match the required string format.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows performs specific checks and returns the custom error.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}
	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}
	return nil
}