package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var randChan = make(chan int)
	var averageChan = make(chan float64)
	var nums []int

	for i := 0; i <= 100; i++ {
		go func() {
			randChan <- rand.Intn(100)
		}()

		go func() {
			var sum int

			for num := range randChan {
				nums = append(nums, num)
				for _, val := range nums {
					sum += val
				}
				averageChan <- float64(sum) / float64(len(nums))
			}
		}()

		go func() {
			fmt.Printf("\nСереднє значення: %.3f", <-averageChan)
		}()
	}
}
