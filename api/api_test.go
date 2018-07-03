package api

import (
	"net/http"
	"testing"

	"github.com/pkg/errors"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		params   Config
		expected error
	}{
		{
			params:   Config{},
			expected: errors.New("invalid port"),
		},
		{
			params:   Config{Port: ":8080"},
			expected: errors.New("empty handlers map"),
		},
		{
			params: Config{
				Port: ":8080",
				Handlers: map[string]func(w http.ResponseWriter, r *http.Request){
					"/whatever": func(w http.ResponseWriter, r *http.Request) {},
				}},
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
