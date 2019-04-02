package main

import (
	bb "botball"
	"fmt"
	"math"
	wb "wallaby"
)

type State func() State

const (
	SENSOR_DIST  = 3
	SENSOR_FRONT = 4
	SENSOR_BACK  = 5
)

var (
	state_DrivingToLine        State = driveToLine
	state_AligningWithLine     State = alignWithLine
	state_DrivingAlongLine     State = driveAlongLine
	state_TakingWaterCollector State = takeWaterCollector
	state_CollectingFlufs      State = collectFlufs
	state_DrivingToUtilityZone State = driveToUtilityZone
	state_testDrivingFast      State = test_driveFast
	state_testRotatingExact    State = test_rotateExact
	state_testDrivingBack      State = test_driveBack
	state_testCompass          State = test_compass
	state_testAligningWithLine State = test_alignWithLine
	state_demoCollection       State = demo_collect
	state_END                  State = nil
)

var drive bb.DifferentialDrive
var state = state_AligningWithLine //state_DrivingToLine

const BLACK_THRESH = 1000
const EDGE_THRESH = 2800

func main() {
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
	drive = bb.DifferentialDrive{
		LeftWheel:  leftWheel,
		RightWheel: rightWheel,
		TrackWidth: 13.9, //13.5
	}
	wb.Enable_servos()
	for state != nil {
		state = state()
	}
	fmt.Println("DONE!")
	wb.Disable_servos()
	drive.Off()
}

func test_alignWithLine() State {
	<-drive.AlignWithLine(0.2, -800, SENSOR_FRONT)
	return state_END
}

func test_compass() State {
	wb.Wait_for_any_button()
	if wb.Left_button() == 1 {
		wb.Calibrate_compass()
	}
	wb.Set_compass_params(107, 73.79, 17.06, -0.12, 0.09, 1.01, 1.02)
	wb.Wait_for_any_button()
	for wb.Any_button() != 0 {
	}
	for wb.Any_button() == 0 {
		rad := wb.Get_compass_angle()
		deg := rad * float32(180./math.Pi)
		fmt.Printf("Angle: %v\n", deg)
		wb.Msleep(10)
	}
	return state_END
}

func test_driveFast() State {
	fmt.Println("Testing driving fast")
	<-drive.DriveRegulated(1.4, 200, 50)
	return state_END
}

func test_driveBack() State {
	<-drive.DriveRegulated(-0.4, -3, 150)
	return state_END
}

func demo_collect() State {
	wb.Enable_servos()
	fmt.Println("Demoing collecting")
	wb.Set_servo_position(1, 2047)
	wb.Msleep(1000)
	<-drive.DriveRegulated(0.5, 10, 100)
	for i := 0; i < 800; i++ {
		wb.Set_servo_position(1, 2047-i)
		wb.Msleep(1)
	}
	wb.Wait_for_any_button()
	fmt.Println("continuing...")
	wb.Msleep(2000)
	for i := 0; i < 800; i++ {
		wb.Set_servo_position(1, 1247-i)
		wb.Msleep(1)
	}
	wb.Msleep(1000)
	return state_END
}

func test_rotateExact() State {
	fmt.Println("Testing exact turning")
	fmt.Println("-25 deg turn")
	<-drive.RotateDeg(-25)
	drive.Freeze()
	fmt.Println("-30 deg turn")
	wb.Wait_for_any_button()
	<-drive.RotateDeg(-30)
	drive.Freeze()
	fmt.Println("-35 deg turn")
	wb.Wait_for_any_button()
	<-drive.RotateDeg(-35)
	drive.Freeze()
	fmt.Println("-40 deg turn")
	wb.Wait_for_any_button()
	<-drive.RotateDeg(-40)
	drive.Freeze()
	return state_END
}

func driveToLine() State {
	fmt.Println("Changed state to driveToLine")
	fmt.Println("Turnining -45deg with r=~50")
	<-drive.TurnDeg(-100, 45)
	wb.Msleep(500)
	fmt.Println("Driving until line with back sensor")
	stopChan := drive.DriveRegulated(0.3, 9999, 100)
	for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
	}
	fmt.Println("reached line")
	stopChan <- true
	wb.Msleep(500)
	fmt.Println("Turning 60 deg with r=~15")
	<-drive.TurnDeg(60, 18.5)
	wb.Msleep(500)
	fmt.Println("Driving ~5 cm")
	<-drive.DriveRegulated(0.7, 5, 100)
	wb.Msleep(500)
	fmt.Println("Turning 30 deg with r=~15")
	<-drive.TurnDeg(30, 18.5)
	wb.Msleep(500)
	fmt.Println("Driving until line with back sensor")
	stopChan = drive.DriveRegulated(0.7, 9999, 100)
	for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
	}
	stopChan <- true
	wb.Msleep(500)
	fmt.Println("Driving ~5cm")
	drive.DriveRegulated(0.7, 6, 100)
	wb.Msleep(500)
	fmt.Println("Turning 90deg with r=~5")
	<-drive.TurnDeg(90, 5)
	wb.Msleep(500)
	return state_AligningWithLine
	/*
		stopChan = drive.DriveRegulated(0.7, 9999, 100)
		fmt.Println("Driving to line, waiting until Analog(SENSOR_FRONT) > BLACK_THRESH")
		for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
		}
		for wb.Analog(SENSOR_FRONT) > BLACK_THRESH {
		}
		fmt.Println("Passed first line")
		for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
		}
		for wb.Analog(SENSOR_FRONT) > BLACK_THRESH {
		}
		fmt.Println("Passed second line")
		for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
		}
		fmt.Println("Reached third line")
		stopChan <- true
		wb.Msleep(500)
		fmt.Println("Reached line, rotating 70 deg")
		<-drive.RotateDeg(80)
		return state_AligningWithLine*/
}

func alignWithLine() State {
	fmt.Println("Changed state to alignWithLine")
	fmt.Println("Driving ~-10 cm")
	<-drive.DriveRegulated(0.7, -10, 100)
	wb.Msleep(500)
	//fmt.Println("Shovel down")
	//wb.Set_servo_position(1, 2047)
	wb.Set_servo_position(1, 300)
	wb.Msleep(250)
	controlChan := drive.DriveRegulated(0.7, 12.5, 100)
	fmt.Println("Driving until line")
	for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
	}
	fmt.Println("Reached line")
	controlChan <- true
	drive.Freeze()
	wb.Msleep(2000)
	//fmt.Println("Driving ~3 cm")
	//<-drive.DriveRegulated(0.7, 2, 100)
	fmt.Println("Aligning with line")
	<-drive.AlignWithLine(0.5, -800, SENSOR_FRONT)
	return state_END
	wb.Msleep(500)
	for i := 0; i < 600; i++ {
		wb.Set_servo_position(1, 2047-600)
		wb.Msleep(5)
	}
	<-drive.AlignWithLine(0.5, BLACK_THRESH, SENSOR_FRONT)
	wb.Msleep(500)
	spd := 800
	wb.Cmpc(0)
	wb.Cmpc(1)
	wb.Mav(0, spd)
	wb.Mav(1, -spd)
	fmt.Println("Tunring until line")
	for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
	}
	fmt.Println("Line detected")
	wb.Mav(0, -spd)
	wb.Mav(1, spd)
	wb.Msleep(500)
	wb.Cmpc(0)
	wb.Cmpc(1)
	fmt.Println("Turning back until not line")
	for wb.Analog(SENSOR_FRONT) > BLACK_THRESH {
		fmt.Printf("%v\n", wb.Analog(SENSOR_FRONT))
	}
	fmt.Println("Line undetected")
	wb.Msleep(500)
	fmt.Println("Turning back until line")
	for wb.Analog(SENSOR_FRONT) < BLACK_THRESH {
	}
	fmt.Println("Line detected")
	wb.Mav(1, 0)
	wb.Mav(0, 0)
	delta0 := wb.Gmpc(0)
	delta1 := wb.Gmpc(1)
	fmt.Printf("Deltas: %v, %v\n", delta0, delta1)
	wb.Msleep(500)
	fmt.Println("Turning back")
	wb.Mtp(0, spd, delta0/2)
	wb.Mtp(1, spd, delta1/2)
	wb.Bmd(0)
	if wb.Get_motor_done(1) == 0 {
		wb.Bmd(1)
	}
	return state_END
	return state_DrivingAlongLine
}

func driveAlongLine() State {
	fmt.Println("Changed state to drivingAlongLine")
	fmt.Println("Driving back ~10cm")
	<-drive.DriveRegulated(-0.7, -10, 100)
	wb.Msleep(500)
	fmt.Println("Shovel down")
	wb.Set_servo_position(1, 2047)
	wb.Msleep(500)
	fmt.Println("Driving along line")
	<-drive.DriveRegulated(0.7, 110, 50)
	return state_TakingWaterCollector
}

func takeWaterCollector() State {
	wb.Msleep(500)
	fmt.Println("Opening gate")
	wb.Set_servo_position(3, 2047)
	wb.Msleep(500)
	fmt.Println("Driving ~10cm")
	<-drive.DriveRegulated(0.5, 10, 100)
	wb.Msleep(500)
	fmt.Println("Turning -45 deg")
	<-drive.TurnDeg(-45, drive.TrackWidth/2)
	wb.Msleep(500)
	fmt.Println("Closing gate")
	wb.Set_servo_position(3, 760)
	wb.Msleep(500)
	fmt.Println("Turning 45 deg")
	<-drive.TurnDeg(45, drive.TrackWidth/2)
	return state_END
	return state_CollectingFlufs
}

func collectFlufs() State {
	wb.Msleep(500)
	wb.Set_servo_position(3, 740)
	wb.Set_servo_position(1, 2047)
	wb.Msleep(500)
	shakeShovel := func() {
		startPos := wb.Get_servo_position(1)
		fmt.Printf("Start Pos: %v\n", startPos)
		for i := 0; i < 8; i++ {
			if i%2 == 0 {
				wb.Set_servo_position(1, startPos+110)
				wb.Msleep(100)
			} else {
				wb.Set_servo_position(1, startPos-10)
				wb.Msleep(100)
			}
		}
		wb.Set_servo_position(1, startPos)
	}
	fmt.Println("Starting shoveling")
	shovelDown := 2047
	shovelTilt := 200
	shovelUp := 1600
	fmt.Println("Shovel down")
	wb.Set_servo_position(1, shovelDown)
	for i := 0; i < 4; i++ {
		wb.Msleep(250)
		fmt.Println("Driving 10 cm foreward")
		<-drive.DriveRegulated(0.4, 10, 50)
		wb.Msleep(250)
		fmt.Println("Shovel tilt")
		for i := 0; i < shovelTilt; i++ {
			wb.Set_servo_position(1, shovelDown-i)
			wb.Msleep(1)
		}
		wb.Msleep(250)
		fmt.Println("Driving 4 cm back")
		<-drive.DriveRegulated(-0.4, -4, 50)
		wb.Msleep(250)
		fmt.Println("Shovel up")
		for i := 0; i < shovelUp; i++ {
			wb.Set_servo_position(1, shovelDown-shovelTilt-i)
			wb.Msleep(1)
		}
		wb.Msleep(250)
		fmt.Println("Shaking shovel")
		shakeShovel()
		wb.Msleep(250)
		fmt.Println("Shovel down")
		wb.Set_servo_position(1, shovelDown)
		wb.Msleep(500)
		fmt.Println("Pressing")
		wb.Set_servo_position(2, 960)
		wb.Msleep(500)
		wb.Set_servo_position(2, 0)
		wb.Msleep(500)
		fmt.Println("Driving 6 cm back")
		<-drive.DriveRegulated(-0.4, -6, 50)
		wb.Msleep(250)
	}
	fmt.Println("Pressing")
	wb.Set_servo_position(2, 960)
	wb.Msleep(500)
	wb.Set_servo_position(2, 0)
	wb.Msleep(500)
	return state_DrivingToUtilityZone
}

func driveToUtilityZone() State {
	wb.Msleep(500)
	fmt.Println("Rotating 90deg")
	drive.RotateDeg(90)
	wb.Msleep(500)
	fmt.Println("Driving 80cm")
	drive.DriveRegulated(0.7, 80, 100)
	return state_END
}
