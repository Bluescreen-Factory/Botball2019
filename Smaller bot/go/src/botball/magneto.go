package botball

/*
import (
	bb "botball",
	"math"
)

//Calibration value
type calibVal struct {
	min int,
	max int
}

var calX = calibVal{
	0,
	100
}

var calX = calibVal{
	0,
	100
}

const MagnetoSensor = struct {
	Yaw float32
}{
	0.0
}

func GetMagnetoCalibrationValues() {
	minX := math.MaxFloat32
	maxX := - math.MaxFloat32
	minY := math.MaxFloat32
	maxY := - math.MaxFloat32

	rightWheel := bb.PoweredWheel{
		Diameter: 6.9,
		Motor: bb.Motor{
			Port: 0,
			TPR:  2045,
		},
	}
	leftWheel := bb.PoweredWheel{
		Diameter: 6.9,
		Motor: bb.Motor{
			Port: 1,
			TPR:  1898,
		},
	}
	drive := bb.DifferentialDrive{
		LeftWheel:  leftWheel,
		RightWheel: rightWheel,
		TrackWidth: 13.5,
	}

	for i := 0; i < 4; i++ {
		doneChan := drive.RotateDeg(90)
		stop := false
		go func() {
			for !stop {

			}
		}()
		<-doneChan
		stop = true
	}
}*/
