package main

import (
	bb "botball"
	"wallaby"
)

func main() {
	leftWheel := bb.PoweredWheel{6.9, bb.Motor{1, 2020}}
	rightWheel := bb.PoweredWheel{6.9, bb.Motor{0, 1852}}
	drive := bb.DifferentialDrive{leftWheel, rightWheel, 13.5}
	defer wallaby.Ao()
	<-drive.Drive(10)
	<-drive.RotateDeg(90)
	<-drive.Drive(10)
	<-drive.RotateDeg(90)
	<-drive.Drive(10)
	<-drive.RotateDeg(90)
	<-drive.Drive(10)
	<-drive.RotateDeg(90)
}
