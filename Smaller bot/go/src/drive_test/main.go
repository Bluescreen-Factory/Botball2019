package main

import wlby "wallaby"

func main() {
	for wlby.Right_button() != 1 {

	}
	wlby.Motor(0, 100)
	wlby.Msleep(500)
	wlby.Motor(0, 0)
}
