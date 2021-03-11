package main

import (
	"fmt"
	"sort"
	"time"
)

func backgroundSort(s []int, job func(), c chan int) {
	start := time.Now()
	go func() {
		start := time.Now()
		fmt.Println("Start sort at ", start)
		sort.Ints(s)
		duration := time.Since(start)
		end := time.Now()
		fmt.Println("End sort at ", end)
		fmt.Println("Sort duration ", duration)
		c <- 1
	}()

	job()
	<-c
	duration := time.Since(start)
	fmt.Println("backgroundSorts duration ", duration)
}

func handleRequestLimited(criticalCapacity int, requestQueue []string) {
	sem := make(chan int, criticalCapacity)
	stop := make(chan int, len(requestQueue))

	for i, req := range requestQueue {
		fmt.Println("Waiting for critical section for index ", i)
		sem <- 1
		fmt.Println("Entered critical section for index ", i)
		go func(req string) {
			fmt.Println("Called subroutine for ", req)
			process(req)
			<-sem
			stop <- 1
		}(req)
	}

	for i := 0; i < len(requestQueue); i++ {
		<-stop
	}
}

func process(req string) {
	time.Sleep(time.Second * 3)
	fmt.Println("Processed ", req)
}
