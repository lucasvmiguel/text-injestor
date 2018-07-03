package textanalyzer

import (
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

const text1 = `Mention is made in many alchemical writings of a mythical personage named Hermes
Trismegistus, who is said to have lived a little later than the time of Moses.
Representations of Hermes Trismegistus are found on ancient Egyptian monuments. We
are told that Alexander the Great found his tomb near Hebron; and that the tomb
contained a slab of emerald whereon thirteen sentences were written. The eighth
sentence is rendered many alchemical books as follows:
"Ascend with the greatest sagacity from the earth to heaven, and then again
descend to the earth, and unite together the powers of things superior and things
inferior. Thus you will obtain the glory of the whole world, and obscurity will fly
away from you."
This sentence evidently teaches the unity of things in heaven and things on
earth, and asserts the possibility of gaining, not merely a theoretical, but also a
practical, knowledge of the essential characters of all things. Moreover, the
sentence implies that this fruitful knowledge is be obtained by examining

nature, using as guide the fundamental similarity supposed to exist between things
above and things beneath.
The alchemical writers constantly harp on this theme: follow nature; provided
you never lose the clue, which is simplicity and similarity.
The author of The Only Way (1677) beseeches his readers "to enlist under the
standard of that method which proceeds in strict obedience to the teaching of
nature ... in short, the method which nature herself pursues in the bowels of the
earth."`

const text2 = `test test 2
test
	test

test`

const text3 = "Test,. 1 test 2"
const text4 = "test 1 test"
const text5 = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled to make a type specimen book. It has survived not only five centuries, but also the leap into electronic, remaining essentially unchanged. was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. lorem the lorem ipsum"

func TestNew(t *testing.T) {
	var tests = []struct {
		params   string
		expected error
	}{
		{
			params:   "",
			expected: errors.New("empty text"),
		},
		{
			params:   text4,
			expected: nil,
		},
	}

	for _, test := range tests {
		_, actual := New(test.params)

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
			params:   text1,
			expected: 1273,
		},
		{
			params:   text2,
			expected: 21,
		},
		{
			params:   text3,
			expected: 12,
		},
		{
			params:   text4,
			expected: 9,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params)
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
			params:   text1,
			expected: 247,
		},
		{
			params:   text2,
			expected: 6,
		},
		{
			params:   text3,
			expected: 4,
		},
		{
			params:   text4,
			expected: 3,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params)
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
			params:   text1,
			expected: 22,
		},
		{
			params:   text2,
			expected: 4,
		},
		{
			params:   text3,
			expected: 1,
		},
		{
			params:   text4,
			expected: 1,
		},
	}

	for _, test := range tests {
		client, _ := New(test.params)
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
			params:   text1,
			expected: []string{"the", "of", "and", "things", "to"},
		},
		{
			params:   text4,
			expected: []string{"test", "1"},
		},
		{
			params:   text5,
			expected: []string{"the", "lorem", "ipsum", "of", "and"},
		},
	}

	for _, test := range tests {
		client, _ := New(test.params)
		actual := client.FiveMostUsedWords()

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("result: %#v | expected: %#v", actual, test.expected)
		}
	}
}
