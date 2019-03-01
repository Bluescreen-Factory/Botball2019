package main

import (
	"fmt"
	"math"
	"time"
	wb "wallaby"
)

func main() {
	for wb.A_button() == 1 {
	}
	distance := 1.
	for {
		fmt.Println("Configure distance. left/right A")
		for {
			if wb.Right_button() == 1 {
				distance += 0.1
				fmt.Printf("Dist: %.1f rev\n", distance)
			}
			if wb.Left_button() == 1 {
				distance -= 0.1
				fmt.Printf("Dist: %.1f rev\n", distance)
			}
			if wb.A_button() == 1 {
				break
			}
			wb.Msleep(100)
		}
		distance = math.Round(distance*10) / 10
		tpr0 := 1898.
		tpr1 := 2045.
		wb.Cmpc(0)
		wb.Cmpc(1)
		wb.Mtp(0, int(1500.*(tpr0/tpr1)), int(tpr0*distance))
		wb.Mtp(1, 1500, int(tpr1*distance))
		var done0 time.Time
		ch0 := make(chan bool)
		go func() {
			wb.Bmd(0)
			wb.Freeze(0)
			if wb.Get_motor_done(1) == 0 {
				wb.Freeze(1)
			}
			done0 = time.Now()
			ch0 <- true
		}()
		ch1 := make(chan bool)
		var done1 time.Time
		go func() {
			wb.Bmd(1)
			wb.Freeze(1)
			if wb.Get_motor_done(0) == 0 {
				wb.Freeze(0)
			}
			done1 = time.Now()
			ch1 <- true
		}()
		<-ch0
		<-ch1
		fmt.Printf("0: %v/%v 1: %v/%v\n", wb.Gmpc(0), int(tpr0*distance), wb.Gmpc(1), int(tpr1*distance))
		fmt.Printf("delta 0-1: %v (ms)\n", int64(done0.Sub(done1)/time.Millisecond))
		fmt.Println("Stop (left), Again (right)")
		for {
			if wb.Right_button() == 1 {
				break
			} else if wb.Left_button() == 1 {
				return
			}
		}
	}
}
