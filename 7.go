package main

import (
	"fmt"
	"sync"
	"time"
)

// func wait(wg *sync.WaitGroup) {

// 	fmt.Printf("wait")
// }
// func wait(wg *sync.WaitGroup){

// }
func work1(waitGroup *sync.WaitGroup) {
	fmt.Println("work1")
	waitGroup.Done()
}

func work2(waitGroup *sync.WaitGroup) {
	fmt.Println("work2")
	waitGroup.Done()
}

func main() {
	//var wag sync.WaitGroup
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	//wag.Add(2)
	// work := func(id int) {
	// 	defer wag.Done()
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Printf("Рутина %d\n", id)
	// }
	wg1.Add(1)
	wg2.Add(1)
	go work1(&wg1)
	go work2(&wg2)
	wg1.Wait()
	wg2.Wait()

	//wag.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("finish")
}
