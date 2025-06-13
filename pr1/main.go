package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
)

/*
Напиши программу на Go, которая:
Принимает слайс []int, например: []int{23, 567, 4322}.
Каждое число обрабатывается в отдельной горутине:
Конвертируется в строку.
Строка разбивается на отдельные цифры.
Цифры конвертируются обратно в числа (strconv.Atoi).
Суммируются.
Каждая горутина добавляет свой результат в общую сумму (result) — используя sync.Mutex для защиты.
В конце программа выводит общую сумму всех цифр всех чисел. */

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	numbers := []int{23, 647, 2, 44, 56, 22}
	result := 0

	for _, v := range numbers {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			mu.Lock()
			result += sum(v)
			mu.Unlock()
		}(v)
	}

	wg.Wait()

	fmt.Println("Total sum of all single values after split: ", result)

}

func sum(nums int) int {
	sum := 0
	numsStr := strconv.Itoa(nums)
	values := strings.Split(numsStr, "")

	fmt.Printf("Converted: %v\t", values)

	for _, v := range values {
		v, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		sum += v
	}

	fmt.Printf("Each value in slice sum : %d\n", sum)

	return sum
}
