# About 

This is a small application written to test some features provided by new AWS Graviton2 processor type. 
The idea is to write a few identical application (one calculates fibonacci numbers, another just echoer) and deploy them under different archs: 
- x86 
- ARM (Graviton2)

## Project structure 

- cmd/echo-lambda - echoer source code. It reads ``foo`` argument from URL and returns back in response body;
- cmd/fibo-lambda - fibonacci calculator source code. It reads number position from ENV (30 by default) and calculates the number.

To test applications' performance we should somehow send requests to lambda URL and do it for some period of time. 
The easiest way to do it is to write a Golang application which sends HTTP requests to given lambda and splits traffic in half between x86 and ARM processors. 

This logic implemented in the following folders: 
- cmd/echo-lambda-runner 
- cmd-fibo-lambda runner 

Lambdas' code is deployed to cloud automatically using AWS SAM tooling. You can find config in ``template.yml`` in the root folder. 

Also, there are Makefile to simplify deployment process. 

## Makefile 

- ``make build`` command prepares binaries for all applications and architectures;
- ``make deploy`` automatically deploys those binaries and configures everything on AWS side; 
- ``make run-fibonacci-lambda`` triggers fibonacci lambda calculator;
- ``make run-echo-lambda`` triggers echoer lambda;

## Configuration 

Once lambdas are deployed you'll notice an URL in the CLI output with host id. Please copy it and use as ``IDENTIFIER`` env variable for
``make run-fibonacci-lambda`` and ``make run-echo-lambda`` commands. Also both of them supports ``ITERATIONS`` env variable to configure amount of function invocations. 