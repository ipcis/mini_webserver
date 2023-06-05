package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Directory where downloadable files are located
	dir := "./"

	// Handler function for file download
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		// Path to the requested download
		filePath := dir + r.URL.Path[len("/download/"):]

		// Open the file
		file, err := http.Dir(dir).Open(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		// Get file size
		stat, err := file.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileSize := stat.Size()

		// Set HTTP headers
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", stat.Name()))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))

		// Serve the file as an HTTP response
		http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
	})

	// Handler function for file upload
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request uses HTTP POST multipart/form-data
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Process the uploaded file form field
		file, handler, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Create the destination directory if it doesn't exist
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Create the destination file
		dstPath := filepath.Join(dir, handler.Filename)
		dstFile, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dstFile.Close()

		// Copy the uploaded file content to the destination file
		if _, err := io.Copy(dstFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Successful response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "File '%s' successfully uploaded.\n", handler.Filename)
	})

	// Start the server on port 8080
	fmt.Println("The web server has started. You can download files from http://localhost:8080/download/ and upload files to http://localhost:8080/upload.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
