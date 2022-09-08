package gotypo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/tarumes/goblin"
)

func (c *Client) Save(filename string) error {
	if strings.HasSuffix(filename, ".gob") {
		return goblin.GOBWrite(filename, c.dictionary)
	}

	if strings.HasSuffix(filename, ".json") {
		data, err := json.MarshalIndent(c.dictionary, "", "\t")
		if err != nil {
			return err
		}
		f, err := os.Create(filename)
		if err != nil {
			return err
		}
		if _, err = f.Write(data); err != nil {
			return err
		}
		f.Close()
		return nil
	}

	return fmt.Errorf("no valid filetype (.gob, .json) => %s", filename)
}

func (c *Client) Load(filename string) error {
	if strings.HasSuffix(filename, ".gob") {
		return goblin.GOBRead(filename, &c.dictionary)
	}

	if strings.HasSuffix(filename, ".json") {
		data, err := os.ReadFile(filename)
		if err != nil {
			return err
		}
		return json.Unmarshal(data, &c.dictionary)
	}

	return fmt.Errorf("no valid filetype (.gob, .json) => %s", filename)
}
