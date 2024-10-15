package model

import "time"

type Entity struct {
	Name    string
	Address Address
}

type Client struct {
	Entity
	Birthdate time.Time
}

func (client *Client) GetAge() int {
	now := time.Now().Year()

	return now - client.Birthdate.Year()
}

func NewClient(name string, address Address, birthdate time.Time) Client {
	return Client{
		Entity: Entity{
			Name:    name,
			Address: address,
		},
		Birthdate: birthdate,
	}
}

type LegalEntity struct {
	Entity
	CNPJ string
}
