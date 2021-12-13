package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const (
	// defaultNumber means we want to get 30th number in Fibonacci sequence by default
	defaultNumber = 30

	runTypeRecursive              = "RECURSIVE"
	runTypeTailOptimisedRecursion = "TAIL_OPTIMISED"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var n = defaultNumber
	// try to override default config
	if val, err := strconv.Atoi(os.Getenv("NUMBER")); err == nil && val > 0 {
		n = val
	}

	if mode := os.Getenv("CALL_TYPE"); mode == runTypeRecursive {
		fibonacciRecursive(n)
	} else if mode == runTypeTailOptimisedRecursion {
		fibonacciTailOptimised(n)
	} else {
		fibonacciIterative(n)
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}

func fibonacciIterative(n int) (result int) {
	for i, first, second := 0, 0, 1; i <= n; i, first, second = i+1, first+second, first {
		if i == n {
			result = first
		}
	}
	return result
}

func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciRecursive(n-2) + fibonacciRecursive(n-1)
}

func fibonacciTailOptimised(n int) int {
	return fiboT(n, 0, 1)
}

func fiboT(n, first, second int) int {
	if n == 0 {
		return first
	}
	return fiboT(n-1, second, first+second)
}
