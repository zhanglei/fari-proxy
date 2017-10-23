package http

import (
	"testing"
	"net"
	"time"
	"bufio"
	"fmt"
	"log"
	"net/http"
)

var serverAddr = "127.0.0.1:20010"
var ciphertext = "uzon57jd0v869t7w"

func handle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if  ciphertext != r.Header.Get("Authorization") {
		log.Printf("ciphertext:%s - Authorization: %s", ciphertext, r.Header.Get("Authorization"))
	} else {
		log.Printf("match")
	}
	fmt.Fprintf(w, "hello")
}

func httpServer() {
	http.HandleFunc("/blog.html", handle)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func httpClient(plaintext string) {
	httpBody := NewHttp([]byte(plaintext))
	tcpAddr, _ := net.ResolveTCPAddr("tcp", serverAddr)

	conn, _:= net.DialTCP("tcp", nil, tcpAddr)

	conn.SetDeadline(time.Now().Add(1 * time.Minute))


	conn.Write([]byte(httpBody))

	reader := bufio.NewReader(conn)
	reply := make([]byte, 1024)
	reader.Read(reply)
}

func TestNewHttp(t *testing.T) {
	go httpClient("uzon57jd0v869t7w")
	httpServer()
}

func TestParseHttp(t *testing.T) {
	bodyLength := len(body)
	log.Printf("http header length %d\n", bodyLength)
	httpBody := NewHttp([]byte(ciphertext))
	parserd := string(ParseHttp(httpBody))
	if parserd != ciphertext {
		t.Errorf("http parse filaed.")
	} else {
		log.Printf("http parse success.")
	}
}