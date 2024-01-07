package models

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	//Time allowed to write a message to the peer
	writeWait = 10*time.Second

	//Time allowed to read the next pong message from the peer
	pongWait = 60*time.Second

	//Send pings to peer with this period .Must be less than pongwait
	pingPeriod = (pongWait * 9)/10

	//Maxium message size allowed from peer
	maxMessageSize = 512

)

/*
Type Consistency:
By using byte slices, you maintain consistency with other byte slice operations. 
For example, if you're working with byte slices for message manipulation, 
it's consistent to represent single bytes as byte slices rather than strings.

Flexibility:
Byte slices can represent more than just ASCII characters. 
If you need to work with binary data or non-ASCII characters, 
byte slices offer flexibility. 
Using []byte{'\n'} allows you to represent any byte value, 
not just ASCII characters.

Byte vs. String Semantics:
Bytes and strings have different semantics in Go. 
A byte slice represents a sequence of bytes, 
and each element is a numeric value in the range 0-255. 
A string, on the other hand, represents a sequence of Unicode code points. 
If you're working specifically with bytes (as in network protocols), 
it makes sense to use byte slices directly.

Avoiding String Allocation:
When you use a string literal like "\n" or " ", 
it creates a new string each time. 
In performance-critical sections, 
using byte slices directly can help avoid unnecessary string allocations.
*/
var (
	newline = []byte{'\n'}
	space = []byte{' '}
)


type Client struct{
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	WebsocketConn *websocket.Conn `json:"-"`
	hub *Hub
	send chan []byte
}

func NewClient(id int , username , email string , conn *websocket.Conn) *Client{
	return &Client{
		ID : id,
		Username: username,
		Email : email , 
		WebsocketConn: conn,
	}
}


