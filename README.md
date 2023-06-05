## Minimal Webserver
Minimal webserver in go

<code>go run mini_webserver.go</code>

<code>curl -O localhost:8080/download/test.txt</code>


## Compile (Windows)
SET GOOS=windows

SET GOARCH=amd64

go build -o bin/app-amd64.exe app.go

## Compile (Linux)
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64 app.go
