package main

type Room struct {
	Clients []*Client
}

func (r *Room) AddClient(client *Client) {
	r.Clients = append(r.Clients, client)
}

func (r *Room) GetClients() []Client {
	var clients []Client

	for _, c := range r.Clients {
		clients = append(clients, *c)
	}
	return clients
}

func (r *Room) Broadcast(msg []byte) error {
	for _, c := range r.Clients {
		if err := c.Send(msg); err != nil {
			return err
		}
	}
	return nil
}
