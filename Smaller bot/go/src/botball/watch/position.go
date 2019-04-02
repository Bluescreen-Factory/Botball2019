package watch

import (
	"fmt"
	"math"
	"time"
	wb "wallaby"
)

func (w *Watcher) WatchPosition(wheelDiameter, trackWidth float32, tpr0, tpr1 int) {
	go func() {
		var x float32
		var y float32
		var heading float32
		pos0 := wb.Gmpc(0)
		pos1 := wb.Gmpc(1)
		fmt.Println("Watching position")
		for w.running {
			<-time.After(time.Duration(int32(w.interval)) * time.Millisecond)
			if !w.running {
				break
			}
			nPos0 := wb.Gmpc(0)
			nPos1 := wb.Gmpc(1)
			d0 := nPos0 - pos0
			d1 := nPos1 - pos1
			dRev0 := float32(d0) / float32(tpr0)
			dRev1 := float32(d1) / float32(tpr1)
			dCm0 := dRev0 * wheelDiameter
			dCm1 := dRev1 * wheelDiameter

			if d0-d1 == 0 {
				x = x + dCm0*float32(math.Cos(float64(heading)))
				y = y + dCm1*float32(math.Sin(float64(heading)))
			} else {
				r := trackWidth * (dCm0 + dCm1) / (2 * (dCm1 - dCm0))
				wd := (dCm1 - dCm0) / trackWidth
				x = x + r*float32(math.Sin(float64(wd+heading))) - r*float32(math.Sin(float64(heading)))
				y = y - r*float32(math.Cos(float64(wd+heading))) + r*float32(math.Cos(float64(heading)))
				heading = heading + wd
			}
			w.chanMain <- fmt.Sprintf("Position:%v;%v", x, y)
		}
	}()
}
