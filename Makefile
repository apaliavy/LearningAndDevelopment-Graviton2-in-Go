.PHONY: build deploy

LAMBDA_IDENTIFIER ?= fill-me-please
ITERATIONS        ?= 1500

build:
	sam build

deploy:
	sam deploy --guided

run-fibonacci-lambda:
	LAMBDA_IDENTIFIER=$(LAMBDA_IDENTIFIER) ITERATIONS=$(ITERATIONS) go run cmd/fibo-lambda-runner/main.go

run-echo-lambda:
	LAMBDA_IDENTIFIER=$(LAMBDA_IDENTIFIER) ITERATIONS=$(ITERATIONS) go run cmd/echo-lambda-runner/main.go
