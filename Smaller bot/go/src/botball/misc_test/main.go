package main

import (
	"fmt"
	"math"
	wb "wallaby"
)

/*
	Überdrehungsfunktion: y = 0.0837x - 5.3568
	Achung! wwerte unter 100 sind unter 0
	Und Freeze verwenden
*/

func main() {
	/*wb.Cmpc(0)
	//wb.Mav(0, 1500)
	wb.Motor_power(0, 100)
	prev := 0
	start := time.Now()
	for {
		wb.Msleep(100)
		curr := wb.Gmpc(0)
		delta := time.Since(start)
		start = time.Now()
		fmt.Printf("%v;%v;%v\n", curr, (curr-prev)*(1000/(int(delta/time.Millisecond))), delta/time.Millisecond)
		if (curr-prev)*(1000/(int(delta/time.Millisecond))) >= 1500 {
			break
		}
		prev = curr
	}
	fmt.Println()
	start = time.Now()
	start2 := time.Now()
	for time.Since(start2) < time.Second {
		wb.Msleep(100)
		curr := wb.Gmpc(0)
		delta := time.Since(start)
		start = time.Now()
		fmt.Printf("%v;%v;%v\n", curr, (curr-prev)*(1000/(int(delta/time.Millisecond))), delta/time.Millisecond)
		prev = curr
	}
	fmt.Println()
	wb.Motor_power(0, 0)
	//wb.Mav(0, 0)
	prev = wb.Gmpc(0)
	start = time.Now()
	for {
		wb.Msleep(100)
		curr := wb.Gmpc(0)
		delta := time.Since(start)
		start = time.Now()
		fmt.Printf("%v;%v;%v\n", curr, (curr-prev)*(1000/(int(delta/time.Millisecond))), delta/time.Millisecond)
		if (curr-prev)*(1000/(int(delta/time.Millisecond))) == 0 {
			break
		}
		prev = curr
	}
	wb.Freeze(0)
	wb.Msleep(500)
	wb.Ao()*/
	/*for v := 100; v <= 1500; v += 100 {
		wb.Cmpc(0)
		wb.Mav(0, v)
		for wb.Gmpc(0) <= 1000 {
		}
		stopPos := wb.Gmpc(0)
		wb.Mav(0, 0)
		wb.Msleep(750)
		endPos := wb.Gmpc(0)
		fmt.Printf("%v;%v;%v;%v\n", v, endPos-stopPos, stopPos, endPos)
	}
	fmt.Println()*/
	for v := 100; v <= 1500; v += 100 {
		wb.Cmpc(0)
		wb.Mav(0, v)
		for float32(wb.Gmpc(0)) <= float32(3000)-float32(math.Max(float64(v)*0.0837-5.3568, 0)) {
		}
		stopPos := wb.Gmpc(0)
		wb.Freeze(0)
		wb.Msleep(1000)
		endPos := wb.Gmpc(0)
		fmt.Printf("%v;%v;%v;%v\n", v, endPos-stopPos, stopPos, endPos)
	}
	wb.Ao()
}
