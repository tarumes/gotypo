package gotypo

import (
	"regexp"
	"sync"
)

type Client struct {
	dictionary dict
	lock       sync.Locker

	strip regexp.Regexp
}

type dict struct {
	Wordlist        []string `json:"wordlist"`
	StripFromString string   `json:"strip"`
}

func New() *Client {
	reply := &Client{
		dictionary: dict{
			Wordlist: []string{},
		},
		lock:  &sync.Mutex{},
		strip: *regexp.MustCompile(``),
	}

	return reply
}
