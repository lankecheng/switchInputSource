package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
	"github.com/go-vgo/robotgo"
)

const RAWCODE_RIGHT_SHIFT = 60

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("select-input-source broken down")
			fmt.Println(r)
		}
	}()

	fmt.Println("Please press right shift to select the input source.")

	
	eventChan := hook.Start()
	defer hook.End()

	for e := range eventChan {
		// fmt.Println("hook: ", e)

		if e.Kind == hook.KeyHold && e.Rawcode == RAWCODE_RIGHT_SHIFT {
			robotgo.KeyTap("space", "ctrl")
		}
	}
}