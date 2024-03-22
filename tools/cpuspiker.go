// Datadog CPU spiker
// Use this program to test detections for malware hashes and crypto mining
package main

import (
	"flag"
	"fmt"
)

func cpuTask(id int, data chan int) {
	for {
	}
}

func main() {

	priorityPtr := flag.Int("cpu-priority", 0, "Mining threads priority")
	noSignalPtr := flag.Bool("no-signal", false, "Execute without triggering a crypto mining signal")
	flag.Parse()

	if *priorityPtr != 0 || *noSignalPtr {
		fmt.Print("Spiking CPU to test Datadog security detections...\n")
		channel := make(chan int)

		for i := 0; i < 10; i++ {
			go cpuTask(i, channel)
		}

		for i := 0; i < 10; i++ {
			channel <- i
		}
	} else {
		flag.PrintDefaults()
	}
}
