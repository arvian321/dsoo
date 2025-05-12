package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func main() {
	prosesOrderTicker := time.NewTicker(5 * time.Second)
	dokterAcceptTicker := time.NewTicker(1 * time.Minute)

	defer prosesOrderTicker.Stop()
	defer dokterAcceptTicker.Stop()

	for {
		select {
		case <-prosesOrderTicker.C:
			go hitProsesOrder()
		case <-dokterAcceptTicker.C:
			go hitDokterAccept()
		}
	}
}

func hitProsesOrder() {
	url := "https://dokteronline.k24.co.id/apiDokter/prosesOrder"

	// Contoh payload, bisa dikosongkan jika tidak perlu
	jsonPayload := []byte(`{"order_id":"12345"}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	//fmt.Printf("POST /prosesOrder Status: %s\n", resp.Status)
}

func hitDokterAccept() {
	url := "https://dokteronline.k24.co.id/apiDokter/dokterAceept"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	//fmt.Printf("GET /dokterAceept Status: %s\n", resp.Status)
}
