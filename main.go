package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var randChan = make(chan int)
	var averageChan = make(chan float64)
	var nums []int

	go func() {
		for i := 0; i < 3; i++ {
			randChan <- rand.Intn(100)
		}
	}()

	go func() {
		var sum int

		for {
			select {
			case num := <-randChan:
				nums = append(nums, num)

				for _, val := range nums {
					sum += val
				}

				averageChan <- float64(sum) / float64(len(nums))

			default:
				continue
			}
		}
	}()

	go func() {
		for {
			select {
			case av := <-averageChan:
				fmt.Printf("\nСереднє значення: %.3f", av)
			default:
				continue
			}
		}
	}()

	time.Sleep(3 * time.Second)
}
