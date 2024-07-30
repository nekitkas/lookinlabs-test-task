package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	fmt.Println("Hello, World!")
	problemSolver(1000000) // Increase the data amount to make the leak more apparent
}

func problemSolver(dataAmount int) {
	var data []int
	for {
		data = append(data, make([]int, dataAmount)...) // Append a large slice to increase memory usage
		time.Sleep(10 * time.Millisecond)
		fmt.Println("Length of slice:", len(data))
	}
}
