package watch

import (
	"fmt"
	"net"
)

type Watcher struct {
	running  bool
	chanMain chan string
	chansOut []chan<- string
	interval int
}

func NewWatcher() Watcher {
	return Watcher{false, make(chan string), make([]chan<- string, 0), 100}
}

func (w *Watcher) Start() {
	fmt.Println("Watcher started")
	if w.running {
		fmt.Println("Watcher already started")
		return
	}
	w.running = true
	l, err := net.Listen("tcp4", ":55555")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for w.running {
			c, err := l.Accept()
			if !w.running {
				break
			}
			if err != nil {
				fmt.Println(err)
				return
			}
			ch := make(chan string, 0)
			w.chansOut = append(w.chansOut, ch)
			go handleConnection(c, ch)
		}
	}()
	go func() {
		for w.running {
			for data := range w.chanMain {
				if !w.running {
					return
				}
				for _, ch := range w.chansOut {
					ch <- data + "\x00"
				}
			}
		}
	}()
}

/*func (w *Watcher) Stop() {
	w.running = false
	close(w.chanMain)
	for _, ch := range w.chansIn {
		close(ch)
	}
	for _, ch := range w.chansOut {
		close(ch)
	}
}*/

func handleConnection(c net.Conn, out <-chan string) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for data := range out {
		c.Write([]byte(data))
	}
	fmt.Printf("Disconnected %s\n", c.RemoteAddr().String())
	c.Close()
}
