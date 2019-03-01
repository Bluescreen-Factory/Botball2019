package main

import (
	bb "botball"
	"fmt"
)

type State func() State

var (
	state_DrivingToLine State = driveToLine
	state_END           State = nil
)

var drive bb.DifferentialDrive
var state = state_DrivingToLine

func main() {
	rightWheel := bb.PoweredWheel{
		Diameter: 6.9,
		bb.Motor{
			Port: 0,
			TPR:  1898,
		},
	}
	leftWheel := bb.PoweredWheel{
		Diameter: 6.9,
		bb.Motor{
			Port: 1,
			TPR:  2045,
		},
	}
	drive = bb.DifferentialDrive{
		LeftWheel:  leftWheel,
		RightWheel: rightWheel,
		TrackWidth: 13.5,
	}

	for state != state_END {
		state()
	}
	fmt.Println("DONE!")
}

func driveToLine() State {
	<-drive.RotateDeg(-30)
	<-drive.Drive(100)
	return state_END
}
