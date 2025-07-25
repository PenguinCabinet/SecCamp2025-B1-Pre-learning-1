package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func http_server() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Success! GET Method by Server"))
	})

	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)

		w.Write([]byte(fmt.Sprintf("Success! POST Method.Body: %v by Server", string(buf.Bytes()))))
	})

	fmt.Printf("Server: The http server is open at localhost:8080\n")

	http.ListenAndServe("localhost:8080", nil)
}

func http_request() {
	fmt.Println("")

	/*Get request*/
	req, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatalf("Client: %s\n", err.Error())
	}
	data, err := io.ReadAll(req.Body)
	if err == nil {
		fmt.Printf("Client: Get request is sent.Result data: %v\n", string(data))
	} else {
		log.Fatalf("Client: %s\n", err.Error())
	}

	/*Post request*/
	req, err = http.Post("http://localhost:8080", "application/json", bytes.NewReader([]byte("{\"data\":\"test\"}")))
	if err != nil {
		log.Fatalf("Client: %s\n", err.Error())
	}
	data, err = io.ReadAll(req.Body)
	if err == nil {
		fmt.Printf("Client: Post request is sent.Result data: %v\n", string(data))
	} else {
		log.Fatalf("Client: %s\n", err.Error())
	}

}

func main() {
	go http_server()

	/*サーバーが起動するまで待機*/
	for {
		_, err := http.Get("http://localhost:8080")
		time.Sleep(100)
		if err == nil {
			break
		}
	}

	http_request()

}
