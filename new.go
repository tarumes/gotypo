package gotypo

import "sync"

type Client struct {
	dictionary dict
	lock       sync.Locker
}

type dict struct {
	Wordlist []string `json:"wordlist"`
}

func New() *Client {
	return &Client{
		dictionary: dict{
			Wordlist: []string{},
		},
		lock: &sync.Mutex{},
	}
}
