package textanalyzer

import (
	"testing"

	"github.com/pkg/errors"
)

const exampleText = `Mention is made in many alchemical writings of a mythical personage named Hermes
Trismegistus, who is said to have lived a little later than the time of Moses.
Representations of Hermes Trismegistus are found on ancient Egyptian monuments. We
are told that Alexander the Great found his tomb near Hebron; and that the tomb
contained a slab of emerald whereon thirteen sentences were written. The eighth
sentence is rendered in many alchemical books as follows:
"Ascend with the greatest sagacity from the earth to heaven, and then again
descend to the earth, and unite together the powers of things superior and things
inferior. Thus you will obtain the glory of the whole world, and obscurity will fly
away from you."
This sentence evidently teaches the unity of things in heaven and things on
earth, and asserts the possibility of gaining, not merely a theoretical, but also a
practical, knowledge of the essential characters of all things. Moreover, the
sentence implies that this fruitful knowledge is to be obtained by examining

nature, using as guide the fundamental similarity supposed to exist between things
above and things beneath.
The alchemical writers constantly harp on this theme: follow nature; provided
you never lose the clue, which is simplicity and similarity.
The author of The Only Way (1677) beseeches his readers "to enlist under the
standard of that method which proceeds in strict obedience to the teaching of
nature ... in short, the method which nature herself pursues in the bowels of the
earth."`

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
			params:   "test 1 test 2",
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
			params:   "test 1 test 2",
			expected: 10,
		},
		{
			params:   "test,. 1 test 2",
			expected: 12,
		},
		{
			params: `test test 2
			test
			test
			
			test`,
			expected: 21,
		},
		{
			params:   exampleText,
			expected: 1277,
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
