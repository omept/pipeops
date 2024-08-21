package main

import (
	"os"
	"testing"
	"time"
)

// TestGenerator verifies that the generator produces the expected number of random integers.
func TestGenerator(t *testing.T) {
	num := 10
	out := make(chan int, num)
	expectedCount := num

	go generator(num, out)

	count := 0
	for range out {
		count++
	}

	if count != expectedCount {
		t.Errorf("Expected %d numbers, got %d", expectedCount, count)
	}
}

// TestSquare verifies that the square function correctly squares the input integers.
func TestSquare(t *testing.T) {
	in := make(chan int, 5)
	out := make(chan int, 5)

	numbers := []int{2, 3, 4, 5}
	expectedSquares := []int{4, 9, 16, 25}

	for _, n := range numbers {
		in <- n
	}
	close(in)

	go square(in, out)

	i := 0
	for result := range out {
		if result != expectedSquares[i] {
			t.Errorf("Expected square of %d to be %d, but got %d", numbers[i], expectedSquares[i], result)
		}
		i++
	}

	if i != len(expectedSquares) {
		t.Errorf("Expected %d squared numbers, but got %d", len(expectedSquares), i)
	}
}

// TestSum verifies that the sum function correctly sums the input integers.
func TestSum(t *testing.T) {
	in := make(chan int, 4)
	out := make(chan int)

	numbers := []int{4, 9, 16, 25}
	expectedSum := 54

	go func() {
		for _, n := range numbers {
			in <- n
		}
		close(in)
	}()

	go sum(in, out)

	result := <-out

	if result != expectedSum {
		t.Errorf("Expected sum to be %d, but got %d", expectedSum, result)
	}
}

// TestPipelineIntegration verifies the integration of the entire pipeline.
func TestPipelineIntegration(t *testing.T) {
	num := 5
	genChan := make(chan int, num)
	squareChan := make(chan int, num)
	totalChan := make(chan int)

	go generator(num, genChan)
	go square(genChan, squareChan)
	go sum(squareChan, totalChan)

	// Wait for the total to be calculated
	total := <-totalChan

	// There's no predefined expected total because random numbers are generated,
	// but we ensure the pipeline completes without deadlock.
	if total <= 0 {
		t.Error("Expected total to be greater than 0, but got 0 or less")
	}
}

func TestMain(m *testing.M) {
	// Run all tests
	exitVal := m.Run()

	// Exit with the correct status code
	time.Sleep(1 * time.Second)
	os.Exit(exitVal)
}
