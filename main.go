// +build ignore

package main

import (
	"flag"
	"fmt"
	"time"
	"github.com/kidoman/embd"
        "github.com/stuphi/scroll-phat-go/scrollphat"

	_ "github.com/kidoman/embd/host/all"
        "os"
        "os/signal"
	"syscall"

)

func zero(s scrollphat.ScrollPhat) {
    s.Buffer = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    s.Offset = 0
    s.Update()
}

func blink() {
    var sf scrollphat.ScrollPhat
    var x, y uint

    sf.Init()
    zero(sf)

    for x = 0; x < 11; x++ { 
	for y = 0; y < 5; y++ {
            sf.SetPixel(x,y,1)
            sf.Update()
            time.Sleep(10 * time.Millisecond)
        }
    }
    for x = 0; x < 11; x++ { 
        for y = 0; y < 5; y++ {
            sf.SetPixel(x,y,0)
            sf.Update()
            time.Sleep(10 * time.Millisecond)
        }
    }
}

func captureCtrlC(btn embd.DigitalPin) {
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    go func() {
	for sig := range c {
	    fmt.Println(fmt.Sprintf("Captured '%v', exiting..", sig))
	    btn.Close()
	    os.Exit(1)
	}
    }()
}

func main() {
	flag.Parse()

	if err := embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()

	btn, err := embd.NewDigitalPin(10)
	if err != nil {
		panic(err)
	}
	defer btn.Close()
	captureCtrlC(btn)
	fmt.Println("start")
	if err := btn.SetDirection(embd.In); err != nil {
		panic(err)
	}
	btn.ActiveLow(false)
	
	quit := make(chan interface{})
	err = btn.Watch(embd.EdgeFalling, func(btn embd.DigitalPin) {
	    quit <- btn
	    fmt.Printf("Button %v was pressed.\n", btn)
	})
	if err != nil {
		panic(err)
	}
	for {
	    select {
	    case btn := <-quit:
                fmt.Println("!!!!", btn)
		blink()
	    }
	}
}
