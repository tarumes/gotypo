package gotypo

func (c *Client) LearnWord(input []string) error {
	c.lock.Lock()
	c.dictionary.Wordlist = append(c.dictionary.Wordlist, input...)
	c.lock.Unlock()
	return nil
}
