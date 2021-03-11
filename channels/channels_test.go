package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestBackgroundJob(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, 10000000)
	for i := range s {
		s[i] = rand.Int()
	}

	fakeJob := func() {
		start := time.Now()
		fmt.Println("Start fake job at ", start)
		time.Sleep(time.Second)
		duration := time.Since(start)
		end := time.Now()
		fmt.Println("End fake job at ", end)
		fmt.Println("Fake job duration ", duration)
	}
	c := make(chan int)
	backgroundSort(s, fakeJob, c)
}

func TestHandleRequestLimited(t *testing.T) {
	queue := []string{"/users/", "/orders", "/orders&id=1", "/paymentMethods", "/login"}
	handleRequestLimited(2, queue)

	handleRequestLimited(4, queue)
}
