## Minimal Webserver
(Minimal webserver in go)

The Go-based web server is a minimalistic file server that allows clients to download files and upload files to the server. It provides a simple and straightforward way to share files over HTTP. Clients can download files by accessing the /download/ endpoint followed by the desired file's path. The server retrieves the file from the specified directory and sends it as a download to the client. Additionally, clients can upload files to the server using the /upload endpoint. The program handles file uploads through HTTP POST requests, allowing users to easily upload files either through cURL or a web browser. Uploaded files are saved in a designated directory on the server. The program logs information such as the client's IP address and the filename whenever a file is downloaded, providing a convenient way to keep track of file access. With its simplicity and functionality, the Go web server is a versatile tool for sharing and managing files over HTTP.

<code>go run mini_webserver.go</code>

<code>curl -O localhost:8080/download/test.txt</code>

<code>curl -X POST -F "file=@AnyDesk.exe" http://localhost:8080/upload</code>


## Compile (Windows)
SET GOOS=windows

SET GOARCH=amd64

go build -o bin/app-amd64.exe app.go

## Compile (Linux)
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64 app.go
