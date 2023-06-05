## Minimal Webserver
(Minimal webserver in go)

The program is a minimal web server written in Go. It allows clients to upload and download files via HTTP. When a client requests a file, the server retrieves the file from the specified directory and sends it as a download to the client. The program also logs the client's IP address and the filename whenever a file is downloaded.

<code>go run mini_webserver.go</code>

<code>curl -O localhost:8080/download/test.txt</code>


## Compile (Windows)
SET GOOS=windows

SET GOARCH=amd64

go build -o bin/app-amd64.exe app.go

## Compile (Linux)
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64 app.go
