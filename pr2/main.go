package main

import (
	"fmt"
	"sync"
)

/*
	Напиши программу, которая запускает 10 горутин, каждая из которых увеличивает общее значение переменной-счётчика.
	Используй sync.Mutex, чтобы избежать гонки.
	Добавь sync.WaitGroup, чтобы программа корректно дождалась завершения всех горутин.
	Добавь defer для Unlock внутри каждой горутины.
*/

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(10)
	for i := range 10 {
		go IncreaseCounter(i, &wg, &mu)
	}

	wg.Wait()
}

func IncreaseCounter(counter int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	counter += 1
	mu.Unlock()
	fmt.Println("counter: ", counter)
}
