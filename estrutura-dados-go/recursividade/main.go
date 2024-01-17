package main

import (
	"fmt"
)

func Soma(array []int, item int) int {
	if len(array) == 1 {
		return item
	}
	newarray := array[1:]
	return item + Soma(newarray, newarray[0])
}

func Fatorial(x int) int{
	if x == 1 {
		return x
	}
	return x * Fatorial(x - 1)
}

func InverterString(array []string) (string){
	if len(array) == 1 {
		return array[0]
	}
	qtd := len(array) - 1
	str := array[len(array) - 1]
	newarray := append(array[:qtd], array[qtd+1:]...)
	return str + InverterString(newarray)
}

func TamanhoLista(array []int) int {
	if len(array) == 1 {
		return 1
	}
	return 1 + TamanhoLista(array[1:])
}

func AcharMaior(array []int) int {
	if len(array) == 1 {
		return array[0]
	}
	if array[0] >= array[1] {
		return AcharMaior(append(array[:1], array[2:]...))
	}
	return AcharMaior(array[1:])
}


func main() {
	array := []int{1000, 20, 130, 70, 90, 40}

	//str := "joao"
	//chars := strings.Split(str, "")
	
	//x := Soma(array[:], array[0])

	//fmt.Println(x)
	//a := InverterString(chars)
	//fmt.Println(a)

	//fmt.Println(TamanhoLista(array[:]))

	fmt.Println(AcharMaior(array[:]))

	//fmt.Println(Fatorial(6))

	//fmt.Println(nCr(6,4))
}

func nCr(n, r int) int {
	if r == 0 || n == r {
		return 1
	}

	return Fatorial(n) / (Fatorial(r) * Fatorial(n-r))
}