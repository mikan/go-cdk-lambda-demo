# go-cdk-lambda-demo

This is a demo of deploying Lambda written in Go using Go CDK.

## Build and deploy

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello cmd/hello/main.go
zip hello hello.zip
cdk deploy
```

## License

[Apache License 2.0](LICENSE)

## Author

[mikan](https://github.com/mikan)
