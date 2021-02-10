Hello world example service using [runtime](https://github.com/g-wilson/runtime), a framework for writing JSON-RPC services in Go and hosting them on AWS Lambda with API Gateway.

### How to run locally

Start the development server:

```
go run debug/main.go
```

Send a request:

```
http --json POST http://localhost:3030/helloworld/greet name='george'
```

### How to run on AWS

Package the service:

```
GOARCH=amd64 GOOS=linux go build ./cmd/main.go && zip helloworld.zip main
```

Create a Lambda function and upload the zip file. Set the environment variables.

Create an HTTP API on API Gateway.

Create a route like `/helloworld/v1/{method+}` and set up an integration with the Lambda function you just created.

That's it!
