package textanalyzer

import (
	"reflect"
	"strings"
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		params   string
		expected error
	}{
		{
			params:   "",
			expected: errors.New("error empty text"),
		},
		{
			params:   testText4,
			expected: nil,
		},
	}

	for _, test := range tests {
		_, actual := New(test.params, false)

		if actual != test.expected && actual.Error() != test.expected.Error() {
			t.Errorf("result: %s | expected: %s", actual, test.expected)
		}
	}
}

func TestNumberOfChars(t *testing.T) {
	var tests = []struct {
		params   string
		expected int
	}{
		{
			params:   testText1,
			expected: 1273,
		},
		{
			params:   testText2,
			expected: 21,
		},
		{
			params:   testText3,
			expected: 12,
		},
		{
			params:   testText4,
			expected: 9,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params, false)
		actual := client.NumberOfChars()

		if actual != test.expected {
			t.Errorf("result: %d | expected: %d", actual, test.expected)
		}
	}
}

func TestNumberOfWords(t *testing.T) {
	var tests = []struct {
		params   string
		expected int
	}{
		{
			params:   testText1,
			expected: 247,
		},
		{
			params:   testText2,
			expected: 6,
		},
		{
			params:   testText3,
			expected: 4,
		},
		{
			params:   testText4,
			expected: 3,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params, false)
		actual := client.NumberOfWords()

		if actual != test.expected {
			t.Errorf("result: %d | expected: %d", actual, test.expected)
		}
	}
}

func TestNumberOfLines(t *testing.T) {
	var tests = []struct {
		params   string
		expected int
	}{
		{
			params:   testText1,
			expected: 22,
		},
		{
			params:   testText2,
			expected: 4,
		},
		{
			params:   testText3,
			expected: 1,
		},
		{
			params:   testText4,
			expected: 1,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params, false)
		actual := client.NumberOfLines()

		if actual != test.expected {
			t.Errorf("result: %d | expected: %d", actual, test.expected)
		}
	}
}

func TestFiveMostUsedWords(t *testing.T) {
	var tests = []struct {
		params   string
		expected []string
	}{
		{
			params:   testText1,
			expected: []string{"the", "of", "and", "things", "to"},
		},
		{
			params:   testText4,
			expected: []string{"test", "1"},
		},
		{
			params:   testText5,
			expected: []string{"the", "lorem", "ipsum", "of", "and"},
		},
	}

	for _, test := range tests {
		client, _ := New(test.params, false)
		actual := client.FiveMostUsedWords()

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("result: %#v | expected: %#v", actual, test.expected)
		}
	}
}

func BenchmarkNewWithSmallText(b *testing.B) {
	smallText := strings.Repeat(testText1, 10)

	for n := 0; n < b.N; n++ {
		New(smallText, false)
	}
}

func BenchmarkNewWithMediumText(b *testing.B) {
	mediumText := strings.Repeat(testText1, 100)

	for n := 0; n < b.N; n++ {
		New(mediumText, false)
	}
}

func BenchmarkNewWithHugeText(b *testing.B) {
	hugeText := strings.Repeat(testText1, 1000)

	for n := 0; n < b.N; n++ {
		New(hugeText, false)
	}
}
