package main

import (
	"testing"
)

// example of what we expect from the shout function
var shoutTests = []struct {
	word     string
	shout    bool
	expected string
}{
	{"hello", true, "HELLO"},
	{"hello", false, "hello"},
	{"HEllo", true, "HELLO"},
	{"HEllo", false, "HEllo"},
}

// examples of what we expect from the repeat function
var repeatTests = []struct {
	word     string
	reps     int
	expected string
}{
	{"hello", 0, ""},
	{"hello", 1, "hello"},
	{"hello", 2, "hello hello"},
	{"GoodBye", 3, "GoodBye GoodBye GoodBye"},
}

func TestShouting(t *testing.T) {
	var actual string
	for _, curr := range shoutTests {
		actual = adjustVolume(curr.word, curr.shout)
		if actual != curr.expected {
			t.Errorf("Error when adjusting volume. word:%s shout:%t expected:%s actual:%s", curr.word, curr.shout, curr.expected, actual)
		}
	}
}

func TestRepeating(t *testing.T) {
	var actual string
	for _, curr := range repeatTests {
		actual = repeatWord(curr.word, curr.reps)
		if actual != curr.expected {
			t.Errorf("Error when repeating. word:%s reps:%d expected:%s actual:%s", curr.word, curr.reps, curr.expected, actual)
		}
	}
}

