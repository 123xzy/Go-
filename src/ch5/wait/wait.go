package main

import (
	"log"
	"fmt"
	"time"
)

func WaitForServer(url string) error{

	const timeout := time.Minute * 1
	deadline := time.Now().Add(timeout)
	for tries := 0;time.Now().Before(deadline);tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
