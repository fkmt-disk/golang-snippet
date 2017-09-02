/*
blink led

platform: RaspberryPi2
output: GPIO21
stop: Ctrl-C
*/
package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("start")

	err := rpio.Open()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pin := rpio.Pin(21)
	pin.Output()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)
	loop := true

	for loop {
		select {
		case s := <-ch:
			fmt.Printf("signal receive: %v\n", s)
			if s == syscall.SIGINT {
				loop = false
				pin.Low()
			}
		default:
			pin.Toggle()
			time.Sleep(1 * time.Second)
		}
	}

	rpio.Close()
	fmt.Println("stop")
}
