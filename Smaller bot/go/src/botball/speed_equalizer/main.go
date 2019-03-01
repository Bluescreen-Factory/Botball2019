package main

import (
	"fmt"
	"time"
	wb "wallaby"
)

const revs int = 5

func main() {
	wb.Mav(0, 150)
	for wb.Analog(0) > 1000 {
	}
	for wb.Analog(0) < 1000 {
	}
	wb.Mav(0, 0)
	wb.Msleep(250)
	fmt.Println("0 synced")
	wb.Mav(1, 150)
	for wb.Analog(1) > 1000 {
	}
	for wb.Analog(1) < 1000 {
	}
	wb.Mav(1, 0)
	wb.Msleep(250)
	fmt.Println("1 synced")
	vel0 := 400
	vel1 := 400
	wb.Mav(0, vel0)
	wb.Mav(1, vel1)
	ch0 := make(chan int)
	ch1 := make(chan int)
	firstFirst := -1
	for {
		wb.Cmpc(0)
		wb.Cmpc(1)
		stop := false
		go func() {
			for i := 0; i < 5*revs && !stop; i++ {
				for wb.Analog(0) > 1000 && !stop {
				}
				for wb.Analog(0) < 1000 && !stop {
				}
			}
			ch0 <- 0
		}()

		go func() {
			for i := 0; i < 5*revs && !stop; i++ {
				for wb.Analog(1) > 1000 && !stop {
				}
				for wb.Analog(1) < 1000 && !stop {
				}
			}
			ch1 <- 1
		}()
		fmt.Println("Selecting")
		var first int
		select {
		case first = <-ch0:
			break
		case first = <-ch1:
			break
		}

		stop = true

		fmt.Printf("First: %v\n", first)

		if firstFirst == -1 {
			firstFirst = first
		}

		if firstFirst == first {
			if first == 0 {
				vel1 += 5
				wb.Mav(1, vel1)
			} else {
				vel0 += 5
				wb.Mav(0, vel0)
			}
		}
		fmt.Printf("Ticks: %v %v\n", wb.Gmpc(0), wb.Gmpc(1))
		fmt.Printf("vels: %v %v\n", vel0, vel1)
		fmt.Printf("Waiting for: %v\n", 1-first)
		if first == 0 {
			<-ch1
		} else {
			<-ch0
		}
		if firstFirst != first {
			break
		}
	}
	wb.Ao()
	fmt.Println("DONE")
	fmt.Printf("vels: %v %v\n", vel0, vel1)
	wb.Msleep(500)
	fmt.Println("Checking results...")
	wb.Mav(0, vel0)
	wb.Mav(1, vel1)
	stop := false
	revs0 := 0.
	go func() {
		for !stop {
			for wb.Analog(0) > 1000 && !stop {
			}
			fmt.Println("0w")
			revs0 += 0.1
			for wb.Analog(0) < 1000 && !stop {
			}
			fmt.Println("0b")
			revs0 += 0.1
		}
	}()
	revs1 := 0.
	go func() {
		for !stop {
			for wb.Analog(1) > 1000 && !stop {
			}
			fmt.Println("1w")
			revs1 += 0.1
			for wb.Analog(1) < 1000 && !stop {
			}
			fmt.Println("1b")
			revs1 += 0.1
		}
	}()
	<-time.After(60 * time.Second)
	stop = true
	fmt.Printf("Revs: %v %v\n", revs0, revs1)
	wb.Ao()
	fmt.Println("DONE")

}
