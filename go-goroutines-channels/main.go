package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string 
	Age int
}

func sumAges(wg *sync.WaitGroup,people []Person, i int, partSize int, resultChannel chan int) {
	defer wg.Done()
	start := i * partSize
	end := (i + 1) * partSize

	sum := 0
	for j := start; j < end; j++ {
		sum += people[j].Age
	}

	resultChannel <- sum
	
}

func main() {
	people := []Person{{Name: "Joao", Age: 10}, {Name: "Gabriel", Age: 23},{Name: "Luis", Age: 49}, {Name: "Gabriela", Age: 73}}

	numGoroutines := 2

	resultChannel := make(chan int, numGoroutines)

	var wg sync.WaitGroup

	partSize := len(people) / numGoroutines

	for i := 0; i < numGoroutines ; i++ {
		wg.Add(1)

		go sumAges(&wg, people, i, partSize, resultChannel)
	}

	wg.Wait()

	close(resultChannel)

	total := 0
	for partialSum := range resultChannel {
		total += partialSum
	}

	fmt.Println(total)
}