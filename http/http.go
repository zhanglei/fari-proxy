package http

import (
	"bytes"
	"strconv"
	"fmt"
)

var body = "GET /blog.html HTTP/1.1\r\n" +
		"Accept:image/gif.image/jpeg,*/*\r\n" +
		"Accept-Language:zh-cn\r\n" +
		"Connection:Keep-Alive\r\n" +
		"Host:localhost\r\n" +
		"User-Agent:Mozila/4.0(compatible;MSIE5.01;Window NT5.0)\r\n" +
		"Accept-Encoding:gzip,deflate\r\n" +
		"Content-Length:"

func NewHttp(ciphertext []byte) []byte {
	httpBody := body + strconv.Itoa(len(ciphertext)) + "\r\n\r\n" + string(ciphertext)
	return []byte(httpBody)
}

func ParseHttp(msg []byte) []byte{
	header := bytes.Split(msg, []byte("\r\n"))
	fmt.Printf("%d\r\n", len(header))
	if (len(header) != 10) {
		fmt.Printf("error?")
	}
	lengthName := bytes.Split(header[7], []byte(":"))[0]
	length := bytes.Split(header[7], []byte(":"))[1]
	if string(lengthName) == "Content-Length" {
		contentLength, _ := strconv.Atoi(string(length))
		return msg[222 + len(length) + 4 : 222 + len(length) + 4 + contentLength]
	}
	return msg
}