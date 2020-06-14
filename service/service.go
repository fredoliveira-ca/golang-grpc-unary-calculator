package service

import "errors"

// Calculate the result of an operation
func Calculate(fisrtNumber, secondNumber float32, operation int) (float32, error) {
	switch operation {
	case 1:
		return sum(fisrtNumber, secondNumber)
	case 2:
		return subtract(fisrtNumber, secondNumber)
	case 3:
		return multiply(fisrtNumber, secondNumber)
	case 4:
		return divide(fisrtNumber, secondNumber)
	default:
		return 0, nil
	}
}

func sum(fisrtNumber, secondNumber float32) (float32, error) {
	return fisrtNumber + secondNumber, nil
}

func subtract(fisrtNumber, secondNumber float32) (float32, error) {
	return fisrtNumber - secondNumber, nil
}

func multiply(fisrtNumber, secondNumber float32) (float32, error) {
	return fisrtNumber * secondNumber, nil
}

func divide(fisrtNumber, secondNumber float32) (float32, error) {
	var err error
	
	if secondNumber == 0 {
		err = errors.New("Could'n divide by zero")
	}

	return fisrtNumber / secondNumber, err
}
