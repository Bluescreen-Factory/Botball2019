package main

import (
	"botball/watch"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"time"
	wb "wallaby"
)

type Motor struct {
	TPR  int
	Port int
}

func main() {
	watcher := watch.NewWatcher()
	watcher.Start()
	watcher.WatchMotors()
	motorL := Motor{
		TPR:  2045,
		Port: 1,
	}
	motorR := Motor{
		TPR:  1895,
		Port: 0,
	}
	watcher.WatchPosition(6.9, 13.5, motorL.TPR, motorR.TPR)
	var useMouse bool
	flag.BoolVar(&useMouse, "mouse", false, "save mouse position")
	flag.Parse()
	var mouseChan chan bool
start:
	fmt.Println("Press any button to start")
	wb.Wait_for_any_button()
	if useMouse {
		output, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		output += fmt.Sprintf("/positions-%v.txt", time.Now().Format("20016102150405"))
		fmt.Printf("Saving mouse positions at: %q \n", output)
		cmd := exec.Command("/home/root/Documents/KISS/Default User/tm_mouse/bin/main.py")
		outR, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		outW, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY, 644)
		if err != nil {
			panic(err)
		}
		if mouseChan == nil {
			mouseChan = make(chan bool)
		} else {
			mouseChan <- true
			<-mouseChan
		}
		go func(outR io.ReadCloser, outW io.WriteCloser) {
			go func() {
				io.Copy(outW, outR)
			}()
			<-mouseChan
			outR.Close()
			outW.Close()
			mouseChan <- true
		}(outR, outW)
		if err != nil {
			panic(err)
		}
		err = cmd.Start()
		if err != nil {
			panic(err)
		}
	}
	spd := float32(0.7) //rps
	dist := float32(8.) //rev
	tpsR := int(spd * float32(motorR.TPR))
	tpsL := int(spd * float32(motorL.TPR))
	distR := int(dist * float32(motorR.TPR))
	distL := int(dist * float32(motorL.TPR))
	interval := 100
	fmt.Printf("Starting values: Speed: %f, Distance: %f\n", spd, dist)
	fmt.Printf("Right: tsp: %v, dist: %v\n", tpsR, distR)
	fmt.Printf("Left: tsp: %v, dist: %v\n", tpsL, distL)
	fmt.Println()
	<-time.After(150 * time.Millisecond)
	for {
		if tpsR > 1500 {
			fmt.Printf("Right motor ist turning %vtps / %vrps too fast (%v/1500)\n", tpsR-1500, float32(tpsR-1500)/float32(motorR.TPR), tpsR)
		}
		if tpsL > 1500 {
			fmt.Printf("Left motor ist turning %vtps / %vrps too fast (%v/1500)\n", tpsL-1500, float32(tpsL-1500)/float32(motorL.TPR), tpsL)
		}
		wb.Cmpc(motorR.Port)
		wb.Cmpc(motorL.Port)
		if distR > 0 {
			wb.Mtp(motorR.Port, tpsR, distR)
		}
		if distL > 0 {
			wb.Mtp(motorL.Port, tpsL, distL)
		}
		<-time.After(time.Duration(interval) * time.Millisecond)
		ticksR := wb.Gmpc(motorR.Port)
		ticksL := wb.Gmpc(motorL.Port)
		revR := float32(ticksR) / float32(motorR.TPR)
		revL := float32(ticksL) / float32(motorL.TPR)
		rpsR := revR * 1000. / float32(interval)
		rpsL := revL * 1000. / float32(interval)
		fmt.Printf("Right drove %v ticks / %.3f revs at %.3f rps\n", ticksR, revR, rpsR)
		fmt.Printf("Left  drove %v ticks / %.3f revs at %.3f rps\n", ticksL, revL, rpsL)
		deltaR := rpsR - spd
		deltaL := rpsL - spd
		dTpsR := int(deltaR * float32(motorR.TPR))
		dTpsL := int(deltaL * float32(motorL.TPR))
		adjustR := int(math.Round(float64(dTpsR) / 20.))
		adjustL := int(math.Round(float64(dTpsL) / 20.))
		tpsR -= adjustR
		tpsL -= adjustL
		fmt.Printf("Right was %v / %.3f wrong, adjusting by %v. New Speed: %v\n", dTpsR, deltaR, adjustR, tpsR)
		fmt.Printf("Left  was %v / %.3f wrong, adjusting by %v. New Speed: %v\n", dTpsL, deltaL, adjustL, tpsL)
		distR -= ticksR
		distL -= ticksL
		fmt.Printf("Right covered %v and still needs to go %v\n", ticksR, distR)
		fmt.Printf("Left  covered %v and still needs to go %v\n", ticksL, distL)
		fmt.Println()
		if distR <= 0 && distL <= 0 {
			break
		}
	}
	fmt.Println("Exit, Again? (Left, Right)")
	for (wb.Left_button() + wb.Right_button()) == 0 {
	}
	if wb.Left_button() == 1 {
		mouseChan <- true
		<-mouseChan
		return
	}
	goto start
}
