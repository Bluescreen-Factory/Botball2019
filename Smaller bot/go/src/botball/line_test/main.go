package main

import (
	"fmt"
	wb "wallaby"
)

func main() {
	spd := 800
	wb.Mav(0, spd)
	wb.Mav(1, -spd)
	for wb.Analog(0) < 1500 {
	}
	wb.Mav(0, -spd)
	wb.Mav(1, spd)
	wb.Cmpc(0)
	wb.Cmpc(1)
	for wb.Analog(0) > 1500 {
	}
	for wb.Analog(0) < 1500 {
	}
	wb.Mav(1, 0)
	wb.Mav(0, 0)
	delta0 := wb.Gmpc(0)
	delta1 := wb.Gmpc(1)
	fmt.Printf("Deltas: %v, %v\n", delta0, delta1)
	wb.Mtp(0, spd, delta0/2)
	wb.Mtp(1, spd, delta1/2)
	wb.Bmd(0)
	if wb.Get_motor_done(1) == 0 {
		wb.Bmd(1)
	}
}
