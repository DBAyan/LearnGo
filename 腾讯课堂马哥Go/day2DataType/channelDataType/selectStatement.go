package main

import (
	"fmt"
	"time"
)

func main() {

	healthTick1 := time.Tick(1 * time.Second)
	healthTick2 := time.Tick(2 * time.Second)
	healthTick3 := time.Tick(3 * time.Second)

	for {
		select {
		//case <-healthTick1:
		//	fmt.Println("healthTick1 is exec!")
		//case <-healthTick2:
		//	fmt.Println("healthTick2 is exec!")
		//case <-healthTick3:
		//	fmt.Println("healthTick3 is exec!")

		case <- healthTick1:
			go func() {
				fmt.Println("healthTick1 is exec!")
			}()
		case <- healthTick2:
			go func() {
				fmt.Println("healthTick2 is exec!")
			}()

		case <- healthTick3:
			go func() {
				fmt.Println("healthTick3 is exec!")
			}()
		}
		fmt.Println("----")
	}


}



