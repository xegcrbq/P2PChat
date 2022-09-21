package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	port := 0
	fmt.Print("Введите порт для приёма сообщений: ")
	fmt.Scan(&port)
	go server(strconv.Itoa(port))
	fmt.Print("Введите порт для отправки сообщений: ")
	fmt.Scan(&port)
	client(strconv.Itoa(port))
}

func server(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", messageReceiver)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, mux))

}
func client(port string) {
	URL := "http://127.0.0.1:" + port + "/send"
	for {
		messageToSend := ""
		fmt.Scan(&messageToSend)
		res, _ := http.Post(URL, "application/json", strings.NewReader(messageToSend))
		answer, _ := io.ReadAll(res.Body)
		fmt.Println(string(answer))
	}
}
func messageReceiver(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	fmt.Fprintf(w, `your message "%v" received`, string(data))
	fmt.Println("message: ", string(data))
}
