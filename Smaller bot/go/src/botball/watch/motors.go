package watch

import (
	"fmt"
	"time"
	wb "wallaby"
)

func (w *Watcher) WatchMotors() {
	go func() {
		fmt.Println("Watching motor")
		for w.running {
			<-time.After(time.Duration(int32(w.interval)) * time.Millisecond)
			if !w.running {
				break
			}
			w.chanMain <- fmt.Sprintf("MotorTicks:%v;%v;%v;%v", wb.Gmpc(0), wb.Gmpc(1), wb.Gmpc(2), wb.Gmpc(3))
		}
	}()
}
