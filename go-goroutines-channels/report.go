package main

import (
	"fmt"
	"sync"
)

const (
	good   = 2
	medium = 1
	bad    = 0
)

type Report struct {
	Name   string
	Status int
}

func sumGood(reports []Report, wg *sync.WaitGroup, i int, partSize int, resultChannel chan int) {
	defer wg.Done() //goroutine feita

	start, end := sumEndAndStart(i, partSize)

	resultChannel <- count(start, end, good, reports)
}

func sumMedium(reports []Report, wg *sync.WaitGroup, i int, partSize int, resultChannel chan int) {
	defer wg.Done() //goroutine feita

	start, end := sumEndAndStart(i, partSize)

	resultChannel <- count(start, end, medium, reports)
}

func count(start int, end int, status int, reports []Report) int {
	sum := 0
	for j := start; j < end; j++ {
		if reports[j].Status == status {
			sum += 1
		}
	}

	return sum
}

func sumBad(reports []Report, wg *sync.WaitGroup, i int, partSize int, resultChannel chan int) {
	defer wg.Done() //goroutine feita

	start, end := sumEndAndStart(i, partSize)

	resultChannel <- count(start, end, bad, reports)
}

func sumEndAndStart(i int, partSize int) (int, int) {
	return i * partSize, (i + 1) * partSize
}

func getSum(channel chan int) int {
	total := 0
	for partialSum := range channel {
		total += partialSum
	}

	return total
}

func main() {
	reports := []Report{{Name: "R1", Status: 0},
		{Name: "R2", Status: 1},
		{Name: "R3", Status: 2},
		{Name: "R4", Status: 0}}

	numGoroutines := 2

	goodChannel := make(chan int, numGoroutines)
	mediumChannel := make(chan int, numGoroutines)
	badChannel := make(chan int, numGoroutines)

	partSize := len(reports) / numGoroutines

	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(3) //adiciona 3 ao contador do wait para iniciar que estamos usando 3 goroutines

		go sumGood(reports, &wg, i, partSize, goodChannel)
		go sumMedium(reports, &wg, i, partSize, mediumChannel)
		go sumBad(reports, &wg, i, partSize, badChannel)
	}

	//aguarda conclusao de goroutines
	wg.Wait()

	close(goodChannel)
	close(mediumChannel)
	close(badChannel)

	fmt.Printf("Good: %d\n", getSum(goodChannel))
	fmt.Printf("Medium: %d\n", getSum(mediumChannel))
	fmt.Printf("Bad: %d\n", getSum(badChannel))
}
