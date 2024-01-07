package models

import "time"

type Message struct{
	ID int `json:"id"`
	Sender *User `json:"sender"`
	Receiver *User `json:"receiver"`
	Content string `json:"sender"`
	SentAt  time.Time `json:"sent_at"`
	ReadStatus bool `json:"read_status"`
}

func NewMessage(id int , sender , receiver *User , content string , sentAt time.Time , readStatus bool) *Message{
	return &Message{
		ID: id,
		Sender: sender,
		Receiver: receiver,
		Content: content,
		SentAt: sentAt,
		ReadStatus: readStatus,
	}
}