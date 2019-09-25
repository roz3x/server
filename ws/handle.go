package ws

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	fin        = 1 << 7
	rsv1       = 1 << 6
	rsv2       = 1 << 5
	rsv3       = 1 << 4
	textmode   = 1
	binarymode = 2
)

var (
	buff    *bufio.ReadWriter
	newConn net.Conn
)

//Ws serves ws
func Ws(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v\n\n", r.Header)
	hiJ, ok := w.(http.Hijacker)
	if !ok {
		fmt.Printf("error")
	}
	conn, _, err := hiJ.Hijack()
	if err != nil {
		fmt.Printf("error")
	}
	hash := sha1.New()
	io.WriteString(hash, r.Header["Sec-Websocket-Key"][0])
	io.WriteString(hash, "258EAFA5-E914-47DA-95CA-C5AB0DC85B11")
	hashsum := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	resp := make([]byte, 0)
	resp = append(resp, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\nSec-WebSocket-Accept: "...)
	resp = append(resp, hashsum...)
	resp = append(resp, "\r\n\r\n"...)
	_, err = conn.Write(resp)
	if err != nil {
		fmt.Printf("error")
	}
	conn.SetDeadline(time.Time{})

	payload := "hello firefox i am having fun"
	length := len(payload)
	msg := make([]byte, 0, 1024)
	msg = append(msg, byte(fin|textmode), byte(length))
	msg = append(msg, payload...)
	conn.Write(msg)
	fmt.Printf("%v\n", msg)
	_, err = conn.Read(msg)
	fmt.Printf("%v\n", msg)
	time.Sleep(3 * time.Second)
	conn.Close()
}
