# Golang 1.8beta1 plugin example

For more informations about GoLang plugins check the official [documentation](https://beta.golang.org/pkg/plugin/).

### Build and run the example using Docker
```bash
docker run -it --entrypoint bash golang:1.8beta1-wheezy
go get github.com/devalecs/go-plugin-example && cd $GOPATH/src/github.com/devalecs/go-plugin-example
go build -buildmode=plugin -o hello.so hello_plugin.go
go run main.go
```

### Output
```
Hello Alex!
```
