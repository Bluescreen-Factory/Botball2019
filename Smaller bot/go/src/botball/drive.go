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
		wb.Bmd(m.Port)
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
	lSpd, rSpd := EqualSpeed(d.LeftWheel.Motor, d.RightWheel.Motor, 1500)
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
	innerDist := (middleRadius - d.TrackWidth/2) * float32(math.Abs(float64(deg))) / 360
	outerDist := (middleRadius + d.TrackWidth/2) * float32(math.Abs(float64(deg))) / 360
	var distL float32
	var distR float32
	var spdL int
	var spdR int
	if deg > 0 {
		distL = outerDist
		distR = innerDist
		spdL = 1500
		spdR = int(float32(spdL) * (innerDist / outerDist))
	} else {
		distR = outerDist
		distL = innerDist
		spdR = 1500
		spdL = int(float32(spdR) * (innerDist / outerDist))
	}
	revL := distL / d.LeftWheel.Diameter
	revR := distR / d.RightWheel.Diameter
	slp1 := d.LeftWheel.TurnExact(spdL, int(revL*float32(d.LeftWheel.TPR)))
	slp2 := d.RightWheel.TurnExact(spdR, int(revR*float32(d.RightWheel.TPR)))
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

func (d DifferentialDrive) Off() {
	wb.Off(d.LeftWheel.Motor.Port)
	wb.Off(d.RightWheel.Motor.Port)
}
