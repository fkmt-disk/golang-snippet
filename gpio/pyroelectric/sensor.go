/*
	pyroelectoric sensor

	platform: RaspberryPi2
	input: GPIO14
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

const ReadPinNo = 14

func main() {
	fmt.Println("start")

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin := rpio.Pin(ReadPinNo)
	pin.Input()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT)

	loop := true

	for loop {
		select {
		case s := <-ch:
			fmt.Printf("signal receive: %v\n", s)
			if s == syscall.SIGINT {
				loop = false
			}
		default:
			fmt.Println(pin.Read())
			// if pin.Read() == rpio.High {
			// 	fmt.Println("!")
			// } else {
			// 	fmt.Println(".")
			// }
			time.Sleep(time.Second)
		}
	}

	fmt.Println("stop")
}
