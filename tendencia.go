package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mediana(array []float64) float64 {
	arr := array
	if len(arr)%2 == 0 {
		if len(arr) != 2 {
			arr = arr[:len(arr)]
			arr = arr[1:]
			fmt.Println(arr)
			mediana(arr)
		} else {
			return media(arr)
		}
	} else {
		if len(arr) != 1 {
			arr = arr[:len(arr)]
			arr = arr[1:]
			mediana(arr)
			fmt.Println(arr)
		} else {
			return media(arr)
		}

	}
	return arr[0]
}
func media(array []float64) float64 {
	var sumas float64 = 0
	var result float64 = 0
	for i := 0; i < len(array); i++ {
		sumas = sumas + array[i]
	}
	result = sumas / float64(len(array))
	return float64(result)
}

func converter(s string) []float64 {
	array := []float64{}
	m := strings.Split(s, ",") //split the numbers
	for i := 0; i < len(m); i++ {
		input, _ := strconv.ParseFloat(m[i], 64) /*print then of convert into a float64*/
		array = append(array, input)
	}
	sort.Float64s(array)
	return array
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)
	m := converter(text) //convert into a float64 aarray
	n := media(m)
	l := mediana(m)
	fmt.Println("la mediana es \n", l)
	fmt.Println("la media es  \n", n)
	fmt.Println(m)
}
