package main

import (
	"errors"
	"fmt"
	"strconv"
)

func PesquisaBinaria(array []int, item int) (int, error) {
	baixo := 0
	alto := len(array) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := array[meio]

		if chute == item {
			return meio, nil
		}
		if chute > item {
			alto = meio - 1
		}
		if chute < item {
			baixo = meio + 1
		}
	}
	return -1,errors.ErrUnsupported
}

func main() {
	array := [5]int{10, 20, 30, 70, 90}
	var item = 10

	resultado, err := PesquisaBinaria(array[:], item)

	if err != nil {
		panic("erro, nao encontrado")
	}

	fmt.Println(strconv.Itoa(resultado))
}