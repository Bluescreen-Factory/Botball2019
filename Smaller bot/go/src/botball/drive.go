package botball

import (
	"fmt"
	"math"
	"time"
	wb "wallaby"
)

type Motor struct {
	Port int
	TPR  int //ticks per revolution
}

func newSlpChan() chan interface{} {
	return make(chan interface{}, 1)
}

func clamp(x, min, max float32) float32 {
	return float32(math.Max(math.Min(float64(x), float64(max)), float64(min)))
}

func (m Motor) TurnMs(speed, ms int) <-chan interface{} {
	slp := newSlpChan()
	wb.Motor(m.Port, speed)
	go func() {
		slp <- time.After(time.Duration(ms) * time.Millisecond)
		wb.Motor(m.Port, 0)
	}()
	return slp
}

func (m Motor) TurnExact(tps, ticks int) <-chan interface{} {
	slp := newSlpChan()
	wb.Cmpc(m.Port)
	//19/12
	ticks -= int(clamp((19./12.)*(float32(ticks)-60.), 0, 115))
	wb.Mrp(m.Port, tps, ticks)
	go func() {
		if ticks != 0 && tps != 0 {
			wb.Bmd(m.Port)
		}
		slp <- true
	}()
	return slp
}

func EqualSpeed(m1, m2 Motor, tps int) (tps1, tps2 int) {
	d := (float32(m1.TPR)/float32(m2.TPR))*float32(tps) - float32(tps)
	tps1 = tps - int(d/2)
	tps2 = tps + int(d/2)
	return
}

/*int motor_equalSpeed(struct motor *dest, struct motor *subj, int tps) {
	return (int) (((float) (subj->tpr) / dest->tpr)*tps)
}*/

type PoweredWheel struct {
	Diameter float32
	Motor
}

func (pw PoweredWheel) Perimeter() float32 {
	return pw.Diameter * math.Pi
}

type DifferentialDrive struct {
	LeftWheel  PoweredWheel
	RightWheel PoweredWheel
	TrackWidth float32
}

func (d DifferentialDrive) AlignWithLine(rps float32, thresh int, sensPort int) <-chan interface{} {
	slp := newSlpChan()
	cmpThresh := func(val, thresh int) bool {
		if thresh < 0 {
			if val < -thresh {
				return true
			}
			return false
		} else {
			if val > thresh {
				return true
			}
			return false
		}
	}
	go func() {
		spdR := int(rps * float32(d.RightWheel.TPR))
		spdL := int(rps * float32(d.LeftWheel.TPR))
		wb.Cmpc(d.RightWheel.Port)
		wb.Cmpc(d.LeftWheel.Port)
		wb.Mav(d.RightWheel.Port, spdR)
		wb.Mav(d.LeftWheel.Port, -spdL)
		fmt.Println("Tunring until line")
		for !cmpThresh(wb.Analog(sensPort), thresh) {
		}
		fmt.Println("Line detected")
		wb.Mav(d.RightWheel.Port, -spdR)
		wb.Mav(d.LeftWheel.Port, spdL)
		wb.Msleep(200)
		fmt.Println("Turning back until not line")
		for cmpThresh(wb.Analog(sensPort), thresh) {
		}
		fmt.Println("Line undetected")
		wb.Cmpc(d.RightWheel.Port)
		wb.Cmpc(d.LeftWheel.Port)
		wb.Msleep(200)
		fmt.Println("Turning back until line")
		for !cmpThresh(wb.Analog(sensPort), thresh) {
		}
		fmt.Println("Line detected")
		wb.Mav(d.RightWheel.Port, 0)
		wb.Mav(d.LeftWheel.Port, 0)
		deltaR := wb.Gmpc(d.RightWheel.Port)
		deltaL := wb.Gmpc(d.LeftWheel.Port)
		fmt.Printf("Deltas: %v, %v\n", deltaR, deltaL)
		wb.Msleep(200)
		fmt.Println("Turning back")
		wb.Cmpc(d.RightWheel.Port)
		wb.Cmpc(d.LeftWheel.Port)
		wb.Mrp(d.RightWheel.Port, -spdR, int(float32(-deltaR)/2.35)) //int(math.Abs(float64(deltaR/2))))
		wb.Mrp(d.LeftWheel.Port, -spdL, int(float32(-deltaL)/2.35))  //int(math.Abs(float64(deltaL/2))))
		slpR := make(chan bool)
		go func() {
			wb.Bmd(d.RightWheel.Port)
			slpR <- true
		}()
		slpL := make(chan bool)
		go func() {
			wb.Bmd(d.LeftWheel.Port)
			slpL <- true
		}()
		select {
		case <-slpR:
			break
		case <-slpL:
			break
		}
		wb.Mav(d.RightWheel.Port, 0)
		wb.Mav(d.LeftWheel.Port, 0)
		deltaR = wb.Gmpc(d.RightWheel.Port)
		deltaL = wb.Gmpc(d.LeftWheel.Port)
		fmt.Printf("Deltas: %v, %v\n", deltaR, deltaL)
		slp <- true
	}()
	return slp
}

func (d DifferentialDrive) Perimeter() float32 {
	return d.TrackWidth * math.Pi
}

func (d DifferentialDrive) RotateDeg(deg float32) <-chan interface{} {
	slp := newSlpChan()
	dist := d.Perimeter() * deg / 360.
	if d.LeftWheel.Diameter != d.RightWheel.Diameter {
		panic(fmt.Errorf("RotateDeg only works with equal wheel diameter"))
	}
	rev := dist / d.LeftWheel.Perimeter()
	spd := 800 + int(500*math.Min(float64(deg/180), 1))
	lSpd, rSpd := EqualSpeed(d.LeftWheel.Motor, d.RightWheel.Motor, spd)
	slp1 := d.LeftWheel.TurnExact(lSpd, int(rev*float32(d.LeftWheel.TPR)))
	slp2 := d.RightWheel.TurnExact(-rSpd, -int(rev*float32(d.RightWheel.TPR)))
	go func() {
		<-slp1
		<-slp2
		wb.Mav(d.RightWheel.Port, 0)
		wb.Mav(d.LeftWheel.Port, 0)
		slp <- true
	}()
	return slp
}

func (d DifferentialDrive) TurnDeg(deg, middleRadius float32) <-chan interface{} {
	slp := newSlpChan()
	if d.LeftWheel.Diameter != d.RightWheel.Diameter {
		panic(fmt.Errorf("TurnDeg only works with equal wheel diameter"))
	}
	innerDist := (middleRadius - d.TrackWidth/2) * float32(math.Abs(float64(deg))) / 180
	outerDist := (middleRadius + d.TrackWidth/2) * float32(math.Abs(float64(deg))) / 180
	var distL float32
	var distR float32
	var spdL int
	var spdR int
	if deg > 0 {
		distL = outerDist
		distR = innerDist
		spdL = 1000 + int(500*math.Min(float64(deg/180), 1))
		spdR = int(float32(spdL) * (innerDist / outerDist) / 1.1)
	} else {
		distR = outerDist
		distL = innerDist
		spdR = -1000 + int(500*math.Max(float64(deg/180), -1))
		spdL = int(float32(spdR) * (innerDist / outerDist) / 1.1)
	}
	revL := distL / d.LeftWheel.Diameter
	revR := distR / d.RightWheel.Diameter
	slp1 := d.LeftWheel.TurnExact(spdL, int(revL*float32(d.LeftWheel.TPR)))
	slp2 := d.RightWheel.TurnExact(spdR, int(revR*float32(d.RightWheel.TPR)))
	fmt.Printf("spdR: %v, spdL: %v, distR: %v, distL: %v\n", spdR, spdL, distR, distL)
	go func() {
		<-slp1
		<-slp2
		wb.Mav(d.RightWheel.Port, 0)
		wb.Mav(d.LeftWheel.Port, 0)
		slp <- true
	}()
	return slp
}

func (d DifferentialDrive) Drive(cm float32) <-chan interface{} {
	slp := newSlpChan()
	if d.LeftWheel.Diameter != d.RightWheel.Diameter {
		panic(fmt.Errorf("Drive only works with equal wheel diameter"))
	}
	rev := cm / d.LeftWheel.Perimeter()
	spdL, spdR := EqualSpeed(d.LeftWheel.Motor, d.RightWheel.Motor, 1500)
	fmt.Printf("%v %v\n", spdL, spdR)
	slp1 := d.LeftWheel.TurnExact(spdL, int(rev*float32(d.LeftWheel.TPR)))
	slp2 := d.RightWheel.TurnExact(spdR, int(rev*float32(d.RightWheel.TPR)))
	go func() {
		<-slp1
		<-slp2
		wb.Mav(d.RightWheel.Port, 0)
		wb.Mav(d.LeftWheel.Port, 0)
		slp <- true
	}()
	return slp
}

func (d DifferentialDrive) DriveRegulated(rps, cm float32, interval int) chan interface{} {
	controlChan := newSlpChan()
	spd := rps //rps
	if d.LeftWheel.Diameter != d.RightWheel.Diameter {
		panic(fmt.Errorf("Drive only works with equal wheel diameter"))
	}
	dist := cm / (float32(math.Pi) * d.LeftWheel.Diameter) //rev
	tpsR := int(spd * float32(d.RightWheel.TPR))
	tpsL := int(spd * float32(d.LeftWheel.TPR))
	distR := int(dist * float32(d.RightWheel.TPR))
	distL := int(dist * float32(d.LeftWheel.TPR))
	/*
		fmt.Printf("Starting values: Speed: %f, Distance: %f\n", spd, dist)
		fmt.Printf("Right: tsp: %v, dist: %v\n", tpsR, distR)
		fmt.Printf("Left: tsp: %v, dist: %v\n", tpsL, distL)
		fmt.Println()
		//*/
	stopChan := make(chan interface{})
	stopped := false
	go func() {
		select {
		case <-stopChan:
			break
		case <-controlChan:
			stopChan <- true
			break
		}
		stopped = true
	}()
	go func() {
		for {
			/*if tpsR > 1500 {
				fmt.Printf("Right motor ist turning %vtps / %vrps too fast (%v/1500)\n", tpsR-1500, float32(tpsR-1500)/float32(d.RightWheel.TPR), tpsR)
			}
			if tpsL > 1500 {
				fmt.Printf("Left motor ist turning %vtps / %vrps too fast (%v/1500)\n", tpsL-1500, float32(tpsL-1500)/float32(d.LeftWheel.TPR), tpsL)
			}*/
			wb.Cmpc(d.RightWheel.Port)
			wb.Cmpc(d.LeftWheel.Port)
			if (cm > 0 && distR > 0) || (cm < 0 && distR < 0) {
				wb.Mtp(d.RightWheel.Port, tpsR, distR)
			}
			if (cm > 0 && distL > 0) || (cm < 0 && dist < 0) {
				wb.Mtp(d.LeftWheel.Port, tpsL, distL)
			}
			select {
			case <-time.After(time.Duration(interval) * time.Millisecond):
				break
			case <-stopChan:
				return
				break
			}
			ticksR := wb.Gmpc(d.RightWheel.Port)
			ticksL := wb.Gmpc(d.LeftWheel.Port)
			revR := float32(ticksR) / float32(d.RightWheel.TPR)
			revL := float32(ticksL) / float32(d.LeftWheel.TPR)
			rpsR := revR * 1000. / float32(interval)
			rpsL := revL * 1000. / float32(interval)
			/*
				fmt.Printf("Right drove %v ticks / %.3f revs at %.3f rps\n", ticksR, revR, rpsR)
				fmt.Printf("Left  drove %v ticks / %.3f revs at %.3f rps\n", ticksL, revL, rpsL)
				//*/
			deltaR := rpsR - spd
			deltaL := rpsL - spd
			dTpsR := int(deltaR * float32(d.RightWheel.TPR))
			dTpsL := int(deltaL * float32(d.LeftWheel.TPR))
			adjustR := int(math.Round(float64(dTpsR) / 20.))
			adjustL := int(math.Round(float64(dTpsL) / 20.))
			tpsR -= adjustR
			tpsL -= adjustL
			/*
				fmt.Printf("Right was %v / %.3f wrong, adjusting by %v. New Speed: %v\n", dTpsR, deltaR, adjustR, tpsR)
				fmt.Printf("Left  was %v / %.3f wrong, adjusting by %v. New Speed: %v\n", dTpsL, deltaL, adjustL, tpsL)
				//*/
			distR -= ticksR
			distL -= ticksL
			fmt.Printf("Right covered %v and still needs to go %v\n", ticksR, distR)
			fmt.Printf("Left  covered %v and still needs to go %v\n", ticksL, distL)
			if (cm > 0 && distR <= 0 && distL <= 0) || (cm < 0 && distR >= 0 && distL >= 0) {
				if !stopped {
					stopChan <- true
					controlChan <- true
				}
				break
			}
		}
	}()
	return controlChan
}

func (d DifferentialDrive) DriveVel(rps float32) {
	tpsR := int(rps * float32(d.RightWheel.TPR))
	tpsL := int(rps * float32(d.LeftWheel.TPR))
	wb.Mav(d.RightWheel.Port, tpsR)
	wb.Mav(d.LeftWheel.Port, tpsL)
}

func (d DifferentialDrive) Off() {
	wb.Off(d.LeftWheel.Motor.Port)
	wb.Off(d.RightWheel.Motor.Port)
}

func (d DifferentialDrive) Freeze() {
	wb.Freeze(d.LeftWheel.Motor.Port)
	wb.Freeze(d.RightWheel.Motor.Port)
}
