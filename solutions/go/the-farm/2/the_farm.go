package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	amount, err := calc.FodderAmount(numberOfCows)
	if err != nil {
		return 0, err
	}
	factor, err := calc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return (amount * factor) / float64(numberOfCows), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, numberOfCows int) (float64, error) {
	if numberOfCows < 1 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(calc, numberOfCows)
}

func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows == 0 {
		return &invalidCowsError{"no cows don't need food", numberOfCows}
	} else if numberOfCows < 1 {
		return &invalidCowsError{"there are no negative cows", numberOfCows}
	}
	return nil
}

type invalidCowsError struct {
	s            string
	numberOfCows int
}

func (e *invalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.numberOfCows, e.s)
}
