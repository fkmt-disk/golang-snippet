/*
blink led

platform: RaspberryPi2
output: GPIO21
stop: Ctrl-C
*/
package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"
  "time"
  "github.com/stianeikeland/go-rpio"
)

func main() {
  fmt.Println("start")

  err := rpio.Open()

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  stop_chan := make(chan bool, 1)

  signal_chan := make(chan os.Signal, 1)
  signal.Notify(signal_chan, syscall.SIGINT)

  go func() {
    s := <-signal_chan
    switch s {
      case syscall.SIGINT:
        fmt.Println("SIGINT")
        stop_chan <- true
      }
  }()

  pin := rpio.Pin(21)

  pin.Output()

  LOOP: for {
    select {
      case stop := <-stop_chan:
        if stop {
          break LOOP
        }
      default:
        pin.Toggle()
    }
    time.Sleep(1 * time.Second)
  }

  rpio.Close()
  fmt.Println("stop")
}
