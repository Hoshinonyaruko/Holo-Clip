package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.design/x/clipboard"
)

func main() {
	var targetIP string
	var port string
	flag.StringVar(&targetIP, "ip", "localhost", "IP address of the target machine")
	flag.StringVar(&port, "port", "localhost", "IP address of the target machine")
	flag.Parse()

	err := clipboard.Init()
	if err != nil {
		log.Fatalf("Clipboard init failed: %v", err)
	}

	go func() {
		var lastTextData, lastImageData []byte
		for {
			textData := clipboard.Read(clipboard.FmtText)
			imageData := clipboard.Read(clipboard.FmtImage)
			if !bytes.Equal(textData, lastTextData) {
				sendClipboardData(targetIP, port, "text", textData)
				lastTextData = textData
			}
			if !bytes.Equal(imageData, lastImageData) {
				sendClipboardData(targetIP, port, "image", imageData)
				lastImageData = imageData
			}
			time.Sleep(time.Second)
		}
	}()

	http.HandleFunc("/sync", func(w http.ResponseWriter, r *http.Request) {
		formatStr := r.URL.Query().Get("format")
		log.Println("formatStr:", formatStr)
		format := stringToClipboardFormat(formatStr)
		encoded, _ := io.ReadAll(r.Body)
		decoded, _ := base64.StdEncoding.DecodeString(string(encoded))
		currentData := clipboard.Read(format)
		if !bytes.Equal(decoded, currentData) {
			clipboard.Write(format, decoded)
			log.Println("Clipboard updated from the other machine.")
		}
	})

	http.ListenAndServe(":"+port, nil)
}

func sendClipboardData(targetIP string, port string, safeFormat string, data []byte) {
	encoded := base64.StdEncoding.EncodeToString(data)

	targetURL := fmt.Sprintf("http://%s:%s/sync?format=%s", targetIP, port, safeFormat)

	resp, err := http.Post(targetURL, "application/octet-stream", bytes.NewReader([]byte(encoded)))

	if err != nil {
		log.Printf("Failed to sync: %v", err)
		return
	}

	defer resp.Body.Close()
}

func stringToClipboardFormat(str string) clipboard.Format {
	switch str {
	case "text":
		log.Println("text")
		return clipboard.FmtText
	case "image":
		log.Println("image")
		return clipboard.FmtImage
	default:
		log.Println("default")
		return clipboard.FmtText
	}
}
