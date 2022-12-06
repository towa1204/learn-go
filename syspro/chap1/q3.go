package main

// curl -v --compressed http://localhost:8080

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}

	gzipWriter := gzip.NewWriter(w)
	encoder := json.NewEncoder(io.MultiWriter(os.Stdout, gzipWriter))
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
	gzipWriter.Flush()
	gzipWriter.Close()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
