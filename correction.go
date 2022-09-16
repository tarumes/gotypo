package gotypo

import (
	"regexp"
	"strings"
	"sync"
)

var space regexp.Regexp = *regexp.MustCompile(`(\s|⠀)+`)
var numeric regexp.Regexp = *regexp.MustCompile(`(\$|€|)([0-9]{1,}\.[0-9]{1,}|[0-9]{1,})(\$|€|)`)

type Corrections struct {
	Result  string
	Matches [][]string
}

func (c *Client) TypoCorrections(sentence string) Corrections {
	var result Corrections
	var splitted []string = strings.Split(space.ReplaceAllString(sentence, " "), " ")
	var await sync.WaitGroup

	tempresult := make([][]string, len(splitted))
	for i, item := range splitted {
		await.Add(1)
		go func(v string, iter int) {
			var tmp []string = []string{}
			correct := c.TypoCorrection(v)
			tmp = append(tmp, correct.Best)
			tmp = append(tmp, correct.Same...)
			tempresult[iter] = tmp
			await.Done()
		}(item, i)
	}
	await.Wait()
	result.Matches = make([][]string, len(tempresult))
	for i, v := range tempresult {
		result.Result = result.Result + v[0] + " "
		result.Matches[i] = append(result.Matches[i], v[1:]...)
	}
	result.Result = strings.TrimSpace(result.Result)
	return result
}

type Correction struct {
	Best string
	Same []string
	Rank int
}

func (c *Client) TypoCorrection(input string) Correction {
	var result Correction = Correction{
		Best: input,
		Same: []string{},
		Rank: 100,
	}

	if numeric.MatchString(input) || input == "" {
		return result
	}

	input = c.strip.ReplaceAllString(input, "")

	for _, v := range c.dictionary.Wordlist {
		if v == input {
			result.Best = input
			result.Rank = 1
		} else if strings.EqualFold(v, input) {
			result.Best = v
			result.Rank = 1
		}
	}

	var others []string = []string{}
	for _, word := range c.dictionary.Wordlist {
		if word[0] == input[0] {
			rank(input, word, &result)
		} else if strings.EqualFold(word, input) {
			rank(input, word, &result)
		} else {
			others = append(others, word)
		}
	}

	if result.Rank == 1 && len(result.Same) == 0 || result.Rank > 1 {
		for _, word := range others {
			rank(input, word, &result)
		}
	}

	return result
}
