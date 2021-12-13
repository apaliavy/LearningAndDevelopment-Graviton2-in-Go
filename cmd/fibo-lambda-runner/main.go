package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	// timeout between lambda calls
	timeout = time.Millisecond * 500
	// defaultIterationsCnt describes how many calls do we want to make
	defaultIterationsCnt = 1500
)

func main() {
	id := os.Getenv("LAMBDA_IDENTIFIER")
	if id == "" {
		log.Fatalf("failed to run application - lambda id is missing")
	}

	iterations := defaultIterationsCnt
	if i, err := strconv.Atoi(os.Getenv("ITERATIONS")); err == nil {
		iterations = i
	}

	for i := 0; i < iterations; i++ {
		url := fmt.Sprintf("https://%s.execute-api.us-east-1.amazonaws.com/Prod/x86/number", id)
		if i%2 == 0 {
			url = fmt.Sprintf("https://%s.execute-api.us-east-1.amazonaws.com/Prod/arm64/number", id)
		}

		if _, err := http.Get(url); err != nil {
			log.Printf("failed to call %s - received an error %s\n", url, err)
		}

		time.Sleep(timeout)
	}
}
