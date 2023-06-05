package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Verzeichnis, in dem sich die herunterladbaren Dateien befinden
	dir := "./"

	// Handler-Funktion für den Download von Dateien
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		// Pfad zum angeforderten Download
		filePath := dir + r.URL.Path[len("/download/"):]

		// Ermitteln der Client-IP-Adresse
		clientIP := r.RemoteAddr
		if index := strings.LastIndex(clientIP, ":"); index != -1 {
			clientIP = clientIP[:index]
		}

		// Protokollieren der heruntergeladenen Datei und der Client-IP-Adresse
		fmt.Printf("Client %s hat die Datei %s heruntergeladen.\n", clientIP, filePath)

		// Öffnen der Datei
		file, err := http.Dir(dir).Open(filePath)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer file.Close()

		// Ermitteln der Dateigröße
		stat, err := file.Stat()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fileSize := stat.Size()

		// Einstellen des HTTP-Headers
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", stat.Name()))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", fileSize))

		// Senden der Datei als HTTP-Response
		http.ServeContent(w, r, stat.Name(), stat.ModTime(), file)
	})

	// Starten des Servers auf Port 8080
	fmt.Println("Der Webserver ist gestartet. Sie können die Dateien unter http://localhost:8080/download/ herunterladen.")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
