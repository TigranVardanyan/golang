package main

import (
	"fmt"
	"time"
)
import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	work := func(id int) {
		defer wg.Done()
		fmt.Println("Goroutine", id,  "run")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine", id, "end")
	}

	go work(1)
	go work(2)

	wg.Wait()        // wait for end all group goroutines
	fmt.Println("Горутины завершили выполнение")
}
