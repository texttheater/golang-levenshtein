package levenshtein_test

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"testing"
)

var testCases = []struct {
	source   string
	target   string
	distance int
}{
	{"", "世", 1},
	{"世", "世世", 1},
	{"世", "世世世", 2},
	{"", "", 0},
	{"世", "界", 2},
	{"世世世", "世界世", 2},
	{"世世世", "世界", 3},
	{"世", "世", 0},
	{"世界", "世界", 0},
	{"世", "", 1},
	{"世世", "世", 1},
	{"世世世", "世", 2},
}

func TestLevenshtein(t *testing.T) {
	for _, testCase := range testCases {
		distance := levenshtein.DistanceForStrings(
			[]rune(testCase.source),
			[]rune(testCase.target),
			levenshtein.DefaultOptions)
		if distance != testCase.distance {
			t.Log(
				"Distance between",
				testCase.source,
				"and",
				testCase.target,
				"computed as",
				distance,
				", should be",
				testCase.distance)
			t.Fail()
		}
	}
}
