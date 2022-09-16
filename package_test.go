package gotypo

import (
	"os"
	"testing"
)

var learn []string = []string{
	"hello", "mother", "im", "home", "tree", "elf", "goblin", "soup", "women",
	"pace", "symptom", "tycoon", "undermine", "egg", "gem", "financial", "merit",
	"linger", "orgy", "menu", "sphere", "presidency", "literacy", "general", "bar",
	"transaction", "executrix", "knowledge", "strict", "frame", "glimpse", "debut",
	"tired", "contain", "personal", "wrong", "dump", "wheat", "carrot",
	"ritual", "admire", "friendly", "solve", "interference", "change", "man",
	"registration", "nationalism", "program", "valid", "connection", "president",
	"carve", "slime", "reproduce", "grow", "favourite", "version", "prosecute",
}

func TestCorrection(t *testing.T) {
	tc := New()
	err := tc.LearnFilter(`[\.:,;_\-+*!"/()=?\\\]\[]`)
	if err != nil {
		t.Error(err)
	}
	err = tc.LearnWord(learn)
	if err != nil {
		t.Error(err)
	}

	var result string = "hello mother im home"
	cor := tc.TypoCorrections(result)
	if cor.Result != result {
		t.Error("expected output not match", cor.Result, result)
	}

	cor = tc.TypoCorrections(`Henlo muddde, in hone`)
	if cor.Result != result {
		t.Error("expected output not match", cor.Result, result)
	}

	cor = tc.TypoCorrections(" ")
	if cor.Result != "" {
		t.Error("expected output not match", cor.Result, result)
	}

}

func TestSaveLoad(t *testing.T) {
	tc := New()
	err := tc.LearnFilter(`[\.:,;_\-+*!"/()=?\\\]\[]`)
	if err != nil {
		t.Error(err)
	}

	err = tc.LearnWord(learn)
	if err != nil {
		t.Error(err)
	}

	var filename string = "testfile_843579632sdugfhiae"

	os.Remove(filename + ".gob")
	err = tc.Save(filename + ".gob")
	if err != nil {
		t.Error(err)
	}
	err = tc.Load(filename + ".gob")
	if err != nil {
		t.Error(err)
	}
	err = os.Remove(filename + ".gob")
	if err != nil {
		t.Error(err)
	}

	os.Remove(filename + ".json")
	err = tc.Save(filename + ".json")
	if err != nil {
		t.Error(err)
	}
	err = tc.Load(filename + ".json")
	if err != nil {
		t.Error(err)
	}
	err = os.Remove(filename + ".json")
	if err != nil {
		t.Error(err)
	}

	err = tc.Load(filename)
	if err == nil {
		t.Error("invalid file load")
	}
	err = tc.Save(filename)
	if err == nil {
		t.Error("invalid file load")
	}
}

type test_lexensteindist struct {
	WordA    string
	WordB    string
	Distance int
}

func TestLevenstein(t *testing.T) {
	var teststrings []test_lexensteindist = []test_lexensteindist{
		{
			WordA:    "house",
			WordB:    "house",
			Distance: 0,
		},
		{
			WordA:    "house",
			WordB:    "mouse",
			Distance: 1,
		},
		{
			WordA:    "house",
			WordB:    "pneumonoultramicroscopicsilicovolcanoconiosis",
			Distance: 42,
		},
	}

	for _, v := range teststrings {
		if d := levensteinDistance(v.WordA, v.WordB); d != v.Distance {
			t.Error("distance is not ", v.Distance, "distance is", d)
		}
	}
}
