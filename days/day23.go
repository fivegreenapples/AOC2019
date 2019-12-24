package days

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fivegreenapples/AOC2019/intcode"
)

func (r *Runner) Day23Part1(in string) string {

	nic := intcode.NewFromString(in)
	nic.SetNonBlockingInput(-1)

	inputs := make([]chan int, 50)
	outputs := make([]chan int, 50)
	inputTwoFiveFive := make(chan int)

	// Create all input and output channels
	for addr := 0; addr <= 49; addr++ {

		// Make comms channel buffered with an arbitrary capacity, hopefully enough to avoid isses
		inputs[addr] = make(chan int, 100)
		outputs[addr] = make(chan int, 100)

		// put addr in input chan so it is ready for when nic runs
		inputs[addr] <- addr
	}

	// Fire up goroutines to watch output channels and send to input
	for addr := 0; addr <= 49; addr++ {

		thisAddress := addr
		go func() {
			for {
				sendAddress, valid := <-outputs[thisAddress]
				if !valid {
					return
				}
				if sendAddress == 255 {
					inputTwoFiveFive <- <-outputs[thisAddress]
					inputTwoFiveFive <- <-outputs[thisAddress]
				} else {
					if r.verbose {
						fmt.Println("Sending from", thisAddress, "to", sendAddress)
					}
					inputs[sendAddress] <- <-outputs[thisAddress]
					inputs[sendAddress] <- <-outputs[thisAddress]
				}
			}
		}()
	}

	// Fire up goroutines to run all NICs
	for addr := 0; addr <= 49; addr++ {

		thisAddress := addr
		go func() {
			nic.Run(inputs[thisAddress], outputs[thisAddress])
			close(outputs[thisAddress])
		}()
	}

	<-inputTwoFiveFive
	y := <-inputTwoFiveFive

	return strconv.Itoa(y)
}

func (r *Runner) Day23Part2(in string) string {
	nic := intcode.NewFromString(in)
	nic.SetNonBlockingInput(-1)

	inputs := make([]chan int, 50)
	outputs := make([]chan int, 50)
	natInput := make(chan int)
	natSensorChannel := make(chan bool)
	natResultChannel := make(chan int)

	// Create all input and output channels
	for addr := 0; addr <= 49; addr++ {

		// Make comms channel buffered with an arbitrary capacity, hopefully enough to avoid isses
		inputs[addr] = make(chan int, 100)
		outputs[addr] = make(chan int, 100)

		// put addr in input chan so it is ready for when nic runs
		inputs[addr] <- addr
	}

	// Fire up goroutines to watch output channels and send to input
	for addr := 0; addr <= 49; addr++ {

		thisAddress := addr
		go func() {
			for {
				sendAddress, valid := <-outputs[thisAddress]
				if !valid {
					return
				}
				natSensorChannel <- true
				if sendAddress == 255 {
					natInput <- <-outputs[thisAddress]
					natInput <- <-outputs[thisAddress]
				} else {
					inputs[sendAddress] <- <-outputs[thisAddress]
					inputs[sendAddress] <- <-outputs[thisAddress]
				}
			}
		}()
	}

	// Fire up goroutines for NAT
	go func() {
		var natX, natY int
		sentY := map[int]bool{}
		for {
			timer := time.NewTimer(100 * time.Millisecond)
			select {
			case natX = <-natInput:
				natY = <-natInput
			case <-natSensorChannel:
				if !timer.Stop() {
					<-timer.C
				}
			case <-timer.C:
				// indicates a period of inactivity on sending packets. So probably idle
				// so we send to NIC 0

				if sentY[natY] {
					if r.verbose {
						fmt.Printf("First repeat of a Y value sent from NAT. Repeating %d\n", natY)
					}
					natResultChannel <- natY
				}

				if r.verbose {
					fmt.Printf("NAT is sending (%d,%d) to NIC-0\n", natX, natY)
				}
				inputs[0] <- natX
				inputs[0] <- natY
				sentY[natY] = true

			}
		}

	}()

	// Fire up goroutines to run all NICs
	for addr := 0; addr <= 49; addr++ {

		thisAddress := addr
		go func() {
			nic.Run(inputs[thisAddress], outputs[thisAddress])
			close(outputs[thisAddress])
		}()
	}

	repeatedY := <-natResultChannel
	return strconv.Itoa(repeatedY)
}
