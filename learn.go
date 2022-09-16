package gotypo

import (
	"fmt"
	"regexp"
)

func (c *Client) LearnWord(input []string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, v := range input {
		if contains(c.dictionary.Wordlist, v) == -1 {
			c.dictionary.Wordlist = append(c.dictionary.Wordlist, input...)
		}
	}
	fmt.Println(len(c.dictionary.Wordlist))
	return nil
}

func (c *Client) LearnFilter(input string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	reg, err := regexp.Compile(input)
	if err != nil {
		return err
	}
	c.dictionary.StripFromString = input
	c.strip = *reg
	return nil
}
