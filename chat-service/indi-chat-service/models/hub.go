package models


type Hub struct {
	//Registered clients
	clients map[*Client]bool

	//Inbound messages from the clients
	broadcast chan []byte

	//Register reqs from the clients
	register chan *Client 

	//Unregister reqs from the clients
	unregister chan *Client
}

func newHub() *Hub{
	return &Hub{
		broadcast: make(chan []byte),
		register: make(chan *Client),
		unregister: make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (h *Hub) run(){
	for{
		//select statements listen for events on mutliple channels 
		/*
		The select statement in Go is used to wait on multiple communication operations. 
		It allows a goroutine to wait on multiple communication channels and proceed with the first channel that is ready for communication.
		*/
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:

			//h.clients[client] attempts to access the value associated with the key "client" int "h.clients" map
			if _,ok := h.clients[client]; ok {
				delete(h.clients,client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients,client)
				}
			}
		}
	}
}