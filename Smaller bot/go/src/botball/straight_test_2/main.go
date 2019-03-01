package main

import (
	"fmt"
	"sync"
	wb "wallaby"
)

const (
	tprR   = 1898
	tprL   = 2045
	speedK = 1.0825
)

func main() {
	for {
		fmt.Println("Exit: Left, Continue: Right")
		for wb.Right_button() == 0 && wb.Left_button() == 0 {
		}
		if wb.Left_button() == 1 {
			return
		}
		wb.Cmpc(0)
		wb.Cmpc(1)
		revs := 10
		spdR := 1000
		spdL := int(float32(spdR) * speedK)
		fmt.Printf("Speeds (L, R): %v %v\n", spdL, spdR)
		//y = 0.0837x - 5.3568
		ticksR := int(float32(tprR*revs) - (float32(spdR)*0.0837 - 5.53568))
		ticksL := int(float32(tprL*revs) - (float32(spdL)*0.0837 - 5.53568))
		fmt.Printf("Ticks (L, R): %v %v\n", ticksL, ticksR)
		wb.Mtp(0, spdR, ticksR)
		wb.Mtp(1, spdL, ticksL)
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			//wb.Bmd(0)
			for wb.Gmpc(0) <= ticksR {
			}
			fmt.Printf("Right before freeze: %v\n", wb.Gmpc(0))
			wb.Freeze(0)
			fmt.Printf("Right done: %v = %v revs\n", wb.Gmpc(0), float32(wb.Gmpc(0))/float32(tprR))
			wg.Done()
		}()
		go func() {
			/*if wb.Get_motor_done(1) == 1 {
				wb.Bmd(1)
			}*/
			for wb.Gmpc(1) <= ticksL {
			}
			fmt.Printf("Left before freeze: %v\n", wb.Gmpc(1))
			wb.Freeze(1)
			fmt.Printf("Left done: %v = %v revs\n", wb.Gmpc(1), float32(wb.Gmpc(1))/float32(tprL))
			wg.Done()
		}()
		wg.Wait()
		wb.Msleep(500)
		wb.Ao()
	}
}
