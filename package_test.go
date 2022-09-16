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
	cor := tc.TypoCorrections(`Henlo muddde, in hone`)
	if cor.Result != result {
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

	err = tc.Save("test.gob")
	if err != nil {
		t.Error(err)
	}
	err = tc.Load("test.gob")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("test.gob")

	err = tc.Save("test.json")
	if err != nil {
		t.Error(err)
	}
	err = tc.Load("test.json")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove("test.json")

}
