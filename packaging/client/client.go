package client

type Client struct {
	id   int
	name string
}

func NewClient(id int, name string) *Client {
	return &Client{id, name}
}

func (c *Client) GetId() int {
	return c.id
}

func (c *Client) GetName() string {
	return c.name
}
