
# Minimalistic Go Web Server

The Minimalistic Go Web Server is a lightweight file server that allows clients to download files and upload files to the server using HTTP. It provides a simple and convenient way to share and manage files over the web.

## Features

- Download files by accessing the `/download/` endpoint followed by the desired file's path.
- Upload files to the server using the `/upload` endpoint.
- Supports file uploads through HTTP POST requests, making it compatible with cURL and web browsers.
- Logs client IP addresses and filenames whenever a file is downloaded.
- Lightweight and minimalistic design for easy deployment and usage.

## Requirements

- Go programming language (version 1.15 or later)

## Usage

1. Clone the repository:

   ```shell
   git clone https://github.com/ipcis/mini_webserver.git
   ```

2. Change into the project directory:

   ```shell
   cd mini_webserver
   ```

3. Start the server:

   ```shell
   go run mini_webserver.go
   ```
   or
     ```shell
   go run mini_webserver.go -port 8888
   ```

4. The server is now running on `http://localhost:8080`. You can access the following endpoints:

   - Download files: `http://localhost:8080/download/`
   - Upload files: `http://localhost:8080/upload`



## Client-side usage
   ```shell
   curl -O localhost:8080/download/test.txt
   ```
   ```shell
   curl -X POST -F "file=@AnyDesk.exe" http://localhost:8080/upload
   ```


## Compile (Windows)
```shell
SET GOOS=windows
```
```shell
SET GOARCH=amd64
```
```shell
go build -o bin/app-amd64.exe mini_webserver.go
```

## Compile (Linux)
 ```shell
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64 mini_webserver.go
   ```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).



