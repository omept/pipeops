package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	defaultNumberOfIntegers = 10000
)

// Generator function to produce random integers
func generator(num int, out chan<- int) {
	defer close(out)
	for i := 0; i < num; i++ {
		randNum := rand.Intn(100) + 1
		out <- rand.Intn(100) + randNum
	}
}

// Square function to square the integers
func square(in <-chan int, out chan<- int) {
	for num := range in {
		randNumSqr := num * num
		fmt.Printf("Squared Random Number(%d): %d \n", num, randNumSqr)
		out <- randNumSqr
	}
	close(out)
}

// Sum function to sum the squared integers
func sum(in <-chan int, out chan<- int) {
	total := 0
	for num := range in {
		total += num
	}
	out <- total
}

func main() {
	start := time.Now()

	// Parse command-line flag
	num := flag.Int("n", defaultNumberOfIntegers, "Number of random integers to generate")
	flag.Parse()
	fmt.Printf("Random numbers to generate: %d \n", *num)
	// Channels
	genChan := make(chan int, *num)
	squareChan := make(chan int, *num)
	totalChan := make(chan int)
	var total int

	// Start pipeline stages
	go generator(*num, genChan)

	go square(genChan, squareChan)

	go sum(squareChan, totalChan)

	total = <-totalChan

	fmt.Printf("Sum of squared numbers: %d\n", total)
	fmt.Printf("Time taken: %v\n", time.Since(start))
}
