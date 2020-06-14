package tests

import (
	"testing"

	"github.com/fredoliveira-ca/golang-grpc/calculator/service"
)

func TestSumOnePlusOne(t *testing.T) {
	sumOperator := 1
	actual, _ := service.Calculate(1, 1, sumOperator)
	expected := float32(2.0)

	if expected != actual {
		t.Errorf("Calculate(1, 1, 1) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestSumOnePlusTwo(t *testing.T) {
	sumOperator := 1
	actual, _ := service.Calculate(1, 2, sumOperator)
	expected := float32(3.0)

	if expected != actual {
		t.Errorf("Calculate(1, 2, 1) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestSubtractionOneMinusTwo(t *testing.T) {
	subtractOperator := 2
	actual, _ := service.Calculate(1, 2, subtractOperator)
	expected := float32(-1.0)

	if expected != actual {
		t.Errorf("Calculate(1, 2, 2) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestSubtractionFourMinusTwo(t *testing.T) {
	subtractOperator := 2
	actual, _ := service.Calculate(4, 2, subtractOperator)
	expected := float32(2.0)

	if expected != actual {
		t.Errorf("Calculate(4, 2, 2) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestMultilication4567TimesTwo(t *testing.T) {
	multiplicationOperator := 3
	actual, _ := service.Calculate(4567, 2, multiplicationOperator)
	expected := float32(9134.0)

	if expected != actual {
		t.Errorf("Calculate(4, 2, 2) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestFourDividedByTwo(t *testing.T) {
	divideOperator := 4
	actual, _ := service.Calculate(4, 2, divideOperator)
	expected := float32(2.0)
	if expected != actual {
		t.Errorf("Calculate(4, 2, 4) \n actual: %v \n expected: %v", actual, expected)
	}
}

func TestFourDividedByZeroCauseError(t *testing.T) {
	divideOperator := 4
	_, err := service.Calculate(4, 0, divideOperator)

	if err == nil {
		t.Errorf("It should return an error!")
	}
}

func TestInexistentOperator(t *testing.T) {
	inexistentOperator := 5
	actual, _ := service.Calculate(4, 2, inexistentOperator)
	expected := float32(0.0)
	if expected != actual {
		t.Errorf("Calculate(4, 2, 4) \n actual: %v \n expected: %v", actual, expected)
	}
}
